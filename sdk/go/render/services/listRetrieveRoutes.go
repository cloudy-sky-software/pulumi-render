// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRetrieveRoutes(ctx *pulumi.Context, args *ListRetrieveRoutesArgs, opts ...pulumi.InvokeOption) (*ListRetrieveRoutesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListRetrieveRoutesResult
	err := ctx.Invoke("render:services:listRetrieveRoutes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRetrieveRoutesArgs struct {
}

type ListRetrieveRoutesResult struct {
	Items []ListRetrieveRoutesItemProperties `pulumi:"items"`
}

func ListRetrieveRoutesOutput(ctx *pulumi.Context, args ListRetrieveRoutesOutputArgs, opts ...pulumi.InvokeOption) ListRetrieveRoutesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListRetrieveRoutesResult, error) {
			args := v.(ListRetrieveRoutesArgs)
			r, err := ListRetrieveRoutes(ctx, &args, opts...)
			var s ListRetrieveRoutesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListRetrieveRoutesResultOutput)
}

type ListRetrieveRoutesOutputArgs struct {
}

func (ListRetrieveRoutesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRetrieveRoutesArgs)(nil)).Elem()
}

type ListRetrieveRoutesResultOutput struct{ *pulumi.OutputState }

func (ListRetrieveRoutesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRetrieveRoutesResult)(nil)).Elem()
}

func (o ListRetrieveRoutesResultOutput) ToListRetrieveRoutesResultOutput() ListRetrieveRoutesResultOutput {
	return o
}

func (o ListRetrieveRoutesResultOutput) ToListRetrieveRoutesResultOutputWithContext(ctx context.Context) ListRetrieveRoutesResultOutput {
	return o
}

func (o ListRetrieveRoutesResultOutput) Items() ListRetrieveRoutesItemPropertiesArrayOutput {
	return o.ApplyT(func(v ListRetrieveRoutesResult) []ListRetrieveRoutesItemProperties { return v.Items }).(ListRetrieveRoutesItemPropertiesArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRetrieveRoutesResultOutput{})
}
