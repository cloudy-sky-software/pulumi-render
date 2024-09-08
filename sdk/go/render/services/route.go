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

type Route struct {
	pulumi.CustomResourceState

	Destination pulumi.StringOutput `pulumi:"destination"`
	// Redirect and Rewrite Rules are applied in priority order starting at 0
	Priority pulumi.IntOutput    `pulumi:"priority"`
	Source   pulumi.StringOutput `pulumi:"source"`
	Type     TypeOutput          `pulumi:"type"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Route
	err := ctx.RegisterResource("render:services:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("render:services:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
}

type RouteState struct {
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	Destination string `pulumi:"destination"`
	// Redirect and Rewrite Rules are applied in priority order starting at 0. Defaults to last in the priority list.
	Priority *int `pulumi:"priority"`
	// The ID of the service
	ServiceId *string `pulumi:"serviceId"`
	Source    string  `pulumi:"source"`
	Type      Type    `pulumi:"type"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	Destination pulumi.StringInput
	// Redirect and Rewrite Rules are applied in priority order starting at 0. Defaults to last in the priority list.
	Priority pulumi.IntPtrInput
	// The ID of the service
	ServiceId pulumi.StringPtrInput
	Source    pulumi.StringInput
	Type      TypeInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

type RouteOutput struct{ *pulumi.OutputState }

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func (o RouteOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Destination }).(pulumi.StringOutput)
}

// Redirect and Rewrite Rules are applied in priority order starting at 0
func (o RouteOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *Route) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

func (o RouteOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

func (o RouteOutput) Type() TypeOutput {
	return o.ApplyT(func(v *Route) TypeOutput { return v.Type }).(TypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteInput)(nil)).Elem(), &Route{})
	pulumi.RegisterOutputType(RouteOutput{})
}