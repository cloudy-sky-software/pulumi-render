package provider

import (
	"context"

	"github.com/getkin/kin-openapi/openapi3"
)

const (
	isResource         = "x-isResource"
	isReadOnlyResource = "x-isReadOnlyResource"
)

func main() {
	doc, err := openapi3.NewLoader().LoadFromFile("openapi.yml")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	if err := doc.Validate(ctx); err != nil {
		panic(err)
	}
}
