// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRedis(ctx *pulumi.Context, args *ListRedisArgs, opts ...pulumi.InvokeOption) (*ListRedisResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListRedisResult
	err := ctx.Invoke("render:redis:listRedis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRedisArgs struct {
}

type ListRedisResult struct {
	Items []RedisWithCursor `pulumi:"items"`
}

func ListRedisOutput(ctx *pulumi.Context, args ListRedisOutputArgs, opts ...pulumi.InvokeOption) ListRedisResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListRedisResultOutput, error) {
			args := v.(ListRedisArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("render:redis:listRedis", args, ListRedisResultOutput{}, options).(ListRedisResultOutput), nil
		}).(ListRedisResultOutput)
}

type ListRedisOutputArgs struct {
}

func (ListRedisOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRedisArgs)(nil)).Elem()
}

type ListRedisResultOutput struct{ *pulumi.OutputState }

func (ListRedisResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRedisResult)(nil)).Elem()
}

func (o ListRedisResultOutput) ToListRedisResultOutput() ListRedisResultOutput {
	return o
}

func (o ListRedisResultOutput) ToListRedisResultOutputWithContext(ctx context.Context) ListRedisResultOutput {
	return o
}

func (o ListRedisResultOutput) Items() RedisWithCursorArrayOutput {
	return o.ApplyT(func(v ListRedisResult) []RedisWithCursor { return v.Items }).(RedisWithCursorArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRedisResultOutput{})
}
