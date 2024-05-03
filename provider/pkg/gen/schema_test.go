package gen

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cloudy-sky-software/pulumi-provider-framework/openapi"
	"github.com/stretchr/testify/assert"
)

func TestPulumiSchema(t *testing.T) {
	b, err := os.ReadFile(filepath.Join("..", "..", "cmd", "pulumi-gen-render", "openapi.yml"))
	if err != nil {
		t.Fatalf("Failed reading openapi.yml: %v", err)
	}

	oaSpec := openapi.GetOpenAPISpec(b)
	err = FixOpenAPIDoc(oaSpec)
	assert.Nil(t, err)
	PulumiSchema(*oaSpec)
}
