// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupJobResult
	err := ctx.Invoke("render:services:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	// The ID of the job
	JobId string `pulumi:"jobId"`
	// The ID of the service
	ServiceId string `pulumi:"serviceId"`
}

type LookupJobResult struct {
	CreatedAt    string     `pulumi:"createdAt"`
	FinishedAt   *string    `pulumi:"finishedAt"`
	Id           string     `pulumi:"id"`
	PlanId       string     `pulumi:"planId"`
	ServiceId    string     `pulumi:"serviceId"`
	StartCommand string     `pulumi:"startCommand"`
	StartedAt    *string    `pulumi:"startedAt"`
	Status       *JobStatus `pulumi:"status"`
}

func LookupJobOutput(ctx *pulumi.Context, args LookupJobOutputArgs, opts ...pulumi.InvokeOption) LookupJobResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupJobResultOutput, error) {
			args := v.(LookupJobArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("render:services:getJob", args, LookupJobResultOutput{}, options).(LookupJobResultOutput), nil
		}).(LookupJobResultOutput)
}

type LookupJobOutputArgs struct {
	// The ID of the job
	JobId pulumi.StringInput `pulumi:"jobId"`
	// The ID of the service
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (LookupJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobArgs)(nil)).Elem()
}

type LookupJobResultOutput struct{ *pulumi.OutputState }

func (LookupJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobResult)(nil)).Elem()
}

func (o LookupJobResultOutput) ToLookupJobResultOutput() LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) ToLookupJobResultOutputWithContext(ctx context.Context) LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) FinishedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.FinishedAt }).(pulumi.StringPtrOutput)
}

func (o LookupJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.PlanId }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) StartCommand() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.StartCommand }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) StartedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.StartedAt }).(pulumi.StringPtrOutput)
}

func (o LookupJobResultOutput) Status() JobStatusPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *JobStatus { return v.Status }).(JobStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobResultOutput{})
}
