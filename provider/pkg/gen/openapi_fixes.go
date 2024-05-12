package gen

import (
	"fmt"
	"net/url"
	"slices"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"

	pulschema_pkg "github.com/cloudy-sky-software/pulschema/pkg"

	"github.com/pkg/errors"
)

func getResourceFromPath(path string) (string, error) {
	u, err := url.Parse(path)
	if err != nil {
		return "", fmt.Errorf("invalid path: %s", path)
	}

	pathParts := strings.Split(u.Path, "/")
	var resourceName string
	var segmentIndex int
	// If the path ends with a path parameter,
	// we should use the path segment before
	// the last one.
	if strings.HasSuffix(path, "}") {
		segmentIndex = len(pathParts) - 2
	} else {
		// Otherwise, use the last segment.
		segmentIndex = len(pathParts) - 1
	}

	resourceName = pathParts[segmentIndex]
	return pulschema_pkg.ToPascalCase(resourceName), nil
}

func patchEnvVarsPutEndpoint(openAPIDoc *openapi3.T) error {
	// The actual API for updating/replacing env vars
	//actually just takes a top-level array of objects.
	// We've nested it under an envVars property to help
	// with resource construction via Pulumi and handle
	// the conversion in the OnPreCreate provider callback.
	pathItem := openAPIDoc.Paths.Find("/services/{serviceId}/env-vars")
	if pathItem == nil {
		return errors.New("expected to find /services/{serviceId}/env-vars")
	}

	contract.Assertf(pathItem.Put != nil, "put operation is nil")

	reqBodySchema := pathItem.Put.RequestBody.Value.Content.Get("application/json")
	originalSchema := reqBodySchema.Schema.Value
	reqBodySchema.Schema.Value = openapi3.NewSchema().WithProperties(map[string]*openapi3.Schema{
		"envVars": originalSchema,
	})

	return nil
}

// Render's API operations do not have an operationId,
// so we'll need to generate them based on the resource
// in the operation path.
func ensureOperationId(openAPIDoc *openapi3.T) error {
	for _, path := range openAPIDoc.Paths.InMatchingOrder() {
		pathItem := openAPIDoc.Paths.Find(path)
		if pathItem == nil {
			return errors.Errorf("path item not found for path %s", path)
		}

		resourceName, err := getResourceFromPath(path)
		if err != nil {
			return errors.Wrap(err, "getting resource name")
		}

		endsWithPathParam := strings.HasSuffix(path, "}")

		if pathItem.Get != nil && pathItem.Get.OperationID == "" {
			if endsWithPathParam {
				pathItem.Get.OperationID = "get" + resourceName
			} else {
				pathItem.Get.OperationID = "list" + resourceName
			}
		}

		if pathItem.Post != nil && pathItem.Post.OperationID == "" {
			pathItem.Post.OperationID = "create" + resourceName
		}

		if pathItem.Put != nil && pathItem.Put.OperationID == "" {
			// PUT request methods are special.
			// If the current path does not also
			// have a POST request endpoint, then
			// the PUT can be used to create the
			// resource but only if it also has a
			// DELETE endpoint. Otherwise, it's a
			// one-shot "replace" type of operation
			// on the current resource.
			if pathItem.Post == nil && pathItem.Delete != nil {
				pathItem.Put.OperationID = "create" + resourceName
			} else {
				pathItem.Put.OperationID = "replace" + resourceName
			}
		}

		if pathItem.Patch != nil && pathItem.Patch.OperationID == "" {
			pathItem.Patch.OperationID = "update" + resourceName
		}

		if pathItem.Delete != nil && pathItem.Delete.OperationID == "" {
			pathItem.Delete.OperationID = "delete" + resourceName
		}
	}

	return nil
}

func addServiceDiscriminator(openAPIDoc *openapi3.T) error {
	getDiscriminator := func(suffix string) *openapi3.Discriminator {
		return &openapi3.Discriminator{
			PropertyName: "type",
			Mapping: map[string]string{
				"static_site":       "#/components/schemas/staticSiteDetails" + suffix,
				"web_service":       "#/components/schemas/webServiceDetails" + suffix,
				"private_service":   "#/components/schemas/privateServiceDetails" + suffix,
				"background_worker": "#/components/schemas/backgroundWorkerDetails" + suffix,
				"cron_job":          "#/components/schemas/cronJobDetails" + suffix,
			},
		}
	}

	for _, suffix := range []string{"", "POST", "PATCH"} {
		schema, ok := openAPIDoc.Components.Schemas["service"+suffix]
		if !ok {
			return errors.New(fmt.Sprintf("service%s schema type not found", suffix))
		}

		schema.Value.Properties["serviceDetails"].Value.Discriminator = getDiscriminator(suffix)

		indexOfTypeProp := -1
		for i, prop := range schema.Value.Required {
			if prop == "type" {
				indexOfTypeProp = i
				break
			}
		}

		if indexOfTypeProp == -1 {
			continue
		}

		schema.Value.Required = slices.Delete(schema.Value.Required, indexOfTypeProp, indexOfTypeProp)
	}
	return nil
}

func addReadOnlyDiscriminatedProperty(openAPIDoc *openapi3.T, schemaName string, typeName string, defaultValue *string) error {
	typ, ok := openAPIDoc.Components.Schemas[schemaName]
	if !ok {
		return errors.New(fmt.Sprintf("schema type %s not found", schemaName))
	}

	strSchema := openapi3.NewSchemaRef("", openapi3.NewSchema())
	strSchema.Value.ReadOnly = true
	strSchema.Value.Type = &openapi3.Types{openapi3.TypeString}
	if defaultValue != nil {
		strSchema.Value.Default = *defaultValue
	}
	typ.Value.Properties[typeName] = strSchema
	return nil
}

func FixOpenAPIDoc(openAPIDoc *openapi3.T) error {
	if err := ensureOperationId(openAPIDoc); err != nil {
		return err
	}

	if err := patchEnvVarsPutEndpoint(openAPIDoc); err != nil {
		return err
	}

	if err := addServiceDiscriminator(openAPIDoc); err != nil {
		return err
	}

	schemas := map[string]string{"staticSiteDetails": "static_site", "webServiceDetails": "web_service", "privateServiceDetails": "private_service", "backgroundWorkerDetails": "background_worker", "cronJobDetails": "cron_job"}
	for schema, defaultDiscriminatedValue := range schemas {
		defaultDiscriminatedValuePtr := defaultDiscriminatedValue
		addReadOnlyDiscriminatedProperty(openAPIDoc, schema, "type", &defaultDiscriminatedValuePtr)
	}

	return nil
}
