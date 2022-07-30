// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomDomain(ctx *pulumi.Context, args *LookupCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupCustomDomainResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupCustomDomainResult
	err := ctx.Invoke("render:services:getCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomDomainArgs struct {
	// (Required) The ID or name of the custom domain
	Id string `pulumi:"id"`
	// (Required) The ID of the service
	ServiceId string `pulumi:"serviceId"`
}

type LookupCustomDomainResult struct {
	Items CustomDomainType `pulumi:"items"`
}

func LookupCustomDomainOutput(ctx *pulumi.Context, args LookupCustomDomainOutputArgs, opts ...pulumi.InvokeOption) LookupCustomDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomDomainResult, error) {
			args := v.(LookupCustomDomainArgs)
			r, err := LookupCustomDomain(ctx, &args, opts...)
			var s LookupCustomDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomDomainResultOutput)
}

type LookupCustomDomainOutputArgs struct {
	// (Required) The ID or name of the custom domain
	Id pulumi.StringInput `pulumi:"id"`
	// (Required) The ID of the service
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (LookupCustomDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomDomainArgs)(nil)).Elem()
}

type LookupCustomDomainResultOutput struct{ *pulumi.OutputState }

func (LookupCustomDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomDomainResult)(nil)).Elem()
}

func (o LookupCustomDomainResultOutput) ToLookupCustomDomainResultOutput() LookupCustomDomainResultOutput {
	return o
}

func (o LookupCustomDomainResultOutput) ToLookupCustomDomainResultOutputWithContext(ctx context.Context) LookupCustomDomainResultOutput {
	return o
}

func (o LookupCustomDomainResultOutput) Items() CustomDomainTypeOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) CustomDomainType { return v.Items }).(CustomDomainTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomDomainResultOutput{})
}