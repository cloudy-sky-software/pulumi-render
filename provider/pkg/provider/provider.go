package provider

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/pkg/v3/resource/provider"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	fwCallback "github.com/cloudy-sky-software/pulumi-provider-framework/callback"
	fwRest "github.com/cloudy-sky-software/pulumi-provider-framework/rest"
	"github.com/cloudy-sky-software/pulumi-provider-framework/state"
)

type renderProvider struct {
	fwCallback.UnimplementedProviderCallback

	name    string
	version string

	apiKey string

	clearCacheOnServiceUpdateDeployments string
}

var (
	handler  *fwRest.Provider
	callback fwCallback.ProviderCallback
)

const (
	envVarResourceTypeToken       = "render:services:EnvVarsForService"
	customDomainResourceTypeToken = "render:services:CustomDomain"
	secretFilesResourceTypeToken  = "render:services:SecretFilesForService"
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

func (p *renderProvider) OnPreInvoke(_ context.Context, _ *pulumirpc.InvokeRequest, _ *http.Request) error {
	return nil
}

func (p *renderProvider) OnPostInvoke(_ context.Context, _ *pulumirpc.InvokeRequest, _ interface{}) (map[string]interface{}, error) {
	return nil, nil
}

// OnConfigure is called by the provider framework when Pulumi calls Configure on
// the resource provider server.
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

// OnDiff checks what impacts a hypothetical update will have on the resource's properties.
func (p *renderProvider) OnDiff(_ context.Context, _ *pulumirpc.DiffRequest, resourceTypeToken string, diff *resource.ObjectDiff, jsonReq *openapi3.MediaType) (*pulumirpc.DiffResponse, error) {
	if len(jsonReq.Schema.Value.OneOf) == 0 {
		return nil, nil
	}

	changes := pulumirpc.DiffResponse_DIFF_SOME
	var replaces []string
	var diffs []string

	var allOfSchemaRefs openapi3.SchemaRefs
	// Taking a shortcut to handle service type-specific updates.
	switch resourceTypeToken {
	case "render:services:BackgroundWorker":
		allOfSchemaRefs = handler.GetOpenAPIDoc().Components.Schemas["backgroundWorkerPatch"].Value.AllOf
	case "render:services:CronJob":
		allOfSchemaRefs = handler.GetOpenAPIDoc().Components.Schemas["cronJobPatch"].Value.AllOf
	case "render:services:PrivateService":
		allOfSchemaRefs = handler.GetOpenAPIDoc().Components.Schemas["privateServicePatch"].Value.AllOf
	case "render:services:StaticSite":
		allOfSchemaRefs = handler.GetOpenAPIDoc().Components.Schemas["staticSitePatch"].Value.AllOf
	case "render:services:WebService":
		allOfSchemaRefs = handler.GetOpenAPIDoc().Components.Schemas["webServicePatch"].Value.AllOf
	}

	allProps := make(openapi3.Schemas, 0)
	for _, schemaRef := range allOfSchemaRefs {
		for k, v := range schemaRef.Value.Properties {
			allProps[k] = v
		}
	}

	replaces, diffs = p.determineDiffsAndReplacements(diff, allProps)

	logging.V(3).Infof("Diff response: replaces: %v; diffs: %v", replaces, diffs)

	return &pulumirpc.DiffResponse{
		Changes:  changes,
		Replaces: replaces,
		Diffs:    diffs,
	}, nil
}

func (p *renderProvider) OnPreCreate(_ context.Context, req *pulumirpc.CreateRequest, httpReq *http.Request) error {
	resourceTypeToken := fwRest.GetResourceTypeToken(req.GetUrn())
	if _, ok := arrayResourceTypeTokenToRootObjectName[resourceTypeToken]; !ok {
		return nil
	}

	logging.V(3).Infof("handling pre-create callback for %s", resourceTypeToken)

	modifiedRequestBody, err := unWrapArrayRequest(httpReq, resourceTypeToken)
	if err != nil {
		return errors.Wrap(err, "marshaling modified body in OnPreCreate")
	}
	httpReq.Body = io.NopCloser(bytes.NewBuffer(modifiedRequestBody))
	httpReq.ContentLength = int64(len(modifiedRequestBody))
	return nil
}

// OnPostCreate allocates a new instance of the provided resource and returns its unique ID afterwards.
func (p *renderProvider) OnPostCreate(_ context.Context, req *pulumirpc.CreateRequest, outputs interface{}) (map[string]interface{}, error) {
	resourceTypeToken := fwRest.GetResourceTypeToken(req.GetUrn())
	var outputsMap map[string]interface{}

	if _, ok := arrayResourceTypeTokenToRootObjectName[resourceTypeToken]; ok {
		outputsMap = wrapArrayResponse(outputs, resourceTypeToken)

		id := sha256.New()
		resp := outputs.([]interface{})
		b, _ := json.Marshal(resp)
		id.Write(b)
		// Array-type responses do not have a provider-generated "id".
		// We'll fake one for the sake of serializing the resource
		// into the state.
		outputsMap["id"] = fmt.Sprintf("%x", id.Sum(nil))

		return outputsMap, nil
	}

	// Creating a custom domain resource returns a list response of all custom domains.
	// We only need to return the custom domain object that this current resource
	// represents as the output map. The remaining custom domains are either implicit
	// domains (like apex records) or perhaps tied to other custom domain resources
	// the user may have created which would be associated with those resources.
	if resourceTypeToken == customDomainResourceTypeToken {
		inputs, err := plugin.UnmarshalProperties(req.GetProperties(), state.DefaultUnmarshalOpts)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal input properties as propertymap")
		}

		custDomResp := outputs.([]interface{})

		for _, c := range custDomResp {
			cMap := c.(map[string]interface{})
			if cMap["name"] == inputs["name"].StringValue() {
				outputsMap = cMap
				break
			}
		}

		return outputsMap, nil
	}

	// Service resources return an object that has the deployment ID
	// for the newly-created resources, as well as the newly-created service
	// object.
	outputsMap = outputs.(map[string]interface{})
	if service, serviceOk := outputsMap["service"]; serviceOk {
		logging.V(3).Info("Found service object in the response. Using that as the output result.")
		outputsMap = service.(map[string]interface{})
	}

	return outputsMap, nil
}

func (p *renderProvider) OnPreRead(_ context.Context, _ *pulumirpc.ReadRequest, _ *http.Request) error {
	return nil
}

func (p *renderProvider) OnPostRead(_ context.Context, req *pulumirpc.ReadRequest, outputs interface{}) (map[string]interface{}, error) {
	resourceTypeToken := fwRest.GetResourceTypeToken(req.GetUrn())

	if _, ok := arrayResourceTypeTokenToRootObjectName[resourceTypeToken]; ok {
		outputsMap := wrapArrayResponse(outputs, resourceTypeToken)
		// Re-use the previously-generated id from ReadRequest.
		outputsMap["id"] = req.Id
		return outputsMap, nil
	}

	return outputs.(map[string]interface{}), nil
}

func (p *renderProvider) OnPreUpdate(_ context.Context, req *pulumirpc.UpdateRequest, httpReq *http.Request) error {
	resourceTypeToken := fwRest.GetResourceTypeToken(req.GetUrn())
	if _, ok := arrayResourceTypeTokenToRootObjectName[resourceTypeToken]; !ok {
		return nil
	}

	logging.V(3).Infof("OnePreUpdate: Modifying request body for %s", resourceTypeToken)

	modifiedRequestBody, err := unWrapArrayRequest(httpReq, resourceTypeToken)
	if err != nil {
		return errors.Wrap(err, "marshaling modified body in OnPreUpdate")
	}
	httpReq.Body = io.NopCloser(bytes.NewBuffer(modifiedRequestBody))
	httpReq.ContentLength = int64(len(modifiedRequestBody))
	return nil
}

func (p *renderProvider) OnPostUpdate(ctx context.Context, req *pulumirpc.UpdateRequest, httpReq http.Request, outputs interface{}) (map[string]interface{}, error) {
	resourceTypeToken := fwRest.GetResourceTypeToken(req.GetUrn())
	var outputsMap map[string]interface{}

	if _, ok := arrayResourceTypeTokenToRootObjectName[resourceTypeToken]; ok {
		outputsMap = wrapArrayResponse(outputs, resourceTypeToken)

		id := sha256.New()
		resp := outputs.([]interface{})
		b, _ := json.Marshal(resp)
		id.Write(b)
		// Array-type responses do not have a provider-generated "id".
		// We'll fake one for the sake of serializing the resource
		// into the state.
		outputsMap["id"] = fmt.Sprintf("%x", id.Sum(nil))
	} else {
		outputsMap = outputs.(map[string]interface{})
	}

	if resourceTypeToken != "render:services:StaticSite" &&
		resourceTypeToken != "render:services:WebService" &&
		resourceTypeToken != "render:services:PrivateService" &&
		resourceTypeToken != "render:services:BackgroundWorker" &&
		resourceTypeToken != "render:services:CronJob" &&
		resourceTypeToken != envVarResourceTypeToken {
		return outputsMap, nil
	}

	// When a service or env var is updated via the API,
	// we should trigger a deployment.
	logging.V(3).Infof("Triggering a deployment since resource type %s was updated", resourceTypeToken)
	urlPath := httpReq.URL.Path
	urlPath = strings.ReplaceAll(urlPath, "/v1", "")
	urlPath = strings.ReplaceAll(urlPath, "/env-vars", "")
	urlPath = strings.TrimPrefix(urlPath, "/")

	parts := strings.Split(urlPath, "/")
	serviceID := parts[1]

	// TODO: We are sort of cheating here by not using the API spec to derive the
	// required input for the cache clear request. A fancier approach can be
	// done in the future.
	//
	// Check if we should also request to clear the cache.
	var reqBody []byte
	if p.clearCacheOnServiceUpdateDeployments == "clear" {
		reqBody, _ = json.Marshal(map[string]string{"clearCache": "clear"})
	} else {
		reqBody, _ = json.Marshal(map[string]string{"clearCache": "do_not_clear"})
	}

	inputs := resource.NewPropertyMapFromMap(map[string]interface{}{
		"serviceId": serviceID,
	})
	clearCacheHTTPReq, createReqErr := handler.CreatePostRequest(ctx, "/services/{serviceId}/deploys", reqBody, inputs)
	if createReqErr != nil {
		logging.Warningf("Failed to create POST request object to clear the cache: %v", createReqErr)
		return outputsMap, nil
	}

	resp, err := handler.GetHTTPClient().Do(clearCacheHTTPReq)
	if err != nil {
		logging.Warningf("Service was updated successfully but triggering a deployment failed: %v. Trigger a deployment manually using the Render dashboard.", err)
	}

	if resp == nil {
		return outputsMap, nil
	}

	logging.V(3).Infof("Response for triggering a deployment (code: %d)", resp.StatusCode)

	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	logging.V(3).Infof("Response body for triggering a deployment: %s", string(respBody))

	return outputsMap, nil
}

func (p *renderProvider) executeResumeSerivce(ctx context.Context, serviceID string) error {
	httpReq, _ := handler.CreatePostRequest(ctx, handler.GetBaseURL()+"/services"+serviceID+"/resume", nil, nil)

	resp, err := handler.GetHTTPClient().Do(httpReq)
	if err != nil {
		return errors.Wrapf(err, "calling resume service endpoint %s", serviceID)
	}

	resp.Body.Close()
	return nil
}

func (p *renderProvider) OnPreDelete(_ context.Context, _ *pulumirpc.DeleteRequest, _ *http.Request) error {
	return nil
}

func (p *renderProvider) OnPostDelete(ctx context.Context, req *pulumirpc.DeleteRequest) error {
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), state.DefaultUnmarshalOpts)
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
