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
	Items []PostgresWithCursor `pulumi:"items"`
}

func ListPostgresOutput(ctx *pulumi.Context, args ListPostgresOutputArgs, opts ...pulumi.InvokeOption) ListPostgresResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListPostgresResultOutput, error) {
			args := v.(ListPostgresArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("render:postgres:listPostgres", args, ListPostgresResultOutput{}, options).(ListPostgresResultOutput), nil
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

func (o ListPostgresResultOutput) Items() PostgresWithCursorArrayOutput {
	return o.ApplyT(func(v ListPostgresResult) []PostgresWithCursor { return v.Items }).(PostgresWithCursorArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListPostgresResultOutput{})
}
