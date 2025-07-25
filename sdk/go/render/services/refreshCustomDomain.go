// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RefreshCustomDomain struct {
	pulumi.CustomResourceState
}

// NewRefreshCustomDomain registers a new resource with the given unique name, arguments, and options.
func NewRefreshCustomDomain(ctx *pulumi.Context,
	name string, args *RefreshCustomDomainArgs, opts ...pulumi.ResourceOption) (*RefreshCustomDomain, error) {
	if args == nil {
		args = &RefreshCustomDomainArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RefreshCustomDomain
	err := ctx.RegisterResource("render:services:RefreshCustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRefreshCustomDomain gets an existing RefreshCustomDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRefreshCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RefreshCustomDomainState, opts ...pulumi.ResourceOption) (*RefreshCustomDomain, error) {
	var resource RefreshCustomDomain
	err := ctx.ReadResource("render:services:RefreshCustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RefreshCustomDomain resources.
type refreshCustomDomainState struct {
}

type RefreshCustomDomainState struct {
}

func (RefreshCustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*refreshCustomDomainState)(nil)).Elem()
}

type refreshCustomDomainArgs struct {
	// The ID or name of the custom domain
	CustomDomainIdOrName *string `pulumi:"customDomainIdOrName"`
	// The ID of the service
	ServiceId *string `pulumi:"serviceId"`
}

// The set of arguments for constructing a RefreshCustomDomain resource.
type RefreshCustomDomainArgs struct {
	// The ID or name of the custom domain
	CustomDomainIdOrName pulumi.StringPtrInput
	// The ID of the service
	ServiceId pulumi.StringPtrInput
}

func (RefreshCustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*refreshCustomDomainArgs)(nil)).Elem()
}

type RefreshCustomDomainInput interface {
	pulumi.Input

	ToRefreshCustomDomainOutput() RefreshCustomDomainOutput
	ToRefreshCustomDomainOutputWithContext(ctx context.Context) RefreshCustomDomainOutput
}

func (*RefreshCustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**RefreshCustomDomain)(nil)).Elem()
}

func (i *RefreshCustomDomain) ToRefreshCustomDomainOutput() RefreshCustomDomainOutput {
	return i.ToRefreshCustomDomainOutputWithContext(context.Background())
}

func (i *RefreshCustomDomain) ToRefreshCustomDomainOutputWithContext(ctx context.Context) RefreshCustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RefreshCustomDomainOutput)
}

type RefreshCustomDomainOutput struct{ *pulumi.OutputState }

func (RefreshCustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RefreshCustomDomain)(nil)).Elem()
}

func (o RefreshCustomDomainOutput) ToRefreshCustomDomainOutput() RefreshCustomDomainOutput {
	return o
}

func (o RefreshCustomDomainOutput) ToRefreshCustomDomainOutputWithContext(ctx context.Context) RefreshCustomDomainOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RefreshCustomDomainInput)(nil)).Elem(), &RefreshCustomDomain{})
	pulumi.RegisterOutputType(RefreshCustomDomainOutput{})
}
