package provider

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

const (
	authSchemePrefix = "Bearer"
)

func (p *renderProvider) determineDiffsAndReplacements(d *resource.ObjectDiff, properties openapi3.Schemas) ([]string, []string) {
	replaces := make([]string, 0)
	diffs := make([]string, 0)

	for propKey := range d.Adds {
		prop := string(propKey)
		// If the added property is not part of the PATCH operation schema,
		// then suggest a replacement triggered by this property.
		if _, ok := properties[prop]; !ok {
			replaces = append(replaces, prop)
		} else {
			diffs = append(diffs, prop)
		}
	}

	for propKey := range d.Updates {
		prop := string(propKey)
		// If the updated property is not part of the PATCH operation schema,
		// then suggest a replacement triggered by this property.
		if _, ok := properties[prop]; !ok {
			replaces = append(replaces, prop)
		} else {
			diffs = append(diffs, prop)
		}
	}

	for propKey := range d.Deletes {
		prop := string(propKey)
		// If the deleted property is not part of the PATCH operation schema,
		// then suggest a replacement triggered by this property.
		if _, ok := properties[prop]; !ok {
			replaces = append(replaces, prop)
		} else {
			diffs = append(diffs, prop)
		}
	}

	return replaces, diffs
}
