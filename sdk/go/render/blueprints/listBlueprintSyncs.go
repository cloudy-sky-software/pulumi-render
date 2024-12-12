// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blueprints

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBlueprintSyncs(ctx *pulumi.Context, args *ListBlueprintSyncsArgs, opts ...pulumi.InvokeOption) (*ListBlueprintSyncsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListBlueprintSyncsResult
	err := ctx.Invoke("render:blueprints:listBlueprintSyncs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBlueprintSyncsArgs struct {
	// The ID of the blueprint
	BlueprintId string `pulumi:"blueprintId"`
}

type ListBlueprintSyncsResult struct {
	Items []SyncWithCursor `pulumi:"items"`
}

func ListBlueprintSyncsOutput(ctx *pulumi.Context, args ListBlueprintSyncsOutputArgs, opts ...pulumi.InvokeOption) ListBlueprintSyncsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListBlueprintSyncsResultOutput, error) {
			args := v.(ListBlueprintSyncsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("render:blueprints:listBlueprintSyncs", args, ListBlueprintSyncsResultOutput{}, options).(ListBlueprintSyncsResultOutput), nil
		}).(ListBlueprintSyncsResultOutput)
}

type ListBlueprintSyncsOutputArgs struct {
	// The ID of the blueprint
	BlueprintId pulumi.StringInput `pulumi:"blueprintId"`
}

func (ListBlueprintSyncsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBlueprintSyncsArgs)(nil)).Elem()
}

type ListBlueprintSyncsResultOutput struct{ *pulumi.OutputState }

func (ListBlueprintSyncsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBlueprintSyncsResult)(nil)).Elem()
}

func (o ListBlueprintSyncsResultOutput) ToListBlueprintSyncsResultOutput() ListBlueprintSyncsResultOutput {
	return o
}

func (o ListBlueprintSyncsResultOutput) ToListBlueprintSyncsResultOutputWithContext(ctx context.Context) ListBlueprintSyncsResultOutput {
	return o
}

func (o ListBlueprintSyncsResultOutput) Items() SyncWithCursorArrayOutput {
	return o.ApplyT(func(v ListBlueprintSyncsResult) []SyncWithCursor { return v.Items }).(SyncWithCursorArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListBlueprintSyncsResultOutput{})
}
