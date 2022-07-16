package openapi

import (
	_ "embed"

	"context"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

//go:embed openapi.yml
var openapi []byte

// GetOpenAPISpec read the openapi.yml spec file,
// validates it and returns the parsed doc.
func GetOpenAPISpec() *openapi3.T {
	doc, err := openapi3.NewLoader().LoadFromData(openapi)
	if err != nil {
		contract.Failf("Failed to load openapi.yml: %v", err)
	}

	ctx := context.Background()
	if err := doc.Validate(ctx); err != nil {
		contract.Failf("OpenAPI spec failed validation: %v", err)
	}

	return doc
}
