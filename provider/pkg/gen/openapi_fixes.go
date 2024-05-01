package gen

import (
	"fmt"
	"net/url"
	"strings"

	pulschema_pkg "github.com/cloudy-sky-software/pulschema/pkg"
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/pkg/errors"
)

func getResourceFromPath(path string) (string, error) {
	u, err := url.Parse(path)
	if err != nil {
		return "", fmt.Errorf("invalid path: %s", path)
	}

	pathParts := strings.Split(u.Path, "/")
	var resourceName string
	var segmentIndex int
	// If the path ends with a path parameter,
	// we should use the path segment before
	// the last one.
	if strings.HasSuffix(path, "}") {
		segmentIndex = len(pathParts) - 2
	} else {
		// Otherwise, use the last segment.
		segmentIndex = len(pathParts) - 1
	}

	resourceName = pathParts[segmentIndex]
	return pulschema_pkg.ToPascalCase(resourceName), nil
}

// Render's API operations do not have an operationId,
// so we'll need to generate them based on the resource
// in the operation path.
func FixOpenAPIDoc(openAPIDoc *openapi3.T) error {
	for _, path := range openAPIDoc.Paths.InMatchingOrder() {
		pathItem := openAPIDoc.Paths.Find(path)
		if pathItem == nil {
			return errors.Errorf("path item not found for path %s", path)
		}

		resourceName, err := getResourceFromPath(path)
		if err != nil {
			return errors.Wrap(err, "getting resource name")
		}

		endsWithPathParam := strings.HasSuffix(path, "}")

		if pathItem.Get != nil {
			if endsWithPathParam {
				pathItem.Get.OperationID = "get" + resourceName
			} else {
				pathItem.Get.OperationID = "list" + resourceName
			}
		}

		if pathItem.Post != nil {
			pathItem.Post.OperationID = "create" + resourceName
		}

		if pathItem.Put != nil {
			// PUT request methods are special.
			// If the current path does not also
			// have a POST request endpoint, then
			// the PUT can be used to create the
			// resource but only if it also has a
			// DELETE endpoint. Otherwise, it's a
			// one-shot "replace" type of operation
			// on the current resource.
			if pathItem.Post == nil && pathItem.Delete != nil {
				pathItem.Put.OperationID = "create" + resourceName
			} else {
				pathItem.Put.OperationID = "replace" + resourceName
			}
		}

		if pathItem.Patch != nil {
			pathItem.Patch.OperationID = "update" + resourceName
		}

		if pathItem.Delete != nil {
			pathItem.Delete.OperationID = "delete" + resourceName
		}
	}

	return nil
}
