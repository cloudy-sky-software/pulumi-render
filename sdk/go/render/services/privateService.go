// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A private service
type PrivateService struct {
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
	Repo           pulumi.StringPtrOutput         `pulumi:"repo"`
	SecretFiles    SecretFileArrayOutput          `pulumi:"secretFiles"`
	ServiceDetails PrivateServiceDetailsPtrOutput `pulumi:"serviceDetails"`
	Slug           pulumi.StringPtrOutput         `pulumi:"slug"`
	Suspended      ServiceSuspendedPtrOutput      `pulumi:"suspended"`
	Suspenders     pulumi.StringArrayOutput       `pulumi:"suspenders"`
	Type           pulumi.StringPtrOutput         `pulumi:"type"`
	UpdatedAt      pulumi.StringPtrOutput         `pulumi:"updatedAt"`
}

// NewPrivateService registers a new resource with the given unique name, arguments, and options.
func NewPrivateService(ctx *pulumi.Context,
	name string, args *PrivateServiceArgs, opts ...pulumi.ResourceOption) (*PrivateService, error) {
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
		args.ServiceDetails = args.ServiceDetails.ToPrivateServiceDetailsPtrOutput().ApplyT(func(v *PrivateServiceDetails) *PrivateServiceDetails { return v.Defaults() }).(PrivateServiceDetailsPtrOutput)
	}
	if isZero(args.Type) {
		args.Type = pulumi.StringPtr("private_service")
	}
	opts = pkgResourceDefaultOpts(opts)
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
	Repo           string                 `pulumi:"repo"`
	SecretFiles    []SecretFile           `pulumi:"secretFiles"`
	ServiceDetails *PrivateServiceDetails `pulumi:"serviceDetails"`
	Slug           *string                `pulumi:"slug"`
	Suspended      *ServiceSuspended      `pulumi:"suspended"`
	Suspenders     []string               `pulumi:"suspenders"`
	Type           *string                `pulumi:"type"`
	UpdatedAt      *string                `pulumi:"updatedAt"`
}

// The set of arguments for constructing a PrivateService resource.
type PrivateServiceArgs struct {
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
	ServiceDetails PrivateServiceDetailsPtrInput
	Slug           pulumi.StringPtrInput
	Suspended      ServiceSuspendedPtrInput
	Suspenders     pulumi.StringArrayInput
	Type           pulumi.StringPtrInput
	UpdatedAt      pulumi.StringPtrInput
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

// Whether to auto deploy the service or not upon git push.
func (o PrivateServiceOutput) AutoDeploy() ServiceAutoDeployPtrOutput {
	return o.ApplyT(func(v *PrivateService) ServiceAutoDeployPtrOutput { return v.AutoDeploy }).(ServiceAutoDeployPtrOutput)
}

// If left empty, this will fall back to the default branch of the repository.
func (o PrivateServiceOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o PrivateServiceOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o PrivateServiceOutput) EnvVars() EnvVarKeyValueArrayOutput {
	return o.ApplyT(func(v *PrivateService) EnvVarKeyValueArrayOutput { return v.EnvVars }).(EnvVarKeyValueArrayOutput)
}

func (o PrivateServiceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The notification setting for this service upon deployment failure.
func (o PrivateServiceOutput) NotifyOnFail() ServiceNotifyOnFailPtrOutput {
	return o.ApplyT(func(v *PrivateService) ServiceNotifyOnFailPtrOutput { return v.NotifyOnFail }).(ServiceNotifyOnFailPtrOutput)
}

// The id of the owner (user/team).
func (o PrivateServiceOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.OwnerId }).(pulumi.StringPtrOutput)
}

// Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
func (o PrivateServiceOutput) Repo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.Repo }).(pulumi.StringPtrOutput)
}

func (o PrivateServiceOutput) SecretFiles() SecretFileArrayOutput {
	return o.ApplyT(func(v *PrivateService) SecretFileArrayOutput { return v.SecretFiles }).(SecretFileArrayOutput)
}

func (o PrivateServiceOutput) ServiceDetails() PrivateServiceDetailsPtrOutput {
	return o.ApplyT(func(v *PrivateService) PrivateServiceDetailsPtrOutput { return v.ServiceDetails }).(PrivateServiceDetailsPtrOutput)
}

func (o PrivateServiceOutput) Slug() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.Slug }).(pulumi.StringPtrOutput)
}

func (o PrivateServiceOutput) Suspended() ServiceSuspendedPtrOutput {
	return o.ApplyT(func(v *PrivateService) ServiceSuspendedPtrOutput { return v.Suspended }).(ServiceSuspendedPtrOutput)
}

func (o PrivateServiceOutput) Suspenders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringArrayOutput { return v.Suspenders }).(pulumi.StringArrayOutput)
}

func (o PrivateServiceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o PrivateServiceOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateServiceInput)(nil)).Elem(), &PrivateService{})
	pulumi.RegisterOutputType(PrivateServiceOutput{})
}
