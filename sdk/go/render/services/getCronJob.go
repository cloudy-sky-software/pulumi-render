// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCronJob(ctx *pulumi.Context, args *LookupCronJobArgs, opts ...pulumi.InvokeOption) (*LookupCronJobResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCronJobResult
	err := ctx.Invoke("render:services:getCronJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupCronJobArgs struct {
	// The ID of the service
	ServiceId string `pulumi:"serviceId"`
}

type LookupCronJobResult struct {
	Items GetCronJobType `pulumi:"items"`
}

// Defaults sets the appropriate defaults for LookupCronJobResult
func (val *LookupCronJobResult) Defaults() *LookupCronJobResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func LookupCronJobOutput(ctx *pulumi.Context, args LookupCronJobOutputArgs, opts ...pulumi.InvokeOption) LookupCronJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCronJobResult, error) {
			args := v.(LookupCronJobArgs)
			r, err := LookupCronJob(ctx, &args, opts...)
			var s LookupCronJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCronJobResultOutput)
}

type LookupCronJobOutputArgs struct {
	// The ID of the service
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (LookupCronJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCronJobArgs)(nil)).Elem()
}

type LookupCronJobResultOutput struct{ *pulumi.OutputState }

func (LookupCronJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCronJobResult)(nil)).Elem()
}

func (o LookupCronJobResultOutput) ToLookupCronJobResultOutput() LookupCronJobResultOutput {
	return o
}

func (o LookupCronJobResultOutput) ToLookupCronJobResultOutputWithContext(ctx context.Context) LookupCronJobResultOutput {
	return o
}

func (o LookupCronJobResultOutput) Items() GetCronJobTypeOutput {
	return o.ApplyT(func(v LookupCronJobResult) GetCronJobType { return v.Items }).(GetCronJobTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCronJobResultOutput{})
}