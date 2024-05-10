// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"errors"
	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RollbackDeploy struct {
	pulumi.CustomResourceState

	Commit    CommitPropertiesPtrOutput `pulumi:"commit"`
	CreatedAt pulumi.StringPtrOutput    `pulumi:"createdAt"`
	// The ID of the deploy to rollback to
	DeployId   pulumi.StringOutput    `pulumi:"deployId"`
	FinishedAt pulumi.StringPtrOutput `pulumi:"finishedAt"`
	// Image information used when creating the deploy. Not present for Git-backed deploys
	Image     ImagePropertiesPtrOutput `pulumi:"image"`
	Status    StatusPtrOutput          `pulumi:"status"`
	Trigger   TriggerPtrOutput         `pulumi:"trigger"`
	UpdatedAt pulumi.StringPtrOutput   `pulumi:"updatedAt"`
}

// NewRollbackDeploy registers a new resource with the given unique name, arguments, and options.
func NewRollbackDeploy(ctx *pulumi.Context,
	name string, args *RollbackDeployArgs, opts ...pulumi.ResourceOption) (*RollbackDeploy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeployId == nil {
		return nil, errors.New("invalid value for required argument 'DeployId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RollbackDeploy
	err := ctx.RegisterResource("render:services:RollbackDeploy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRollbackDeploy gets an existing RollbackDeploy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRollbackDeploy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RollbackDeployState, opts ...pulumi.ResourceOption) (*RollbackDeploy, error) {
	var resource RollbackDeploy
	err := ctx.ReadResource("render:services:RollbackDeploy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RollbackDeploy resources.
type rollbackDeployState struct {
}

type RollbackDeployState struct {
}

func (RollbackDeployState) ElementType() reflect.Type {
	return reflect.TypeOf((*rollbackDeployState)(nil)).Elem()
}

type rollbackDeployArgs struct {
	// The ID of the deploy to rollback to
	DeployId string `pulumi:"deployId"`
}

// The set of arguments for constructing a RollbackDeploy resource.
type RollbackDeployArgs struct {
	// The ID of the deploy to rollback to
	DeployId pulumi.StringInput
}

func (RollbackDeployArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rollbackDeployArgs)(nil)).Elem()
}

type RollbackDeployInput interface {
	pulumi.Input

	ToRollbackDeployOutput() RollbackDeployOutput
	ToRollbackDeployOutputWithContext(ctx context.Context) RollbackDeployOutput
}

func (*RollbackDeploy) ElementType() reflect.Type {
	return reflect.TypeOf((**RollbackDeploy)(nil)).Elem()
}

func (i *RollbackDeploy) ToRollbackDeployOutput() RollbackDeployOutput {
	return i.ToRollbackDeployOutputWithContext(context.Background())
}

func (i *RollbackDeploy) ToRollbackDeployOutputWithContext(ctx context.Context) RollbackDeployOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RollbackDeployOutput)
}

type RollbackDeployOutput struct{ *pulumi.OutputState }

func (RollbackDeployOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RollbackDeploy)(nil)).Elem()
}

func (o RollbackDeployOutput) ToRollbackDeployOutput() RollbackDeployOutput {
	return o
}

func (o RollbackDeployOutput) ToRollbackDeployOutputWithContext(ctx context.Context) RollbackDeployOutput {
	return o
}

func (o RollbackDeployOutput) Commit() CommitPropertiesPtrOutput {
	return o.ApplyT(func(v *RollbackDeploy) CommitPropertiesPtrOutput { return v.Commit }).(CommitPropertiesPtrOutput)
}

func (o RollbackDeployOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollbackDeploy) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The ID of the deploy to rollback to
func (o RollbackDeployOutput) DeployId() pulumi.StringOutput {
	return o.ApplyT(func(v *RollbackDeploy) pulumi.StringOutput { return v.DeployId }).(pulumi.StringOutput)
}

func (o RollbackDeployOutput) FinishedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollbackDeploy) pulumi.StringPtrOutput { return v.FinishedAt }).(pulumi.StringPtrOutput)
}

// Image information used when creating the deploy. Not present for Git-backed deploys
func (o RollbackDeployOutput) Image() ImagePropertiesPtrOutput {
	return o.ApplyT(func(v *RollbackDeploy) ImagePropertiesPtrOutput { return v.Image }).(ImagePropertiesPtrOutput)
}

func (o RollbackDeployOutput) Status() StatusPtrOutput {
	return o.ApplyT(func(v *RollbackDeploy) StatusPtrOutput { return v.Status }).(StatusPtrOutput)
}

func (o RollbackDeployOutput) Trigger() TriggerPtrOutput {
	return o.ApplyT(func(v *RollbackDeploy) TriggerPtrOutput { return v.Trigger }).(TriggerPtrOutput)
}

func (o RollbackDeployOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollbackDeploy) pulumi.StringPtrOutput { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RollbackDeployInput)(nil)).Elem(), &RollbackDeploy{})
	pulumi.RegisterOutputType(RollbackDeployOutput{})
}
