// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CancelJob struct {
	pulumi.CustomResourceState

	CreatedAt    pulumi.StringOutput      `pulumi:"createdAt"`
	FinishedAt   pulumi.StringPtrOutput   `pulumi:"finishedAt"`
	PlanId       pulumi.StringOutput      `pulumi:"planId"`
	ServiceId    pulumi.StringOutput      `pulumi:"serviceId"`
	StartCommand pulumi.StringOutput      `pulumi:"startCommand"`
	StartedAt    pulumi.StringPtrOutput   `pulumi:"startedAt"`
	Status       CancelJobStatusPtrOutput `pulumi:"status"`
}

// NewCancelJob registers a new resource with the given unique name, arguments, and options.
func NewCancelJob(ctx *pulumi.Context,
	name string, args *CancelJobArgs, opts ...pulumi.ResourceOption) (*CancelJob, error) {
	if args == nil {
		args = &CancelJobArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CancelJob
	err := ctx.RegisterResource("render:services:CancelJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCancelJob gets an existing CancelJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCancelJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CancelJobState, opts ...pulumi.ResourceOption) (*CancelJob, error) {
	var resource CancelJob
	err := ctx.ReadResource("render:services:CancelJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CancelJob resources.
type cancelJobState struct {
}

type CancelJobState struct {
}

func (CancelJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*cancelJobState)(nil)).Elem()
}

type cancelJobArgs struct {
	// The ID of the job
	JobId *string `pulumi:"jobId"`
	// The ID of the service
	ServiceId *string `pulumi:"serviceId"`
}

// The set of arguments for constructing a CancelJob resource.
type CancelJobArgs struct {
	// The ID of the job
	JobId pulumi.StringPtrInput
	// The ID of the service
	ServiceId pulumi.StringPtrInput
}

func (CancelJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cancelJobArgs)(nil)).Elem()
}

type CancelJobInput interface {
	pulumi.Input

	ToCancelJobOutput() CancelJobOutput
	ToCancelJobOutputWithContext(ctx context.Context) CancelJobOutput
}

func (*CancelJob) ElementType() reflect.Type {
	return reflect.TypeOf((**CancelJob)(nil)).Elem()
}

func (i *CancelJob) ToCancelJobOutput() CancelJobOutput {
	return i.ToCancelJobOutputWithContext(context.Background())
}

func (i *CancelJob) ToCancelJobOutputWithContext(ctx context.Context) CancelJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CancelJobOutput)
}

type CancelJobOutput struct{ *pulumi.OutputState }

func (CancelJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CancelJob)(nil)).Elem()
}

func (o CancelJobOutput) ToCancelJobOutput() CancelJobOutput {
	return o
}

func (o CancelJobOutput) ToCancelJobOutputWithContext(ctx context.Context) CancelJobOutput {
	return o
}

func (o CancelJobOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CancelJob) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o CancelJobOutput) FinishedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CancelJob) pulumi.StringPtrOutput { return v.FinishedAt }).(pulumi.StringPtrOutput)
}

func (o CancelJobOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *CancelJob) pulumi.StringOutput { return v.PlanId }).(pulumi.StringOutput)
}

func (o CancelJobOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CancelJob) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

func (o CancelJobOutput) StartCommand() pulumi.StringOutput {
	return o.ApplyT(func(v *CancelJob) pulumi.StringOutput { return v.StartCommand }).(pulumi.StringOutput)
}

func (o CancelJobOutput) StartedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CancelJob) pulumi.StringPtrOutput { return v.StartedAt }).(pulumi.StringPtrOutput)
}

func (o CancelJobOutput) Status() CancelJobStatusPtrOutput {
	return o.ApplyT(func(v *CancelJob) CancelJobStatusPtrOutput { return v.Status }).(CancelJobStatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CancelJobInput)(nil)).Elem(), &CancelJob{})
	pulumi.RegisterOutputType(CancelJobOutput{})
}
