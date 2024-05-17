// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CronJob struct {
	pulumi.CustomResourceState

	AutoDeploy     CronJobServiceAutoDeployPtrOutput       `pulumi:"autoDeploy"`
	Branch         pulumi.StringPtrOutput                  `pulumi:"branch"`
	BuildFilter    BuildFilterPtrOutput                    `pulumi:"buildFilter"`
	CreatedAt      pulumi.StringPtrOutput                  `pulumi:"createdAt"`
	ImagePath      pulumi.StringPtrOutput                  `pulumi:"imagePath"`
	Name           pulumi.StringPtrOutput                  `pulumi:"name"`
	NotifyOnFail   CronJobServiceNotifyOnFailPtrOutput     `pulumi:"notifyOnFail"`
	OwnerId        pulumi.StringPtrOutput                  `pulumi:"ownerId"`
	Repo           pulumi.StringPtrOutput                  `pulumi:"repo"`
	RootDir        pulumi.StringPtrOutput                  `pulumi:"rootDir"`
	ServiceDetails CronJobDetailsCreatePtrOutput           `pulumi:"serviceDetails"`
	Slug           pulumi.StringPtrOutput                  `pulumi:"slug"`
	Suspended      CronJobServiceSuspendedPtrOutput        `pulumi:"suspended"`
	Suspenders     CronJobServiceSuspendersItemArrayOutput `pulumi:"suspenders"`
	Type           pulumi.StringPtrOutput                  `pulumi:"type"`
	UpdatedAt      pulumi.StringPtrOutput                  `pulumi:"updatedAt"`
}

// NewCronJob registers a new resource with the given unique name, arguments, and options.
func NewCronJob(ctx *pulumi.Context,
	name string, args *CronJobArgs, opts ...pulumi.ResourceOption) (*CronJob, error) {
	if args == nil {
		args = &CronJobArgs{}
	}

	if args.Type == nil {
		args.Type = pulumi.StringPtr("cron_job")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CronJob
	err := ctx.RegisterResource("render:services:CronJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCronJob gets an existing CronJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCronJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CronJobState, opts ...pulumi.ResourceOption) (*CronJob, error) {
	var resource CronJob
	err := ctx.ReadResource("render:services:CronJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CronJob resources.
type cronJobState struct {
}

type CronJobState struct {
}

func (CronJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*cronJobState)(nil)).Elem()
}

type cronJobArgs struct {
	ServiceDetails *CronJobDetailsCreate `pulumi:"serviceDetails"`
	Type           *string               `pulumi:"type"`
}

// The set of arguments for constructing a CronJob resource.
type CronJobArgs struct {
	ServiceDetails CronJobDetailsCreatePtrInput
	Type           pulumi.StringPtrInput
}

func (CronJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cronJobArgs)(nil)).Elem()
}

type CronJobInput interface {
	pulumi.Input

	ToCronJobOutput() CronJobOutput
	ToCronJobOutputWithContext(ctx context.Context) CronJobOutput
}

func (*CronJob) ElementType() reflect.Type {
	return reflect.TypeOf((**CronJob)(nil)).Elem()
}

func (i *CronJob) ToCronJobOutput() CronJobOutput {
	return i.ToCronJobOutputWithContext(context.Background())
}

func (i *CronJob) ToCronJobOutputWithContext(ctx context.Context) CronJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobOutput)
}

type CronJobOutput struct{ *pulumi.OutputState }

func (CronJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CronJob)(nil)).Elem()
}

func (o CronJobOutput) ToCronJobOutput() CronJobOutput {
	return o
}

func (o CronJobOutput) ToCronJobOutputWithContext(ctx context.Context) CronJobOutput {
	return o
}

func (o CronJobOutput) AutoDeploy() CronJobServiceAutoDeployPtrOutput {
	return o.ApplyT(func(v *CronJob) CronJobServiceAutoDeployPtrOutput { return v.AutoDeploy }).(CronJobServiceAutoDeployPtrOutput)
}

func (o CronJobOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringPtrOutput { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o CronJobOutput) BuildFilter() BuildFilterPtrOutput {
	return o.ApplyT(func(v *CronJob) BuildFilterPtrOutput { return v.BuildFilter }).(BuildFilterPtrOutput)
}

func (o CronJobOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o CronJobOutput) ImagePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringPtrOutput { return v.ImagePath }).(pulumi.StringPtrOutput)
}

func (o CronJobOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CronJobOutput) NotifyOnFail() CronJobServiceNotifyOnFailPtrOutput {
	return o.ApplyT(func(v *CronJob) CronJobServiceNotifyOnFailPtrOutput { return v.NotifyOnFail }).(CronJobServiceNotifyOnFailPtrOutput)
}

func (o CronJobOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringPtrOutput { return v.OwnerId }).(pulumi.StringPtrOutput)
}

func (o CronJobOutput) Repo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringPtrOutput { return v.Repo }).(pulumi.StringPtrOutput)
}

func (o CronJobOutput) RootDir() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringPtrOutput { return v.RootDir }).(pulumi.StringPtrOutput)
}

func (o CronJobOutput) ServiceDetails() CronJobDetailsCreatePtrOutput {
	return o.ApplyT(func(v *CronJob) CronJobDetailsCreatePtrOutput { return v.ServiceDetails }).(CronJobDetailsCreatePtrOutput)
}

func (o CronJobOutput) Slug() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringPtrOutput { return v.Slug }).(pulumi.StringPtrOutput)
}

func (o CronJobOutput) Suspended() CronJobServiceSuspendedPtrOutput {
	return o.ApplyT(func(v *CronJob) CronJobServiceSuspendedPtrOutput { return v.Suspended }).(CronJobServiceSuspendedPtrOutput)
}

func (o CronJobOutput) Suspenders() CronJobServiceSuspendersItemArrayOutput {
	return o.ApplyT(func(v *CronJob) CronJobServiceSuspendersItemArrayOutput { return v.Suspenders }).(CronJobServiceSuspendersItemArrayOutput)
}

func (o CronJobOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o CronJobOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CronJob) pulumi.StringPtrOutput { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CronJobInput)(nil)).Elem(), &CronJob{})
	pulumi.RegisterOutputType(CronJobOutput{})
}
