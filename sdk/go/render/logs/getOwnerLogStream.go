// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logs

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOwnerLogStream(ctx *pulumi.Context, args *GetOwnerLogStreamArgs, opts ...pulumi.InvokeOption) (*GetOwnerLogStreamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOwnerLogStreamResult
	err := ctx.Invoke("render:logs:getOwnerLogStream", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetOwnerLogStreamArgs struct {
	// The ID of the owner (team or personal user) whose log streams should be returned
	OwnerId string `pulumi:"ownerId"`
}

// Owner log stream settings
type GetOwnerLogStreamResult struct {
	// The endpoint to stream logs to.
	Endpoint *string `pulumi:"endpoint"`
	// The ID of the owner.
	OwnerId *string `pulumi:"ownerId"`
	// Whether to send logs or drop them.
	Preview *GetOwnerLogStreamPropertiesPreview `pulumi:"preview"`
}

func GetOwnerLogStreamOutput(ctx *pulumi.Context, args GetOwnerLogStreamOutputArgs, opts ...pulumi.InvokeOption) GetOwnerLogStreamResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetOwnerLogStreamResultOutput, error) {
			args := v.(GetOwnerLogStreamArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("render:logs:getOwnerLogStream", args, GetOwnerLogStreamResultOutput{}, options).(GetOwnerLogStreamResultOutput), nil
		}).(GetOwnerLogStreamResultOutput)
}

type GetOwnerLogStreamOutputArgs struct {
	// The ID of the owner (team or personal user) whose log streams should be returned
	OwnerId pulumi.StringInput `pulumi:"ownerId"`
}

func (GetOwnerLogStreamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOwnerLogStreamArgs)(nil)).Elem()
}

// Owner log stream settings
type GetOwnerLogStreamResultOutput struct{ *pulumi.OutputState }

func (GetOwnerLogStreamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOwnerLogStreamResult)(nil)).Elem()
}

func (o GetOwnerLogStreamResultOutput) ToGetOwnerLogStreamResultOutput() GetOwnerLogStreamResultOutput {
	return o
}

func (o GetOwnerLogStreamResultOutput) ToGetOwnerLogStreamResultOutputWithContext(ctx context.Context) GetOwnerLogStreamResultOutput {
	return o
}

// The endpoint to stream logs to.
func (o GetOwnerLogStreamResultOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOwnerLogStreamResult) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

// The ID of the owner.
func (o GetOwnerLogStreamResultOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOwnerLogStreamResult) *string { return v.OwnerId }).(pulumi.StringPtrOutput)
}

// Whether to send logs or drop them.
func (o GetOwnerLogStreamResultOutput) Preview() GetOwnerLogStreamPropertiesPreviewPtrOutput {
	return o.ApplyT(func(v GetOwnerLogStreamResult) *GetOwnerLogStreamPropertiesPreview { return v.Preview }).(GetOwnerLogStreamPropertiesPreviewPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOwnerLogStreamResultOutput{})
}
