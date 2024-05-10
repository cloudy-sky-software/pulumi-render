// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListJob(ctx *pulumi.Context, args *ListJobArgs, opts ...pulumi.InvokeOption) (*ListJobResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListJobResult
	err := ctx.Invoke("render:services:listJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListJobArgs struct {
}

type ListJobResult struct {
	Items []ListJobItemProperties `pulumi:"items"`
}

func ListJobOutput(ctx *pulumi.Context, args ListJobOutputArgs, opts ...pulumi.InvokeOption) ListJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListJobResult, error) {
			args := v.(ListJobArgs)
			r, err := ListJob(ctx, &args, opts...)
			var s ListJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListJobResultOutput)
}

type ListJobOutputArgs struct {
}

func (ListJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListJobArgs)(nil)).Elem()
}

type ListJobResultOutput struct{ *pulumi.OutputState }

func (ListJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListJobResult)(nil)).Elem()
}

func (o ListJobResultOutput) ToListJobResultOutput() ListJobResultOutput {
	return o
}

func (o ListJobResultOutput) ToListJobResultOutputWithContext(ctx context.Context) ListJobResultOutput {
	return o
}

func (o ListJobResultOutput) Items() ListJobItemPropertiesArrayOutput {
	return o.ApplyT(func(v ListJobResult) []ListJobItemProperties { return v.Items }).(ListJobItemPropertiesArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListJobResultOutput{})
}
