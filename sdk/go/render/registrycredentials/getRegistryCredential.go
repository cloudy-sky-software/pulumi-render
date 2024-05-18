// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package registrycredentials

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistryCredential(ctx *pulumi.Context, args *LookupRegistryCredentialArgs, opts ...pulumi.InvokeOption) (*LookupRegistryCredentialResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRegistryCredentialResult
	err := ctx.Invoke("render:registrycredentials:getRegistryCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegistryCredentialArgs struct {
	// The ID of the registry credential
	RegistryCredentialId string `pulumi:"registryCredentialId"`
}

type LookupRegistryCredentialResult struct {
	Items RegistryCredentialType `pulumi:"items"`
}

func LookupRegistryCredentialOutput(ctx *pulumi.Context, args LookupRegistryCredentialOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryCredentialResult, error) {
			args := v.(LookupRegistryCredentialArgs)
			r, err := LookupRegistryCredential(ctx, &args, opts...)
			var s LookupRegistryCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryCredentialResultOutput)
}

type LookupRegistryCredentialOutputArgs struct {
	// The ID of the registry credential
	RegistryCredentialId pulumi.StringInput `pulumi:"registryCredentialId"`
}

func (LookupRegistryCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryCredentialArgs)(nil)).Elem()
}

type LookupRegistryCredentialResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryCredentialResult)(nil)).Elem()
}

func (o LookupRegistryCredentialResultOutput) ToLookupRegistryCredentialResultOutput() LookupRegistryCredentialResultOutput {
	return o
}

func (o LookupRegistryCredentialResultOutput) ToLookupRegistryCredentialResultOutputWithContext(ctx context.Context) LookupRegistryCredentialResultOutput {
	return o
}

func (o LookupRegistryCredentialResultOutput) Items() RegistryCredentialTypeOutput {
	return o.ApplyT(func(v LookupRegistryCredentialResult) RegistryCredentialType { return v.Items }).(RegistryCredentialTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryCredentialResultOutput{})
}