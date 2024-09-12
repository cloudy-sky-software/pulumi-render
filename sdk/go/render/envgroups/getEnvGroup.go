// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package envgroups

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnvGroup(ctx *pulumi.Context, args *LookupEnvGroupArgs, opts ...pulumi.InvokeOption) (*LookupEnvGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEnvGroupResult
	err := ctx.Invoke("render:env-groups:getEnvGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvGroupArgs struct {
	// Filter for resources that belong to an environment group
	EnvGroupId string `pulumi:"envGroupId"`
}

type LookupEnvGroupResult struct {
	CreatedAt     string       `pulumi:"createdAt"`
	EnvVars       []EnvVar     `pulumi:"envVars"`
	EnvironmentId *string      `pulumi:"environmentId"`
	Id            string       `pulumi:"id"`
	Name          string       `pulumi:"name"`
	OwnerId       string       `pulumi:"ownerId"`
	SecretFiles   []SecretFile `pulumi:"secretFiles"`
	// List of serviceIds linked to the envGroup
	ServiceLinks []ServiceLink `pulumi:"serviceLinks"`
	UpdatedAt    string        `pulumi:"updatedAt"`
}

func LookupEnvGroupOutput(ctx *pulumi.Context, args LookupEnvGroupOutputArgs, opts ...pulumi.InvokeOption) LookupEnvGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvGroupResultOutput, error) {
			args := v.(LookupEnvGroupArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupEnvGroupResult
			secret, err := ctx.InvokePackageRaw("render:env-groups:getEnvGroup", args, &rv, "", opts...)
			if err != nil {
				return LookupEnvGroupResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupEnvGroupResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupEnvGroupResultOutput), nil
			}
			return output, nil
		}).(LookupEnvGroupResultOutput)
}

type LookupEnvGroupOutputArgs struct {
	// Filter for resources that belong to an environment group
	EnvGroupId pulumi.StringInput `pulumi:"envGroupId"`
}

func (LookupEnvGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvGroupArgs)(nil)).Elem()
}

type LookupEnvGroupResultOutput struct{ *pulumi.OutputState }

func (LookupEnvGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvGroupResult)(nil)).Elem()
}

func (o LookupEnvGroupResultOutput) ToLookupEnvGroupResultOutput() LookupEnvGroupResultOutput {
	return o
}

func (o LookupEnvGroupResultOutput) ToLookupEnvGroupResultOutputWithContext(ctx context.Context) LookupEnvGroupResultOutput {
	return o
}

func (o LookupEnvGroupResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvGroupResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupEnvGroupResultOutput) EnvVars() EnvVarArrayOutput {
	return o.ApplyT(func(v LookupEnvGroupResult) []EnvVar { return v.EnvVars }).(EnvVarArrayOutput)
}

func (o LookupEnvGroupResultOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvGroupResult) *string { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

func (o LookupEnvGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEnvGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEnvGroupResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvGroupResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

func (o LookupEnvGroupResultOutput) SecretFiles() SecretFileArrayOutput {
	return o.ApplyT(func(v LookupEnvGroupResult) []SecretFile { return v.SecretFiles }).(SecretFileArrayOutput)
}

// List of serviceIds linked to the envGroup
func (o LookupEnvGroupResultOutput) ServiceLinks() ServiceLinkArrayOutput {
	return o.ApplyT(func(v LookupEnvGroupResult) []ServiceLink { return v.ServiceLinks }).(ServiceLinkArrayOutput)
}

func (o LookupEnvGroupResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvGroupResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvGroupResultOutput{})
}
