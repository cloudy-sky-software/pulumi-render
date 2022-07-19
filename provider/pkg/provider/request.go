package provider

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/pkg/errors"
)

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
	}
	if err := openapi3filter.ValidateRequest(ctx, requestValidationInput); err != nil {
		return errors.Wrap(err, "request validation failed")
	}

	return nil
}

func (p *renderProvider) getPathParamsMap(resourceTypeToken, apiPath string, inputs resource.PropertyMap) (map[string]string, error) {
	pathParams := make(map[string]string)

	for _, param := range p.openapiDoc.Paths[apiPath].Post.Parameters {
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
