// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blueprints

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBlueprint(ctx *pulumi.Context, args *GetBlueprintArgs, opts ...pulumi.InvokeOption) (*GetBlueprintResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBlueprintResult
	err := ctx.Invoke("render:blueprints:getBlueprint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetBlueprintArgs struct {
	// The ID of the blueprint
	BlueprintId string `pulumi:"blueprintId"`
}

type GetBlueprintResult struct {
	// Automatically sync changes to render.yaml
	AutoSync  bool                                            `pulumi:"autoSync"`
	Branch    string                                          `pulumi:"branch"`
	Id        string                                          `pulumi:"id"`
	LastSync  *string                                         `pulumi:"lastSync"`
	Name      string                                          `pulumi:"name"`
	Repo      string                                          `pulumi:"repo"`
	Resources []GetBlueprintPropertiesResourcesItemProperties `pulumi:"resources"`
	Status    GetBlueprintPropertiesStatus                    `pulumi:"status"`
}

func GetBlueprintOutput(ctx *pulumi.Context, args GetBlueprintOutputArgs, opts ...pulumi.InvokeOption) GetBlueprintResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetBlueprintResultOutput, error) {
			args := v.(GetBlueprintArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("render:blueprints:getBlueprint", args, GetBlueprintResultOutput{}, options).(GetBlueprintResultOutput), nil
		}).(GetBlueprintResultOutput)
}

type GetBlueprintOutputArgs struct {
	// The ID of the blueprint
	BlueprintId pulumi.StringInput `pulumi:"blueprintId"`
}

func (GetBlueprintOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBlueprintArgs)(nil)).Elem()
}

type GetBlueprintResultOutput struct{ *pulumi.OutputState }

func (GetBlueprintResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBlueprintResult)(nil)).Elem()
}

func (o GetBlueprintResultOutput) ToGetBlueprintResultOutput() GetBlueprintResultOutput {
	return o
}

func (o GetBlueprintResultOutput) ToGetBlueprintResultOutputWithContext(ctx context.Context) GetBlueprintResultOutput {
	return o
}

// Automatically sync changes to render.yaml
func (o GetBlueprintResultOutput) AutoSync() pulumi.BoolOutput {
	return o.ApplyT(func(v GetBlueprintResult) bool { return v.AutoSync }).(pulumi.BoolOutput)
}

func (o GetBlueprintResultOutput) Branch() pulumi.StringOutput {
	return o.ApplyT(func(v GetBlueprintResult) string { return v.Branch }).(pulumi.StringOutput)
}

func (o GetBlueprintResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBlueprintResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBlueprintResultOutput) LastSync() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBlueprintResult) *string { return v.LastSync }).(pulumi.StringPtrOutput)
}

func (o GetBlueprintResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetBlueprintResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetBlueprintResultOutput) Repo() pulumi.StringOutput {
	return o.ApplyT(func(v GetBlueprintResult) string { return v.Repo }).(pulumi.StringOutput)
}

func (o GetBlueprintResultOutput) Resources() GetBlueprintPropertiesResourcesItemPropertiesArrayOutput {
	return o.ApplyT(func(v GetBlueprintResult) []GetBlueprintPropertiesResourcesItemProperties { return v.Resources }).(GetBlueprintPropertiesResourcesItemPropertiesArrayOutput)
}

func (o GetBlueprintResultOutput) Status() GetBlueprintPropertiesStatusOutput {
	return o.ApplyT(func(v GetBlueprintResult) GetBlueprintPropertiesStatus { return v.Status }).(GetBlueprintPropertiesStatusOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBlueprintResultOutput{})
}
