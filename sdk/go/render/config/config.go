// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/services"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// The Render API key
func GetApiKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "render:apiKey")
}

// When a service is updated, a deployment is automatically triggered. This variable controls whether or not the service cache should be cleared upon deployment.
func GetClearCacheOnServiceUpdateDeployments(ctx *pulumi.Context) string {
	return config.Get(ctx, "render:clearCacheOnServiceUpdateDeployments")
}