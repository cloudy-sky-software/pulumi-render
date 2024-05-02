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

// A static website service
//
// ## Example Usage
type ServicesStaticSite struct {
	pulumi.CustomResourceState

	// Whether to auto deploy the service or not upon git push.
	AutoDeploy ServiceAutoDeployPtrOutput `pulumi:"autoDeploy"`
	// If left empty, this will fall back to the default branch of the repository.
	Branch    pulumi.StringPtrOutput    `pulumi:"branch"`
	CreatedAt pulumi.StringPtrOutput    `pulumi:"createdAt"`
	EnvVars   EnvVarKeyValueArrayOutput `pulumi:"envVars"`
	Name      pulumi.StringPtrOutput    `pulumi:"name"`
	// The notification setting for this service upon deployment failure.
	NotifyOnFail ServiceNotifyOnFailPtrOutput `pulumi:"notifyOnFail"`
	// The id of the owner (user/team).
	OwnerId pulumi.StringPtrOutput `pulumi:"ownerId"`
	// Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
	Repo           pulumi.StringPtrOutput            `pulumi:"repo"`
	SecretFiles    SecretFileArrayOutput             `pulumi:"secretFiles"`
	ServiceDetails StaticSiteServiceDetailsPtrOutput `pulumi:"serviceDetails"`
	Slug           pulumi.StringPtrOutput            `pulumi:"slug"`
	Suspended      ServiceSuspendedPtrOutput         `pulumi:"suspended"`
	Suspenders     pulumi.StringArrayOutput          `pulumi:"suspenders"`
	Type           pulumi.StringPtrOutput            `pulumi:"type"`
	UpdatedAt      pulumi.StringPtrOutput            `pulumi:"updatedAt"`
}

// NewServicesStaticSite registers a new resource with the given unique name, arguments, and options.
func NewServicesStaticSite(ctx *pulumi.Context,
	name string, args *ServicesStaticSiteArgs, opts ...pulumi.ResourceOption) (*ServicesStaticSite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.OwnerId == nil {
		return nil, errors.New("invalid value for required argument 'OwnerId'")
	}
	if args.Repo == nil {
		return nil, errors.New("invalid value for required argument 'Repo'")
	}
	if args.AutoDeploy == nil {
		args.AutoDeploy = ServiceAutoDeploy("no")
	}
	if args.ServiceDetails != nil {
		args.ServiceDetails = args.ServiceDetails.ToStaticSiteServiceDetailsPtrOutput().ApplyT(func(v *StaticSiteServiceDetails) *StaticSiteServiceDetails { return v.Defaults() }).(StaticSiteServiceDetailsPtrOutput)
	}
	if args.Type == nil {
		args.Type = pulumi.StringPtr("static_site")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServicesStaticSite
	err := ctx.RegisterResource("render:services:ServicesStaticSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicesStaticSite gets an existing ServicesStaticSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicesStaticSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicesStaticSiteState, opts ...pulumi.ResourceOption) (*ServicesStaticSite, error) {
	var resource ServicesStaticSite
	err := ctx.ReadResource("render:services:ServicesStaticSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicesStaticSite resources.
type servicesStaticSiteState struct {
}

type ServicesStaticSiteState struct {
}

func (ServicesStaticSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicesStaticSiteState)(nil)).Elem()
}

type servicesStaticSiteArgs struct {
	// Whether to auto deploy the service or not upon git push.
	AutoDeploy *ServiceAutoDeploy `pulumi:"autoDeploy"`
	// If left empty, this will fall back to the default branch of the repository.
	Branch    *string          `pulumi:"branch"`
	CreatedAt *string          `pulumi:"createdAt"`
	EnvVars   []EnvVarKeyValue `pulumi:"envVars"`
	Name      string           `pulumi:"name"`
	// The notification setting for this service upon deployment failure.
	NotifyOnFail *ServiceNotifyOnFail `pulumi:"notifyOnFail"`
	// The id of the owner (user/team).
	OwnerId string `pulumi:"ownerId"`
	// Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
	Repo           string                    `pulumi:"repo"`
	SecretFiles    []SecretFile              `pulumi:"secretFiles"`
	ServiceDetails *StaticSiteServiceDetails `pulumi:"serviceDetails"`
	Slug           *string                   `pulumi:"slug"`
	Suspended      *ServiceSuspended         `pulumi:"suspended"`
	Suspenders     []string                  `pulumi:"suspenders"`
	Type           *string                   `pulumi:"type"`
	UpdatedAt      *string                   `pulumi:"updatedAt"`
}

// The set of arguments for constructing a ServicesStaticSite resource.
type ServicesStaticSiteArgs struct {
	// Whether to auto deploy the service or not upon git push.
	AutoDeploy ServiceAutoDeployPtrInput
	// If left empty, this will fall back to the default branch of the repository.
	Branch    pulumi.StringPtrInput
	CreatedAt pulumi.StringPtrInput
	EnvVars   EnvVarKeyValueArrayInput
	Name      pulumi.StringInput
	// The notification setting for this service upon deployment failure.
	NotifyOnFail ServiceNotifyOnFailPtrInput
	// The id of the owner (user/team).
	OwnerId pulumi.StringInput
	// Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
	Repo           pulumi.StringInput
	SecretFiles    SecretFileArrayInput
	ServiceDetails StaticSiteServiceDetailsPtrInput
	Slug           pulumi.StringPtrInput
	Suspended      ServiceSuspendedPtrInput
	Suspenders     pulumi.StringArrayInput
	Type           pulumi.StringPtrInput
	UpdatedAt      pulumi.StringPtrInput
}

func (ServicesStaticSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicesStaticSiteArgs)(nil)).Elem()
}

type ServicesStaticSiteInput interface {
	pulumi.Input

	ToServicesStaticSiteOutput() ServicesStaticSiteOutput
	ToServicesStaticSiteOutputWithContext(ctx context.Context) ServicesStaticSiteOutput
}

func (*ServicesStaticSite) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesStaticSite)(nil)).Elem()
}

func (i *ServicesStaticSite) ToServicesStaticSiteOutput() ServicesStaticSiteOutput {
	return i.ToServicesStaticSiteOutputWithContext(context.Background())
}

func (i *ServicesStaticSite) ToServicesStaticSiteOutputWithContext(ctx context.Context) ServicesStaticSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesStaticSiteOutput)
}

type ServicesStaticSiteOutput struct{ *pulumi.OutputState }

func (ServicesStaticSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesStaticSite)(nil)).Elem()
}

func (o ServicesStaticSiteOutput) ToServicesStaticSiteOutput() ServicesStaticSiteOutput {
	return o
}

func (o ServicesStaticSiteOutput) ToServicesStaticSiteOutputWithContext(ctx context.Context) ServicesStaticSiteOutput {
	return o
}

// Whether to auto deploy the service or not upon git push.
func (o ServicesStaticSiteOutput) AutoDeploy() ServiceAutoDeployPtrOutput {
	return o.ApplyT(func(v *ServicesStaticSite) ServiceAutoDeployPtrOutput { return v.AutoDeploy }).(ServiceAutoDeployPtrOutput)
}

// If left empty, this will fall back to the default branch of the repository.
func (o ServicesStaticSiteOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesStaticSite) pulumi.StringPtrOutput { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o ServicesStaticSiteOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesStaticSite) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o ServicesStaticSiteOutput) EnvVars() EnvVarKeyValueArrayOutput {
	return o.ApplyT(func(v *ServicesStaticSite) EnvVarKeyValueArrayOutput { return v.EnvVars }).(EnvVarKeyValueArrayOutput)
}

func (o ServicesStaticSiteOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesStaticSite) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The notification setting for this service upon deployment failure.
func (o ServicesStaticSiteOutput) NotifyOnFail() ServiceNotifyOnFailPtrOutput {
	return o.ApplyT(func(v *ServicesStaticSite) ServiceNotifyOnFailPtrOutput { return v.NotifyOnFail }).(ServiceNotifyOnFailPtrOutput)
}

// The id of the owner (user/team).
func (o ServicesStaticSiteOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesStaticSite) pulumi.StringPtrOutput { return v.OwnerId }).(pulumi.StringPtrOutput)
}

// Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
func (o ServicesStaticSiteOutput) Repo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesStaticSite) pulumi.StringPtrOutput { return v.Repo }).(pulumi.StringPtrOutput)
}

func (o ServicesStaticSiteOutput) SecretFiles() SecretFileArrayOutput {
	return o.ApplyT(func(v *ServicesStaticSite) SecretFileArrayOutput { return v.SecretFiles }).(SecretFileArrayOutput)
}

func (o ServicesStaticSiteOutput) ServiceDetails() StaticSiteServiceDetailsPtrOutput {
	return o.ApplyT(func(v *ServicesStaticSite) StaticSiteServiceDetailsPtrOutput { return v.ServiceDetails }).(StaticSiteServiceDetailsPtrOutput)
}

func (o ServicesStaticSiteOutput) Slug() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesStaticSite) pulumi.StringPtrOutput { return v.Slug }).(pulumi.StringPtrOutput)
}

func (o ServicesStaticSiteOutput) Suspended() ServiceSuspendedPtrOutput {
	return o.ApplyT(func(v *ServicesStaticSite) ServiceSuspendedPtrOutput { return v.Suspended }).(ServiceSuspendedPtrOutput)
}

func (o ServicesStaticSiteOutput) Suspenders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServicesStaticSite) pulumi.StringArrayOutput { return v.Suspenders }).(pulumi.StringArrayOutput)
}

func (o ServicesStaticSiteOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesStaticSite) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ServicesStaticSiteOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesStaticSite) pulumi.StringPtrOutput { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServicesStaticSiteInput)(nil)).Elem(), &ServicesStaticSite{})
	pulumi.RegisterOutputType(ServicesStaticSiteOutput{})
}
