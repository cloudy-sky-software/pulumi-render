// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"errors"
	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Job struct {
	pulumi.CustomResourceState

	CreatedAt    pulumi.StringOutput    `pulumi:"createdAt"`
	FinishedAt   pulumi.StringPtrOutput `pulumi:"finishedAt"`
	PlanId       pulumi.StringOutput    `pulumi:"planId"`
	ServiceId    pulumi.StringOutput    `pulumi:"serviceId"`
	StartCommand pulumi.StringOutput    `pulumi:"startCommand"`
	StartedAt    pulumi.StringPtrOutput `pulumi:"startedAt"`
	Status       JobStatusPtrOutput     `pulumi:"status"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StartCommand == nil {
		return nil, errors.New("invalid value for required argument 'StartCommand'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Job
	err := ctx.RegisterResource("render:services:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("render:services:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
}

type JobState struct {
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	PlanId *string `pulumi:"planId"`
	// The ID of the service
	ServiceId    *string `pulumi:"serviceId"`
	StartCommand string  `pulumi:"startCommand"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	PlanId pulumi.StringPtrInput
	// The ID of the service
	ServiceId    pulumi.StringPtrInput
	StartCommand pulumi.StringInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

func (o JobOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o JobOutput) FinishedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.FinishedAt }).(pulumi.StringPtrOutput)
}

func (o JobOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.PlanId }).(pulumi.StringOutput)
}

func (o JobOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

func (o JobOutput) StartCommand() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.StartCommand }).(pulumi.StringOutput)
}

func (o JobOutput) StartedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.StartedAt }).(pulumi.StringPtrOutput)
}

func (o JobOutput) Status() JobStatusPtrOutput {
	return o.ApplyT(func(v *Job) JobStatusPtrOutput { return v.Status }).(JobStatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobInput)(nil)).Elem(), &Job{})
	pulumi.RegisterOutputType(JobOutput{})
}
