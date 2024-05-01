// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListServices(ctx *pulumi.Context, args *ListServicesArgs, opts ...pulumi.InvokeOption) (*ListServicesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListServicesResult
	err := ctx.Invoke("render:services:listServices", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListServicesArgs struct {
}

type ListServicesResult struct {
	Items []ListServiceResponse `pulumi:"items"`
}

func ListServicesOutput(ctx *pulumi.Context, args ListServicesOutputArgs, opts ...pulumi.InvokeOption) ListServicesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListServicesResult, error) {
			args := v.(ListServicesArgs)
			r, err := ListServices(ctx, &args, opts...)
			var s ListServicesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListServicesResultOutput)
}

type ListServicesOutputArgs struct {
}

func (ListServicesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServicesArgs)(nil)).Elem()
}

type ListServicesResultOutput struct{ *pulumi.OutputState }

func (ListServicesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServicesResult)(nil)).Elem()
}

func (o ListServicesResultOutput) ToListServicesResultOutput() ListServicesResultOutput {
	return o
}

func (o ListServicesResultOutput) ToListServicesResultOutputWithContext(ctx context.Context) ListServicesResultOutput {
	return o
}

func (o ListServicesResultOutput) Items() ListServiceResponseArrayOutput {
	return o.ApplyT(func(v ListServicesResult) []ListServiceResponse { return v.Items }).(ListServiceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListServicesResultOutput{})
}
