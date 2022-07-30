// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackgroundWorker(ctx *pulumi.Context, args *LookupBackgroundWorkerArgs, opts ...pulumi.InvokeOption) (*LookupBackgroundWorkerResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupBackgroundWorkerResult
	err := ctx.Invoke("render:services:getBackgroundWorker", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupBackgroundWorkerArgs struct {
	// (Required) The ID of the service
	Id string `pulumi:"id"`
}

type LookupBackgroundWorkerResult struct {
	Items GetBackgroundWorkerType `pulumi:"items"`
}

// Defaults sets the appropriate defaults for LookupBackgroundWorkerResult
func (val *LookupBackgroundWorkerResult) Defaults() *LookupBackgroundWorkerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func LookupBackgroundWorkerOutput(ctx *pulumi.Context, args LookupBackgroundWorkerOutputArgs, opts ...pulumi.InvokeOption) LookupBackgroundWorkerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackgroundWorkerResult, error) {
			args := v.(LookupBackgroundWorkerArgs)
			r, err := LookupBackgroundWorker(ctx, &args, opts...)
			var s LookupBackgroundWorkerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackgroundWorkerResultOutput)
}

type LookupBackgroundWorkerOutputArgs struct {
	// (Required) The ID of the service
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupBackgroundWorkerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackgroundWorkerArgs)(nil)).Elem()
}

type LookupBackgroundWorkerResultOutput struct{ *pulumi.OutputState }

func (LookupBackgroundWorkerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackgroundWorkerResult)(nil)).Elem()
}

func (o LookupBackgroundWorkerResultOutput) ToLookupBackgroundWorkerResultOutput() LookupBackgroundWorkerResultOutput {
	return o
}

func (o LookupBackgroundWorkerResultOutput) ToLookupBackgroundWorkerResultOutputWithContext(ctx context.Context) LookupBackgroundWorkerResultOutput {
	return o
}

func (o LookupBackgroundWorkerResultOutput) Items() GetBackgroundWorkerTypeOutput {
	return o.ApplyT(func(v LookupBackgroundWorkerResult) GetBackgroundWorkerType { return v.Items }).(GetBackgroundWorkerTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackgroundWorkerResultOutput{})
}