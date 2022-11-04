package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/pkg/v3/resource/provider"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	fwCallback "github.com/cloudy-sky-software/pulumi-provider-framework/callback"
	fwRest "github.com/cloudy-sky-software/pulumi-provider-framework/rest"

	"github.com/golang/glog"
)

type renderProvider struct {
	name    string
	version string

	apiKey string

	clearCacheOnServiceUpdateDeployments string
}

var (
	handler  *fwRest.Provider
	callback fwCallback.ProviderCallback
)

func makeProvider(host *provider.HostClient, name, version string, pulumiSchemaBytes, openapiDocBytes, metadataBytes []byte) (pulumirpc.ResourceProviderServer, error) {
	p := &renderProvider{
		name:    name,
		version: version,
	}

	callback = p
	rp, err := fwRest.MakeProvider(host, name, version, pulumiSchemaBytes, openapiDocBytes, metadataBytes, callback)

	handler = rp.(*fwRest.Provider)

	return rp, err
}

func (p *renderProvider) GetAuthorizationHeader() string {
	return fmt.Sprintf("%s %s", authSchemePrefix, p.apiKey)
}

func (p *renderProvider) OnInvoke(ctx context.Context, req *pulumirpc.InvokeRequest) (*pulumirpc.InvokeResponse, error) {
	return nil, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (p *renderProvider) OnConfigure(_ context.Context, req *pulumirpc.ConfigureRequest) (*pulumirpc.ConfigureResponse, error) {
	apiKey, ok := req.GetVariables()["render:config:apiKey"]
	if !ok {
		// Check if it's set as an env var.
		envVarNames := handler.GetSchemaSpec().Provider.InputProperties["apiKey"].DefaultInfo.Environment
		for _, n := range envVarNames {
			v := os.Getenv(n)
			if v != "" {
				apiKey = v
			}
		}

		// Return an error if the API key is still empty.
		if apiKey == "" {
			return nil, errors.New("api key is required")
		}
	}

	logging.V(3).Info("Configuring Render API key")
	p.apiKey = apiKey

	clearCacheOnServiceUpdateDeployments, ok := req.GetVariables()["render:config:clearCacheOnServiceUpdateDeployments"]
	if ok {
		p.clearCacheOnServiceUpdateDeployments = clearCacheOnServiceUpdateDeployments
	}

	return &pulumirpc.ConfigureResponse{
		AcceptSecrets: true,
	}, nil
}

// Diff checks what impacts a hypothetical update will have on the resource's properties.
func (p *renderProvider) OnDiff(ctx context.Context, req *pulumirpc.DiffRequest, resourceTypeToken string, diff *resource.ObjectDiff, jsonReq *openapi3.MediaType) (*pulumirpc.DiffResponse, error) {
	if len(jsonReq.Schema.Value.AnyOf) == 0 {
		return nil, nil
	}

	changes := pulumirpc.DiffResponse_DIFF_SOME
	var replaces []string
	var diffs []string

	// HACK! Taking a shortcut to handle service type-specific updates.
	switch resourceTypeToken {
	case "render:services:StaticSite":
		replaces, diffs = p.determineDiffsAndReplacements(diff, handler.GetOpenAPIDoc().Components.Schemas["patchStaticSite"].Value.Properties)
	case "render:services:WebService":
		replaces, diffs = p.determineDiffsAndReplacements(diff, handler.GetOpenAPIDoc().Components.Schemas["patchWebService"].Value.Properties)
	}

	logging.V(3).Infof("Diff response: replaces: %v; diffs: %v", replaces, diffs)

	return &pulumirpc.DiffResponse{
		Changes:  changes,
		Replaces: replaces,
		Diffs:    diffs,
	}, nil
}

func (p *renderProvider) OnPreCreate(ctx context.Context, req *pulumirpc.CreateRequest, httpReq *http.Request) error {
	return nil
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.
func (p *renderProvider) OnPostCreate(ctx context.Context, req *pulumirpc.CreateRequest, outputs map[string]interface{}) (map[string]interface{}, error) {
	// Service resources return an object that has the deployment ID
	// for the newly-created resources, as well as the newly-created service
	// object.
	if service, serviceOk := outputs["service"]; serviceOk {
		glog.V(3).Info("Found service object in the response. Using that as the output result.")
		outputs = service.(map[string]interface{})
	}

	return outputs, nil
}

func (p *renderProvider) OnPreRead(ctx context.Context, req *pulumirpc.ReadRequest, httpReq *http.Request) error {
	return nil
}

func (p *renderProvider) OnPostRead(ctx context.Context, req *pulumirpc.ReadRequest, outputs map[string]interface{}) (map[string]interface{}, error) {
	return outputs, nil
}

func (p *renderProvider) OnPreUpdate(ctx context.Context, req *pulumirpc.UpdateRequest, httpReq *http.Request) error {
	return nil
}

func (p *renderProvider) OnPostUpdate(ctx context.Context, req *pulumirpc.UpdateRequest, httpReq http.Request, outputs map[string]interface{}) (map[string]interface{}, error) {
	resourceTypeToken := fwRest.GetResourceTypeToken(req.GetUrn())
	url := httpReq.URL.String()

	if resourceTypeToken != "render:services:Service" {
		return outputs, nil
	}

	// When a service is updated via the API, we should trigger a deployment.
	// Check if we should also request to clear the cache.
	// We are sort of cheating here by not using the API spec to derive the
	// required input for the cache clear request. A fancier approach can be
	// done in the future.
	var reqBody []byte
	if p.clearCacheOnServiceUpdateDeployments == "clear" {
		reqBody, _ = json.Marshal(map[string]string{"clearCache": "clear"})
	}
	clearCacheHTTPReq, createReqErr := handler.CreatePostRequest(ctx, url+"/deploys", reqBody, nil)
	if createReqErr != nil {
		glog.Warningf("Failed to create POST request object to clear the Render Service cache")
		return outputs, nil
	}

	resp, err := handler.GetHttpClient().Do(clearCacheHTTPReq)
	if err != nil {
		glog.Warningf("Service was updated successfully but triggering a deployment failed: %v. Trigger a deployment manually using the Render dashboard.", err)
	}

	if resp == nil {
		return outputs, nil
	}

	resp.Body.Close()
	return outputs, nil
}

func (p *renderProvider) executeResumeSerivce(ctx context.Context, serviceID string) error {
	httpReq, _ := handler.CreatePostRequest(ctx, handler.GetBaseURL()+"/services"+serviceID+"/resume", nil, nil)

	resp, err := handler.GetHttpClient().Do(httpReq)
	if err != nil {
		return errors.Wrapf(err, "calling resume service endpoint %s", serviceID)
	}

	resp.Body.Close()
	return nil
}

func (p *renderProvider) OnPreDelete(ctx context.Context, req *pulumirpc.DeleteRequest, httpReq *http.Request) error {
	return nil
}

func (p *renderProvider) OnPostDelete(ctx context.Context, req *pulumirpc.DeleteRequest) error {
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), defaultUnmarshalOpts)
	if err != nil {
		return errors.Wrap(err, "unmarshal input properties as propertymap")
	}

	resourceTypeToken := fwRest.GetResourceTypeToken(req.GetUrn())

	// Resume the service when the Suspend resource is deleted.
	if resourceTypeToken == "render:services:Suspend" {
		if err := p.executeResumeSerivce(ctx, inputs["serviceId"].StringValue()); err != nil {
			return errors.Wrap(err, "resuming service")
		}
	}

	return nil
}
