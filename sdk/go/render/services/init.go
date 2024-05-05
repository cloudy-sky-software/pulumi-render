// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "render:services:AutoScaling":
		r = &AutoScaling{}
	case "render:services:BackgroundWorker":
		r = &BackgroundWorker{}
	case "render:services:CronJob":
		r = &CronJob{}
	case "render:services:CustomDomains":
		r = &CustomDomains{}
	case "render:services:Deploys":
		r = &Deploys{}
	case "render:services:EnvVars":
		r = &EnvVars{}
	case "render:services:Jobs":
		r = &Jobs{}
	case "render:services:PrivateService":
		r = &PrivateService{}
	case "render:services:Scale":
		r = &Scale{}
	case "render:services:StaticSite":
		r = &StaticSite{}
	case "render:services:Suspend":
		r = &Suspend{}
	case "render:services:WebService":
		r = &WebService{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"render",
		"services",
		&module{version},
	)
}
