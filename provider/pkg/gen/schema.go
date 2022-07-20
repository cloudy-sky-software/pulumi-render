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
	"github.com/golang/glog"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
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

	getParentPath := func(path string) string {
		parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
		lastPathPart := parts[len(parts)-1]
		if !strings.HasPrefix(lastPathPart, "{") && !strings.HasSuffix(lastPathPart, "}") {
			return path
		}

		// Skip the last path part which contains a path parameter.
		return "/" + strings.Join(parts[0:len(parts)-1], "/")
	}

	resourceCRUDMap := make(map[string]*CRUDOperationsMap)

	for path, pathSchema := range openapiDoc.Paths {
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

		if pathSchema.Get != nil {
			parentPath := getParentPath(currentPath)
			glog.V(3).Infof("GET: Parent path for %s is %s\n", currentPath, parentPath)

			jsonReq := pathSchema.Get.Responses.Get(200).Value.Content.Get("application/json")
			if jsonReq.Schema.Value == nil {
				contract.Failf("Path %s has no schema definition for status code 200", currentPath)
			}

			resourceType := jsonReq.Schema.Value
			if resourceType.Type != "array" {
				typeToken := fmt.Sprintf("%s:%s:%s", packageName, module, resourceType.Title)
				if existing, ok := resourceCRUDMap[typeToken]; ok {
					existing.R = &currentPath
				} else {
					resourceCRUDMap[typeToken] = &CRUDOperationsMap{
						R: &currentPath,
					}
				}
			}
			// TODO: We'll should also add this as get*/list* provider functions.
		}

		if pathSchema.Patch != nil {
			parentPath := getParentPath(currentPath)
			glog.V(3).Infof("PATCH: Parent path for %s is %s\n", currentPath, parentPath)

			jsonReq := pathSchema.Patch.RequestBody.Value.Content.Get("application/json")
			if jsonReq.Schema.Value == nil {
				contract.Failf("Path %s has no schema definition for Patch method", currentPath)
			}

			resourceType := jsonReq.Schema.Value
			typeToken := fmt.Sprintf("%s:%s:%s", packageName, module, resourceType.Title)
			if existing, ok := resourceCRUDMap[typeToken]; ok {
				existing.U = &currentPath
			} else {
				resourceCRUDMap[typeToken] = &CRUDOperationsMap{
					U: &currentPath,
				}
			}
		}

		if pathSchema.Put != nil {
			parentPath := getParentPath(currentPath)
			glog.V(3).Infof("PUT: Parent path for %s is %s\n", currentPath, parentPath)

			jsonReq := pathSchema.Put.RequestBody.Value.Content.Get("application/json")
			if jsonReq.Schema.Value == nil {
				contract.Failf("Path %s has no schema definition for Put method", currentPath)
			}

			resourceType := jsonReq.Schema.Value
			typeToken := fmt.Sprintf("%s:%s:%s", packageName, module, resourceType.Title)
			if existing, ok := resourceCRUDMap[typeToken]; ok {
				existing.P = &currentPath
			} else {
				resourceCRUDMap[typeToken] = &CRUDOperationsMap{
					P: &currentPath,
				}
			}
		}

		if pathSchema.Delete != nil {
			parentPath := getParentPath(currentPath)
			glog.V(3).Infof("DELETE: Parent path for %s is %s\n", currentPath, parentPath)

			jsonReq := openapiDoc.Paths[parentPath].Post.RequestBody.Value.Content.Get("application/json")
			if jsonReq.Schema.Value == nil {
				contract.Failf("Path %s has no schema definition for Delete method", parentPath)
			}

			resourceType := jsonReq.Schema.Value
			typeToken := fmt.Sprintf("%s:%s:%s", packageName, module, resourceType.Title)

			if existing, ok := resourceCRUDMap[typeToken]; ok {
				existing.D = &currentPath
			} else {
				resourceCRUDMap[typeToken] = &CRUDOperationsMap{
					D: &currentPath,
				}
			}
		}

		if pathSchema.Post == nil {
			continue
		}

		jsonReq := pathSchema.Post.RequestBody.Value.Content.Get("application/json")
		if jsonReq.Schema.Value == nil {
			glog.Warningf("Path %s has no schema definition for Post method", currentPath)
			continue
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
			propSpec := pkgCtx.genPropertySpec(ToPascalCase(propName), *prop)
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

		// If this endpoint path has path parameters,
		// then those should be required inputs too.
		for _, param := range pathSchema.Parameters {
			if param.Value.In != "path" {
				continue
			}

			paramName := param.Value.Name
			inputProperties[paramName] = pschema.PropertySpec{
				Description: param.Value.Description,
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			}
			requiredInputs.Add(paramName)
		}

		typeToken := fmt.Sprintf("%s:%s:%s", packageName, module, resourceType.Title)
		if existing, ok := resourceCRUDMap[typeToken]; ok {
			existing.C = &currentPath
		} else {
			resourceCRUDMap[typeToken] = &CRUDOperationsMap{
				C: &currentPath,
			}
		}

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
	})

	metadata := ProviderMetadata{
		ResourceCRUDMap: resourceCRUDMap,
	}
	return pkg, metadata
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

func rawMessage(v interface{}) pschema.RawMessage {
	var out bytes.Buffer
	encoder := json.NewEncoder(&out)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	contract.Assert(err == nil)
	return out.Bytes()
}
