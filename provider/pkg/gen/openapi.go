package gen

import (
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/golang/glog"

	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

type openAPIContext struct {
	doc             openapi3.T
	pkg             *pschema.PackageSpec
	resourceCRUDMap map[string]*CRUDOperationsMap
}

func getRootPath(path string) string {
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	return parts[0]
}

func getParentPath(path string) string {
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	lastPathPart := parts[len(parts)-1]
	if !strings.HasPrefix(lastPathPart, "{") && !strings.HasSuffix(lastPathPart, "}") {
		return path
	}

	// Skip the last path part which contains a path parameter.
	return "/" + strings.Join(parts[0:len(parts)-1], "/")
}

func (o *openAPIContext) gatherResourcesFromAPI(csharpNamespaces map[string]string) error {
	for path, pathItem := range o.doc.Paths {
		// Capture the iteration variable `path` because we use its pointer
		// in the crudMap.
		currentPath := path
		module := getRootPath(currentPath)

		// TODO: TEMPORARY!
		if currentPath == "/services/{serviceId}/resume" ||
			strings.HasPrefix(currentPath, "/services/{serviceId}/jobs") ||
			currentPath == "/services/{serviceId}/custom-domains/{customDomainIdOrName}/verify" {
			continue
		}
		//

		glog.V(3).Infof("Processing path %s\n", currentPath)

		if pathItem.Get != nil {
			parentPath := getParentPath(currentPath)
			glog.V(3).Infof("GET: Parent path for %s is %s\n", currentPath, parentPath)

			jsonReq := pathItem.Get.Responses.Get(200).Value.Content.Get("application/json")
			if jsonReq.Schema.Value == nil {
				contract.Failf("Path %s has no schema definition for status code 200", currentPath)
			}

			resourceType := jsonReq.Schema.Value
			if resourceType.Type != "array" {
				typeToken := fmt.Sprintf("%s:%s:%s", packageName, module, resourceType.Title)
				if existing, ok := o.resourceCRUDMap[typeToken]; ok {
					existing.R = &currentPath
				} else {
					o.resourceCRUDMap[typeToken] = &CRUDOperationsMap{
						R: &currentPath,
					}
				}
			}
			// TODO: We'll should also add this as get*/list* provider functions.
		}

		if pathItem.Patch != nil {
			parentPath := getParentPath(currentPath)
			glog.V(3).Infof("PATCH: Parent path for %s is %s\n", currentPath, parentPath)

			jsonReq := pathItem.Patch.RequestBody.Value.Content.Get("application/json")
			if jsonReq.Schema.Value == nil {
				contract.Failf("Path %s has no schema definition for Patch method", currentPath)
			}

			resourceType := jsonReq.Schema.Value
			typeToken := fmt.Sprintf("%s:%s:%s", packageName, module, resourceType.Title)
			if existing, ok := o.resourceCRUDMap[typeToken]; ok {
				existing.U = &currentPath
			} else {
				o.resourceCRUDMap[typeToken] = &CRUDOperationsMap{
					U: &currentPath,
				}
			}
		}

		if pathItem.Put != nil {
			parentPath := getParentPath(currentPath)
			glog.V(3).Infof("PUT: Parent path for %s is %s\n", currentPath, parentPath)

			jsonReq := pathItem.Put.RequestBody.Value.Content.Get("application/json")
			if jsonReq.Schema.Value == nil {
				contract.Failf("Path %s has no schema definition for Put method", currentPath)
			}

			resourceType := jsonReq.Schema.Value
			typeToken := fmt.Sprintf("%s:%s:%s", packageName, module, resourceType.Title)
			if existing, ok := o.resourceCRUDMap[typeToken]; ok {
				existing.P = &currentPath
			} else {
				o.resourceCRUDMap[typeToken] = &CRUDOperationsMap{
					P: &currentPath,
				}
			}
		}

		if pathItem.Delete != nil {
			parentPath := getParentPath(currentPath)
			glog.V(3).Infof("DELETE: Parent path for %s is %s\n", currentPath, parentPath)

			jsonReq := o.doc.Paths[parentPath].Post.RequestBody.Value.Content.Get("application/json")
			if jsonReq.Schema.Value == nil {
				contract.Failf("Path %s has no schema definition for Delete method", parentPath)
			}

			resourceType := jsonReq.Schema.Value
			typeToken := fmt.Sprintf("%s:%s:%s", packageName, module, resourceType.Title)

			if existing, ok := o.resourceCRUDMap[typeToken]; ok {
				existing.D = &currentPath
			} else {
				o.resourceCRUDMap[typeToken] = &CRUDOperationsMap{
					D: &currentPath,
				}
			}
		}

		if pathItem.Post == nil {
			continue
		}

		if err := o.gatherResourceFromAPIPath(currentPath, *pathItem, module); err != nil {
			return errors.Wrapf(err, "generating resource for api path %s", currentPath)
		}

		csharpNamespaces[module] = ToPascalCase(module)
	}

	return nil
}

func (o *openAPIContext) gatherResourceFromAPIPath(apiPath string, pathItem openapi3.PathItem, module string) error {
	jsonReq := pathItem.Post.RequestBody.Value.Content.Get("application/json")
	if jsonReq.Schema.Value == nil {
		return errors.Errorf("path %s has no api schema definition for post method", apiPath)
	}

	resourceType := jsonReq.Schema.Value

	var resourceTypeToken *string
	var err error

	if resourceType.Discriminator != nil {
		for _, mappingRef := range resourceType.Discriminator.Mapping {
			schemaName := strings.TrimPrefix(mappingRef, "#/components/schemas/")
			typeSchema, ok := o.doc.Components.Schemas[schemaName]
			if !ok {
				return errors.Errorf("%s not found in api schemas for discriminated type in path %s", schemaName, apiPath)
			}

			resourceTypeToken, err = o.gatherResourceFromAPISchema(*typeSchema.Value, apiPath, module)
		}
	} else {
		resourceTypeToken, err = o.gatherResourceFromAPISchema(*resourceType, apiPath, module)
	}

	if err != nil {
		return errors.Wrapf(err, "gathering resource from api path %s", apiPath)
	}

	resource := o.pkg.Resources[*resourceTypeToken]
	requiredInputs := codegen.NewStringSet(resource.RequiredInputs...)

	// If this endpoint path has path parameters,
	// then those should be required inputs too.
	for _, param := range pathItem.Parameters {
		if param.Value.In != "path" {
			continue
		}

		paramName := param.Value.Name
		resource.InputProperties[paramName] = pschema.PropertySpec{
			Description: param.Value.Description,
			TypeSpec:    pschema.TypeSpec{Type: "string"},
		}
		requiredInputs.Add(paramName)
	}

	o.pkg.Resources[*resourceTypeToken] = resource

	return nil
}

func (o *openAPIContext) gatherResourceFromAPISchema(apiSchema openapi3.Schema, apiPath, module string) (*string, error) {
	pkgCtx := &resourceContext{
		mod:               module,
		pkg:               o.pkg,
		resourceName:      apiSchema.Title,
		openapiComponents: o.doc.Components,
		visitedTypes:      codegen.NewStringSet(),
	}

	inputProperties := make(map[string]pschema.PropertySpec)
	properties := make(map[string]pschema.PropertySpec)
	requiredInputs := codegen.NewStringSet()
	for propName, prop := range apiSchema.Properties {
		propSpec := pkgCtx.genPropertySpec(ToPascalCase(propName), *prop)
		if !prop.Value.ReadOnly {
			inputProperties[propName] = propSpec
		}
		properties[propName] = propSpec
	}

	// Create a set of required inputs for this resource.
	// Filter out required props that are marked as read-only.
	for _, requiredProp := range apiSchema.Required {
		propSchema := apiSchema.Properties[requiredProp]
		if propSchema.Value.ReadOnly {
			continue
		}

		requiredInputs.Add(requiredProp)
	}

	if len(apiSchema.AllOf) > 0 {
		parentName := ToPascalCase(apiSchema.Title)
		var types []pschema.TypeSpec
		for _, schemaRef := range apiSchema.AllOf {
			typ, err := pkgCtx.propertyTypeSpec(parentName, *schemaRef)
			if err != nil {
				return nil, errors.Wrapf(err, "generating property type spec from allOf schema for %s", apiSchema.Title)
			}
			types = append(types, *typ)
		}

		// Now that all of the types have been added to schema's Types,
		// gather all of their properties and smash them together into
		// a new type and get rid of those top-level ones.
		typeToken := fmt.Sprintf("%s:%s:%s", packageName, module, parentName)
		properties := make(map[string]pschema.PropertySpec)
		for _, t := range types {
			refType := pkgCtx.pkg.Types[strings.TrimPrefix(t.Ref, "#/types/")]

			for name, propSpec := range refType.Properties {
				properties[name] = propSpec
			}

			for _, r := range refType.Required {
				if requiredInputs.Has(r) {
					continue
				}
				requiredInputs.Add(r)
			}
		}

		if existing, ok := o.resourceCRUDMap[typeToken]; ok {
			existing.C = &apiPath
		} else {
			o.resourceCRUDMap[typeToken] = &CRUDOperationsMap{
				C: &apiPath,
			}
		}

		o.pkg.Resources[typeToken] = pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Description: apiSchema.Description,
				Type:        "object",
				Properties:  properties,
				Required:    apiSchema.Required,
			},
			InputProperties: inputProperties,
			RequiredInputs:  requiredInputs.SortedValues(),
		}

		return &typeToken, nil
	}

	typeToken := fmt.Sprintf("%s:%s:%s", packageName, module, apiSchema.Title)
	if existing, ok := o.resourceCRUDMap[typeToken]; ok {
		existing.C = &apiPath
	} else {
		o.resourceCRUDMap[typeToken] = &CRUDOperationsMap{
			C: &apiPath,
		}
	}

	o.pkg.Resources[typeToken] = pschema.ResourceSpec{
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: apiSchema.Description,
			Type:        "object",
			Properties:  properties,
			Required:    apiSchema.Required,
		},
		InputProperties: inputProperties,
		RequiredInputs:  requiredInputs.SortedValues(),
	}

	return &typeToken, nil
}

func (ctx *resourceContext) genPropertySpec(propName string, p openapi3.SchemaRef) pschema.PropertySpec {
	propertySpec := pschema.PropertySpec{
		Description: p.Value.Description,
	}
	if p.Value.Default != nil {
		propertySpec.Default = p.Value.Default
	}
	languageName := strings.ToUpper(propName[:1]) + propName[1:]
	if languageName == ctx.resourceName {
		// .NET does not allow properties to be the same as the enclosing class - so special case these
		propertySpec.Language = map[string]pschema.RawMessage{
			"csharp": rawMessage(dotnetgen.CSharpPropertyInfo{
				Name: languageName + "Value",
			}),
		}
	}
	// JSONSchema type includes `$ref` and `$schema` properties, and $ is an invalid character in
	// the generated names. Replace them with `Ref` and `Schema`.
	if strings.HasPrefix(propName, "$") {
		propertySpec.Language = map[string]pschema.RawMessage{
			"csharp": rawMessage(map[string]interface{}{
				"name": strings.ToUpper(propName[1:2]) + propName[2:],
			}),
		}
	}

	typeSpec, err := ctx.propertyTypeSpec(propName, p)
	if err != nil {
		contract.Failf("Failed to generate type spec (resource: %s; prop %s): %v", ctx.resourceName, propName, err)
	}

	propertySpec.TypeSpec = *typeSpec

	return propertySpec
}

// propertyTypeSpec converts a JSON type definition to a Pulumi type definition.
func (ctx *resourceContext) propertyTypeSpec(parentName string, propSchema openapi3.SchemaRef) (*pschema.TypeSpec, error) {
	// References to other type definitions.
	if propSchema.Ref != "" {
		schemaName := strings.TrimPrefix(propSchema.Ref, "#/components/schemas/")
		typName := ToPascalCase(schemaName)
		if !strings.HasPrefix(schemaName, ctx.resourceName) {
			typName = fmt.Sprintf("%s%s", ctx.resourceName, typName)
		}
		tok := fmt.Sprintf("%s:%s:%s", packageName, ctx.mod, typName)

		typeSchema, ok := ctx.openapiComponents.Schemas[schemaName]
		if !ok {
			return nil, errors.Errorf("definition %s not found in resource %s", schemaName, ctx.resourceName)
		}

		if !ctx.visitedTypes.Has(tok) {
			ctx.visitedTypes.Add(tok)
			specs, requiredSpecs, err := ctx.genProperties(typName, *typeSchema.Value)
			if err != nil {
				return nil, err
			}

			ctx.pkg.Types[tok] = pschema.ComplexTypeSpec{
				ObjectTypeSpec: pschema.ObjectTypeSpec{
					Description: typeSchema.Value.Description,
					Type:        "object",
					Properties:  specs,
					Required:    requiredSpecs.SortedValues(),
				},
			}
		}

		referencedTypeName := fmt.Sprintf("#/types/%s", tok)
		return &pschema.TypeSpec{Ref: referencedTypeName}, nil
	}

	// Inline properties.
	if len(propSchema.Value.Properties) > 0 {
		typName := parentName + "Properties"
		tok := fmt.Sprintf("%s:%s:%s", packageName, ctx.mod, typName)
		specs, requiredSpecs, err := ctx.genProperties(typName, *propSchema.Value)
		if err != nil {
			return nil, err
		}

		ctx.pkg.Types[tok] = pschema.ComplexTypeSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Description: propSchema.Value.Description,
				Type:        "object",
				Properties:  specs,
				Required:    requiredSpecs.SortedValues(),
			},
		}
		referencedTypeName := fmt.Sprintf("#/types/%s", tok)
		return &pschema.TypeSpec{Ref: referencedTypeName}, nil
	}

	// Union types.
	if len(propSchema.Value.OneOf) > 0 {
		var types []pschema.TypeSpec
		for _, schemaRef := range propSchema.Value.OneOf {
			typ, err := ctx.propertyTypeSpec(parentName, *schemaRef)
			if err != nil {
				return nil, err
			}
			types = append(types, *typ)
		}

		var discriminator *pschema.DiscriminatorSpec
		if propSchema.Value.Discriminator != nil {
			discriminator = &pschema.DiscriminatorSpec{
				PropertyName: propSchema.Value.Discriminator.PropertyName,
				Mapping:      propSchema.Value.Discriminator.Mapping,
			}
		}

		return &pschema.TypeSpec{
			OneOf:         types,
			Discriminator: discriminator,
		}, nil
	}

	// TODO: allOf
	if len(propSchema.Value.AllOf) > 0 {
		var types []pschema.TypeSpec
		for _, schemaRef := range propSchema.Value.AllOf {
			typ, err := ctx.propertyTypeSpec(parentName, *schemaRef)
			if err != nil {
				return nil, err
			}
			types = append(types, *typ)
		}

		// Now that all of the types have been added to schema's Types,
		// gather all of their properties and smash them together into
		// a new type and get rid of those top-level ones.
		tok := fmt.Sprintf("%s:%s:%s", packageName, ctx.mod, ToPascalCase(parentName))
		properties := make(map[string]pschema.PropertySpec)
		requiredSpecs := codegen.NewStringSet()
		for _, t := range types {
			refType := ctx.pkg.Types[t.Ref]

			for name, propSpec := range refType.Properties {
				properties[name] = propSpec
			}

			for _, r := range refType.Required {
				if requiredSpecs.Has(r) {
					continue
				}
				requiredSpecs.Add(r)
			}
		}

		ctx.pkg.Types[tok] = pschema.ComplexTypeSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Description: propSchema.Value.Description,
				Type:        "object",
				Properties:  properties,
				Required:    requiredSpecs.SortedValues(),
			},
		}
		referencedTypeName := fmt.Sprintf("#/types/%s", tok)
		return &pschema.TypeSpec{
			Ref: referencedTypeName,
		}, nil
	}

	if len(propSchema.Value.Enum) > 0 {
		enum, err := ctx.genEnumType(parentName, *propSchema.Value)
		if err != nil {
			return nil, err
		}
		if enum != nil {
			return enum, nil
		}
	}

	// All other types.
	switch propSchema.Value.Type {
	case openapi3.TypeInteger:
		return &pschema.TypeSpec{Type: "integer"}, nil
	case openapi3.TypeString:
		return &pschema.TypeSpec{Type: "string"}, nil
	case openapi3.TypeBoolean:
		return &pschema.TypeSpec{Type: "boolean"}, nil
	case openapi3.TypeNumber:
		return &pschema.TypeSpec{Type: "number"}, nil
	case openapi3.TypeObject:
		return &pschema.TypeSpec{Ref: "pulumi.json#/Any"}, nil
	case openapi3.TypeArray:
		elementType, err := ctx.propertyTypeSpec(parentName+"Item", *propSchema.Value.Items)
		if err != nil {
			return nil, err
		}
		return &pschema.TypeSpec{
			Type:  "array",
			Items: elementType,
		}, nil
	}

	return nil, errors.Errorf("failed to generate property types for %+v", propSchema)
}

func (ctx *resourceContext) genProperties(parentName string, typeSchema openapi3.Schema) (map[string]pschema.PropertySpec, codegen.StringSet, error) {
	specs := map[string]pschema.PropertySpec{}
	requiredSpecs := codegen.NewStringSet()
	for _, name := range codegen.SortedKeys(typeSchema.Properties) {
		value := typeSchema.Properties[name]
		sdkName := ToSdkName(name)

		typeSpec, err := ctx.propertyTypeSpec(parentName+ToPascalCase(name), *value)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "property %s", name)
		}
		propertySpec := pschema.PropertySpec{
			Description: value.Value.Description,
			TypeSpec:    *typeSpec,
		}
		if value.Value.Default != nil {
			propertySpec.Default = value.Value.Default
		}

		specs[sdkName] = propertySpec
	}
	for _, name := range typeSchema.Required {
		sdkName := ToSdkName(name)
		if _, has := specs[sdkName]; has {
			requiredSpecs.Add(sdkName)
		}
	}
	return specs, requiredSpecs, nil
}

// genEnumType generates the enum type for a given schema.
func (ctx *resourceContext) genEnumType(enumName string, propSchema openapi3.Schema) (*pschema.TypeSpec, error) {
	if len(propSchema.Type) == 0 {
		return nil, nil
	}
	if propSchema.Type != openapi3.TypeString {
		return nil, nil
	}

	typName := ToPascalCase(enumName)
	if !strings.HasPrefix(enumName, ctx.resourceName) {
		typName = ctx.resourceName + enumName
	}

	tok := fmt.Sprintf("%s:%s:%s", packageName, ctx.mod, typName)

	enumSpec := &pschema.ComplexTypeSpec{
		Enum: []pschema.EnumValueSpec{},
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: propSchema.Description,
			Type:        "string",
		},
	}

	values := codegen.NewStringSet()
	for _, val := range propSchema.Enum {
		str := ToPascalCase(val.(string))
		if values.Has(str) {
			continue
		}
		values.Add(str)
		enumVal := pschema.EnumValueSpec{
			Value: val,
			Name:  str,
		}
		enumSpec.Enum = append(enumSpec.Enum, enumVal)
	}

	// Make sure that the type name we composed doesn't clash with another type
	// already defined in the schema earlier. The same enum does show up in multiple
	// places of specs, so we want to error only if they a) have the same name
	// b) the list of values does not match.
	if other, ok := ctx.pkg.Types[tok]; ok {
		same := len(enumSpec.Enum) == len(other.Enum)
		for _, val := range other.Enum {
			same = same && values.Has(val.Name)
		}
		if !same {
			return nil, errors.Errorf("duplicate enum %q: %+v vs. %+v", tok, enumSpec.Enum, other.Enum)
		}
	}
	ctx.pkg.Types[tok] = *enumSpec

	referencedTypeName := fmt.Sprintf("#/types/%s", tok)
	return &pschema.TypeSpec{
		Ref: referencedTypeName,
	}, nil
}
