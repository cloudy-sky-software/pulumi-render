// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgres

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetPostgresConnectionInfo(ctx *pulumi.Context, args *GetPostgresConnectionInfoArgs, opts ...pulumi.InvokeOption) (*GetPostgresConnectionInfoResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPostgresConnectionInfoResult
	err := ctx.Invoke("render:postgres:getPostgresConnectionInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetPostgresConnectionInfoArgs struct {
	PostgresId string `pulumi:"postgresId"`
}

type GetPostgresConnectionInfoResult struct {
	Items PostgresConnectionInfo `pulumi:"items"`
}

func GetPostgresConnectionInfoOutput(ctx *pulumi.Context, args GetPostgresConnectionInfoOutputArgs, opts ...pulumi.InvokeOption) GetPostgresConnectionInfoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPostgresConnectionInfoResult, error) {
			args := v.(GetPostgresConnectionInfoArgs)
			r, err := GetPostgresConnectionInfo(ctx, &args, opts...)
			var s GetPostgresConnectionInfoResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPostgresConnectionInfoResultOutput)
}

type GetPostgresConnectionInfoOutputArgs struct {
	PostgresId pulumi.StringInput `pulumi:"postgresId"`
}

func (GetPostgresConnectionInfoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPostgresConnectionInfoArgs)(nil)).Elem()
}

type GetPostgresConnectionInfoResultOutput struct{ *pulumi.OutputState }

func (GetPostgresConnectionInfoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPostgresConnectionInfoResult)(nil)).Elem()
}

func (o GetPostgresConnectionInfoResultOutput) ToGetPostgresConnectionInfoResultOutput() GetPostgresConnectionInfoResultOutput {
	return o
}

func (o GetPostgresConnectionInfoResultOutput) ToGetPostgresConnectionInfoResultOutputWithContext(ctx context.Context) GetPostgresConnectionInfoResultOutput {
	return o
}

func (o GetPostgresConnectionInfoResultOutput) Items() PostgresConnectionInfoOutput {
	return o.ApplyT(func(v GetPostgresConnectionInfoResult) PostgresConnectionInfo { return v.Items }).(PostgresConnectionInfoOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPostgresConnectionInfoResultOutput{})
}
