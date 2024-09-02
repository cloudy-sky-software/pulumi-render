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

type AutoscaleService struct {
	pulumi.CustomResourceState

	Criteria CriteriaPropertiesOutput `pulumi:"criteria"`
	Enabled  pulumi.BoolOutput        `pulumi:"enabled"`
	// The maximum number of instances for the service
	Max pulumi.IntOutput `pulumi:"max"`
	// The minimum number of instances for the service
	Min pulumi.IntOutput `pulumi:"min"`
}

// NewAutoscaleService registers a new resource with the given unique name, arguments, and options.
func NewAutoscaleService(ctx *pulumi.Context,
	name string, args *AutoscaleServiceArgs, opts ...pulumi.ResourceOption) (*AutoscaleService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Criteria == nil {
		return nil, errors.New("invalid value for required argument 'Criteria'")
	}
	if args.Max == nil {
		return nil, errors.New("invalid value for required argument 'Max'")
	}
	if args.Min == nil {
		return nil, errors.New("invalid value for required argument 'Min'")
	}
	args.Criteria = args.Criteria.ToCriteriaPropertiesOutput().ApplyT(func(v CriteriaProperties) CriteriaProperties { return *v.Defaults() }).(CriteriaPropertiesOutput)
	if args.Enabled == nil {
		args.Enabled = pulumi.Bool(false)
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AutoscaleService
	err := ctx.RegisterResource("render:services:AutoscaleService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoscaleService gets an existing AutoscaleService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoscaleService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoscaleServiceState, opts ...pulumi.ResourceOption) (*AutoscaleService, error) {
	var resource AutoscaleService
	err := ctx.ReadResource("render:services:AutoscaleService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoscaleService resources.
type autoscaleServiceState struct {
}

type AutoscaleServiceState struct {
}

func (AutoscaleServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscaleServiceState)(nil)).Elem()
}

type autoscaleServiceArgs struct {
	Criteria CriteriaProperties `pulumi:"criteria"`
	Enabled  bool               `pulumi:"enabled"`
	// The maximum number of instances for the service
	Max int `pulumi:"max"`
	// The minimum number of instances for the service
	Min int `pulumi:"min"`
	// The ID of the service
	ServiceId *string `pulumi:"serviceId"`
}

// The set of arguments for constructing a AutoscaleService resource.
type AutoscaleServiceArgs struct {
	Criteria CriteriaPropertiesInput
	Enabled  pulumi.BoolInput
	// The maximum number of instances for the service
	Max pulumi.IntInput
	// The minimum number of instances for the service
	Min pulumi.IntInput
	// The ID of the service
	ServiceId pulumi.StringPtrInput
}

func (AutoscaleServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscaleServiceArgs)(nil)).Elem()
}

type AutoscaleServiceInput interface {
	pulumi.Input

	ToAutoscaleServiceOutput() AutoscaleServiceOutput
	ToAutoscaleServiceOutputWithContext(ctx context.Context) AutoscaleServiceOutput
}

func (*AutoscaleService) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleService)(nil)).Elem()
}

func (i *AutoscaleService) ToAutoscaleServiceOutput() AutoscaleServiceOutput {
	return i.ToAutoscaleServiceOutputWithContext(context.Background())
}

func (i *AutoscaleService) ToAutoscaleServiceOutputWithContext(ctx context.Context) AutoscaleServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleServiceOutput)
}

type AutoscaleServiceOutput struct{ *pulumi.OutputState }

func (AutoscaleServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleService)(nil)).Elem()
}

func (o AutoscaleServiceOutput) ToAutoscaleServiceOutput() AutoscaleServiceOutput {
	return o
}

func (o AutoscaleServiceOutput) ToAutoscaleServiceOutputWithContext(ctx context.Context) AutoscaleServiceOutput {
	return o
}

func (o AutoscaleServiceOutput) Criteria() CriteriaPropertiesOutput {
	return o.ApplyT(func(v *AutoscaleService) CriteriaPropertiesOutput { return v.Criteria }).(CriteriaPropertiesOutput)
}

func (o AutoscaleServiceOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *AutoscaleService) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// The maximum number of instances for the service
func (o AutoscaleServiceOutput) Max() pulumi.IntOutput {
	return o.ApplyT(func(v *AutoscaleService) pulumi.IntOutput { return v.Max }).(pulumi.IntOutput)
}

// The minimum number of instances for the service
func (o AutoscaleServiceOutput) Min() pulumi.IntOutput {
	return o.ApplyT(func(v *AutoscaleService) pulumi.IntOutput { return v.Min }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutoscaleServiceInput)(nil)).Elem(), &AutoscaleService{})
	pulumi.RegisterOutputType(AutoscaleServiceOutput{})
}
