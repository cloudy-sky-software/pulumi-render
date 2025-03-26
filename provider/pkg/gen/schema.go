// Copyright 2022, Cloudy Sky Software LLC.

package gen

import (
	"bytes"
	"encoding/json"

	"github.com/getkin/kin-openapi/openapi3"

	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	nodejsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"

	openapigen "github.com/cloudy-sky-software/pulschema/pkg"

	"github.com/cloudy-sky-software/pulumi-render/provider/pkg/gen/examples"
)

const packageName = "render"

// PulumiSchema will generate a Pulumi schema for the given k8s schema.
func PulumiSchema(openAPIDoc openapi3.T) (pschema.PackageSpec, openapigen.ProviderMetadata, openapi3.T) {
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
					TypeSpec:    pschema.TypeSpec{Ref: "#/types/render:services:ClearCache"},
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

		PluginDownloadURL: "github://api.github.com/cloudy-sky-software/pulumi-render",
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

	openAPICtx := &openapigen.OpenAPIContext{
		Doc: openAPIDoc,
		Pkg: &pkg,
		ExcludedPaths: []string{
			"/services/{serviceId}/resume",
			"/services/{serviceId}/custom-domains/{id}/verify",
			"/services/{serviceId}/events",
			"/logs",
			"/logs/subscribe",
			"/logs/values",
			// The GET endpoint response for /disks/{diskId}/snapshots
			// has an invalid "201 Created" response code despite
			// the endpoint being a list endpoint to list all snapshots.
			"/disks/{diskId}/snapshots",
			"/events/{eventId}",
			"/maintenance",
			"/maintenance/{maintenanceRunParam}",
			"/maintenance/{maintenanceRunParam}/trigger",
			"/metrics/cpu",
			"/metrics/cpu-limit",
			"/metrics/cpu-target",
			"/metrics/memory",
			"/metrics/memory-limit",
			"/metrics/memory-target",
			"/metrics/http-requests",
			"/metrics/http-latency",
			"/metrics/bandwidth",
			"/metrics/disk-usage",
			"/metrics/disk-capacity",
			"/metrics/instance-count",
			"/metrics/filters/application",
			"/metrics/filters/http",
			"/metrics/filters/path",
			"/metrics-stream/{ownerId}",
		},
	}

	providerMetadata, updatedOpenAPIDoc, err := openAPICtx.GatherResourcesFromAPI(csharpNamespaces)
	if err != nil {
		contract.Failf("generating resources from OpenAPI spec: %v", err)
	}

	// Add examples to resources
	for k, v := range examples.ResourceExample {
		if r, ok := pkg.Resources[k]; ok {
			r.Description += "\n\n" + v
			pkg.Resources[k] = r
		}
	}

	// Override the namespace for registrycredentials.
	csharpNamespaces["registrycredentials"] = "RegistryCredentials"

	pkg.Language["csharp"] = rawMessage(map[string]interface{}{
		"rootNamespace": "Pulumi",
		"packageReferences": map[string]string{
			"Pulumi": "3.*",
		},
		"namespaces": csharpNamespaces,
		// TODO: What does this enable?
		// "dictionaryConstructors": true,
	})

	pkg.Language["go"] = rawMessage(gogen.GoPackageInfo{
		ImportBasePath: "github.com/cloudy-sky-software/pulumi-render/sdk/go/render",
		ModuleToPackage: map[string]string{
			"cron-jobs":             "cronJobs",
			"env-groups":            "envGroups",
			"notification-settings": "notificationSettings",
		},
	})
	pkg.Language["nodejs"] = rawMessage(nodejsgen.NodePackageInfo{
		PackageName: "@cloudyskysoftware/pulumi-render",
		ModuleToPackage: map[string]string{
			"cron-jobs":             "cronjobs",
			"env-groups":            "envgroups",
			"notification-settings": "notificationsettings",
		},
	})
	pkg.Language["python"] = rawMessage(map[string]interface{}{
		"packageName": "pulumi_render",
		"requires": map[string]string{
			"pulumi": ">=3.0.0,<4.0.0",
		},
		"pyproject": map[string]bool{
			"enabled": true,
		},
	})

	// The EnvVars resource's GET operation looks like a list method,
	// but it is also the endpoint to read the current set of env vars
	// for a Render service. So we'll manually set the read endpoint
	// for that type token here.
	// TODO: Can pulschema handle this instead?
	envVarsReadEndpoint := "/services/{serviceId}/env-vars"
	providerMetadata.ResourceCRUDMap["render:services:EnvVarsForService"].R = &envVarsReadEndpoint
	secretFilesReadEndpoint := "/services/{serviceId}/secret-files"
	providerMetadata.ResourceCRUDMap["render:services:SecretFilesForService"].R = &secretFilesReadEndpoint

	metadata := openapigen.ProviderMetadata{
		ResourceCRUDMap:  providerMetadata.ResourceCRUDMap,
		AutoNameMap:      providerMetadata.AutoNameMap,
		SDKToAPINameMap:  providerMetadata.SDKToAPINameMap,
		APIToSDKNameMap:  providerMetadata.APIToSDKNameMap,
		PathParamNameMap: providerMetadata.PathParamNameMap,
	}
	return pkg, metadata, updatedOpenAPIDoc
}

func rawMessage(v interface{}) pschema.RawMessage {
	var out bytes.Buffer
	encoder := json.NewEncoder(&out)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	contract.Assert(err == nil)
	return out.Bytes()
}
