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

const jsonMimeType = "application/json"

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
				if ai.RequestValidationInput.Request.Header.Get("Authorization") != "" {
					return nil
				}

				return errors.New("authorization header is required")
			},
		},
	}

	if err := openapi3filter.ValidateRequest(ctx, requestValidationInput); err != nil {
		return errors.Wrap(err, "request validation failed")
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

func (p *renderProvider) getPathParamsMap(resourceTypeToken, apiPath, requestMethod string, inputs resource.PropertyMap) (map[string]string, error) {
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

	for _, param := range parameters {
		if param.Value.In != "path" {
			continue
		}

		paramName := param.Value.Name
		input := inputs[resource.PropertyKey(paramName)]
		if input.IsComputed() {
			pathParams[paramName] = input.Input().Element.StringValue()
		} else {
			pathParams[paramName] = input.StringValue()
		}
	}

	if len(pathParams) == 0 {
		return nil, errors.Errorf("did not find any path parameters in the openapi doc for %s", resourceTypeToken)
	}

	return pathParams, nil
}

func (p *renderProvider) replacePathParams(path string, pathParams map[string]string) string {
	for k, v := range pathParams {
		path = strings.ReplaceAll(path, fmt.Sprintf("{%s}", k), v)
	}

	return path
}
