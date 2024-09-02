// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListHeaders(ctx *pulumi.Context, args *ListHeadersArgs, opts ...pulumi.InvokeOption) (*ListHeadersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListHeadersResult
	err := ctx.Invoke("render:services:listHeaders", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListHeadersArgs struct {
	// The ID of the service
	ServiceId string `pulumi:"serviceId"`
}

type ListHeadersResult struct {
	Items []HeaderWithCursor `pulumi:"items"`
}

func ListHeadersOutput(ctx *pulumi.Context, args ListHeadersOutputArgs, opts ...pulumi.InvokeOption) ListHeadersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListHeadersResult, error) {
			args := v.(ListHeadersArgs)
			r, err := ListHeaders(ctx, &args, opts...)
			var s ListHeadersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListHeadersResultOutput)
}

type ListHeadersOutputArgs struct {
	// The ID of the service
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (ListHeadersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListHeadersArgs)(nil)).Elem()
}

type ListHeadersResultOutput struct{ *pulumi.OutputState }

func (ListHeadersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListHeadersResult)(nil)).Elem()
}

func (o ListHeadersResultOutput) ToListHeadersResultOutput() ListHeadersResultOutput {
	return o
}

func (o ListHeadersResultOutput) ToListHeadersResultOutputWithContext(ctx context.Context) ListHeadersResultOutput {
	return o
}

func (o ListHeadersResultOutput) Items() HeaderWithCursorArrayOutput {
	return o.ApplyT(func(v ListHeadersResult) []HeaderWithCursor { return v.Items }).(HeaderWithCursorArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListHeadersResultOutput{})
}
