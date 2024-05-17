// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebServiceOutput(ctx *pulumi.Context, args *LookupWebServiceOutputArgs, opts ...pulumi.InvokeOption) (*LookupWebServiceOutputResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWebServiceOutputResult
	err := ctx.Invoke("render:services:getWebServiceOutput", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWebServiceOutputArgs struct {
	// The ID of the service
	ServiceId string `pulumi:"serviceId"`
}

type LookupWebServiceOutputResult struct {
	Items GetWebServiceOutput `pulumi:"items"`
}

// Defaults sets the appropriate defaults for LookupWebServiceOutputResult
func (val *LookupWebServiceOutputResult) Defaults() *LookupWebServiceOutputResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func LookupWebServiceOutputOutput(ctx *pulumi.Context, args LookupWebServiceOutputOutputArgs, opts ...pulumi.InvokeOption) LookupWebServiceOutputResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebServiceOutputResult, error) {
			args := v.(LookupWebServiceOutputArgs)
			r, err := LookupWebServiceOutput(ctx, &args, opts...)
			var s LookupWebServiceOutputResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebServiceOutputResultOutput)
}

type LookupWebServiceOutputOutputArgs struct {
	// The ID of the service
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (LookupWebServiceOutputOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebServiceOutputArgs)(nil)).Elem()
}

type LookupWebServiceOutputResultOutput struct{ *pulumi.OutputState }

func (LookupWebServiceOutputResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebServiceOutputResult)(nil)).Elem()
}

func (o LookupWebServiceOutputResultOutput) ToLookupWebServiceOutputResultOutput() LookupWebServiceOutputResultOutput {
	return o
}

func (o LookupWebServiceOutputResultOutput) ToLookupWebServiceOutputResultOutputWithContext(ctx context.Context) LookupWebServiceOutputResultOutput {
	return o
}

func (o LookupWebServiceOutputResultOutput) Items() GetWebServiceOutputOutput {
	return o.ApplyT(func(v LookupWebServiceOutputResult) GetWebServiceOutput { return v.Items }).(GetWebServiceOutputOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebServiceOutputResultOutput{})
}