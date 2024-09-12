// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRedisConnectionInfo(ctx *pulumi.Context, args *GetRedisConnectionInfoArgs, opts ...pulumi.InvokeOption) (*GetRedisConnectionInfoResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRedisConnectionInfoResult
	err := ctx.Invoke("render:redis:getRedisConnectionInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetRedisConnectionInfoArgs struct {
	RedisId string `pulumi:"redisId"`
}

// A Redis instance
type GetRedisConnectionInfoResult struct {
	// The connection string to use from outside Render
	ExternalConnectionString string `pulumi:"externalConnectionString"`
	// The connection string to use from within Render
	InternalConnectionString string `pulumi:"internalConnectionString"`
	// The Redis CLI command to connect to the Redis instance
	RedisCLICommand string `pulumi:"redisCLICommand"`
}

func GetRedisConnectionInfoOutput(ctx *pulumi.Context, args GetRedisConnectionInfoOutputArgs, opts ...pulumi.InvokeOption) GetRedisConnectionInfoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRedisConnectionInfoResultOutput, error) {
			args := v.(GetRedisConnectionInfoArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetRedisConnectionInfoResult
			secret, err := ctx.InvokePackageRaw("render:redis:getRedisConnectionInfo", args, &rv, "", opts...)
			if err != nil {
				return GetRedisConnectionInfoResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetRedisConnectionInfoResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetRedisConnectionInfoResultOutput), nil
			}
			return output, nil
		}).(GetRedisConnectionInfoResultOutput)
}

type GetRedisConnectionInfoOutputArgs struct {
	RedisId pulumi.StringInput `pulumi:"redisId"`
}

func (GetRedisConnectionInfoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRedisConnectionInfoArgs)(nil)).Elem()
}

// A Redis instance
type GetRedisConnectionInfoResultOutput struct{ *pulumi.OutputState }

func (GetRedisConnectionInfoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRedisConnectionInfoResult)(nil)).Elem()
}

func (o GetRedisConnectionInfoResultOutput) ToGetRedisConnectionInfoResultOutput() GetRedisConnectionInfoResultOutput {
	return o
}

func (o GetRedisConnectionInfoResultOutput) ToGetRedisConnectionInfoResultOutputWithContext(ctx context.Context) GetRedisConnectionInfoResultOutput {
	return o
}

// The connection string to use from outside Render
func (o GetRedisConnectionInfoResultOutput) ExternalConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v GetRedisConnectionInfoResult) string { return v.ExternalConnectionString }).(pulumi.StringOutput)
}

// The connection string to use from within Render
func (o GetRedisConnectionInfoResultOutput) InternalConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v GetRedisConnectionInfoResult) string { return v.InternalConnectionString }).(pulumi.StringOutput)
}

// The Redis CLI command to connect to the Redis instance
func (o GetRedisConnectionInfoResultOutput) RedisCLICommand() pulumi.StringOutput {
	return o.ApplyT(func(v GetRedisConnectionInfoResult) string { return v.RedisCLICommand }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRedisConnectionInfoResultOutput{})
}
