// Copyright 2022, Cloudy Sky Software LLC.

package main

import (
	_ "embed"

	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	providerSchemaGen "github.com/cloudy-sky-software/pulumi-render/provider/pkg/gen"
	providerVersion "github.com/cloudy-sky-software/pulumi-render/provider/pkg/version"

	"github.com/cloudy-sky-software/pulumi-provider-framework/openapi"

	"github.com/pkg/errors"

	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	nodejsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pythongen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

//go:embed openapi.yml
var openapiDocBytes []byte

// TemplateDir is the path to the base directory for code generator templates.
var TemplateDir string

// BaseDir is the path to the base pulumi-render directory.
var BaseDir string

// Language is the SDK language.
type Language string

const (
	DotNet Language = "dotnet"
	Go     Language = "go"
	NodeJS Language = "nodejs"
	Python Language = "python"
	Schema Language = "schema"
)

func main() {
	flag.Usage = func() {
		const usageFormat = "Usage: %s <language> <schema-file> <root-pulumi-render-dir>"
		_, err := fmt.Fprintf(flag.CommandLine.Output(), usageFormat, os.Args[0])
		contract.IgnoreError(err)
		flag.PrintDefaults()
	}

	var version string
	flag.StringVar(&version, "version", providerVersion.Version, "the provider version to record in the generated code")

	flag.Parse()
	args := flag.Args()
	if len(args) < 3 {
		flag.Usage()
		return
	}

	language, inputFile := Language(args[0]), args[1]

	BaseDir = args[2]
	TemplateDir = filepath.Join(BaseDir, "provider", "pkg", "gen")
	outdir := filepath.Join(BaseDir, "sdk", string(language))

	var schemaPkg *schema.Package
	if language != Schema {
		schemaPkg = readSchema(inputFile, version)
	}

	switch language {
	case NodeJS:
		writeNodeJSClient(schemaPkg, outdir)
	case Python:
		writePythonClient(schemaPkg, outdir)
	case DotNet:
		writeDotnetClient(schemaPkg, outdir)
	case Go:
		writeGoClient(schemaPkg, outdir)
	case Schema:
		openAPIDoc := openapi.GetOpenAPISpec(openapiDocBytes)

		err := providerSchemaGen.FixOpenAPIDoc(openAPIDoc)
		if err != nil {
			panic(err)
		}

		schemaSpec, metadata, updatedOpenAPIDoc := providerSchemaGen.PulumiSchema(*openAPIDoc)
		providerDir := filepath.Join(".", "provider", "cmd", "pulumi-resource-render")
		mustWritePulumiSchema(schemaSpec, providerDir)

		// Write the metadata.json file as well.
		metadataBytes, _ := json.Marshal(metadata)
		mustWriteFile(providerDir, "metadata.json", metadataBytes)

		updatedOpenAPIDocBytes, _ := yaml.Marshal(updatedOpenAPIDoc)
		// Also copy the raw OpenAPI spec file to the provider dir.
		mustWriteFile(providerDir, "openapi_generated.yml", updatedOpenAPIDocBytes)
	default:
		panic(fmt.Sprintf("Unrecognized language '%s'", language))
	}
}

func mustWritePulumiSchema(pkgSpec schema.PackageSpec, outdir string) {
	pkgSpec.Version = ""
	schemaJSON, err := json.MarshalIndent(pkgSpec, "", "    ")
	if err != nil {
		panic(errors.Wrap(err, "marshaling Pulumi schema"))
	}
	mustWriteFile(outdir, "schema.json", schemaJSON)
}

func readSchema(schemaPath string, version string) *schema.Package {
	// Read in, decode, and import the schema.
	schemaBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		panic(err)
	}

	var pkgSpec schema.PackageSpec
	if err = json.Unmarshal(schemaBytes, &pkgSpec); err != nil {
		panic(err)
	}
	pkgSpec.Version = version

	pkg, err := schema.ImportSpec(pkgSpec, nil)
	if err != nil {
		panic(err)
	}
	return pkg
}

func writeNodeJSClient(pkg *schema.Package, outdir string) {
	_, err := nodejsgen.LanguageResources(pkg)
	if err != nil {
		panic(err)
	}

	overlays := map[string][]byte{}
	files, err := nodejsgen.GeneratePackage("pulumigen", pkg, overlays, nil)
	if err != nil {
		panic(err)
	}

	mustWriteFiles(outdir, files)
}

func writePythonClient(pkg *schema.Package, outdir string) {
	_, err := pythongen.LanguageResources("pulumigen", pkg)
	if err != nil {
		panic(err)
	}

	overlays := map[string][]byte{}

	files, err := pythongen.GeneratePackage("pulumigen", pkg, overlays)
	if err != nil {
		panic(err)
	}

	mustWriteFiles(outdir, files)
}

func writeDotnetClient(pkg *schema.Package, outdir string) {
	_, err := dotnetgen.LanguageResources("pulumigen", pkg)
	if err != nil {
		panic(err)
	}

	overlays := map[string][]byte{}

	files, err := dotnetgen.GeneratePackage("pulumigen", pkg, overlays, nil)
	if err != nil {
		panic(err)
	}

	for filename, contents := range files {
		path := filepath.Join(outdir, filename)

		if err = os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			panic(err)
		}
		// nolint: gosec
		err := os.WriteFile(path, contents, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func writeGoClient(pkg *schema.Package, outdir string) {
	files, err := gogen.GeneratePackage("pulumigen", pkg, nil)
	if err != nil {
		panic(err)
	}

	mustWriteFiles(outdir, files)
}

func mustWriteFiles(rootDir string, files map[string][]byte) {
	for filename, contents := range files {
		mustWriteFile(rootDir, filename, contents)
	}
}

func mustWriteFile(rootDir, filename string, contents []byte) {
	outPath := filepath.Join(rootDir, filename)

	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		panic(err)
	}
	// nolint: gosec
	err := os.WriteFile(outPath, contents, 0644)
	if err != nil {
		panic(err)
	}
}
