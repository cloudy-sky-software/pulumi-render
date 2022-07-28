// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"

	"github.com/pkg/errors"

	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/pkg/v3/resource/provider"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	providerGen "github.com/cloudy-sky-software/pulumi-render/provider/pkg/gen"
	providerOpenAPI "github.com/cloudy-sky-software/pulumi-render/provider/pkg/openapi"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/glog"
	pbempty "github.com/golang/protobuf/ptypes/empty"
)

type renderProvider struct {
	host    *provider.HostClient
	name    string
	version string
	schema  pschema.PackageSpec

	baseURL    string
	openapiDoc openapi3.T
	metadata   providerGen.ProviderMetadata
	router     routers.Router

	httpClient                           *http.Client
	apiKey                               string
	clearCacheOnServiceUpdateDeployments string
}

func defaultTransportDialContext(dialer *net.Dialer) func(context.Context, string, string) (net.Conn, error) {
	return dialer.DialContext
}

func makeProvider(host *provider.HostClient, name, version string, pulumiSchemaBytes, openapiDocBytes, metadataBytes []byte) (pulumirpc.ResourceProviderServer, error) {
	openapiDoc := providerOpenAPI.GetOpenAPISpec(openapiDocBytes)

	router, err := gorillamux.NewRouter(openapiDoc)
	if err != nil {
		return nil, errors.Wrap(err, "creating api router mux")
	}

	var metadata providerGen.ProviderMetadata
	if err := json.Unmarshal(metadataBytes, &metadata); err != nil {
		return nil, errors.Wrap(err, "unmarshaling the metadata bytes to json")
	}

	httpClient := &http.Client{
		// The transport is mostly a copy of the http.DefaultTransport
		// with the exception of ForceAttemptHTTP2 set to false.
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: defaultTransportDialContext(&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}),
			ForceAttemptHTTP2:     false,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return errors.New("unable to handle redirects")
		},
	}

	var pulumiSchema pschema.PackageSpec
	if err := json.Unmarshal(pulumiSchemaBytes, &pulumiSchema); err != nil {
		return nil, errors.Wrap(err, "unmarshaling pulumi schema into its package spec form")
	}

	// Return the new provider
	return &renderProvider{
		host:       host,
		name:       name,
		version:    version,
		schema:     pulumiSchema,
		baseURL:    openapiDoc.Servers[0].URL,
		openapiDoc: *openapiDoc,
		metadata:   metadata,
		router:     router,
		httpClient: httpClient,
	}, nil
}

func getResourceTypeToken(u string) string {
	urn := resource.URN(u)
	return urn.Type().String()
}

// Attach sends the engine address to an already running plugin.
func (p *renderProvider) Attach(context context.Context, req *pulumirpc.PluginAttach) (*pbempty.Empty, error) {
	host, err := provider.NewHostClient(req.GetAddress())
	if err != nil {
		return nil, err
	}
	p.host = host
	return &pbempty.Empty{}, nil
}

// Call dynamically executes a method in the provider associated with a component resource.
func (p *renderProvider) Call(ctx context.Context, req *pulumirpc.CallRequest) (*pulumirpc.CallResponse, error) {
	return nil, status.Error(codes.Unimplemented, "call is not yet implemented")
}

// Construct creates a new component resource.
func (p *renderProvider) Construct(ctx context.Context, req *pulumirpc.ConstructRequest) (*pulumirpc.ConstructResponse, error) {
	return nil, status.Error(codes.Unimplemented, "construct is not yet implemented")
}

// CheckConfig validates the configuration for this provider.
func (p *renderProvider) CheckConfig(ctx context.Context, req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	return &pulumirpc.CheckResponse{Inputs: req.GetNews()}, nil
}

// DiffConfig diffs the configuration for this provider.
func (p *renderProvider) DiffConfig(ctx context.Context, req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	return &pulumirpc.DiffResponse{}, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (p *renderProvider) Configure(_ context.Context, req *pulumirpc.ConfigureRequest) (*pulumirpc.ConfigureResponse, error) {
	apiKey, ok := req.GetVariables()["render:config:apiKey"]
	if !ok {
		// Check if it's set as an env var.
		envVarNames := p.schema.Provider.InputProperties["apiKey"].DefaultInfo.Environment
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

// Invoke dynamically executes a built-in function in the provider.
func (p *renderProvider) Invoke(ctx context.Context, req *pulumirpc.InvokeRequest) (*pulumirpc.InvokeResponse, error) {
	args, err := plugin.UnmarshalProperties(req.Args, defaultUnmarshalOpts)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal input properties as propertymap")
	}

	invokeTypeToken := req.GetTok()

	body, err := p.executeGet(ctx, invokeTypeToken, args)
	if err != nil {
		return nil, errors.Wrap(err, "executing get request for invoke")
	}

	var obj resource.PropertyMap
	if strings.Contains(invokeTypeToken, ":list") {
		var outputs []interface{}
		if err := json.Unmarshal(body, &outputs); err != nil {
			return nil, errors.Wrap(err, "unmarshaling the response")
		}

		m := make(map[string]interface{})
		m["items"] = outputs
		obj = resource.NewPropertyMapFromMap(m)
	} else {
		var outputs map[string]interface{}
		if err := json.Unmarshal(body, &outputs); err != nil {
			return nil, errors.Wrap(err, "unmarshaling the response")
		}

		obj = resource.NewPropertyMapFromMap(outputs)
	}

	outputProperties, err := plugin.MarshalProperties(obj, defaultMarshalOpts)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling the output properties map")
	}

	return &pulumirpc.InvokeResponse{
		Return: outputProperties,
	}, nil
}

// StreamInvoke dynamically executes a built-in function in the provider. The result is streamed
// back as a series of messages.
func (p *renderProvider) StreamInvoke(req *pulumirpc.InvokeRequest, server pulumirpc.ResourceProvider_StreamInvokeServer) error {
	tok := req.GetTok()
	return fmt.Errorf("unknown StreamInvoke token '%s'", tok)
}

// Check validates that the given property bag is valid for a resource of the given type and returns
// the inputs that should be passed to successive calls to Diff, Create, or Update for this
// resource. As a rule, the provider inputs returned by a call to Check should preserve the original
// representation of the properties as present in the program inputs. Though this rule is not
// required for correctness, violations thereof can negatively impact the end-user experience, as
// the provider inputs are used for detecting and rendering diffs.
func (p *renderProvider) Check(ctx context.Context, req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	return &pulumirpc.CheckResponse{Inputs: req.News, Failures: nil}, nil
}

// Diff checks what impacts a hypothetical update will have on the resource's properties.
func (p *renderProvider) Diff(ctx context.Context, req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	oldState, err := plugin.UnmarshalProperties(req.GetOlds(), defaultUnmarshalOpts)
	if err != nil {
		return nil, err
	}

	olds := getOldInputs(oldState)
	if olds == nil {
		return nil, errors.New("fetching old inputs from the state")
	}

	news, err := plugin.UnmarshalProperties(req.GetNews(), defaultUnmarshalOpts)
	if err != nil {
		return nil, err
	}

	logging.V(3).Infof("Calculating diff: olds: %v; news: %v", olds, news)
	d := olds.Diff(news)
	if d == nil || !d.AnyChanges() {
		return &pulumirpc.DiffResponse{Changes: pulumirpc.DiffResponse_DIFF_NONE}, nil
	}

	logging.V(3).Info("Detected some changes...")
	logging.V(4).Infof("ADDS: %v", d.Adds)
	logging.V(4).Infof("DELETES: %v", d.Deletes)
	logging.V(4).Infof("UPDATES: %v", d.Updates)

	resourceTypeToken := getResourceTypeToken(req.GetUrn())
	crudMap, ok := p.metadata.ResourceCRUDMap[resourceTypeToken]
	if !ok {
		return nil, errors.Errorf("unknown resource type %s", resourceTypeToken)
	}
	if crudMap.U == nil {
		return nil, errors.Errorf("resource update endpoint is unknown for %s", resourceTypeToken)
	}

	patchOp := p.openapiDoc.Paths[*crudMap.U].Patch
	if patchOp == nil {
		return nil, errors.Errorf("openapi doc does not have patch endpoint definition for the path %s", *crudMap.U)
	}

	var replaces []string
	var diffs []string
	changes := pulumirpc.DiffResponse_DIFF_SOME
	jsonReq := patchOp.RequestBody.Value.Content[jsonMimeType]

	if len(jsonReq.Schema.Value.AnyOf) > 0 {
		// HACK! Taking a shortcut to handle service type-specific updates.
		switch resourceTypeToken {
		case "render:services:StaticSite":
			replaces, diffs = p.determineDiffsAndReplacements(d, p.openapiDoc.Components.Schemas["patchStaticSite"].Value.Properties)
		case "render:services:WebService":
			replaces, diffs = p.determineDiffsAndReplacements(d, p.openapiDoc.Components.Schemas["patchWebService"].Value.Properties)
		}
	} else if len(jsonReq.Schema.Value.Properties) != 0 {
		replaces, diffs = p.determineDiffsAndReplacements(d, jsonReq.Schema.Value.Properties)
	} else {
		changes = pulumirpc.DiffResponse_DIFF_UNKNOWN
	}

	logging.V(3).Infof("Diff response: replaces: %v; diffs: %v", replaces, diffs)

	return &pulumirpc.DiffResponse{
		Changes:  changes,
		Replaces: replaces,
		Diffs:    diffs,
	}, nil
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.
func (p *renderProvider) Create(ctx context.Context, req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error) {
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), defaultUnmarshalOpts)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal input properties as propertymap")
	}

	resourceTypeToken := getResourceTypeToken(req.GetUrn())
	crudMap, ok := p.metadata.ResourceCRUDMap[resourceTypeToken]
	if !ok {
		return nil, errors.Errorf("unknown resource type %s", resourceTypeToken)
	}
	if crudMap.C == nil {
		return nil, errors.Errorf("resource construction endpoint is unknown for %s", resourceTypeToken)
	}

	httpEndpointPath := *crudMap.C

	b, err := json.Marshal(inputs.Mappable())
	if err != nil {
		return nil, errors.Wrap(err, "marshaling inputs")
	}

	buf := bytes.NewBuffer(b)
	logging.V(3).Infof("REQUEST BODY: %s", string(b))
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, p.baseURL+httpEndpointPath, buf)
	if err != nil {
		return nil, errors.Wrap(err, "initializing request")
	}

	logging.V(3).Infof("URL: %s", httpReq.URL.String())

	// Set the API key in the auth header.
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.apiKey))
	httpReq.Header.Add("Accept", jsonMimeType)
	httpReq.Header.Add("Content-Type", jsonMimeType)

	hasPathParams := strings.Contains(httpEndpointPath, "{")
	var pathParams map[string]string
	// If the endpoint has path params, peek into the OpenAPI doc
	// for the param names.
	if hasPathParams {
		var err error
		pathParams, err = p.getPathParamsMap(resourceTypeToken, httpEndpointPath, http.MethodPost, inputs)
		if err != nil {
			return nil, errors.Wrap(err, "getting path params")
		}
	}

	if err := p.validateRequest(ctx, httpReq, pathParams); err != nil {
		return nil, errors.Wrap(err, "validate http request")
	}

	httpReq.URL.Path = p.replacePathParams(httpReq.URL.Path, pathParams)

	logging.V(3).Info("Executing create resource request")

	// Create the resource.
	httpResp, err := p.httpClient.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "executing http request")
	}

	if httpResp.StatusCode != http.StatusOK && httpResp.StatusCode != http.StatusCreated {
		body, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "http request failed and the error response could not be read")
		}

		httpResp.Body.Close()
		return nil, errors.Errorf("http request failed (status: %s): %s", httpResp.Status, string(body))
	}

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading response body")
	}

	httpResp.Body.Close()

	var outputs map[string]interface{}
	if err := json.Unmarshal(body, &outputs); err != nil {
		return nil, errors.Wrap(err, "unmarshaling the response")
	}

	logging.V(3).Infof("RESPONSE BODY: %v", outputs)

	// HACK! Service resources return an object that has the deployment ID
	// for the newly-created resources, as well as the newly-created service
	// object.
	if service, serviceOk := outputs["service"]; serviceOk {
		glog.V(3).Info("Found service object in the response. Using that as the output result.")
		outputs = service.(map[string]interface{})
	}

	outputProperties, err := plugin.MarshalProperties(getResourceState(outputs, inputs), defaultMarshalOpts)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling the output properties map")
	}

	id, ok := outputs["id"]
	if !ok {
		return nil, errors.New("resource may have been created successfully but the id was not present in the response")
	}

	return &pulumirpc.CreateResponse{
		Id:         id.(string), // ID's in Render are always strings.
		Properties: outputProperties,
	}, nil
}

// Read the current live state associated with a resource.
func (p *renderProvider) Read(ctx context.Context, req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error) {
	oldState, err := plugin.UnmarshalProperties(req.GetProperties(), defaultUnmarshalOpts)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal input properties as propertymap")
	}

	resourceTypeToken := getResourceTypeToken(req.GetUrn())
	crudMap, ok := p.metadata.ResourceCRUDMap[resourceTypeToken]
	if !ok {
		return nil, errors.Errorf("unknown resource type %s", resourceTypeToken)
	}
	if crudMap.R == nil {
		return nil, errors.Errorf("resource read endpoint is unknown for %s", resourceTypeToken)
	}

	httpEndpointPath := *crudMap.R

	httpReq, err := http.NewRequestWithContext(ctx, "GET", p.baseURL+httpEndpointPath, nil)
	if err != nil {
		return nil, errors.Wrap(err, "initializing request")
	}

	// Set the API key in the auth header.
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.apiKey))
	httpReq.Header.Add("Accept", jsonMimeType)
	httpReq.Header.Add("Content-Type", jsonMimeType)

	hasPathParams := strings.Contains(httpEndpointPath, "{")
	var pathParams map[string]string
	// If the endpoint has path params, peek into the OpenAPI doc
	// for the param names.
	if hasPathParams {
		var err error

		pathParams, err = p.getPathParamsMap(resourceTypeToken, httpEndpointPath, http.MethodGet, oldState)
		if err != nil {
			return nil, errors.Wrap(err, "getting path params")
		}
	}

	if err := p.validateRequest(ctx, httpReq, pathParams); err != nil {
		return nil, errors.Wrap(err, "validate http request")
	}

	httpReq.URL.Path = p.replacePathParams(httpReq.URL.Path, pathParams)

	// Read the resource.
	httpResp, err := p.httpClient.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "executing http request")
	}

	if httpResp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "http request failed and the error response could not be read")
		}

		httpResp.Body.Close()
		return nil, errors.Errorf("http request failed (status: %s): %s", httpResp.Status, string(body))
	}

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading response body")
	}

	httpResp.Body.Close()

	var outputs map[string]interface{}
	if err := json.Unmarshal(body, &outputs); err != nil {
		return nil, errors.Wrap(err, "unmarshaling the response")
	}

	inputs := getOldInputs(oldState)
	// If there is no old state, then persist the current outputs as the
	// "old" inputs going forward for this resource.
	if inputs == nil {
		inputs = resource.NewPropertyMapFromMap(outputs)
	} else {
		// Take the values from outputs and apply them to the inputs
		// so that the checkpoint is in-sync with the state in the
		// cloud provider.
		newState := resource.NewPropertyMapFromMap(outputs)
		inputs = applyDiffFromCloudProvider(newState, inputs)
	}

	outputProperties, err := plugin.MarshalProperties(getResourceState(outputs, inputs), defaultMarshalOpts)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling the output properties map")
	}

	id, ok := outputs["id"]
	if !ok {
		return nil, errors.New("looking up id property from the response")
	}

	// Serialize and return the calculated inputs.
	inputsRecord, err := plugin.MarshalProperties(inputs, defaultMarshalOpts)
	if err != nil {
		return nil, err
	}

	return &pulumirpc.ReadResponse{
		Id:         id.(string),
		Inputs:     inputsRecord,
		Properties: outputProperties,
	}, nil
}

// Update updates an existing resource with new values.
func (p *renderProvider) Update(ctx context.Context, req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error) {
	oldState, err := plugin.UnmarshalProperties(req.Olds, defaultUnmarshalOpts)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal olds as propertymap")
	}

	inputs, err := plugin.UnmarshalProperties(req.News, defaultUnmarshalOpts)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal news as propertymap")
	}

	resourceTypeToken := getResourceTypeToken(req.GetUrn())
	crudMap, ok := p.metadata.ResourceCRUDMap[resourceTypeToken]
	if !ok {
		return nil, errors.Errorf("unknown resource type %s", resourceTypeToken)
	}
	if crudMap.U == nil && crudMap.P == nil {
		return nil, errors.Errorf("neither resource update endpoints (update and put) are available for %s", resourceTypeToken)
	}

	var httpEndpointPath string
	var method string
	if crudMap.U != nil {
		httpEndpointPath = *crudMap.U
		method = http.MethodPatch
	} else {
		httpEndpointPath = *crudMap.P
		method = http.MethodPut
	}

	b, err := json.Marshal(inputs.Mappable())
	if err != nil {
		return nil, errors.Wrap(err, "marshaling inputs")
	}

	logging.V(3).Infof("REQUEST BODY: %s", string(b))
	buf := bytes.NewBuffer(b)
	httpReq, err := http.NewRequestWithContext(ctx, method, p.baseURL+httpEndpointPath, buf)
	if err != nil {
		return nil, errors.Wrap(err, "initializing request")
	}

	// Set the API key in the auth header.
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.apiKey))
	httpReq.Header.Add("Accept", jsonMimeType)
	httpReq.Header.Add("Content-Type", jsonMimeType)

	hasPathParams := strings.Contains(httpEndpointPath, "{")
	var pathParams map[string]string
	// If the endpoint has path params, peek into the OpenAPI doc
	// for the param names.
	if hasPathParams {
		var err error

		pathParams, err = p.getPathParamsMap(resourceTypeToken, httpEndpointPath, method, oldState)
		if err != nil {
			return nil, errors.Wrap(err, "getting path params")
		}
	}

	if err := p.validateRequest(ctx, httpReq, pathParams); err != nil {
		return nil, errors.Wrap(err, "validate http request")
	}

	httpReq.URL.Path = p.replacePathParams(httpReq.URL.Path, pathParams)

	// Update the resource.
	httpResp, err := p.httpClient.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "executing http request")
	}

	if httpResp.StatusCode != http.StatusOK && httpResp.StatusCode != http.StatusNoContent {
		return nil, errors.Errorf("http request failed: %v. expected 200 or 204 but got %d", err, httpResp.StatusCode)
	}

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading response body")
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode == http.StatusNoContent {
		return &pulumirpc.UpdateResponse{}, nil
	}

	var outputs map[string]interface{}
	if err := json.Unmarshal(body, &outputs); err != nil {
		return nil, errors.Wrap(err, "unmarshaling the response")
	}

	logging.V(3).Infof("RESPONSE BODY: %v", outputs)

	outputProperties, err := plugin.MarshalProperties(getResourceState(outputs, inputs), defaultMarshalOpts)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling the output properties map")
	}

	p.runPostUpdateAction(ctx, resourceTypeToken, httpReq.URL.String())

	return &pulumirpc.UpdateResponse{
		Properties: outputProperties,
	}, nil
}

func (p *renderProvider) runPostUpdateAction(ctx context.Context, resourceTypeToken, url string) {
	if resourceTypeToken != "render:services:Service" {
		return
	}

	// When a service is updated via the API, we should trigger a deployment.
	// Check if we should also request to clear the cache.
	// We are sort of cheating here by not using the API spec to derive the
	// required input. A fancier approach can be done in the future.
	var buf *bytes.Buffer
	if p.clearCacheOnServiceUpdateDeployments == "clear" {
		b, _ := json.Marshal(map[string]string{"clearCache": "clear"})
		buf = bytes.NewBuffer(b)
	}
	httpReq, _ := http.NewRequestWithContext(ctx, "POST", url+"/deploys", buf)

	// Set the API key in the auth header.
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.apiKey))
	httpReq.Header.Add("Accept", jsonMimeType)
	httpReq.Header.Add("Content-Type", jsonMimeType)

	resp, err := p.httpClient.Do(httpReq)
	if err != nil {
		glog.Warningf("Service was updated successfully. However, triggering a deployment failed: %v. You should trigger a deployment manually using the Render dashboard.", err)
	}

	if resp == nil {
		return
	}

	resp.Body.Close()
}

func (p *renderProvider) executeResumeSerivce(ctx context.Context, serviceID string) error {
	httpReq, _ := http.NewRequestWithContext(ctx, "POST", p.baseURL+"/services"+serviceID+"/resume", nil)

	// Set the API key in the auth header.
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.apiKey))
	httpReq.Header.Add("Accept", jsonMimeType)
	httpReq.Header.Add("Content-Type", jsonMimeType)

	resp, err := p.httpClient.Do(httpReq)
	if err != nil {
		return errors.Wrapf(err, "calling resume service endpoint %s", serviceID)
	}

	resp.Body.Close()
	return nil
}

// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed
// to still exist.
func (p *renderProvider) Delete(ctx context.Context, req *pulumirpc.DeleteRequest) (*pbempty.Empty, error) {
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), defaultUnmarshalOpts)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal input properties as propertymap")
	}

	resourceTypeToken := getResourceTypeToken(req.GetUrn())

	// Nothing to delete if the Suspend resource is being removed.
	// But we do need to resume the service.
	if resourceTypeToken == "render:services:Suspend" {
		if err := p.executeResumeSerivce(ctx, inputs["serviceId"].StringValue()); err != nil {
			return &pbempty.Empty{}, errors.Wrap(err, "resuming service")
		}
		return &pbempty.Empty{}, nil
	}

	crudMap, ok := p.metadata.ResourceCRUDMap[resourceTypeToken]
	if !ok {
		return nil, errors.Errorf("unknown resource type %s", resourceTypeToken)
	}
	if crudMap.D == nil {
		// Nothing to do to delete this resource,
		// simply drop it from the state.
		return &pbempty.Empty{}, nil
	}

	httpEndpointPath := *crudMap.D
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodDelete, p.baseURL+httpEndpointPath, nil)
	if err != nil {
		return nil, errors.Wrap(err, "initializing request")
	}

	// Set the API key in the auth header.
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.apiKey))
	httpReq.Header.Add("Accept", jsonMimeType)
	httpReq.Header.Add("Content-Type", jsonMimeType)

	hasPathParams := strings.Contains(httpEndpointPath, "{")
	var pathParams map[string]string
	// If the endpoint has path params, peek into the OpenAPI doc
	// for the param names.
	if hasPathParams {
		var err error

		pathParams, err = p.getPathParamsMap(resourceTypeToken, httpEndpointPath, http.MethodDelete, inputs)
		if err != nil {
			return nil, errors.Wrap(err, "getting path params")
		}
	}

	if err := p.validateRequest(ctx, httpReq, pathParams); err != nil {
		return nil, errors.Wrap(err, "validate http request")
	}

	httpReq.URL.Path = p.replacePathParams(httpReq.URL.Path, pathParams)

	// Delete the resource.
	httpResp, err := p.httpClient.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "executing http request")
	}

	if httpResp.StatusCode != http.StatusOK && httpResp.StatusCode != http.StatusNoContent {
		return nil, errors.Errorf("http request failed: %v. expected 200 or 204 but got %d", err, httpResp.StatusCode)
	}

	httpResp.Body.Close()

	return &pbempty.Empty{}, nil
}

// GetPluginInfo returns generic information about this plugin, like its version.
func (p *renderProvider) GetPluginInfo(context.Context, *pbempty.Empty) (*pulumirpc.PluginInfo, error) {
	return &pulumirpc.PluginInfo{
		Version: p.version,
	}, nil
}

// GetSchema returns the JSON-serialized schema for the provider.
func (p *renderProvider) GetSchema(ctx context.Context, req *pulumirpc.GetSchemaRequest) (*pulumirpc.GetSchemaResponse, error) {
	b, err := json.Marshal(p.schema)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling the schema")
	}

	return &pulumirpc.GetSchemaResponse{
		Schema: string(b),
	}, nil
}

// Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either
// reutrn a creation error or an initialization error). Since Cancel is advisory and non-blocking,
// it is up to the host to decide how long to wait after Cancel is called before (e.g.)
// hard-closing any gRPC connection.
func (p *renderProvider) Cancel(context.Context, *pbempty.Empty) (*pbempty.Empty, error) {
	return &pbempty.Empty{}, nil
}
