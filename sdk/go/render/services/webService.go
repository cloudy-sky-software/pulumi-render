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

type WebService struct {
	pulumi.CustomResourceState

	AutoDeploy     WebServiceServiceAutoDeployPtrOutput       `pulumi:"autoDeploy"`
	Branch         pulumi.StringPtrOutput                     `pulumi:"branch"`
	BuildFilter    BuildFilterPtrOutput                       `pulumi:"buildFilter"`
	CreatedAt      pulumi.StringPtrOutput                     `pulumi:"createdAt"`
	EnvVars        pulumi.ArrayOutput                         `pulumi:"envVars"`
	Image          ImagePtrOutput                             `pulumi:"image"`
	ImagePath      pulumi.StringPtrOutput                     `pulumi:"imagePath"`
	Name           pulumi.StringPtrOutput                     `pulumi:"name"`
	NotifyOnFail   WebServiceServiceNotifyOnFailPtrOutput     `pulumi:"notifyOnFail"`
	OwnerId        pulumi.StringPtrOutput                     `pulumi:"ownerId"`
	Repo           pulumi.StringPtrOutput                     `pulumi:"repo"`
	RootDir        pulumi.StringPtrOutput                     `pulumi:"rootDir"`
	SecretFiles    SecretFileArrayOutput                      `pulumi:"secretFiles"`
	ServiceDetails WebServiceDetailsOutputPtrOutput           `pulumi:"serviceDetails"`
	Slug           pulumi.StringPtrOutput                     `pulumi:"slug"`
	Suspended      WebServiceServiceSuspendedPtrOutput        `pulumi:"suspended"`
	Suspenders     WebServiceServiceSuspendersItemArrayOutput `pulumi:"suspenders"`
	Type           pulumi.StringPtrOutput                     `pulumi:"type"`
	UpdatedAt      pulumi.StringPtrOutput                     `pulumi:"updatedAt"`
}

// NewWebService registers a new resource with the given unique name, arguments, and options.
func NewWebService(ctx *pulumi.Context,
	name string, args *WebServiceArgs, opts ...pulumi.ResourceOption) (*WebService, error) {
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
		args.AutoDeploy = WebServiceServiceCreateAutoDeploy("yes")
	}
	if args.ServiceDetails != nil {
		args.ServiceDetails = args.ServiceDetails.ToWebServiceDetailsCreatePtrOutput().ApplyT(func(v *WebServiceDetailsCreate) *WebServiceDetailsCreate { return v.Defaults() }).(WebServiceDetailsCreatePtrOutput)
	}
	if args.Type == nil {
		args.Type = pulumi.StringPtr("web_service")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WebService
	err := ctx.RegisterResource("render:services:WebService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebService gets an existing WebService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebServiceState, opts ...pulumi.ResourceOption) (*WebService, error) {
	var resource WebService
	err := ctx.ReadResource("render:services:WebService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebService resources.
type webServiceState struct {
}

type WebServiceState struct {
}

func (WebServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*webServiceState)(nil)).Elem()
}

type webServiceArgs struct {
	// Defaults to "yes"
	AutoDeploy *WebServiceServiceCreateAutoDeploy `pulumi:"autoDeploy"`
	// If left empty, this will fall back to the default branch of the repository
	Branch      *string       `pulumi:"branch"`
	BuildFilter *BuildFilter  `pulumi:"buildFilter"`
	EnvVars     []interface{} `pulumi:"envVars"`
	Image       *Image        `pulumi:"image"`
	Name        string        `pulumi:"name"`
	OwnerId     string        `pulumi:"ownerId"`
	// Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
	Repo           *string                  `pulumi:"repo"`
	RootDir        *string                  `pulumi:"rootDir"`
	SecretFiles    []SecretFile             `pulumi:"secretFiles"`
	ServiceDetails *WebServiceDetailsCreate `pulumi:"serviceDetails"`
	Type           *string                  `pulumi:"type"`
}

// The set of arguments for constructing a WebService resource.
type WebServiceArgs struct {
	// Defaults to "yes"
	AutoDeploy WebServiceServiceCreateAutoDeployPtrInput
	// If left empty, this will fall back to the default branch of the repository
	Branch      pulumi.StringPtrInput
	BuildFilter BuildFilterPtrInput
	EnvVars     pulumi.ArrayInput
	Image       ImagePtrInput
	Name        pulumi.StringInput
	OwnerId     pulumi.StringInput
	// Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
	Repo           pulumi.StringPtrInput
	RootDir        pulumi.StringPtrInput
	SecretFiles    SecretFileArrayInput
	ServiceDetails WebServiceDetailsCreatePtrInput
	Type           pulumi.StringPtrInput
}

func (WebServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webServiceArgs)(nil)).Elem()
}

type WebServiceInput interface {
	pulumi.Input

	ToWebServiceOutput() WebServiceOutput
	ToWebServiceOutputWithContext(ctx context.Context) WebServiceOutput
}

func (*WebService) ElementType() reflect.Type {
	return reflect.TypeOf((**WebService)(nil)).Elem()
}

func (i *WebService) ToWebServiceOutput() WebServiceOutput {
	return i.ToWebServiceOutputWithContext(context.Background())
}

func (i *WebService) ToWebServiceOutputWithContext(ctx context.Context) WebServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServiceOutput)
}

type WebServiceOutput struct{ *pulumi.OutputState }

func (WebServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebService)(nil)).Elem()
}

func (o WebServiceOutput) ToWebServiceOutput() WebServiceOutput {
	return o
}

func (o WebServiceOutput) ToWebServiceOutputWithContext(ctx context.Context) WebServiceOutput {
	return o
}

func (o WebServiceOutput) AutoDeploy() WebServiceServiceAutoDeployPtrOutput {
	return o.ApplyT(func(v *WebService) WebServiceServiceAutoDeployPtrOutput { return v.AutoDeploy }).(WebServiceServiceAutoDeployPtrOutput)
}

func (o WebServiceOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringPtrOutput { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o WebServiceOutput) BuildFilter() BuildFilterPtrOutput {
	return o.ApplyT(func(v *WebService) BuildFilterPtrOutput { return v.BuildFilter }).(BuildFilterPtrOutput)
}

func (o WebServiceOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o WebServiceOutput) EnvVars() pulumi.ArrayOutput {
	return o.ApplyT(func(v *WebService) pulumi.ArrayOutput { return v.EnvVars }).(pulumi.ArrayOutput)
}

func (o WebServiceOutput) Image() ImagePtrOutput {
	return o.ApplyT(func(v *WebService) ImagePtrOutput { return v.Image }).(ImagePtrOutput)
}

func (o WebServiceOutput) ImagePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringPtrOutput { return v.ImagePath }).(pulumi.StringPtrOutput)
}

func (o WebServiceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WebServiceOutput) NotifyOnFail() WebServiceServiceNotifyOnFailPtrOutput {
	return o.ApplyT(func(v *WebService) WebServiceServiceNotifyOnFailPtrOutput { return v.NotifyOnFail }).(WebServiceServiceNotifyOnFailPtrOutput)
}

func (o WebServiceOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringPtrOutput { return v.OwnerId }).(pulumi.StringPtrOutput)
}

func (o WebServiceOutput) Repo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringPtrOutput { return v.Repo }).(pulumi.StringPtrOutput)
}

func (o WebServiceOutput) RootDir() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringPtrOutput { return v.RootDir }).(pulumi.StringPtrOutput)
}

func (o WebServiceOutput) SecretFiles() SecretFileArrayOutput {
	return o.ApplyT(func(v *WebService) SecretFileArrayOutput { return v.SecretFiles }).(SecretFileArrayOutput)
}

func (o WebServiceOutput) ServiceDetails() WebServiceDetailsOutputPtrOutput {
	return o.ApplyT(func(v *WebService) WebServiceDetailsOutputPtrOutput { return v.ServiceDetails }).(WebServiceDetailsOutputPtrOutput)
}

func (o WebServiceOutput) Slug() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringPtrOutput { return v.Slug }).(pulumi.StringPtrOutput)
}

func (o WebServiceOutput) Suspended() WebServiceServiceSuspendedPtrOutput {
	return o.ApplyT(func(v *WebService) WebServiceServiceSuspendedPtrOutput { return v.Suspended }).(WebServiceServiceSuspendedPtrOutput)
}

func (o WebServiceOutput) Suspenders() WebServiceServiceSuspendersItemArrayOutput {
	return o.ApplyT(func(v *WebService) WebServiceServiceSuspendersItemArrayOutput { return v.Suspenders }).(WebServiceServiceSuspendersItemArrayOutput)
}

func (o WebServiceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o WebServiceOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringPtrOutput { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebServiceInput)(nil)).Elem(), &WebService{})
	pulumi.RegisterOutputType(WebServiceOutput{})
}
