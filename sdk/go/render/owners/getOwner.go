// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package owners

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOwner(ctx *pulumi.Context, args *GetOwnerArgs, opts ...pulumi.InvokeOption) (*GetOwnerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOwnerResult
	err := ctx.Invoke("render:owners:getOwner", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetOwnerArgs struct {
}

type GetOwnerResult struct {
	Items Owner `pulumi:"items"`
}

func GetOwnerOutput(ctx *pulumi.Context, args GetOwnerOutputArgs, opts ...pulumi.InvokeOption) GetOwnerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOwnerResult, error) {
			args := v.(GetOwnerArgs)
			r, err := GetOwner(ctx, &args, opts...)
			var s GetOwnerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOwnerResultOutput)
}

type GetOwnerOutputArgs struct {
}

func (GetOwnerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOwnerArgs)(nil)).Elem()
}

type GetOwnerResultOutput struct{ *pulumi.OutputState }

func (GetOwnerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOwnerResult)(nil)).Elem()
}

func (o GetOwnerResultOutput) ToGetOwnerResultOutput() GetOwnerResultOutput {
	return o
}

func (o GetOwnerResultOutput) ToGetOwnerResultOutputWithContext(ctx context.Context) GetOwnerResultOutput {
	return o
}

func (o GetOwnerResultOutput) Items() OwnerOutput {
	return o.ApplyT(func(v GetOwnerResult) Owner { return v.Items }).(OwnerOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOwnerResultOutput{})
}
