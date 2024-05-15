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
func ensureOperationID(openAPIDoc *openapi3.T) error {
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

func addReadOnlyDiscriminatedProperty(openAPIDoc *openapi3.T, schemaName string, typeName string, defaultValue *string) error {
	typ, ok := openAPIDoc.Components.Schemas[schemaName]
	if !ok {
		return fmt.Errorf("schema type %s not found", schemaName)
	}

	strSchema := openapi3.NewSchemaRef("", openapi3.NewSchema())
	// strSchema.Value.ReadOnly = true
	strSchema.Value.Type = &openapi3.Types{openapi3.TypeString}
	if defaultValue != nil {
		strSchema.Value.Default = *defaultValue
	}
	typ.Value.Properties[typeName] = strSchema
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

	for _, suffix := range []string{"POST"} {
		schema, ok := openAPIDoc.Components.Schemas["service"+suffix]
		if !ok {
			return fmt.Errorf("service%s schema type not found", suffix)
		}

		discriminator := getDiscriminator(suffix)
		schema.Value.Properties["serviceDetails"].Value.Discriminator = discriminator

		// Add the `type` property to the individual schemas for each discriminator
		// mapping.
		for discriminatedValue, schemaRefPath := range discriminator.Mapping {
			valuePtr := discriminatedValue
			schemaName := strings.TrimPrefix(schemaRefPath, "#/components/schemas/")
			err := addReadOnlyDiscriminatedProperty(openAPIDoc, schemaName, "type", &valuePtr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// useAnyOfForServicePatchOperation replaces the oneOf definition
// with an anyOf definition due to the lack of a discriminator.
func useAnyOfForServicePatchOperation(openAPIDoc *openapi3.T) error {
	schema, ok := openAPIDoc.Components.Schemas["servicePATCH"]
	if !ok {
		return fmt.Errorf("servicePATCH schema type not found")
	}

	serviceDetails := schema.Value.Properties["serviceDetails"].Value

	schemaRefs := slices.Clone(serviceDetails.OneOf)
	serviceDetails.AnyOf = schemaRefs
	serviceDetails.OneOf = nil

	return nil
}

func pluckServiceObjectFromResponseBody(openAPIDoc *openapi3.T) error {
	pathItem := openAPIDoc.Paths.Find("/services")
	contract.Assertf(pathItem != nil, "endpoint path /services not found")

	pathItem.Post.Responses.Status(201).Value.Content.Get("application/json").Schema.Value.Properties = nil
	pathItem.Post.Responses.Status(201).Value.Content.Get("application/json").Schema.Ref = "#/components/schemas/service"
	return nil
}

func FixOpenAPIDoc(openAPIDoc *openapi3.T) error {
	if err := ensureOperationID(openAPIDoc); err != nil {
		return err
	}

	if err := patchEnvVarsPutEndpoint(openAPIDoc); err != nil {
		return err
	}

	if err := addServiceDiscriminator(openAPIDoc); err != nil {
		return err
	}

	if err := pluckServiceObjectFromResponseBody(openAPIDoc); err != nil {
		return err
	}

	if err := useAnyOfForServicePatchOperation(openAPIDoc); err != nil {
		return err
	}

	return nil
}
