// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package envgroups

import (
	"context"
	"reflect"

	"errors"
	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnvGroup struct {
	pulumi.CustomResourceState

	CreatedAt     pulumi.StringPtrOutput   `pulumi:"createdAt"`
	EnvVars       EnvVarArrayOutput        `pulumi:"envVars"`
	EnvironmentId pulumi.StringPtrOutput   `pulumi:"environmentId"`
	Name          pulumi.StringOutput      `pulumi:"name"`
	OwnerId       pulumi.StringOutput      `pulumi:"ownerId"`
	SecretFiles   SecretFileArrayOutput    `pulumi:"secretFiles"`
	ServiceIds    pulumi.StringArrayOutput `pulumi:"serviceIds"`
	// List of serviceIds linked to the envGroup
	ServiceLinks EnvGroupLinkArrayOutput `pulumi:"serviceLinks"`
	UpdatedAt    pulumi.StringPtrOutput  `pulumi:"updatedAt"`
}

// NewEnvGroup registers a new resource with the given unique name, arguments, and options.
func NewEnvGroup(ctx *pulumi.Context,
	name string, args *EnvGroupArgs, opts ...pulumi.ResourceOption) (*EnvGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvVars == nil {
		return nil, errors.New("invalid value for required argument 'EnvVars'")
	}
	if args.OwnerId == nil {
		return nil, errors.New("invalid value for required argument 'OwnerId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EnvGroup
	err := ctx.RegisterResource("render:env-groups:EnvGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvGroup gets an existing EnvGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvGroupState, opts ...pulumi.ResourceOption) (*EnvGroup, error) {
	var resource EnvGroup
	err := ctx.ReadResource("render:env-groups:EnvGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvGroup resources.
type envGroupState struct {
}

type EnvGroupState struct {
}

func (EnvGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*envGroupState)(nil)).Elem()
}

type envGroupArgs struct {
	EnvVars       []EnvVarInputType     `pulumi:"envVars"`
	EnvironmentId *string               `pulumi:"environmentId"`
	Name          *string               `pulumi:"name"`
	OwnerId       string                `pulumi:"ownerId"`
	SecretFiles   []SecretFileInputType `pulumi:"secretFiles"`
	ServiceIds    []string              `pulumi:"serviceIds"`
}

// The set of arguments for constructing a EnvGroup resource.
type EnvGroupArgs struct {
	EnvVars       EnvVarInputTypeArrayInput
	EnvironmentId pulumi.StringPtrInput
	Name          pulumi.StringPtrInput
	OwnerId       pulumi.StringInput
	SecretFiles   SecretFileInputTypeArrayInput
	ServiceIds    pulumi.StringArrayInput
}

func (EnvGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*envGroupArgs)(nil)).Elem()
}

type EnvGroupInput interface {
	pulumi.Input

	ToEnvGroupOutput() EnvGroupOutput
	ToEnvGroupOutputWithContext(ctx context.Context) EnvGroupOutput
}

func (*EnvGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvGroup)(nil)).Elem()
}

func (i *EnvGroup) ToEnvGroupOutput() EnvGroupOutput {
	return i.ToEnvGroupOutputWithContext(context.Background())
}

func (i *EnvGroup) ToEnvGroupOutputWithContext(ctx context.Context) EnvGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvGroupOutput)
}

type EnvGroupOutput struct{ *pulumi.OutputState }

func (EnvGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvGroup)(nil)).Elem()
}

func (o EnvGroupOutput) ToEnvGroupOutput() EnvGroupOutput {
	return o
}

func (o EnvGroupOutput) ToEnvGroupOutputWithContext(ctx context.Context) EnvGroupOutput {
	return o
}

func (o EnvGroupOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvGroup) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o EnvGroupOutput) EnvVars() EnvVarArrayOutput {
	return o.ApplyT(func(v *EnvGroup) EnvVarArrayOutput { return v.EnvVars }).(EnvVarArrayOutput)
}

func (o EnvGroupOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvGroup) pulumi.StringPtrOutput { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

func (o EnvGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EnvGroupOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvGroup) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

func (o EnvGroupOutput) SecretFiles() SecretFileArrayOutput {
	return o.ApplyT(func(v *EnvGroup) SecretFileArrayOutput { return v.SecretFiles }).(SecretFileArrayOutput)
}

func (o EnvGroupOutput) ServiceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EnvGroup) pulumi.StringArrayOutput { return v.ServiceIds }).(pulumi.StringArrayOutput)
}

// List of serviceIds linked to the envGroup
func (o EnvGroupOutput) ServiceLinks() EnvGroupLinkArrayOutput {
	return o.ApplyT(func(v *EnvGroup) EnvGroupLinkArrayOutput { return v.ServiceLinks }).(EnvGroupLinkArrayOutput)
}

func (o EnvGroupOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvGroup) pulumi.StringPtrOutput { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvGroupInput)(nil)).Elem(), &EnvGroup{})
	pulumi.RegisterOutputType(EnvGroupOutput{})
}
