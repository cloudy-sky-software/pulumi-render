package provider

import (
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
)

const stateKeyInputs = "__inputs"

var defaultMarshalOpts = plugin.MarshalOptions{KeepUnknowns: true, KeepSecrets: true, SkipNulls: true}

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

func applyDiff(newProps resource.PropertyMap, oldProps resource.PropertyMap) resource.PropertyMap {
	diff := oldProps.Diff(newProps)
	if diff == nil {
		return oldProps
	}

	result := resource.PropertyMap{}
	// Maintain inputs that we have that they don't have.
	for name, value := range oldProps {
		if !diff.Deleted(name) {
			result[name] = value
		}
	}
	// Take all the additions and updates from them.
	for key, value := range diff.Adds {
		result[key] = value
	}
	for key, value := range diff.Updates {
		result[key] = value.New
	}
	return result
}
