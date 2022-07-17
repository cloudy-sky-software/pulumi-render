package gen

import (
	"testing"

	"github.com/cloudy-sky-software/pulumi-render/provider/pkg/openapi"
)

func TestPulumiSchema(t *testing.T) {
	oaSpec := openapi.GetOpenAPISpec()

	PulumiSchema(*oaSpec)
}
