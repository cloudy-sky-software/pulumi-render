package provider

import "github.com/pulumi/pulumi/sdk/v3/go/common/resource"

const stateKeyInputs = "__inputs"

func getResourceState(outputs map[string]interface{}, inputs resource.PropertyMap) resource.PropertyMap {
	state := resource.NewPropertyMapFromMap(outputs)
	// Capture the inputs as they were during the creation of the resource
	// so that we can use them during diff if the resource is updated.
	state[stateKeyInputs] = resource.MakeSecret(resource.NewObjectProperty(inputs))
	return state
}

func getOldInputs(state resource.PropertyMap) resource.PropertyMap {
	if v, ok := state[stateKeyInputs]; ok {
		return v.SecretValue().Element.ObjectValue()
	}

	return nil
}
