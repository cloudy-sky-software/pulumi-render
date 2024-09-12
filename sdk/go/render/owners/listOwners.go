// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package owners

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListOwners(ctx *pulumi.Context, args *ListOwnersArgs, opts ...pulumi.InvokeOption) (*ListOwnersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListOwnersResult
	err := ctx.Invoke("render:owners:listOwners", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListOwnersArgs struct {
}

type ListOwnersResult struct {
	Items []ListOwnersItemProperties `pulumi:"items"`
}

func ListOwnersOutput(ctx *pulumi.Context, args ListOwnersOutputArgs, opts ...pulumi.InvokeOption) ListOwnersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListOwnersResultOutput, error) {
			args := v.(ListOwnersArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv ListOwnersResult
			secret, err := ctx.InvokePackageRaw("render:owners:listOwners", args, &rv, "", opts...)
			if err != nil {
				return ListOwnersResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(ListOwnersResultOutput)
			if secret {
				return pulumi.ToSecret(output).(ListOwnersResultOutput), nil
			}
			return output, nil
		}).(ListOwnersResultOutput)
}

type ListOwnersOutputArgs struct {
}

func (ListOwnersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOwnersArgs)(nil)).Elem()
}

type ListOwnersResultOutput struct{ *pulumi.OutputState }

func (ListOwnersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOwnersResult)(nil)).Elem()
}

func (o ListOwnersResultOutput) ToListOwnersResultOutput() ListOwnersResultOutput {
	return o
}

func (o ListOwnersResultOutput) ToListOwnersResultOutputWithContext(ctx context.Context) ListOwnersResultOutput {
	return o
}

func (o ListOwnersResultOutput) Items() ListOwnersItemPropertiesArrayOutput {
	return o.ApplyT(func(v ListOwnersResult) []ListOwnersItemProperties { return v.Items }).(ListOwnersItemPropertiesArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListOwnersResultOutput{})
}
