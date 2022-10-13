// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListJobs(ctx *pulumi.Context, args *ListJobsArgs, opts ...pulumi.InvokeOption) (*ListJobsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv ListJobsResult
	err := ctx.Invoke("render:services:listJobs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListJobsArgs struct {
	// (Required) The ID of the service
	ServiceId string `pulumi:"serviceId"`
}

type ListJobsResult struct {
	Items []ListJobsResponse `pulumi:"items"`
}

func ListJobsOutput(ctx *pulumi.Context, args ListJobsOutputArgs, opts ...pulumi.InvokeOption) ListJobsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListJobsResult, error) {
			args := v.(ListJobsArgs)
			r, err := ListJobs(ctx, &args, opts...)
			var s ListJobsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListJobsResultOutput)
}

type ListJobsOutputArgs struct {
	// (Required) The ID of the service
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (ListJobsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListJobsArgs)(nil)).Elem()
}

type ListJobsResultOutput struct{ *pulumi.OutputState }

func (ListJobsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListJobsResult)(nil)).Elem()
}

func (o ListJobsResultOutput) ToListJobsResultOutput() ListJobsResultOutput {
	return o
}

func (o ListJobsResultOutput) ToListJobsResultOutputWithContext(ctx context.Context) ListJobsResultOutput {
	return o
}

func (o ListJobsResultOutput) Items() ListJobsResponseArrayOutput {
	return o.ApplyT(func(v ListJobsResult) []ListJobsResponse { return v.Items }).(ListJobsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListJobsResultOutput{})
}