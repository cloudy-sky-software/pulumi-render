// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SuspendService struct {
	pulumi.CustomResourceState
}

// NewSuspendService registers a new resource with the given unique name, arguments, and options.
func NewSuspendService(ctx *pulumi.Context,
	name string, args *SuspendServiceArgs, opts ...pulumi.ResourceOption) (*SuspendService, error) {
	if args == nil {
		args = &SuspendServiceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SuspendService
	err := ctx.RegisterResource("render:services:SuspendService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSuspendService gets an existing SuspendService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSuspendService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SuspendServiceState, opts ...pulumi.ResourceOption) (*SuspendService, error) {
	var resource SuspendService
	err := ctx.ReadResource("render:services:SuspendService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SuspendService resources.
type suspendServiceState struct {
}

type SuspendServiceState struct {
}

func (SuspendServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*suspendServiceState)(nil)).Elem()
}

type suspendServiceArgs struct {
	// The ID of the service
	ServiceId *string `pulumi:"serviceId"`
}

// The set of arguments for constructing a SuspendService resource.
type SuspendServiceArgs struct {
	// The ID of the service
	ServiceId pulumi.StringPtrInput
}

func (SuspendServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*suspendServiceArgs)(nil)).Elem()
}

type SuspendServiceInput interface {
	pulumi.Input

	ToSuspendServiceOutput() SuspendServiceOutput
	ToSuspendServiceOutputWithContext(ctx context.Context) SuspendServiceOutput
}

func (*SuspendService) ElementType() reflect.Type {
	return reflect.TypeOf((**SuspendService)(nil)).Elem()
}

func (i *SuspendService) ToSuspendServiceOutput() SuspendServiceOutput {
	return i.ToSuspendServiceOutputWithContext(context.Background())
}

func (i *SuspendService) ToSuspendServiceOutputWithContext(ctx context.Context) SuspendServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuspendServiceOutput)
}

type SuspendServiceOutput struct{ *pulumi.OutputState }

func (SuspendServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SuspendService)(nil)).Elem()
}

func (o SuspendServiceOutput) ToSuspendServiceOutput() SuspendServiceOutput {
	return o
}

func (o SuspendServiceOutput) ToSuspendServiceOutputWithContext(ctx context.Context) SuspendServiceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SuspendServiceInput)(nil)).Elem(), &SuspendService{})
	pulumi.RegisterOutputType(SuspendServiceOutput{})
}
