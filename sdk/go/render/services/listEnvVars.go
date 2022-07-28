// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListEnvVars(ctx *pulumi.Context, args *ListEnvVarsArgs, opts ...pulumi.InvokeOption) (*ListEnvVarsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv ListEnvVarsResult
	err := ctx.Invoke("render:services:listEnvVars", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEnvVarsArgs struct {
	// (Required) The ID of the service
	ServiceId string `pulumi:"serviceId"`
}

type ListEnvVarsResult struct {
	Value [][]ListEnvVarsResponse `pulumi:"value"`
}

func ListEnvVarsOutput(ctx *pulumi.Context, args ListEnvVarsOutputArgs, opts ...pulumi.InvokeOption) ListEnvVarsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListEnvVarsResult, error) {
			args := v.(ListEnvVarsArgs)
			r, err := ListEnvVars(ctx, &args, opts...)
			var s ListEnvVarsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListEnvVarsResultOutput)
}

type ListEnvVarsOutputArgs struct {
	// (Required) The ID of the service
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (ListEnvVarsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEnvVarsArgs)(nil)).Elem()
}

type ListEnvVarsResultOutput struct{ *pulumi.OutputState }

func (ListEnvVarsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEnvVarsResult)(nil)).Elem()
}

func (o ListEnvVarsResultOutput) ToListEnvVarsResultOutput() ListEnvVarsResultOutput {
	return o
}

func (o ListEnvVarsResultOutput) ToListEnvVarsResultOutputWithContext(ctx context.Context) ListEnvVarsResultOutput {
	return o
}

func (o ListEnvVarsResultOutput) Value() ListEnvVarsResponseArrayArrayOutput {
	return o.ApplyT(func(v ListEnvVarsResult) [][]ListEnvVarsResponse { return v.Value }).(ListEnvVarsResponseArrayArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListEnvVarsResultOutput{})
}
