// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListCustomDomains(ctx *pulumi.Context, args *ListCustomDomainsArgs, opts ...pulumi.InvokeOption) (*ListCustomDomainsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv ListCustomDomainsResult
	err := ctx.Invoke("render:services:listCustomDomains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListCustomDomainsArgs struct {
	// (Required) The ID of the service
	ServiceId string `pulumi:"serviceId"`
}

type ListCustomDomainsResult struct {
}

func ListCustomDomainsOutput(ctx *pulumi.Context, args ListCustomDomainsOutputArgs, opts ...pulumi.InvokeOption) ListCustomDomainsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListCustomDomainsResult, error) {
			args := v.(ListCustomDomainsArgs)
			r, err := ListCustomDomains(ctx, &args, opts...)
			var s ListCustomDomainsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListCustomDomainsResultOutput)
}

type ListCustomDomainsOutputArgs struct {
	// (Required) The ID of the service
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (ListCustomDomainsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCustomDomainsArgs)(nil)).Elem()
}

type ListCustomDomainsResultOutput struct{ *pulumi.OutputState }

func (ListCustomDomainsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCustomDomainsResult)(nil)).Elem()
}

func (o ListCustomDomainsResultOutput) ToListCustomDomainsResultOutput() ListCustomDomainsResultOutput {
	return o
}

func (o ListCustomDomainsResultOutput) ToListCustomDomainsResultOutputWithContext(ctx context.Context) ListCustomDomainsResultOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ListCustomDomainsResultOutput{})
}