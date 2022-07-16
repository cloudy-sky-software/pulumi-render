// Copyright 2016-2020, Pulumi Corporation.
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

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	providerSchemaGen "github.com/cloudy-sky-software/pulumi-render/provider/pkg/gen"
	providerOpenAPI "github.com/cloudy-sky-software/pulumi-render/provider/pkg/openapi"
	providerVersion "github.com/cloudy-sky-software/pulumi-render/provider/pkg/version"
	"github.com/pkg/errors"

	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	nodejsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pythongen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

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

	var schema *schema.Package
	if language != Schema {
		schema = readSchema(inputFile, version)
	}

	switch language {
	case NodeJS:
		writeNodeJSClient(schema, outdir)
	case Python:
		writePythonClient(schema, outdir)
	case DotNet:
		writeDotnetClient(schema, outdir)
	case Go:
		writeGoClient(schema, outdir)
	case Schema:
		openapiDoc := providerOpenAPI.GetOpenAPISpec()
		schemaSpec := providerSchemaGen.PulumiSchema(*openapiDoc)
		providerDir := filepath.Join(".", "provider", "cmd", "pulumi-resource-render")
		mustWritePulumiSchema(schemaSpec, providerDir)
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
	schemaBytes, err := ioutil.ReadFile(schemaPath)
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
	files, err := nodejsgen.GeneratePackage("pulumigen", pkg, overlays)
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

	files, err := dotnetgen.GeneratePackage("pulumigen", pkg, overlays)
	if err != nil {
		panic(err)
	}

	for filename, contents := range files {
		path := filepath.Join(outdir, filename)

		if err = os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			panic(err)
		}
		err := ioutil.WriteFile(path, contents, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func writeGoClient(pkg *schema.Package, outdir string) {
	files, err := gogen.GeneratePackage("pulumigen", pkg)
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
	err := ioutil.WriteFile(outPath, contents, 0644)
	if err != nil {
		panic(err)
	}
}
