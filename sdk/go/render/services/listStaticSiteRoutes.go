// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStaticSiteRoutes(ctx *pulumi.Context, args *ListStaticSiteRoutesArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteRoutesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv ListStaticSiteRoutesResult
	err := ctx.Invoke("render:services:listStaticSiteRoutes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteRoutesArgs struct {
	// (Required) The ID of the service
	ServiceId string `pulumi:"serviceId"`
}

type ListStaticSiteRoutesResult struct {
	Value []ListStaticSiteRoutesResponse `pulumi:"value"`
}

func ListStaticSiteRoutesOutput(ctx *pulumi.Context, args ListStaticSiteRoutesOutputArgs, opts ...pulumi.InvokeOption) ListStaticSiteRoutesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStaticSiteRoutesResult, error) {
			args := v.(ListStaticSiteRoutesArgs)
			r, err := ListStaticSiteRoutes(ctx, &args, opts...)
			var s ListStaticSiteRoutesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStaticSiteRoutesResultOutput)
}

type ListStaticSiteRoutesOutputArgs struct {
	// (Required) The ID of the service
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (ListStaticSiteRoutesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteRoutesArgs)(nil)).Elem()
}

type ListStaticSiteRoutesResultOutput struct{ *pulumi.OutputState }

func (ListStaticSiteRoutesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteRoutesResult)(nil)).Elem()
}

func (o ListStaticSiteRoutesResultOutput) ToListStaticSiteRoutesResultOutput() ListStaticSiteRoutesResultOutput {
	return o
}

func (o ListStaticSiteRoutesResultOutput) ToListStaticSiteRoutesResultOutputWithContext(ctx context.Context) ListStaticSiteRoutesResultOutput {
	return o
}

func (o ListStaticSiteRoutesResultOutput) Value() ListStaticSiteRoutesResponseArrayOutput {
	return o.ApplyT(func(v ListStaticSiteRoutesResult) []ListStaticSiteRoutesResponse { return v.Value }).(ListStaticSiteRoutesResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStaticSiteRoutesResultOutput{})
}
