// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Deploy struct {
	pulumi.CustomResourceState

	// If `clear`, Render clears the service's build cache before deploying. This can be useful if you're experiencing issues with your build.
	ClearCache ClearCachePtrOutput       `pulumi:"clearCache"`
	Commit     CommitPropertiesPtrOutput `pulumi:"commit"`
	// The SHA of a specific Git commit to deploy for a service. Defaults to the latest commit on the service's connected branch.
	//
	// Note that deploying a specific commit with this endpoint does not disable autodeploys for the service.
	//
	// You can toggle autodeploys for your service with the [Update service](https://api-docs.render.com/reference/update-service) endpoint or in the Render Dashboard.
	//
	// Not supported for cron jobs.
	CommitId   pulumi.StringPtrOutput `pulumi:"commitId"`
	CreatedAt  pulumi.StringPtrOutput `pulumi:"createdAt"`
	FinishedAt pulumi.StringPtrOutput `pulumi:"finishedAt"`
	// Image information used when creating the deploy. Not present for Git-backed deploys
	Image ImagePropertiesPtrOutput `pulumi:"image"`
	// The URL of the image to deploy for an image-backed service.
	//
	// The host, repository, and image name all must match the currently configured image for the service.
	ImageUrl  pulumi.StringPtrOutput `pulumi:"imageUrl"`
	Status    StatusPtrOutput        `pulumi:"status"`
	Trigger   TriggerPtrOutput       `pulumi:"trigger"`
	UpdatedAt pulumi.StringPtrOutput `pulumi:"updatedAt"`
}

// NewDeploy registers a new resource with the given unique name, arguments, and options.
func NewDeploy(ctx *pulumi.Context,
	name string, args *DeployArgs, opts ...pulumi.ResourceOption) (*Deploy, error) {
	if args == nil {
		args = &DeployArgs{}
	}

	if args.ClearCache == nil {
		args.ClearCache = ClearCache("do_not_clear")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Deploy
	err := ctx.RegisterResource("render:services:Deploy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploy gets an existing Deploy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeployState, opts ...pulumi.ResourceOption) (*Deploy, error) {
	var resource Deploy
	err := ctx.ReadResource("render:services:Deploy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Deploy resources.
type deployState struct {
}

type DeployState struct {
}

func (DeployState) ElementType() reflect.Type {
	return reflect.TypeOf((*deployState)(nil)).Elem()
}

type deployArgs struct {
	// If `clear`, Render clears the service's build cache before deploying. This can be useful if you're experiencing issues with your build.
	ClearCache *ClearCache `pulumi:"clearCache"`
	// The SHA of a specific Git commit to deploy for a service. Defaults to the latest commit on the service's connected branch.
	//
	// Note that deploying a specific commit with this endpoint does not disable autodeploys for the service.
	//
	// You can toggle autodeploys for your service with the [Update service](https://api-docs.render.com/reference/update-service) endpoint or in the Render Dashboard.
	//
	// Not supported for cron jobs.
	CommitId *string `pulumi:"commitId"`
	// The URL of the image to deploy for an image-backed service.
	//
	// The host, repository, and image name all must match the currently configured image for the service.
	ImageUrl *string `pulumi:"imageUrl"`
	// The ID of the service
	ServiceId *string `pulumi:"serviceId"`
}

// The set of arguments for constructing a Deploy resource.
type DeployArgs struct {
	// If `clear`, Render clears the service's build cache before deploying. This can be useful if you're experiencing issues with your build.
	ClearCache ClearCachePtrInput
	// The SHA of a specific Git commit to deploy for a service. Defaults to the latest commit on the service's connected branch.
	//
	// Note that deploying a specific commit with this endpoint does not disable autodeploys for the service.
	//
	// You can toggle autodeploys for your service with the [Update service](https://api-docs.render.com/reference/update-service) endpoint or in the Render Dashboard.
	//
	// Not supported for cron jobs.
	CommitId pulumi.StringPtrInput
	// The URL of the image to deploy for an image-backed service.
	//
	// The host, repository, and image name all must match the currently configured image for the service.
	ImageUrl pulumi.StringPtrInput
	// The ID of the service
	ServiceId pulumi.StringPtrInput
}

func (DeployArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deployArgs)(nil)).Elem()
}

type DeployInput interface {
	pulumi.Input

	ToDeployOutput() DeployOutput
	ToDeployOutputWithContext(ctx context.Context) DeployOutput
}

func (*Deploy) ElementType() reflect.Type {
	return reflect.TypeOf((**Deploy)(nil)).Elem()
}

func (i *Deploy) ToDeployOutput() DeployOutput {
	return i.ToDeployOutputWithContext(context.Background())
}

func (i *Deploy) ToDeployOutputWithContext(ctx context.Context) DeployOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployOutput)
}

type DeployOutput struct{ *pulumi.OutputState }

func (DeployOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Deploy)(nil)).Elem()
}

func (o DeployOutput) ToDeployOutput() DeployOutput {
	return o
}

func (o DeployOutput) ToDeployOutputWithContext(ctx context.Context) DeployOutput {
	return o
}

// If `clear`, Render clears the service's build cache before deploying. This can be useful if you're experiencing issues with your build.
func (o DeployOutput) ClearCache() ClearCachePtrOutput {
	return o.ApplyT(func(v *Deploy) ClearCachePtrOutput { return v.ClearCache }).(ClearCachePtrOutput)
}

func (o DeployOutput) Commit() CommitPropertiesPtrOutput {
	return o.ApplyT(func(v *Deploy) CommitPropertiesPtrOutput { return v.Commit }).(CommitPropertiesPtrOutput)
}

// The SHA of a specific Git commit to deploy for a service. Defaults to the latest commit on the service's connected branch.
//
// Note that deploying a specific commit with this endpoint does not disable autodeploys for the service.
//
// You can toggle autodeploys for your service with the [Update service](https://api-docs.render.com/reference/update-service) endpoint or in the Render Dashboard.
//
// Not supported for cron jobs.
func (o DeployOutput) CommitId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deploy) pulumi.StringPtrOutput { return v.CommitId }).(pulumi.StringPtrOutput)
}

func (o DeployOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deploy) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o DeployOutput) FinishedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deploy) pulumi.StringPtrOutput { return v.FinishedAt }).(pulumi.StringPtrOutput)
}

// Image information used when creating the deploy. Not present for Git-backed deploys
func (o DeployOutput) Image() ImagePropertiesPtrOutput {
	return o.ApplyT(func(v *Deploy) ImagePropertiesPtrOutput { return v.Image }).(ImagePropertiesPtrOutput)
}

// The URL of the image to deploy for an image-backed service.
//
// The host, repository, and image name all must match the currently configured image for the service.
func (o DeployOutput) ImageUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deploy) pulumi.StringPtrOutput { return v.ImageUrl }).(pulumi.StringPtrOutput)
}

func (o DeployOutput) Status() StatusPtrOutput {
	return o.ApplyT(func(v *Deploy) StatusPtrOutput { return v.Status }).(StatusPtrOutput)
}

func (o DeployOutput) Trigger() TriggerPtrOutput {
	return o.ApplyT(func(v *Deploy) TriggerPtrOutput { return v.Trigger }).(TriggerPtrOutput)
}

func (o DeployOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deploy) pulumi.StringPtrOutput { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeployInput)(nil)).Elem(), &Deploy{})
	pulumi.RegisterOutputType(DeployOutput{})
}
