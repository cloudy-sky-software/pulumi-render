// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A static website service
//
// ## Example Usage
type StaticSite struct {
	pulumi.CustomResourceState

	// Whether to auto deploy the service or not upon git push.
	AutoDeploy ServiceAutoDeployPtrOutput `pulumi:"autoDeploy"`
	// If left empty, this will fall back to the default branch of the repository.
	Branch    pulumi.StringPtrOutput                   `pulumi:"branch"`
	CreatedAt pulumi.StringPtrOutput                   `pulumi:"createdAt"`
	EnvVars   EnvVarKeyValueOrGenerateValueArrayOutput `pulumi:"envVars"`
	Name      pulumi.StringPtrOutput                   `pulumi:"name"`
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

// NewStaticSite registers a new resource with the given unique name, arguments, and options.
func NewStaticSite(ctx *pulumi.Context,
	name string, args *StaticSiteArgs, opts ...pulumi.ResourceOption) (*StaticSite, error) {
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
	if isZero(args.AutoDeploy) {
		args.AutoDeploy = ServiceAutoDeploy("no")
	}
	if args.ServiceDetails != nil {
		args.ServiceDetails = args.ServiceDetails.ToStaticSiteServiceDetailsPtrOutput().ApplyT(func(v *StaticSiteServiceDetails) *StaticSiteServiceDetails { return v.Defaults() }).(StaticSiteServiceDetailsPtrOutput)
	}
	if isZero(args.Type) {
		args.Type = pulumi.StringPtr("static_site")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource StaticSite
	err := ctx.RegisterResource("render:services:StaticSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStaticSite gets an existing StaticSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStaticSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticSiteState, opts ...pulumi.ResourceOption) (*StaticSite, error) {
	var resource StaticSite
	err := ctx.ReadResource("render:services:StaticSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StaticSite resources.
type staticSiteState struct {
}

type StaticSiteState struct {
}

func (StaticSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteState)(nil)).Elem()
}

type staticSiteArgs struct {
	// Whether to auto deploy the service or not upon git push.
	AutoDeploy *ServiceAutoDeploy `pulumi:"autoDeploy"`
	// If left empty, this will fall back to the default branch of the repository.
	Branch    *string                         `pulumi:"branch"`
	CreatedAt *string                         `pulumi:"createdAt"`
	EnvVars   []EnvVarKeyValueOrGenerateValue `pulumi:"envVars"`
	Name      string                          `pulumi:"name"`
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

// The set of arguments for constructing a StaticSite resource.
type StaticSiteArgs struct {
	// Whether to auto deploy the service or not upon git push.
	AutoDeploy ServiceAutoDeployPtrInput
	// If left empty, this will fall back to the default branch of the repository.
	Branch    pulumi.StringPtrInput
	CreatedAt pulumi.StringPtrInput
	EnvVars   EnvVarKeyValueOrGenerateValueArrayInput
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

func (StaticSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteArgs)(nil)).Elem()
}

type StaticSiteInput interface {
	pulumi.Input

	ToStaticSiteOutput() StaticSiteOutput
	ToStaticSiteOutputWithContext(ctx context.Context) StaticSiteOutput
}

func (*StaticSite) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSite)(nil)).Elem()
}

func (i *StaticSite) ToStaticSiteOutput() StaticSiteOutput {
	return i.ToStaticSiteOutputWithContext(context.Background())
}

func (i *StaticSite) ToStaticSiteOutputWithContext(ctx context.Context) StaticSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteOutput)
}

type StaticSiteOutput struct{ *pulumi.OutputState }

func (StaticSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSite)(nil)).Elem()
}

func (o StaticSiteOutput) ToStaticSiteOutput() StaticSiteOutput {
	return o
}

func (o StaticSiteOutput) ToStaticSiteOutputWithContext(ctx context.Context) StaticSiteOutput {
	return o
}

// Whether to auto deploy the service or not upon git push.
func (o StaticSiteOutput) AutoDeploy() ServiceAutoDeployPtrOutput {
	return o.ApplyT(func(v *StaticSite) ServiceAutoDeployPtrOutput { return v.AutoDeploy }).(ServiceAutoDeployPtrOutput)
}

// If left empty, this will fall back to the default branch of the repository.
func (o StaticSiteOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringPtrOutput { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o StaticSiteOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o StaticSiteOutput) EnvVars() EnvVarKeyValueOrGenerateValueArrayOutput {
	return o.ApplyT(func(v *StaticSite) EnvVarKeyValueOrGenerateValueArrayOutput { return v.EnvVars }).(EnvVarKeyValueOrGenerateValueArrayOutput)
}

func (o StaticSiteOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The notification setting for this service upon deployment failure.
func (o StaticSiteOutput) NotifyOnFail() ServiceNotifyOnFailPtrOutput {
	return o.ApplyT(func(v *StaticSite) ServiceNotifyOnFailPtrOutput { return v.NotifyOnFail }).(ServiceNotifyOnFailPtrOutput)
}

// The id of the owner (user/team).
func (o StaticSiteOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringPtrOutput { return v.OwnerId }).(pulumi.StringPtrOutput)
}

// Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
func (o StaticSiteOutput) Repo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringPtrOutput { return v.Repo }).(pulumi.StringPtrOutput)
}

func (o StaticSiteOutput) SecretFiles() SecretFileArrayOutput {
	return o.ApplyT(func(v *StaticSite) SecretFileArrayOutput { return v.SecretFiles }).(SecretFileArrayOutput)
}

func (o StaticSiteOutput) ServiceDetails() StaticSiteServiceDetailsPtrOutput {
	return o.ApplyT(func(v *StaticSite) StaticSiteServiceDetailsPtrOutput { return v.ServiceDetails }).(StaticSiteServiceDetailsPtrOutput)
}

func (o StaticSiteOutput) Slug() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringPtrOutput { return v.Slug }).(pulumi.StringPtrOutput)
}

func (o StaticSiteOutput) Suspended() ServiceSuspendedPtrOutput {
	return o.ApplyT(func(v *StaticSite) ServiceSuspendedPtrOutput { return v.Suspended }).(ServiceSuspendedPtrOutput)
}

func (o StaticSiteOutput) Suspenders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringArrayOutput { return v.Suspenders }).(pulumi.StringArrayOutput)
}

func (o StaticSiteOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o StaticSiteOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringPtrOutput { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StaticSiteInput)(nil)).Elem(), &StaticSite{})
	pulumi.RegisterOutputType(StaticSiteOutput{})
}
