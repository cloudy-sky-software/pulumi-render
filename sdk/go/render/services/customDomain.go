// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomDomain struct {
	pulumi.CustomResourceState

	CreatedAt          pulumi.StringPtrOutput               `pulumi:"createdAt"`
	DomainType         CustomDomainDomainTypeOutput         `pulumi:"domainType"`
	Name               pulumi.StringOutput                  `pulumi:"name"`
	PublicSuffix       pulumi.StringPtrOutput               `pulumi:"publicSuffix"`
	RedirectForName    pulumi.StringOutput                  `pulumi:"redirectForName"`
	Server             ServerPropertiesOutput               `pulumi:"server"`
	VerificationStatus CustomDomainVerificationStatusOutput `pulumi:"verificationStatus"`
}

// NewCustomDomain registers a new resource with the given unique name, arguments, and options.
func NewCustomDomain(ctx *pulumi.Context,
	name string, args *CustomDomainArgs, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Server == nil {
		return nil, errors.New("invalid value for required argument 'Server'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CustomDomain
	err := ctx.RegisterResource("render:services:CustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomDomain gets an existing CustomDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomDomainState, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	var resource CustomDomain
	err := ctx.ReadResource("render:services:CustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomDomain resources.
type customDomainState struct {
}

type CustomDomainState struct {
}

func (CustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*customDomainState)(nil)).Elem()
}

type customDomainArgs struct {
	// (Required) The ID of the service
	Id     *string          `pulumi:"id"`
	Name   string           `pulumi:"name"`
	Server ServerProperties `pulumi:"server"`
}

// The set of arguments for constructing a CustomDomain resource.
type CustomDomainArgs struct {
	// (Required) The ID of the service
	Id     pulumi.StringPtrInput
	Name   pulumi.StringInput
	Server ServerPropertiesInput
}

func (CustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customDomainArgs)(nil)).Elem()
}

type CustomDomainInput interface {
	pulumi.Input

	ToCustomDomainOutput() CustomDomainOutput
	ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput
}

func (*CustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (i *CustomDomain) ToCustomDomainOutput() CustomDomainOutput {
	return i.ToCustomDomainOutputWithContext(context.Background())
}

func (i *CustomDomain) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput)
}

type CustomDomainOutput struct{ *pulumi.OutputState }

func (CustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (o CustomDomainOutput) ToCustomDomainOutput() CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o CustomDomainOutput) DomainType() CustomDomainDomainTypeOutput {
	return o.ApplyT(func(v *CustomDomain) CustomDomainDomainTypeOutput { return v.DomainType }).(CustomDomainDomainTypeOutput)
}

func (o CustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CustomDomainOutput) PublicSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringPtrOutput { return v.PublicSuffix }).(pulumi.StringPtrOutput)
}

func (o CustomDomainOutput) RedirectForName() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.RedirectForName }).(pulumi.StringOutput)
}

func (o CustomDomainOutput) Server() ServerPropertiesOutput {
	return o.ApplyT(func(v *CustomDomain) ServerPropertiesOutput { return v.Server }).(ServerPropertiesOutput)
}

func (o CustomDomainOutput) VerificationStatus() CustomDomainVerificationStatusOutput {
	return o.ApplyT(func(v *CustomDomain) CustomDomainVerificationStatusOutput { return v.VerificationStatus }).(CustomDomainVerificationStatusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDomainInput)(nil)).Elem(), &CustomDomain{})
	pulumi.RegisterOutputType(CustomDomainOutput{})
}
