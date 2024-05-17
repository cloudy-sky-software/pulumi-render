package gen

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"

	pulschema_pkg "github.com/cloudy-sky-software/pulschema/pkg"

	"github.com/pkg/errors"
)

const jsonMimeType = "application/json"

var services = map[string]string{"static_site": "staticSite", "web_service": "webService", "private_service": "privateService", "background_worker": "backgroundWorker", "cron_job": "cronJob"}

// Ensure determinism in the order in which the services are processed.
var sortedDiscriminatorValues = []string{"background_worker", "cron_job", "private_service", "static_site", "web_service"}

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

func fixEnvVarsPutEndpoint(openAPIDoc *openapi3.T) error {
	// The actual API for updating/replacing env vars
	// actually just takes a top-level array of objects.
	// We've nested it under an envVars property to help
	// with resource construction via Pulumi and handle
	// the conversion in the OnPreCreate provider callback.
	pathItem := openAPIDoc.Paths.Find("/services/{serviceId}/env-vars")
	if pathItem == nil {
		return errors.New("expected to find /services/{serviceId}/env-vars")
	}

	contract.Assertf(pathItem.Put != nil, "put operation is nil")

	reqBodySchema := pathItem.Put.RequestBody.Value.Content.Get(jsonMimeType)
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

func addDefaultProperty(typ *openapi3.SchemaRef, propName string, defaultValue string) {
	strPropSchema := openapi3.NewSchemaRef("", openapi3.NewSchema())
	strPropSchema.Value.Type = &openapi3.Types{openapi3.TypeString}
	strPropSchema.Value.Default = defaultValue

	typ.Value.Properties[propName] = strPropSchema
}

func addPropertyWithRef(typ *openapi3.SchemaRef, propName, ref string) {
	propSchema := openapi3.NewSchemaRef("#/components/schemas/"+ref, openapi3.NewSchema())

	typ.Value.Properties[propName] = propSchema
}

func getServiceDiscriminator(suffix string) (*openapi3.Discriminator, openapi3.SchemaRefs) {
	refs := make(openapi3.SchemaRefs, 0, len(sortedDiscriminatorValues))
	mapping := make(map[string]string)
	for _, discriminatorValue := range sortedDiscriminatorValues {
		ref := "#/components/schemas/" + services[discriminatorValue] + suffix
		mapping[discriminatorValue] = ref
		refs = append(refs, openapi3.NewSchemaRef(ref, openapi3.NewSchema()))
	}

	return &openapi3.Discriminator{
		PropertyName: "type",
		Mapping:      mapping,
	}, refs
}

func changeOperationSchemas(openAPIDoc *openapi3.T, path, operation string, schemaSuffix string, refs openapi3.SchemaRefs) {
	pathItem := openAPIDoc.Paths.Find(path)
	contract.Assertf(pathItem != nil, "path %s not found", path)

	discriminator, _ := getServiceDiscriminator(schemaSuffix)

	switch operation {
	case http.MethodGet:
		schemaRef := openapi3.NewSchemaRef("", &openapi3.Schema{OneOf: refs})
		schemaRef.Value.Discriminator = discriminator
		pathItem.Get.Responses.Status(200).Value.Content.Get(jsonMimeType).Schema = schemaRef
	case http.MethodPost:
		schemaRef := openapi3.NewSchemaRef("", &openapi3.Schema{OneOf: refs})
		schemaRef.Value.Discriminator = discriminator
		pathItem.Post.RequestBody.Value.Content.Get(jsonMimeType).Schema = schemaRef

		outputDiscriminator, outputRefs := getServiceDiscriminator("Output")
		outputSchemaRef := openapi3.NewSchemaRef("", &openapi3.Schema{OneOf: outputRefs})
		outputSchemaRef.Value.Discriminator = outputDiscriminator
		// POST /services endpoint returns a 201 Created status code.
		pathItem.Post.Responses.Status(201).Value.Content.Get(jsonMimeType).Schema = outputSchemaRef
	case http.MethodPatch:
		schemaRef := openapi3.NewSchemaRef("", &openapi3.Schema{OneOf: refs})
		schemaRef.Value.Discriminator = discriminator
		pathItem.Patch.RequestBody.Value.Content.Get(jsonMimeType).Schema = schemaRef

		outputDiscriminator, outputRefs := getServiceDiscriminator("Output")
		outputSchemaRef := openapi3.NewSchemaRef("", &openapi3.Schema{OneOf: outputRefs})
		outputSchemaRef.Value.Discriminator = outputDiscriminator
		pathItem.Patch.Responses.Status(200).Value.Content.Get(jsonMimeType).Schema = outputSchemaRef
	}
}

type operationAndSuffx struct {
	// The HTTP method.
	operation             string
	baseServiceSchemaName string
	// Flag to indicate if the original schemas in
	// the OpenAPI doc have the HTTP method suffix.
	// For example, servicePOST, servicePATCH,
	// staticSiteDetailsPOST etc.
	schemasSuffixedWithOperation bool
	// The newServiceSchemaSuffix to use for the new per-service type
	// schemas that will be created.
	// This must be in Title-case.
	newServiceSchemaSuffix string
}

func fixServiceSchemas(openAPIDoc *openapi3.T) error {
	pathsAndOperations := map[string][]operationAndSuffx{
		// Ordering is important here. We should process the
		// `/services/{serviceId}` endpoint first and its GET
		// operation first so that we can rename the non-suffixed
		// schemas to be "output" schemas that will be used as
		// response body schemas for the mutating endpoints.
		// Otherwise, we lose that detail when the POST
		// operation is processed and overwrite the
		// non-suffixed schemas with that of the
		// POST request body's.
		"/services/{serviceId}": {
			operationAndSuffx{operation: http.MethodGet, newServiceSchemaSuffix: "Output", schemasSuffixedWithOperation: false, baseServiceSchemaName: "service"},
			operationAndSuffx{operation: http.MethodPatch, newServiceSchemaSuffix: "Patch", schemasSuffixedWithOperation: true, baseServiceSchemaName: "servicePATCH"},
		},
		"/services": {
			// We'll fix the GET /services endpoint separately.
			operationAndSuffx{operation: http.MethodPost, newServiceSchemaSuffix: "", schemasSuffixedWithOperation: true, baseServiceSchemaName: "servicePOST"},
		},
	}

	for path, operationAndSuffixes := range pathsAndOperations {

		for _, operationAndSuffix := range operationAndSuffixes {
			operation := operationAndSuffix.operation
			newServiceSchemaSuffix := operationAndSuffix.newServiceSchemaSuffix
			schemasSuffixedWithOperation := operationAndSuffix.schemasSuffixedWithOperation
			baseServiceSchemaName := operationAndSuffix.baseServiceSchemaName

			baseServiceSchema, ok := openAPIDoc.Components.Schemas[baseServiceSchemaName]
			contract.Assertf(ok, "%s not found", baseServiceSchemaName)
			// Delete the `serviceDetails` and `type` properties from the base
			// service schema for this operation.
			delete(baseServiceSchema.Value.Properties, "serviceDetails")
			delete(baseServiceSchema.Value.Properties, "type")

			operationSchemas := make(openapi3.SchemaRefs, 0, len(services))

			for _, discriminatorValue := range sortedDiscriminatorValues {
				service := services[discriminatorValue]
				serviceDetailsSchemaName := service + "Details" + strings.ToUpper(operation)
				if !schemasSuffixedWithOperation {
					serviceDetailsSchemaName = service + "Details"
				}
				serviceDetailsSchema := openAPIDoc.Components.Schemas[serviceDetailsSchemaName]
				contract.Assertf(serviceDetailsSchema != nil, "schema %s not found", serviceDetailsSchemaName)

				serviceSchema := openapi3.NewAllOfSchema()
				serviceSchema.Title = pulschema_pkg.ToPascalCase(service)

				inlineSchema := openapi3.NewSchemaRef("", openapi3.NewSchema())
				inlineSchema.Value.Type = &openapi3.Types{openapi3.TypeObject}
				inlineSchema.Value.Properties = make(openapi3.Schemas)
				// Service details object should also use the same suffix
				// as the parent service schema object, if it's not empty.
				if newServiceSchemaSuffix != "" && !schemasSuffixedWithOperation {
					openAPIDoc.Components.Schemas[serviceDetailsSchemaName+newServiceSchemaSuffix] = serviceDetailsSchema
					addPropertyWithRef(inlineSchema, "serviceDetails", serviceDetailsSchemaName+newServiceSchemaSuffix)
				} else {
					addPropertyWithRef(inlineSchema, "serviceDetails", serviceDetailsSchemaName)
				}
				addDefaultProperty(inlineSchema, "type", discriminatorValue)

				serviceSchema.AllOf = append(serviceSchema.AllOf, openapi3.NewSchemaRef("#/components/schemas/"+baseServiceSchemaName, baseServiceSchema.Value), inlineSchema)

				openAPIDoc.Components.Schemas[service+newServiceSchemaSuffix] = openapi3.NewSchemaRef("", serviceSchema)

				operationSchemas = append(operationSchemas, openapi3.NewSchemaRef("#/components/schemas/"+service+newServiceSchemaSuffix, nil))
			}

			changeOperationSchemas(openAPIDoc, path, operation, newServiceSchemaSuffix, operationSchemas)
		}
	}

	return nil
}

// Extracts the inline schemas to schema types references
// and fixes the `service` property of the array items
// to use a discriminated oneOf definition for all
// available services.
func fixListServicesEndpoint(openAPIDoc *openapi3.T) {
	pathItem := openAPIDoc.Paths.Find("/services")
	contract.Assertf(pathItem != nil, "path /services not found")

	discriminator, outputRefs := getServiceDiscriminator("Output")
	discriminatedServicesSchemaRef := openapi3.NewSchemaRef("", &openapi3.Schema{OneOf: outputRefs})
	discriminatedServicesSchemaRef.Value.Discriminator = discriminator

	listServicesSchemaRef := pathItem.Get.Responses.Status(200).Value.Content.Get(jsonMimeType).Schema
	itemsSchema := listServicesSchemaRef.Value.Items
	itemsSchema.Value.Properties["service"] = discriminatedServicesSchemaRef

	openAPIDoc.Components.Schemas["listServices"] = listServicesSchemaRef
	openAPIDoc.Components.Schemas["listServicesResponse"] = itemsSchema
	openAPIDoc.Components.Schemas["listServices"].Value.Items = openapi3.NewSchemaRef("#/components/schemas/listServicesResponse", itemsSchema.Value)
	pathItem.Get.Responses.Status(200).Value.WithJSONSchemaRef(openapi3.NewSchemaRef("#/components/schemas/listServices", listServicesSchemaRef.Value))
}

func FixOpenAPIDoc(openAPIDoc *openapi3.T) error {
	if err := ensureOperationID(openAPIDoc); err != nil {
		return err
	}

	if err := fixEnvVarsPutEndpoint(openAPIDoc); err != nil {
		return err
	}

	// copyReadOnlyServiceProperties(openAPIDoc)

	if err := fixServiceSchemas(openAPIDoc); err != nil {
		return err
	}

	// fixGetServiceEndpoint(openAPIDoc)

	fixListServicesEndpoint(openAPIDoc)

	return nil
}
