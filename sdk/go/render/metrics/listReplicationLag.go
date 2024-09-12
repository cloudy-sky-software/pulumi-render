// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package metrics

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListReplicationLag(ctx *pulumi.Context, args *ListReplicationLagArgs, opts ...pulumi.InvokeOption) (*ListReplicationLagResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListReplicationLagResult
	err := ctx.Invoke("render:metrics:listReplicationLag", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListReplicationLagArgs struct {
}

type ListReplicationLagResult struct {
	Items []ListReplicationLagItemProperties `pulumi:"items"`
}

func ListReplicationLagOutput(ctx *pulumi.Context, args ListReplicationLagOutputArgs, opts ...pulumi.InvokeOption) ListReplicationLagResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListReplicationLagResultOutput, error) {
			args := v.(ListReplicationLagArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv ListReplicationLagResult
			secret, err := ctx.InvokePackageRaw("render:metrics:listReplicationLag", args, &rv, "", opts...)
			if err != nil {
				return ListReplicationLagResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(ListReplicationLagResultOutput)
			if secret {
				return pulumi.ToSecret(output).(ListReplicationLagResultOutput), nil
			}
			return output, nil
		}).(ListReplicationLagResultOutput)
}

type ListReplicationLagOutputArgs struct {
}

func (ListReplicationLagOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListReplicationLagArgs)(nil)).Elem()
}

type ListReplicationLagResultOutput struct{ *pulumi.OutputState }

func (ListReplicationLagResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListReplicationLagResult)(nil)).Elem()
}

func (o ListReplicationLagResultOutput) ToListReplicationLagResultOutput() ListReplicationLagResultOutput {
	return o
}

func (o ListReplicationLagResultOutput) ToListReplicationLagResultOutputWithContext(ctx context.Context) ListReplicationLagResultOutput {
	return o
}

func (o ListReplicationLagResultOutput) Items() ListReplicationLagItemPropertiesArrayOutput {
	return o.ApplyT(func(v ListReplicationLagResult) []ListReplicationLagItemProperties { return v.Items }).(ListReplicationLagItemPropertiesArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListReplicationLagResultOutput{})
}
