// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"errors"
	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Input type for creating a Redis instance
type Redis struct {
	pulumi.CustomResourceState

	// The creation time of the Redis instance
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The ID of the environment the Redis instance is associated with
	EnvironmentId pulumi.StringPtrOutput `pulumi:"environmentId"`
	// The IP allow list for the Redis instance
	IpAllowList CidrBlockAndDescriptionArrayOutput `pulumi:"ipAllowList"`
	Maintenance MaintenancePropertiesPtrOutput     `pulumi:"maintenance"`
	// The eviction policy for the Key Value instance
	MaxmemoryPolicy MaxmemoryPolicyPtrOutput `pulumi:"maxmemoryPolicy"`
	// The name of the Redis instance
	Name pulumi.StringOutput `pulumi:"name"`
	// Options for a Redis instance
	Options RedisOptionsOutput `pulumi:"options"`
	Owner   OwnerOutput        `pulumi:"owner"`
	// The ID of the owner of the Redis instance
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	Plan    PlanOutput          `pulumi:"plan"`
	// Defaults to "oregon"
	Region RegionOutput `pulumi:"region"`
	Status StatusOutput `pulumi:"status"`
	// The last updated time of the Redis instance
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The version of Redis
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewRedis registers a new resource with the given unique name, arguments, and options.
func NewRedis(ctx *pulumi.Context,
	name string, args *RedisArgs, opts ...pulumi.ResourceOption) (*Redis, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OwnerId == nil {
		return nil, errors.New("invalid value for required argument 'OwnerId'")
	}
	if args.Plan == nil {
		return nil, errors.New("invalid value for required argument 'Plan'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Redis
	err := ctx.RegisterResource("render:redis:Redis", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRedis gets an existing Redis resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRedis(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RedisState, opts ...pulumi.ResourceOption) (*Redis, error) {
	var resource Redis
	err := ctx.ReadResource("render:redis:Redis", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Redis resources.
type redisState struct {
}

type RedisState struct {
}

func (RedisState) ElementType() reflect.Type {
	return reflect.TypeOf((*redisState)(nil)).Elem()
}

type redisArgs struct {
	EnvironmentId *string                   `pulumi:"environmentId"`
	IpAllowList   []CidrBlockAndDescription `pulumi:"ipAllowList"`
	// The eviction policy for the Key Value instance
	MaxmemoryPolicy *MaxmemoryPolicy `pulumi:"maxmemoryPolicy"`
	// The name of the Redis instance
	Name *string `pulumi:"name"`
	// The ID of the owner of the Redis instance
	OwnerId string `pulumi:"ownerId"`
	Plan    Plan   `pulumi:"plan"`
	// The region where the Redis instance is located
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Redis resource.
type RedisArgs struct {
	EnvironmentId pulumi.StringPtrInput
	IpAllowList   CidrBlockAndDescriptionArrayInput
	// The eviction policy for the Key Value instance
	MaxmemoryPolicy MaxmemoryPolicyPtrInput
	// The name of the Redis instance
	Name pulumi.StringPtrInput
	// The ID of the owner of the Redis instance
	OwnerId pulumi.StringInput
	Plan    PlanInput
	// The region where the Redis instance is located
	Region pulumi.StringPtrInput
}

func (RedisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*redisArgs)(nil)).Elem()
}

type RedisInput interface {
	pulumi.Input

	ToRedisOutput() RedisOutput
	ToRedisOutputWithContext(ctx context.Context) RedisOutput
}

func (*Redis) ElementType() reflect.Type {
	return reflect.TypeOf((**Redis)(nil)).Elem()
}

func (i *Redis) ToRedisOutput() RedisOutput {
	return i.ToRedisOutputWithContext(context.Background())
}

func (i *Redis) ToRedisOutputWithContext(ctx context.Context) RedisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisOutput)
}

type RedisOutput struct{ *pulumi.OutputState }

func (RedisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Redis)(nil)).Elem()
}

func (o RedisOutput) ToRedisOutput() RedisOutput {
	return o
}

func (o RedisOutput) ToRedisOutputWithContext(ctx context.Context) RedisOutput {
	return o
}

// The creation time of the Redis instance
func (o RedisOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the environment the Redis instance is associated with
func (o RedisOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringPtrOutput { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

// The IP allow list for the Redis instance
func (o RedisOutput) IpAllowList() CidrBlockAndDescriptionArrayOutput {
	return o.ApplyT(func(v *Redis) CidrBlockAndDescriptionArrayOutput { return v.IpAllowList }).(CidrBlockAndDescriptionArrayOutput)
}

func (o RedisOutput) Maintenance() MaintenancePropertiesPtrOutput {
	return o.ApplyT(func(v *Redis) MaintenancePropertiesPtrOutput { return v.Maintenance }).(MaintenancePropertiesPtrOutput)
}

// The eviction policy for the Key Value instance
func (o RedisOutput) MaxmemoryPolicy() MaxmemoryPolicyPtrOutput {
	return o.ApplyT(func(v *Redis) MaxmemoryPolicyPtrOutput { return v.MaxmemoryPolicy }).(MaxmemoryPolicyPtrOutput)
}

// The name of the Redis instance
func (o RedisOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Options for a Redis instance
func (o RedisOutput) Options() RedisOptionsOutput {
	return o.ApplyT(func(v *Redis) RedisOptionsOutput { return v.Options }).(RedisOptionsOutput)
}

func (o RedisOutput) Owner() OwnerOutput {
	return o.ApplyT(func(v *Redis) OwnerOutput { return v.Owner }).(OwnerOutput)
}

// The ID of the owner of the Redis instance
func (o RedisOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

func (o RedisOutput) Plan() PlanOutput {
	return o.ApplyT(func(v *Redis) PlanOutput { return v.Plan }).(PlanOutput)
}

// Defaults to "oregon"
func (o RedisOutput) Region() RegionOutput {
	return o.ApplyT(func(v *Redis) RegionOutput { return v.Region }).(RegionOutput)
}

func (o RedisOutput) Status() StatusOutput {
	return o.ApplyT(func(v *Redis) StatusOutput { return v.Status }).(StatusOutput)
}

// The last updated time of the Redis instance
func (o RedisOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The version of Redis
func (o RedisOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RedisInput)(nil)).Elem(), &Redis{})
	pulumi.RegisterOutputType(RedisOutput{})
}
