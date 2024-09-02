// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgres

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListPostgres(ctx *pulumi.Context, args *ListPostgresArgs, opts ...pulumi.InvokeOption) (*ListPostgresResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListPostgresResult
	err := ctx.Invoke("render:postgres:listPostgres", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPostgresArgs struct {
}

type ListPostgresResult struct {
	Items []ListPostgresItemProperties `pulumi:"items"`
}

func ListPostgresOutput(ctx *pulumi.Context, args ListPostgresOutputArgs, opts ...pulumi.InvokeOption) ListPostgresResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListPostgresResult, error) {
			args := v.(ListPostgresArgs)
			r, err := ListPostgres(ctx, &args, opts...)
			var s ListPostgresResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListPostgresResultOutput)
}

type ListPostgresOutputArgs struct {
}

func (ListPostgresOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPostgresArgs)(nil)).Elem()
}

type ListPostgresResultOutput struct{ *pulumi.OutputState }

func (ListPostgresResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPostgresResult)(nil)).Elem()
}

func (o ListPostgresResultOutput) ToListPostgresResultOutput() ListPostgresResultOutput {
	return o
}

func (o ListPostgresResultOutput) ToListPostgresResultOutputWithContext(ctx context.Context) ListPostgresResultOutput {
	return o
}

func (o ListPostgresResultOutput) Items() ListPostgresItemPropertiesArrayOutput {
	return o.ApplyT(func(v ListPostgresResult) []ListPostgresItemProperties { return v.Items }).(ListPostgresItemPropertiesArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListPostgresResultOutput{})
}
