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

type PrivateService struct {
	pulumi.CustomResourceState

	AutoDeploy  ServiceAutoDeployPtrOutput `pulumi:"autoDeploy"`
	Branch      pulumi.StringPtrOutput     `pulumi:"branch"`
	BuildFilter BuildFilterPtrOutput       `pulumi:"buildFilter"`
	CreatedAt   pulumi.StringPtrOutput     `pulumi:"createdAt"`
	// The URL to view the service in the Render Dashboard
	DashboardUrl       pulumi.StringPtrOutput               `pulumi:"dashboardUrl"`
	EnvVars            EnvVarInputTypeArrayOutput           `pulumi:"envVars"`
	EnvironmentId      pulumi.StringPtrOutput               `pulumi:"environmentId"`
	Image              ImagePtrOutput                       `pulumi:"image"`
	ImagePath          pulumi.StringPtrOutput               `pulumi:"imagePath"`
	Name               pulumi.StringPtrOutput               `pulumi:"name"`
	NotifyOnFail       ServiceNotifyOnFailPtrOutput         `pulumi:"notifyOnFail"`
	OwnerId            pulumi.StringPtrOutput               `pulumi:"ownerId"`
	RegistryCredential RegistryCredentialSummaryPtrOutput   `pulumi:"registryCredential"`
	Repo               pulumi.StringPtrOutput               `pulumi:"repo"`
	RootDir            pulumi.StringPtrOutput               `pulumi:"rootDir"`
	SecretFiles        SecretFileInputTypeArrayOutput       `pulumi:"secretFiles"`
	ServiceDetails     PrivateServiceDetailsOutputPtrOutput `pulumi:"serviceDetails"`
	Slug               pulumi.StringPtrOutput               `pulumi:"slug"`
	Suspended          ServiceSuspendedPtrOutput            `pulumi:"suspended"`
	Suspenders         ServiceSuspendersItemArrayOutput     `pulumi:"suspenders"`
	Type               pulumi.StringPtrOutput               `pulumi:"type"`
	UpdatedAt          pulumi.StringPtrOutput               `pulumi:"updatedAt"`
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
	if args.AutoDeploy == nil {
		args.AutoDeploy = ServiceCreateAutoDeploy("yes")
	}
	if args.ServiceDetails != nil {
		args.ServiceDetails = args.ServiceDetails.ToPrivateServiceDetailsCreatePtrOutput().ApplyT(func(v *PrivateServiceDetailsCreate) *PrivateServiceDetailsCreate { return v.Defaults() }).(PrivateServiceDetailsCreatePtrOutput)
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
	AutoDeploy *ServiceCreateAutoDeploy `pulumi:"autoDeploy"`
	// If left empty, this will fall back to the default branch of the repository
	Branch      *string           `pulumi:"branch"`
	BuildFilter *BuildFilter      `pulumi:"buildFilter"`
	EnvVars     []EnvVarInputType `pulumi:"envVars"`
	Image       *Image            `pulumi:"image"`
	Name        string            `pulumi:"name"`
	OwnerId     string            `pulumi:"ownerId"`
	// Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
	Repo           *string                      `pulumi:"repo"`
	RootDir        *string                      `pulumi:"rootDir"`
	SecretFiles    []SecretFileInputType        `pulumi:"secretFiles"`
	ServiceDetails *PrivateServiceDetailsCreate `pulumi:"serviceDetails"`
	Type           *string                      `pulumi:"type"`
}

// The set of arguments for constructing a PrivateService resource.
type PrivateServiceArgs struct {
	AutoDeploy ServiceCreateAutoDeployPtrInput
	// If left empty, this will fall back to the default branch of the repository
	Branch      pulumi.StringPtrInput
	BuildFilter BuildFilterPtrInput
	EnvVars     EnvVarInputTypeArrayInput
	Image       ImagePtrInput
	Name        pulumi.StringInput
	OwnerId     pulumi.StringInput
	// Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
	Repo           pulumi.StringPtrInput
	RootDir        pulumi.StringPtrInput
	SecretFiles    SecretFileInputTypeArrayInput
	ServiceDetails PrivateServiceDetailsCreatePtrInput
	Type           pulumi.StringPtrInput
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

func (o PrivateServiceOutput) AutoDeploy() ServiceAutoDeployPtrOutput {
	return o.ApplyT(func(v *PrivateService) ServiceAutoDeployPtrOutput { return v.AutoDeploy }).(ServiceAutoDeployPtrOutput)
}

func (o PrivateServiceOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o PrivateServiceOutput) BuildFilter() BuildFilterPtrOutput {
	return o.ApplyT(func(v *PrivateService) BuildFilterPtrOutput { return v.BuildFilter }).(BuildFilterPtrOutput)
}

func (o PrivateServiceOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The URL to view the service in the Render Dashboard
func (o PrivateServiceOutput) DashboardUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.DashboardUrl }).(pulumi.StringPtrOutput)
}

func (o PrivateServiceOutput) EnvVars() EnvVarInputTypeArrayOutput {
	return o.ApplyT(func(v *PrivateService) EnvVarInputTypeArrayOutput { return v.EnvVars }).(EnvVarInputTypeArrayOutput)
}

func (o PrivateServiceOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

func (o PrivateServiceOutput) Image() ImagePtrOutput {
	return o.ApplyT(func(v *PrivateService) ImagePtrOutput { return v.Image }).(ImagePtrOutput)
}

func (o PrivateServiceOutput) ImagePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.ImagePath }).(pulumi.StringPtrOutput)
}

func (o PrivateServiceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PrivateServiceOutput) NotifyOnFail() ServiceNotifyOnFailPtrOutput {
	return o.ApplyT(func(v *PrivateService) ServiceNotifyOnFailPtrOutput { return v.NotifyOnFail }).(ServiceNotifyOnFailPtrOutput)
}

func (o PrivateServiceOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.OwnerId }).(pulumi.StringPtrOutput)
}

func (o PrivateServiceOutput) RegistryCredential() RegistryCredentialSummaryPtrOutput {
	return o.ApplyT(func(v *PrivateService) RegistryCredentialSummaryPtrOutput { return v.RegistryCredential }).(RegistryCredentialSummaryPtrOutput)
}

func (o PrivateServiceOutput) Repo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.Repo }).(pulumi.StringPtrOutput)
}

func (o PrivateServiceOutput) RootDir() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.RootDir }).(pulumi.StringPtrOutput)
}

func (o PrivateServiceOutput) SecretFiles() SecretFileInputTypeArrayOutput {
	return o.ApplyT(func(v *PrivateService) SecretFileInputTypeArrayOutput { return v.SecretFiles }).(SecretFileInputTypeArrayOutput)
}

func (o PrivateServiceOutput) ServiceDetails() PrivateServiceDetailsOutputPtrOutput {
	return o.ApplyT(func(v *PrivateService) PrivateServiceDetailsOutputPtrOutput { return v.ServiceDetails }).(PrivateServiceDetailsOutputPtrOutput)
}

func (o PrivateServiceOutput) Slug() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateService) pulumi.StringPtrOutput { return v.Slug }).(pulumi.StringPtrOutput)
}

func (o PrivateServiceOutput) Suspended() ServiceSuspendedPtrOutput {
	return o.ApplyT(func(v *PrivateService) ServiceSuspendedPtrOutput { return v.Suspended }).(ServiceSuspendedPtrOutput)
}

func (o PrivateServiceOutput) Suspenders() ServiceSuspendersItemArrayOutput {
	return o.ApplyT(func(v *PrivateService) ServiceSuspendersItemArrayOutput { return v.Suspenders }).(ServiceSuspendersItemArrayOutput)
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
