// nolint: goconst
package gen

import (
	"bytes"
	"encoding/json"

	"github.com/cloudy-sky-software/pulumi-render/provider/pkg/gen/examples"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

const packageName = "render"

type OperationMap map[string]string

// CRUDOperationsMap identifies the endpoints to perform
// create, read, update and delete (CRUD) operations.
type CRUDOperationsMap struct {
	// C represents the POST (create) endpoint.
	C *string `json:"c,omitempty"`
	// R represents the GET (read) endpoint.
	R *string `json:"r,omitempty"`
	// U represents the PATCH endpoint.
	U *string `json:"u,omitempty"`
	// D represents the DELETE endpoint.
	D *string `json:"d,omitempty"`

	// P represents the PUT (overwrite/update) endpoint.
	P *string `json:"p,omitempty"`
}

type ProviderMetadata struct {
	// ResourceToOperationMap identifies the endpoint that will handle the creation of a resource.
	// ResourceToOperationMap OperationMap `json:"operationMap"`

	// ResourceToOperationMap identifies the endpoint that will
	// handle the CRUD for a given Pulumi resource type token.
	ResourceCRUDMap map[string]*CRUDOperationsMap `json:"crudMap"`
}

type resourceContext struct {
	mod               string
	pkg               *pschema.PackageSpec
	resourceName      string
	openapiComponents openapi3.Components
	visitedTypes      codegen.StringSet
}

// PulumiSchema will generate a Pulumi schema for the given k8s schema.
func PulumiSchema(openapiDoc openapi3.T) (pschema.PackageSpec, ProviderMetadata) {
	pkg := pschema.PackageSpec{
		Name:        packageName,
		Description: "A Pulumi package for creating and managing Render resources.",
		DisplayName: "Render",
		License:     "Apache-2.0",
		Keywords: []string{
			"pulumi",
			packageName,
			"category/cloud",
			"kind/native",
		},
		Homepage:   "https://cloudysky.software",
		Publisher:  "Cloudy Sky Software",
		Repository: "https://github.com/cloudy-sky-software/pulumi-render",

		Config: pschema.ConfigSpec{
			Variables: map[string]pschema.PropertySpec{
				"apiKey": {
					Description: "The Render API key",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Language: map[string]pschema.RawMessage{
						"csharp": rawMessage(map[string]interface{}{
							"name": "ApiKey",
						}),
					},
					Secret: true,
				},
				"clearCacheOnServiceUpdateDeployments": {
					Description: "When a service is updated, a deployment is automatically triggered. This variable controls whether or not the service cache should be cleared upon deployment.",
					TypeSpec:    pschema.TypeSpec{Ref: "#/types/render:services:DeployClearCache"},
					Language: map[string]pschema.RawMessage{
						"csharp": rawMessage(map[string]interface{}{
							"name": "ClearCacheOnServiceUpdateDeployments",
						}),
					},
				},
			},
		},

		Provider: pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Description: "The provider type for the render package.",
				Type:        "object",
			},
			InputProperties: map[string]pschema.PropertySpec{
				"apiKey": {
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"RENDER_APIKEY",
						},
					},
					Description: "The Render API key.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Language: map[string]pschema.RawMessage{
						"csharp": rawMessage(map[string]interface{}{
							"name": "ApiKey",
						}),
					},
					Secret: true,
				},
			},
		},

		// This URL should be the base releases URL for a version,
		// to which the plugin name is suffixed.
		// https://github.com/pulumi/pulumi/blob/c7e17febe5ff0396e25dca24f59ca98f6722edf2/sdk/go/common/workspace/plugins.go#L379
		PluginDownloadURL: "https://github.com/cloudy-sky-software/pulumi-render/releases/download/${VERSION}",
		Types:             map[string]pschema.ComplexTypeSpec{},
		Resources:         map[string]pschema.ResourceSpec{},
		Functions:         map[string]pschema.FunctionSpec{},
		Language:          map[string]pschema.RawMessage{},
	}

	csharpNamespaces := map[string]string{
		"render": "Render",
		// TODO: Is this needed?
		"": "Provider",
	}

	openapiContext := &openAPIContext{
		doc:             openapiDoc,
		pkg:             &pkg,
		resourceCRUDMap: make(map[string]*CRUDOperationsMap),
	}
	if err := openapiContext.gatherResourcesFromAPI(csharpNamespaces); err != nil {
		contract.Failf("generating resources from OpenAPI spec: %v", err)
	}

	// Add examples to resources
	for k, v := range examples.ResourceExample {
		if r, ok := pkg.Resources[k]; ok {
			r.Description += "\n\n" + v
			pkg.Resources[k] = r
		}
	}

	pkg.Language["csharp"] = rawMessage(map[string]interface{}{
		"rootNamespace": "Pulumi",
		"packageReferences": map[string]string{
			"Pulumi": "3.*",
		},
		"namespaces": csharpNamespaces,
		// TODO: What does this enable?
		// "dictionaryConstructors": true,
	})

	pkg.Language["go"] = rawMessage(map[string]interface{}{
		"importBasePath": "github.com/cloudy-sky-software/pulumi-render/sdk/go/render",
	})
	pkg.Language["nodejs"] = rawMessage(map[string]interface{}{
		"packageName": "@cloudyskysoftware/pulumi-render",
		"dependencies": map[string]string{
			"@pulumi/pulumi":    "^3.0.0",
			"shell-quote":       "^1.6.1",
			"tmp":               "^0.0.33",
			"@types/tmp":        "^0.0.33",
			"glob":              "^7.1.2",
			"@types/glob":       "^5.0.35",
			"node-fetch":        "^2.3.0",
			"@types/node-fetch": "^2.1.4",
		},
		"devDependencies": map[string]string{
			"mocha":              "^5.2.0",
			"@types/mocha":       "^5.2.5",
			"@types/shell-quote": "^1.6.0",
		},
	})
	pkg.Language["python"] = rawMessage(map[string]interface{}{
		"packageName": "pulumi_render",
		"requires": map[string]string{
			"pulumi": ">=3.0.0,<4.0.0",
		},
	})

	metadata := ProviderMetadata{
		ResourceCRUDMap: openapiContext.resourceCRUDMap,
	}
	return pkg, metadata
}

func rawMessage(v interface{}) pschema.RawMessage {
	var out bytes.Buffer
	encoder := json.NewEncoder(&out)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	contract.Assert(err == nil)
	return out.Bytes()
}
