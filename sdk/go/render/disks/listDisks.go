// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package disks

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDisks(ctx *pulumi.Context, args *ListDisksArgs, opts ...pulumi.InvokeOption) (*ListDisksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListDisksResult
	err := ctx.Invoke("render:disks:listDisks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDisksArgs struct {
}

type ListDisksResult struct {
	Items []DiskWithCursor `pulumi:"items"`
}

func ListDisksOutput(ctx *pulumi.Context, args ListDisksOutputArgs, opts ...pulumi.InvokeOption) ListDisksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDisksResult, error) {
			args := v.(ListDisksArgs)
			r, err := ListDisks(ctx, &args, opts...)
			var s ListDisksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDisksResultOutput)
}

type ListDisksOutputArgs struct {
}

func (ListDisksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDisksArgs)(nil)).Elem()
}

type ListDisksResultOutput struct{ *pulumi.OutputState }

func (ListDisksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDisksResult)(nil)).Elem()
}

func (o ListDisksResultOutput) ToListDisksResultOutput() ListDisksResultOutput {
	return o
}

func (o ListDisksResultOutput) ToListDisksResultOutputWithContext(ctx context.Context) ListDisksResultOutput {
	return o
}

func (o ListDisksResultOutput) Items() DiskWithCursorArrayOutput {
	return o.ApplyT(func(v ListDisksResult) []DiskWithCursor { return v.Items }).(DiskWithCursorArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDisksResultOutput{})
}