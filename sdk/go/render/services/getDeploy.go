// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

func LookupDeploy(ctx *pulumi.Context, args *LookupDeployArgs, opts ...pulumi.InvokeOption) (*LookupDeployResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDeployResult
	err := ctx.Invoke("render:services:getDeploy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDeployArgs struct {
	// (Required) The ID of the deploy
	Id string `pulumi:"id"`
	// (Required) The ID of the service
	ServiceId string `pulumi:"serviceId"`
}

type LookupDeployResult struct {
	Items DeployType `pulumi:"items"`
}

// Defaults sets the appropriate defaults for LookupDeployResult
func (val *LookupDeployResult) Defaults() *LookupDeployResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func LookupDeployOutput(ctx *pulumi.Context, args LookupDeployOutputArgs, opts ...pulumi.InvokeOption) LookupDeployResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeployResult, error) {
			args := v.(LookupDeployArgs)
			r, err := LookupDeploy(ctx, &args, opts...)
			var s LookupDeployResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeployResultOutput)
}

type LookupDeployOutputArgs struct {
	// (Required) The ID of the deploy
	Id pulumi.StringInput `pulumi:"id"`
	// (Required) The ID of the service
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (LookupDeployOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeployArgs)(nil)).Elem()
}

type LookupDeployResultOutput struct{ *pulumi.OutputState }

func (LookupDeployResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeployResult)(nil)).Elem()
}

func (o LookupDeployResultOutput) ToLookupDeployResultOutput() LookupDeployResultOutput {
	return o
}

func (o LookupDeployResultOutput) ToLookupDeployResultOutputWithContext(ctx context.Context) LookupDeployResultOutput {
	return o
}

func (o LookupDeployResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDeployResult] {
	return pulumix.Output[LookupDeployResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupDeployResultOutput) Items() DeployTypeOutput {
	return o.ApplyT(func(v LookupDeployResult) DeployType { return v.Items }).(DeployTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeployResultOutput{})
}
