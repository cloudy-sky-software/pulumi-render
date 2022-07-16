// Copyright 2016-2021, Pulumi Corporation.
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

//nolint: goconst
package gen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cloudy-sky-software/pulumi-render/provider/pkg/gen/examples"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

const packageName = "render"

// typeOverlays augment the types defined by the kubernetes schema.
var typeOverlays = map[string]pschema.ComplexTypeSpec{}

// resourceOverlays augment the resources defined by the kubernetes schema.
var resourceOverlays = map[string]pschema.ResourceSpec{}

type resourceContext struct {
	mod               string
	pkg               *pschema.PackageSpec
	resourceName      string
	openapiComponents openapi3.Components
	visitedTypes      codegen.StringSet
}

// PulumiSchema will generate a Pulumi schema for the given k8s schema.
func PulumiSchema(openapiDoc openapi3.T) pschema.PackageSpec {
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
				},
			},
		},

		Types:     map[string]pschema.ComplexTypeSpec{},
		Resources: map[string]pschema.ResourceSpec{},
		Functions: map[string]pschema.FunctionSpec{},
		Language:  map[string]pschema.RawMessage{},
	}

	csharpNamespaces := map[string]string{
		"render": "Render",
		// TODO: Is this needed?
		"": "Provider",
	}

	getRootPath := func(path string) string {
		parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
		return parts[0]
	}

	for path, pathSchema := range openapiDoc.Paths {
		if pathSchema.Post == nil {
			continue
		}
		// TODO: TEMPORARY!
		if path != "/services" {
			continue
		}
		//

		module := getRootPath(path)
		jsonReq := pathSchema.Post.RequestBody.Value.Content.Get("application/json")

		if jsonReq.Schema.Value == nil {
			contract.Failf("Path %s has no schema definition", path)
		}

		resourceType := jsonReq.Schema.Value
		pkgCtx := &resourceContext{
			mod:               module,
			pkg:               &pkg,
			resourceName:      resourceType.Title,
			openapiComponents: openapiDoc.Components,
			visitedTypes:      codegen.NewStringSet(),
		}

		inputProperties := make(map[string]pschema.PropertySpec)
		properties := make(map[string]pschema.PropertySpec)
		for propName, prop := range resourceType.Properties {
			propSpec := pkgCtx.genPropertySpec(propName, *prop)
			if !prop.Value.ReadOnly {
				inputProperties[propName] = propSpec
			}
			properties[propName] = propSpec
		}

		requiredInputs := codegen.NewStringSet()
		// Create a set of required inputs for this resource.
		// Filter out required props that are marked as read-only.
		for _, requiredProp := range resourceType.Required {
			propSchema := resourceType.Properties[requiredProp]
			if propSchema.Value.ReadOnly {
				continue
			}

			requiredInputs.Add(requiredProp)
		}

		typeToken := fmt.Sprintf("%s:%s:%s", packageName, module, resourceType.Title)
		pkg.Resources[typeToken] = pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Description: pathSchema.Description,
				Type:        "object",
				Properties:  properties,
				Required:    resourceType.Required,
			},
			InputProperties: inputProperties,
			RequiredInputs:  requiredInputs.SortedValues(),
		}

		csharpNamespaces[module] = ToPascalCase(module)
	}

	// Add examples to resources
	for k, v := range examples.ApiVersionToExample {
		if r, ok := pkg.Resources[k]; ok {
			r.Description += "\n\n" + v
			pkg.Resources[k] = r
		}
	}

	pkg.Language["csharp"] = rawMessage(map[string]interface{}{
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
		"requires": map[string]string{
			"pulumi": ">=3.0.0,<4.0.0",
		},
		// TODO: What does this enable?
		// "usesIOClasses":       true,
	})

	return pkg
}

func (ctx *resourceContext) genPropertySpec(propName string, p openapi3.SchemaRef) pschema.PropertySpec {
	defaultValue := func() *string {
		if p.Value.Default == nil {
			return nil
		}

		// TODO: This typecast will likely fail if the default value
		// is not a string.
		dv := p.Value.Default.(string)
		if len(dv) != 0 {
			return &dv
		}
		return nil
	}

	propertySpec := pschema.PropertySpec{
		Description: p.Value.Description,
	}
	if dv := defaultValue(); dv != nil {
		propertySpec.Default = *dv
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
		typName := schemaName
		if !strings.HasPrefix(schemaName, ctx.resourceName) {
			typName = fmt.Sprintf("%s%s", ctx.resourceName, ToPascalCase(schemaName))
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
	if len(propSchema.Value.AnyOf) > 0 {
		var types []pschema.TypeSpec
		for _, schemaRef := range propSchema.Value.AnyOf {
			typ, err := ctx.propertyTypeSpec(parentName, *schemaRef)
			if err != nil {
				return nil, err
			}
			types = append(types, *typ)
		}
		return &pschema.TypeSpec{
			OneOf: types,
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
		return &schema.TypeSpec{Type: "integer"}, nil
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

		typeSpec, err := ctx.propertyTypeSpec(parentName+name, *value)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "property %s", name)
		}
		propertySpec := pschema.PropertySpec{
			Description: value.Value.Description,
			TypeSpec:    *typeSpec,
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
	typName := enumName
	if !strings.HasPrefix(enumName, ctx.resourceName) {
		typName = ctx.resourceName + ToPascalCase(enumName)
	}
	switch ctx.mod + ":" + typName {
	case "networkfirewall:RuleGroupType":
		typName = "RuleGroupTypeEnum" // Go SDK name conflict vs. RuleGroup resource
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
			switch tok {
			case "aws-native:mediaconnect:FlowSourceProtocol":
				// FlowSourceProtocol is defined in two resources and one is a subset of the other.
				// They seem to be the same type. Pick the longer one here to avoid losing values.
				if len(enumSpec.Enum) < len(other.Enum) {
					enumSpec = &other
				}
			default:
				return nil, errors.Errorf("duplicate enum %q: %+v vs. %+v", tok, enumSpec.Enum, other.Enum)
			}
		}
	}
	ctx.pkg.Types[tok] = *enumSpec

	referencedTypeName := fmt.Sprintf("#/types/%s", tok)
	return &pschema.TypeSpec{
		Ref: referencedTypeName,
	}, nil
}

func rawMessage(v interface{}) pschema.RawMessage {
	var out bytes.Buffer
	encoder := json.NewEncoder(&out)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	contract.Assert(err == nil)
	return out.Bytes()
}
