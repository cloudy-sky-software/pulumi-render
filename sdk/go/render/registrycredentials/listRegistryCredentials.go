// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package registrycredentials

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRegistryCredentials(ctx *pulumi.Context, args *ListRegistryCredentialsArgs, opts ...pulumi.InvokeOption) (*ListRegistryCredentialsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListRegistryCredentialsResult
	err := ctx.Invoke("render:registrycredentials:listRegistryCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRegistryCredentialsArgs struct {
}

type ListRegistryCredentialsResult struct {
	Items []RegistryCredentialType `pulumi:"items"`
}

func ListRegistryCredentialsOutput(ctx *pulumi.Context, args ListRegistryCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListRegistryCredentialsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListRegistryCredentialsResultOutput, error) {
			args := v.(ListRegistryCredentialsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("render:registrycredentials:listRegistryCredentials", args, ListRegistryCredentialsResultOutput{}, options).(ListRegistryCredentialsResultOutput), nil
		}).(ListRegistryCredentialsResultOutput)
}

type ListRegistryCredentialsOutputArgs struct {
}

func (ListRegistryCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRegistryCredentialsArgs)(nil)).Elem()
}

type ListRegistryCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListRegistryCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRegistryCredentialsResult)(nil)).Elem()
}

func (o ListRegistryCredentialsResultOutput) ToListRegistryCredentialsResultOutput() ListRegistryCredentialsResultOutput {
	return o
}

func (o ListRegistryCredentialsResultOutput) ToListRegistryCredentialsResultOutputWithContext(ctx context.Context) ListRegistryCredentialsResultOutput {
	return o
}

func (o ListRegistryCredentialsResultOutput) Items() RegistryCredentialTypeArrayOutput {
	return o.ApplyT(func(v ListRegistryCredentialsResult) []RegistryCredentialType { return v.Items }).(RegistryCredentialTypeArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRegistryCredentialsResultOutput{})
}
