package provider

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"

	"github.com/pkg/errors"
)

const (
	authHeaderName   = "Authorization"
	authSchemePrefix = "Bearer"
	jsonMimeType     = "application/json"
)

func (p *renderProvider) executeGet(
	ctx context.Context,
	typeToken string,
	inputs resource.PropertyMap) ([]byte, error) {
	crudMap, ok := p.metadata.ResourceCRUDMap[typeToken]
	if !ok {
		return nil, errors.Errorf("unknown resource type %s", typeToken)
	}
	if crudMap.R == nil {
		return nil, errors.Errorf("resource read endpoint is unknown for %s", typeToken)
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

		pathParams, err = p.getPathParamsMap(typeToken, httpEndpointPath, http.MethodGet, inputs)
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

	return body, nil
}

func (p *renderProvider) validateRequest(ctx context.Context, httpReq *http.Request, pathParams map[string]string) error {
	route, _, err := p.router.FindRoute(httpReq)
	if err != nil {
		return errors.Wrap(err, "finding route from router")
	}

	// Validate the request.
	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    httpReq,
		PathParams: pathParams,
		Route:      route,
		Options: &openapi3filter.Options{
			AuthenticationFunc: func(ctx context.Context, ai *openapi3filter.AuthenticationInput) error {
				authHeader := ai.RequestValidationInput.Request.Header.Get(authHeaderName)
				if authHeader == "" {
					return errors.New("authorization header is required")
				}
				if !strings.HasPrefix(authHeader, authSchemePrefix) {
					return errors.Errorf("unexpected auth scheme (expected %q)", authSchemePrefix)
				}

				bearerToken := strings.TrimPrefix(authHeader, fmt.Sprintf("%s ", authSchemePrefix))
				if bearerToken == "" {
					return errors.New("bearer token is required")
				}

				return nil
			},
		},
	}

	if err := openapi3filter.ValidateRequest(ctx, requestValidationInput); err != nil {
		return errors.Wrap(err, "request validation failed")
	}

	if httpReq.Body == nil {
		logging.V(3).Info("Request does not have a body. Skipping ContentLength adjustment...")
		return nil
	}

	// Update the original HTTP request's ContentLength since the
	// body might have changed due to default properties getting
	// added to it.
	clonedReq := httpReq.Clone(ctx)
	clonedBody, _ := ioutil.ReadAll(clonedReq.Body)
	newContentLength := int64(len(clonedBody))
	logging.V(3).Infof("REQUEST CONTENT LENGTH: current: %d, new: %d", httpReq.ContentLength, newContentLength)
	httpReq.ContentLength = newContentLength
	logging.V(3).Infof("UPDATED REQUEST BODY: %v", string(clonedBody))
	httpReq.Body = io.NopCloser(bytes.NewBuffer(clonedBody))

	return nil
}

func (p *renderProvider) getPathParamsMap(resourceTypeToken, apiPath, requestMethod string, properties resource.PropertyMap) (map[string]string, error) {
	pathParams := make(map[string]string)

	var parameters openapi3.Parameters

	switch requestMethod {
	case http.MethodGet:
		parameters = p.openapiDoc.Paths[apiPath].Get.Parameters
	case http.MethodPost:
		parameters = p.openapiDoc.Paths[apiPath].Post.Parameters
	case http.MethodPatch:
		parameters = p.openapiDoc.Paths[apiPath].Patch.Parameters
	case http.MethodPut:
		parameters = p.openapiDoc.Paths[apiPath].Put.Parameters
	case http.MethodDelete:
		parameters = p.openapiDoc.Paths[apiPath].Delete.Parameters
	default:
		return pathParams, nil
	}

	oldInputs := getOldInputs(properties)

	logging.V(3).Infof("Process path parameters with %v", properties)
	count := 0
	for _, param := range parameters {
		if param.Value.In != "path" {
			continue
		}

		count++
		paramName := param.Value.Name
		logging.V(3).Infof("Looking for path param %q in resource inputs", paramName)
		property, ok := properties[resource.PropertyKey(paramName)]
		// If the path param is not in the properties, check if
		// we have the old inputs, if we are dealing with the state
		// of an existing resource.
		if !ok {
			if oldInputs == nil {
				return nil, errors.Errorf("did not find value for path param %s in output props (old inputs was nil)", paramName)
			}

			property, ok = oldInputs[resource.PropertyKey(paramName)]
			if !ok {
				return nil, errors.Errorf("did not find value for path param %s in output props and old inputs", paramName)
			}
		}

		if property.IsComputed() {
			pathParams[paramName] = property.Input().Element.StringValue()
		} else if property.IsSecret() {
			pathParams[paramName] = property.SecretValue().Element.StringValue()
		} else {
			pathParams[paramName] = property.StringValue()
		}
	}

	numPathParams := len(pathParams)
	if numPathParams == 0 {
		return nil, errors.Errorf("did not find any path parameters in the openapi doc for %s", resourceTypeToken)
	}

	if numPathParams != count {
		return nil, errors.Errorf("not all path params were found in the properties (expected: %d, found: %d)", count, numPathParams)
	}

	return pathParams, nil
}

func (p *renderProvider) replacePathParams(path string, pathParams map[string]string) string {
	for k, v := range pathParams {
		path = strings.ReplaceAll(path, fmt.Sprintf("{%s}", k), v)
	}

	return path
}

func (p *renderProvider) determineDiffsAndReplacements(d *resource.ObjectDiff, properties openapi3.Schemas) ([]string, []string) {
	replaces := make([]string, 0)
	diffs := make([]string, 0)

	for propKey := range d.Adds {
		prop := string(propKey)
		// If the added property is not part of the PATCH operation schema,
		// then suggest a replacement triggered by this property.
		if _, ok := properties[prop]; !ok {
			replaces = append(replaces, prop)
		} else {
			diffs = append(diffs, prop)
		}
	}

	for propKey := range d.Updates {
		prop := string(propKey)
		// If the updated property is not part of the PATCH operation schema,
		// then suggest a replacement triggered by this property.
		if _, ok := properties[prop]; !ok {
			replaces = append(replaces, prop)
		} else {
			diffs = append(diffs, prop)
		}
	}

	for propKey := range d.Deletes {
		prop := string(propKey)
		// If the deleted property is not part of the PATCH operation schema,
		// then suggest a replacement triggered by this property.
		if _, ok := properties[prop]; !ok {
			replaces = append(replaces, prop)
		} else {
			diffs = append(diffs, prop)
		}
	}

	return replaces, diffs
}
