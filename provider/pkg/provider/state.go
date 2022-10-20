package provider

import (
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
)

var defaultMarshalOpts = plugin.MarshalOptions{KeepUnknowns: true, KeepSecrets: true, SkipNulls: true}
var defaultUnmarshalOpts = plugin.MarshalOptions{KeepUnknowns: true, KeepSecrets: true, SkipNulls: true}
