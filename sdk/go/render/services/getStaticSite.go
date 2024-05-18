// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStaticSite(ctx *pulumi.Context, args *LookupStaticSiteArgs, opts ...pulumi.InvokeOption) (*LookupStaticSiteResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStaticSiteResult
	err := ctx.Invoke("render:services:getStaticSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupStaticSiteArgs struct {
	// The ID of the service
	ServiceId string `pulumi:"serviceId"`
}

type LookupStaticSiteResult struct {
	Items GetStaticSiteType `pulumi:"items"`
}

// Defaults sets the appropriate defaults for LookupStaticSiteResult
func (val *LookupStaticSiteResult) Defaults() *LookupStaticSiteResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func LookupStaticSiteOutput(ctx *pulumi.Context, args LookupStaticSiteOutputArgs, opts ...pulumi.InvokeOption) LookupStaticSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStaticSiteResult, error) {
			args := v.(LookupStaticSiteArgs)
			r, err := LookupStaticSite(ctx, &args, opts...)
			var s LookupStaticSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStaticSiteResultOutput)
}

type LookupStaticSiteOutputArgs struct {
	// The ID of the service
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (LookupStaticSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteArgs)(nil)).Elem()
}

type LookupStaticSiteResultOutput struct{ *pulumi.OutputState }

func (LookupStaticSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteResult)(nil)).Elem()
}

func (o LookupStaticSiteResultOutput) ToLookupStaticSiteResultOutput() LookupStaticSiteResultOutput {
	return o
}

func (o LookupStaticSiteResultOutput) ToLookupStaticSiteResultOutputWithContext(ctx context.Context) LookupStaticSiteResultOutput {
	return o
}

func (o LookupStaticSiteResultOutput) Items() GetStaticSiteTypeOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) GetStaticSiteType { return v.Items }).(GetStaticSiteTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticSiteResultOutput{})
}
