// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateService struct {
	pulumi.CustomResourceState

	AutoDeploy     PrivateServiceAutoDeployOutput          `pulumi:"autoDeploy"`
	Branch         pulumi.StringPtrOutput                  `pulumi:"branch"`
	BuildFilter    BuildFilterPtrOutput                    `pulumi:"buildFilter"`
	CreatedAt      pulumi.StringOutput                     `pulumi:"createdAt"`
	ImagePath      pulumi.StringPtrOutput                  `pulumi:"imagePath"`
	Name           pulumi.StringOutput                     `pulumi:"name"`
	NotifyOnFail   PrivateServiceNotifyOnFailOutput        `pulumi:"notifyOnFail"`
	OwnerId        pulumi.StringOutput                     `pulumi:"ownerId"`
	Repo           pulumi.StringPtrOutput                  `pulumi:"repo"`
	RootDir        pulumi.StringOutput                     `pulumi:"rootDir"`
	ServiceDetails pulumi.AnyOutput                        `pulumi:"serviceDetails"`
	Slug           pulumi.StringOutput                     `pulumi:"slug"`
	Suspended      PrivateServiceSuspendedOutput           `pulumi:"suspended"`
	Suspenders     PrivateServiceSuspendersItemArrayOutput `pulumi:"suspenders"`
	Type           pulumi.StringOutput                     `pulumi:"type"`
	UpdatedAt      pulumi.StringOutput                     `pulumi:"updatedAt"`
}

// NewPrivateService registers a new resource with the given unique name, arguments, and options.
func NewPrivateService(ctx *pulumi.Context,
	name string, args *PrivateServiceArgs, opts ...pulumi.ResourceOption) (*PrivateService, error) {
	if args == nil {
		args = &PrivateServiceArgs{}
	}

	if args.Type == nil {
		args.Type = pulumi.StringPtr("private_service")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PrivateService
	err := ctx.RegisterResource("render:services:PrivateService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateService gets an existing PrivateService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateServiceState, opts ...pulumi.ResourceOption) (*PrivateService, error) {
	var resource PrivateService
	err := ctx.ReadResource("render:services:PrivateService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateService resources.
type privateServiceState struct {
}

type PrivateServiceState struct {
}

func (PrivateServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateServiceState)(nil)).Elem()
}

type privateServiceArgs struct {
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a PrivateService resource.
type PrivateServiceArgs struct {
	Type pulumi.StringPtrInput
}

func (PrivateServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateServiceArgs)(nil)).Elem()
}

type PrivateServiceInput interface {
	pulumi.Input

	ToPrivateServiceOutput() PrivateServiceOutput
	ToPrivateServiceOutputWithContext(ctx context.Context) PrivateServiceOutput
}

func (*PrivateService) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateService)(nil)).Elem()
}

func (i *PrivateService) ToPrivateServiceOutput() PrivateServiceOutput {
	return i.ToPrivateServiceOutputWithContext(context.Background())
}

func (i *PrivateService) ToPrivateServiceOutputWithContext(ctx context.Context) PrivateServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateServiceOutput)
}

type PrivateServiceOutput struct{ *pulumi.OutputState }

func (PrivateServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateService)(nil)).Elem()
}

func (o PrivateServiceOutput) ToPrivateServiceOutput() PrivateServiceOutput {
	return o
}

func (o PrivateServiceOutput) ToPrivateServiceOutputWithContext(ctx context.Context) PrivateServiceOutput {
	return o
}

func (o PrivateServiceOutput) AutoDeploy() PrivateServiceAutoDeployOutput {
	return o.ApplyT(func(v *PrivateService) PrivateServiceAutoDeployOutput { return v.AutoDeploy }).(PrivateServiceAutoDeployOutput)
}

func (o PrivateServiceOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o PrivateServiceOutput) BuildFilter() BuildFilterPtrOutput {
	return o.ApplyT(func(v *PrivateService) BuildFilterPtrOutput { return v.BuildFilter }).(BuildFilterPtrOutput)
}

func (o PrivateServiceOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o PrivateServiceOutput) ImagePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.ImagePath }).(pulumi.StringPtrOutput)
}

func (o PrivateServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateServiceOutput) NotifyOnFail() PrivateServiceNotifyOnFailOutput {
	return o.ApplyT(func(v *PrivateService) PrivateServiceNotifyOnFailOutput { return v.NotifyOnFail }).(PrivateServiceNotifyOnFailOutput)
}

func (o PrivateServiceOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

func (o PrivateServiceOutput) Repo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.Repo }).(pulumi.StringPtrOutput)
}

func (o PrivateServiceOutput) RootDir() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringOutput { return v.RootDir }).(pulumi.StringOutput)
}

func (o PrivateServiceOutput) ServiceDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.AnyOutput { return v.ServiceDetails }).(pulumi.AnyOutput)
}

func (o PrivateServiceOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

func (o PrivateServiceOutput) Suspended() PrivateServiceSuspendedOutput {
	return o.ApplyT(func(v *PrivateService) PrivateServiceSuspendedOutput { return v.Suspended }).(PrivateServiceSuspendedOutput)
}

func (o PrivateServiceOutput) Suspenders() PrivateServiceSuspendersItemArrayOutput {
	return o.ApplyT(func(v *PrivateService) PrivateServiceSuspendersItemArrayOutput { return v.Suspenders }).(PrivateServiceSuspendersItemArrayOutput)
}

func (o PrivateServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PrivateServiceOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateServiceInput)(nil)).Elem(), &PrivateService{})
	pulumi.RegisterOutputType(PrivateServiceOutput{})
}
