package provider

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"

	providerSchemaGen "github.com/cloudy-sky-software/pulumi-render/provider/pkg/gen"
)

var arrayResourceTypeTokenToRootObjectName = map[string]string{
	envVarResourceTypeToken:      providerSchemaGen.EnvVarsRootObjectName,
	secretFilesResourceTypeToken: providerSchemaGen.SecretFilesRootObjectName,
}

func unWrapArrayRequest(httpReq *http.Request, resourceTypeToken string) ([]byte, error) {
	body, err := io.ReadAll(httpReq.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "reading body while handling %s", resourceTypeToken)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(body, &m); err != nil {
		return nil, errors.Wrap(err, "unmarshaling body")
	}

	rootObjectName := arrayResourceTypeTokenToRootObjectName[resourceTypeToken]

	modifiedRequestBody, err := json.Marshal(m[rootObjectName])
	if err != nil {
		return nil, errors.Wrap(err, "marshaling modified body")
	}

	return modifiedRequestBody, nil
}

func wrapArrayResponse(outputs interface{}, resourceTypeToken string) map[string]interface{} {
	var outputsMap map[string]interface{}
	resp := outputs.([]interface{})

	rootObjectName := arrayResourceTypeTokenToRootObjectName[resourceTypeToken]

	outputsMap = map[string]interface{}{
		rootObjectName: resp,
	}
	return outputsMap
}
