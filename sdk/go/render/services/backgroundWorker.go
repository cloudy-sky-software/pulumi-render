// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A background worker service
type BackgroundWorker struct {
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
	Repo           pulumi.StringPtrOutput                  `pulumi:"repo"`
	SecretFiles    SecretFileArrayOutput                   `pulumi:"secretFiles"`
	ServiceDetails BackgroundWorkerServiceDetailsPtrOutput `pulumi:"serviceDetails"`
	Slug           pulumi.StringPtrOutput                  `pulumi:"slug"`
	Suspended      ServiceSuspendedPtrOutput               `pulumi:"suspended"`
	Suspenders     pulumi.StringArrayOutput                `pulumi:"suspenders"`
	Type           pulumi.StringPtrOutput                  `pulumi:"type"`
	UpdatedAt      pulumi.StringPtrOutput                  `pulumi:"updatedAt"`
}

// NewBackgroundWorker registers a new resource with the given unique name, arguments, and options.
func NewBackgroundWorker(ctx *pulumi.Context,
	name string, args *BackgroundWorkerArgs, opts ...pulumi.ResourceOption) (*BackgroundWorker, error) {
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
		args.ServiceDetails = args.ServiceDetails.ToBackgroundWorkerServiceDetailsPtrOutput().ApplyT(func(v *BackgroundWorkerServiceDetails) *BackgroundWorkerServiceDetails { return v.Defaults() }).(BackgroundWorkerServiceDetailsPtrOutput)
	}
	if isZero(args.Type) {
		args.Type = pulumi.StringPtr("background_worker")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BackgroundWorker
	err := ctx.RegisterResource("render:services:BackgroundWorker", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackgroundWorker gets an existing BackgroundWorker resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackgroundWorker(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackgroundWorkerState, opts ...pulumi.ResourceOption) (*BackgroundWorker, error) {
	var resource BackgroundWorker
	err := ctx.ReadResource("render:services:BackgroundWorker", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackgroundWorker resources.
type backgroundWorkerState struct {
}

type BackgroundWorkerState struct {
}

func (BackgroundWorkerState) ElementType() reflect.Type {
	return reflect.TypeOf((*backgroundWorkerState)(nil)).Elem()
}

type backgroundWorkerArgs struct {
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
	Repo           string                          `pulumi:"repo"`
	SecretFiles    []SecretFile                    `pulumi:"secretFiles"`
	ServiceDetails *BackgroundWorkerServiceDetails `pulumi:"serviceDetails"`
	Slug           *string                         `pulumi:"slug"`
	Suspended      *ServiceSuspended               `pulumi:"suspended"`
	Suspenders     []string                        `pulumi:"suspenders"`
	Type           *string                         `pulumi:"type"`
	UpdatedAt      *string                         `pulumi:"updatedAt"`
}

// The set of arguments for constructing a BackgroundWorker resource.
type BackgroundWorkerArgs struct {
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
	ServiceDetails BackgroundWorkerServiceDetailsPtrInput
	Slug           pulumi.StringPtrInput
	Suspended      ServiceSuspendedPtrInput
	Suspenders     pulumi.StringArrayInput
	Type           pulumi.StringPtrInput
	UpdatedAt      pulumi.StringPtrInput
}

func (BackgroundWorkerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backgroundWorkerArgs)(nil)).Elem()
}

type BackgroundWorkerInput interface {
	pulumi.Input

	ToBackgroundWorkerOutput() BackgroundWorkerOutput
	ToBackgroundWorkerOutputWithContext(ctx context.Context) BackgroundWorkerOutput
}

func (*BackgroundWorker) ElementType() reflect.Type {
	return reflect.TypeOf((**BackgroundWorker)(nil)).Elem()
}

func (i *BackgroundWorker) ToBackgroundWorkerOutput() BackgroundWorkerOutput {
	return i.ToBackgroundWorkerOutputWithContext(context.Background())
}

func (i *BackgroundWorker) ToBackgroundWorkerOutputWithContext(ctx context.Context) BackgroundWorkerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackgroundWorkerOutput)
}

type BackgroundWorkerOutput struct{ *pulumi.OutputState }

func (BackgroundWorkerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackgroundWorker)(nil)).Elem()
}

func (o BackgroundWorkerOutput) ToBackgroundWorkerOutput() BackgroundWorkerOutput {
	return o
}

func (o BackgroundWorkerOutput) ToBackgroundWorkerOutputWithContext(ctx context.Context) BackgroundWorkerOutput {
	return o
}

// Whether to auto deploy the service or not upon git push.
func (o BackgroundWorkerOutput) AutoDeploy() ServiceAutoDeployPtrOutput {
	return o.ApplyT(func(v *BackgroundWorker) ServiceAutoDeployPtrOutput { return v.AutoDeploy }).(ServiceAutoDeployPtrOutput)
}

// If left empty, this will fall back to the default branch of the repository.
func (o BackgroundWorkerOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackgroundWorker) pulumi.StringPtrOutput { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o BackgroundWorkerOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackgroundWorker) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o BackgroundWorkerOutput) EnvVars() EnvVarKeyValueOrGenerateValueArrayOutput {
	return o.ApplyT(func(v *BackgroundWorker) EnvVarKeyValueOrGenerateValueArrayOutput { return v.EnvVars }).(EnvVarKeyValueOrGenerateValueArrayOutput)
}

func (o BackgroundWorkerOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackgroundWorker) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The notification setting for this service upon deployment failure.
func (o BackgroundWorkerOutput) NotifyOnFail() ServiceNotifyOnFailPtrOutput {
	return o.ApplyT(func(v *BackgroundWorker) ServiceNotifyOnFailPtrOutput { return v.NotifyOnFail }).(ServiceNotifyOnFailPtrOutput)
}

// The id of the owner (user/team).
func (o BackgroundWorkerOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackgroundWorker) pulumi.StringPtrOutput { return v.OwnerId }).(pulumi.StringPtrOutput)
}

// Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
func (o BackgroundWorkerOutput) Repo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackgroundWorker) pulumi.StringPtrOutput { return v.Repo }).(pulumi.StringPtrOutput)
}

func (o BackgroundWorkerOutput) SecretFiles() SecretFileArrayOutput {
	return o.ApplyT(func(v *BackgroundWorker) SecretFileArrayOutput { return v.SecretFiles }).(SecretFileArrayOutput)
}

func (o BackgroundWorkerOutput) ServiceDetails() BackgroundWorkerServiceDetailsPtrOutput {
	return o.ApplyT(func(v *BackgroundWorker) BackgroundWorkerServiceDetailsPtrOutput { return v.ServiceDetails }).(BackgroundWorkerServiceDetailsPtrOutput)
}

func (o BackgroundWorkerOutput) Slug() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackgroundWorker) pulumi.StringPtrOutput { return v.Slug }).(pulumi.StringPtrOutput)
}

func (o BackgroundWorkerOutput) Suspended() ServiceSuspendedPtrOutput {
	return o.ApplyT(func(v *BackgroundWorker) ServiceSuspendedPtrOutput { return v.Suspended }).(ServiceSuspendedPtrOutput)
}

func (o BackgroundWorkerOutput) Suspenders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackgroundWorker) pulumi.StringArrayOutput { return v.Suspenders }).(pulumi.StringArrayOutput)
}

func (o BackgroundWorkerOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackgroundWorker) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o BackgroundWorkerOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackgroundWorker) pulumi.StringPtrOutput { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackgroundWorkerInput)(nil)).Elem(), &BackgroundWorker{})
	pulumi.RegisterOutputType(BackgroundWorkerOutput{})
}
