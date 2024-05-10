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

type PreviewService struct {
	pulumi.CustomResourceState

	DeployId pulumi.StringPtrOutput `pulumi:"deployId"`
	// Must be either a full URL or the relative path to an image. If a relative path, Render uses the base service's image URL as its root. For example, if the base service's image URL is `docker.io/library/nginx:latest`, then valid values are: `docker.io/library/nginx:<any tag or SHA>`, `library/nginx:<any tag or SHA>`, or `nginx:<any tag or SHA>`. Note that the path must match (only the tag or SHA can vary).
	ImagePath pulumi.StringOutput `pulumi:"imagePath"`
	// A name for the service preview instance. If not specified, Render generates the name using the base service's name and the specified tag or SHA.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
	Plan    PlanPtrOutput        `pulumi:"plan"`
	Service ServiceTypePtrOutput `pulumi:"service"`
}

// NewPreviewService registers a new resource with the given unique name, arguments, and options.
func NewPreviewService(ctx *pulumi.Context,
	name string, args *PreviewServiceArgs, opts ...pulumi.ResourceOption) (*PreviewService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImagePath == nil {
		return nil, errors.New("invalid value for required argument 'ImagePath'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PreviewService
	err := ctx.RegisterResource("render:services:PreviewService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPreviewService gets an existing PreviewService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPreviewService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PreviewServiceState, opts ...pulumi.ResourceOption) (*PreviewService, error) {
	var resource PreviewService
	err := ctx.ReadResource("render:services:PreviewService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PreviewService resources.
type previewServiceState struct {
}

type PreviewServiceState struct {
}

func (PreviewServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*previewServiceState)(nil)).Elem()
}

type previewServiceArgs struct {
	// Must be either a full URL or the relative path to an image. If a relative path, Render uses the base service's image URL as its root. For example, if the base service's image URL is `docker.io/library/nginx:latest`, then valid values are: `docker.io/library/nginx:<any tag or SHA>`, `library/nginx:<any tag or SHA>`, or `nginx:<any tag or SHA>`. Note that the path must match (only the tag or SHA can vary).
	ImagePath string `pulumi:"imagePath"`
	// A name for the service preview instance. If not specified, Render generates the name using the base service's name and the specified tag or SHA.
	Name *string `pulumi:"name"`
	// The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
	Plan *Plan `pulumi:"plan"`
}

// The set of arguments for constructing a PreviewService resource.
type PreviewServiceArgs struct {
	// Must be either a full URL or the relative path to an image. If a relative path, Render uses the base service's image URL as its root. For example, if the base service's image URL is `docker.io/library/nginx:latest`, then valid values are: `docker.io/library/nginx:<any tag or SHA>`, `library/nginx:<any tag or SHA>`, or `nginx:<any tag or SHA>`. Note that the path must match (only the tag or SHA can vary).
	ImagePath pulumi.StringInput
	// A name for the service preview instance. If not specified, Render generates the name using the base service's name and the specified tag or SHA.
	Name pulumi.StringPtrInput
	// The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
	Plan PlanPtrInput
}

func (PreviewServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*previewServiceArgs)(nil)).Elem()
}

type PreviewServiceInput interface {
	pulumi.Input

	ToPreviewServiceOutput() PreviewServiceOutput
	ToPreviewServiceOutputWithContext(ctx context.Context) PreviewServiceOutput
}

func (*PreviewService) ElementType() reflect.Type {
	return reflect.TypeOf((**PreviewService)(nil)).Elem()
}

func (i *PreviewService) ToPreviewServiceOutput() PreviewServiceOutput {
	return i.ToPreviewServiceOutputWithContext(context.Background())
}

func (i *PreviewService) ToPreviewServiceOutputWithContext(ctx context.Context) PreviewServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreviewServiceOutput)
}

type PreviewServiceOutput struct{ *pulumi.OutputState }

func (PreviewServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PreviewService)(nil)).Elem()
}

func (o PreviewServiceOutput) ToPreviewServiceOutput() PreviewServiceOutput {
	return o
}

func (o PreviewServiceOutput) ToPreviewServiceOutputWithContext(ctx context.Context) PreviewServiceOutput {
	return o
}

func (o PreviewServiceOutput) DeployId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PreviewService) pulumi.StringPtrOutput { return v.DeployId }).(pulumi.StringPtrOutput)
}

// Must be either a full URL or the relative path to an image. If a relative path, Render uses the base service's image URL as its root. For example, if the base service's image URL is `docker.io/library/nginx:latest`, then valid values are: `docker.io/library/nginx:<any tag or SHA>`, `library/nginx:<any tag or SHA>`, or `nginx:<any tag or SHA>`. Note that the path must match (only the tag or SHA can vary).
func (o PreviewServiceOutput) ImagePath() pulumi.StringOutput {
	return o.ApplyT(func(v *PreviewService) pulumi.StringOutput { return v.ImagePath }).(pulumi.StringOutput)
}

// A name for the service preview instance. If not specified, Render generates the name using the base service's name and the specified tag or SHA.
func (o PreviewServiceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PreviewService) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
func (o PreviewServiceOutput) Plan() PlanPtrOutput {
	return o.ApplyT(func(v *PreviewService) PlanPtrOutput { return v.Plan }).(PlanPtrOutput)
}

func (o PreviewServiceOutput) Service() ServiceTypePtrOutput {
	return o.ApplyT(func(v *PreviewService) ServiceTypePtrOutput { return v.Service }).(ServiceTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PreviewServiceInput)(nil)).Elem(), &PreviewService{})
	pulumi.RegisterOutputType(PreviewServiceOutput{})
}
