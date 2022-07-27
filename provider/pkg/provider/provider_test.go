package provider

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

func readFileFromProviderResourceDir(t *testing.T, filename string) []byte {
	t.Helper()

	b, err := os.ReadFile(filepath.Join("..", "..", "cmd", "pulumi-resource-render", filename))
	if err != nil {
		t.Fatalf("Failed reading openapi.yml: %v", err)
	}

	return b
}

func TestDiff(t *testing.T) {
	ctx := context.Background()

	openapiBytes := readFileFromProviderResourceDir(t, "openapi_generated.yml")

	pSchemaBytes := readFileFromProviderResourceDir(t, "schema.json")
	metadataBytes := readFileFromProviderResourceDir(t, "metadata.json")

	p, _ := makeProvider(nil, "", "", pSchemaBytes, openapiBytes, metadataBytes)

	outputs := make(map[string]interface{})
	outputs["name"] = "Test"
	oldsStruct, _ := plugin.MarshalProperties(getResourceState(outputs, resource.NewPropertyMapFromMap(outputs)), defaultMarshalOpts)

	news := make(map[string]interface{})
	news["name"] = "Test2"
	newsStruct, _ := plugin.MarshalProperties(resource.NewPropertyMapFromMap(news), defaultMarshalOpts)

	resp, err := p.Diff(ctx, &pulumirpc.DiffRequest{Id: "", Urn: "urn:pulumi:some-stack::some-project::render:services:StaticSite::someResourceName", Olds: oldsStruct, News: newsStruct})
	assert.Nil(t, err)
	assert.Equal(t, pulumirpc.DiffResponse_DIFF_SOME, resp.Changes)
	assert.NotEmpty(t, resp.Diffs)
	assert.Len(t, resp.Diffs, 1)
	assert.Empty(t, resp.Replaces)
}

func TestGetResourceState(t *testing.T) {
	outputs := make(map[string]interface{})
	outputs["name"] = "Test"
	outputs["id"] = "someid"

	inputs := make(map[string]interface{})
	inputs["name"] = "Test"

	state := getResourceState(outputs, resource.NewPropertyMapFromMap(inputs))
	assert.True(t, state.HasValue(resource.PropertyKey(stateKeyInputs)))
}
