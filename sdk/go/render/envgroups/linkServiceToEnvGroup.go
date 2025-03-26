// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package envgroups

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LinkServiceToEnvGroup struct {
	pulumi.CustomResourceState

	CreatedAt     pulumi.StringPtrOutput `pulumi:"createdAt"`
	EnvVars       EnvVarArrayOutput      `pulumi:"envVars"`
	EnvironmentId pulumi.StringPtrOutput `pulumi:"environmentId"`
	Name          pulumi.StringPtrOutput `pulumi:"name"`
	OwnerId       pulumi.StringPtrOutput `pulumi:"ownerId"`
	SecretFiles   SecretFileArrayOutput  `pulumi:"secretFiles"`
	// List of serviceIds linked to the envGroup
	ServiceLinks EnvGroupLinkArrayOutput `pulumi:"serviceLinks"`
	UpdatedAt    pulumi.StringPtrOutput  `pulumi:"updatedAt"`
}

// NewLinkServiceToEnvGroup registers a new resource with the given unique name, arguments, and options.
func NewLinkServiceToEnvGroup(ctx *pulumi.Context,
	name string, args *LinkServiceToEnvGroupArgs, opts ...pulumi.ResourceOption) (*LinkServiceToEnvGroup, error) {
	if args == nil {
		args = &LinkServiceToEnvGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LinkServiceToEnvGroup
	err := ctx.RegisterResource("render:env-groups:LinkServiceToEnvGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkServiceToEnvGroup gets an existing LinkServiceToEnvGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkServiceToEnvGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkServiceToEnvGroupState, opts ...pulumi.ResourceOption) (*LinkServiceToEnvGroup, error) {
	var resource LinkServiceToEnvGroup
	err := ctx.ReadResource("render:env-groups:LinkServiceToEnvGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkServiceToEnvGroup resources.
type linkServiceToEnvGroupState struct {
}

type LinkServiceToEnvGroupState struct {
}

func (LinkServiceToEnvGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkServiceToEnvGroupState)(nil)).Elem()
}

type linkServiceToEnvGroupArgs struct {
	// Filter for resources that belong to an environment group
	EnvGroupId *string `pulumi:"envGroupId"`
	// The ID of the service
	ServiceId *string `pulumi:"serviceId"`
}

// The set of arguments for constructing a LinkServiceToEnvGroup resource.
type LinkServiceToEnvGroupArgs struct {
	// Filter for resources that belong to an environment group
	EnvGroupId pulumi.StringPtrInput
	// The ID of the service
	ServiceId pulumi.StringPtrInput
}

func (LinkServiceToEnvGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkServiceToEnvGroupArgs)(nil)).Elem()
}

type LinkServiceToEnvGroupInput interface {
	pulumi.Input

	ToLinkServiceToEnvGroupOutput() LinkServiceToEnvGroupOutput
	ToLinkServiceToEnvGroupOutputWithContext(ctx context.Context) LinkServiceToEnvGroupOutput
}

func (*LinkServiceToEnvGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkServiceToEnvGroup)(nil)).Elem()
}

func (i *LinkServiceToEnvGroup) ToLinkServiceToEnvGroupOutput() LinkServiceToEnvGroupOutput {
	return i.ToLinkServiceToEnvGroupOutputWithContext(context.Background())
}

func (i *LinkServiceToEnvGroup) ToLinkServiceToEnvGroupOutputWithContext(ctx context.Context) LinkServiceToEnvGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkServiceToEnvGroupOutput)
}

type LinkServiceToEnvGroupOutput struct{ *pulumi.OutputState }

func (LinkServiceToEnvGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkServiceToEnvGroup)(nil)).Elem()
}

func (o LinkServiceToEnvGroupOutput) ToLinkServiceToEnvGroupOutput() LinkServiceToEnvGroupOutput {
	return o
}

func (o LinkServiceToEnvGroupOutput) ToLinkServiceToEnvGroupOutputWithContext(ctx context.Context) LinkServiceToEnvGroupOutput {
	return o
}

func (o LinkServiceToEnvGroupOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkServiceToEnvGroup) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LinkServiceToEnvGroupOutput) EnvVars() EnvVarArrayOutput {
	return o.ApplyT(func(v *LinkServiceToEnvGroup) EnvVarArrayOutput { return v.EnvVars }).(EnvVarArrayOutput)
}

func (o LinkServiceToEnvGroupOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkServiceToEnvGroup) pulumi.StringPtrOutput { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

func (o LinkServiceToEnvGroupOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkServiceToEnvGroup) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LinkServiceToEnvGroupOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkServiceToEnvGroup) pulumi.StringPtrOutput { return v.OwnerId }).(pulumi.StringPtrOutput)
}

func (o LinkServiceToEnvGroupOutput) SecretFiles() SecretFileArrayOutput {
	return o.ApplyT(func(v *LinkServiceToEnvGroup) SecretFileArrayOutput { return v.SecretFiles }).(SecretFileArrayOutput)
}

// List of serviceIds linked to the envGroup
func (o LinkServiceToEnvGroupOutput) ServiceLinks() EnvGroupLinkArrayOutput {
	return o.ApplyT(func(v *LinkServiceToEnvGroup) EnvGroupLinkArrayOutput { return v.ServiceLinks }).(EnvGroupLinkArrayOutput)
}

func (o LinkServiceToEnvGroupOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkServiceToEnvGroup) pulumi.StringPtrOutput { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LinkServiceToEnvGroupInput)(nil)).Elem(), &LinkServiceToEnvGroup{})
	pulumi.RegisterOutputType(LinkServiceToEnvGroupOutput{})
}
