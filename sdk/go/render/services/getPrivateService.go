// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateService(ctx *pulumi.Context, args *LookupPrivateServiceArgs, opts ...pulumi.InvokeOption) (*LookupPrivateServiceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateServiceResult
	err := ctx.Invoke("render:services:getPrivateService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPrivateServiceArgs struct {
	// The ID of the service
	ServiceId string `pulumi:"serviceId"`
}

type LookupPrivateServiceResult struct {
	AutoDeploy  ServiceAutoDeploy `pulumi:"autoDeploy"`
	Branch      *string           `pulumi:"branch"`
	BuildFilter *BuildFilter      `pulumi:"buildFilter"`
	CreatedAt   string            `pulumi:"createdAt"`
	// The URL to view the service in the Render Dashboard
	DashboardUrl       string                       `pulumi:"dashboardUrl"`
	EnvironmentId      *string                      `pulumi:"environmentId"`
	Id                 string                       `pulumi:"id"`
	ImagePath          *string                      `pulumi:"imagePath"`
	Name               string                       `pulumi:"name"`
	NotifyOnFail       ServiceNotifyOnFail          `pulumi:"notifyOnFail"`
	OwnerId            string                       `pulumi:"ownerId"`
	RegistryCredential *RegistryCredentialSummary   `pulumi:"registryCredential"`
	Repo               *string                      `pulumi:"repo"`
	RootDir            string                       `pulumi:"rootDir"`
	ServiceDetails     *PrivateServiceDetailsOutput `pulumi:"serviceDetails"`
	Slug               string                       `pulumi:"slug"`
	Suspended          ServiceSuspended             `pulumi:"suspended"`
	Suspenders         []ServiceSuspendersItem      `pulumi:"suspenders"`
	Type               *string                      `pulumi:"type"`
	UpdatedAt          string                       `pulumi:"updatedAt"`
}

// Defaults sets the appropriate defaults for LookupPrivateServiceResult
func (val *LookupPrivateServiceResult) Defaults() *LookupPrivateServiceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if internal.IsZero(tmp.AutoDeploy) {
		tmp.AutoDeploy = ServiceAutoDeploy("yes")
	}
	tmp.ServiceDetails = tmp.ServiceDetails.Defaults()

	if tmp.Type == nil {
		type_ := "private_service"
		tmp.Type = &type_
	}
	return &tmp
}

func LookupPrivateServiceOutput(ctx *pulumi.Context, args LookupPrivateServiceOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateServiceResultOutput, error) {
			args := v.(LookupPrivateServiceArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupPrivateServiceResult
			secret, err := ctx.InvokePackageRaw("render:services:getPrivateService", args, &rv, "", opts...)
			if err != nil {
				return LookupPrivateServiceResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupPrivateServiceResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupPrivateServiceResultOutput), nil
			}
			return output, nil
		}).(LookupPrivateServiceResultOutput)
}

type LookupPrivateServiceOutputArgs struct {
	// The ID of the service
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (LookupPrivateServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateServiceArgs)(nil)).Elem()
}

type LookupPrivateServiceResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateServiceResult)(nil)).Elem()
}

func (o LookupPrivateServiceResultOutput) ToLookupPrivateServiceResultOutput() LookupPrivateServiceResultOutput {
	return o
}

func (o LookupPrivateServiceResultOutput) ToLookupPrivateServiceResultOutputWithContext(ctx context.Context) LookupPrivateServiceResultOutput {
	return o
}

func (o LookupPrivateServiceResultOutput) AutoDeploy() ServiceAutoDeployOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) ServiceAutoDeploy { return v.AutoDeploy }).(ServiceAutoDeployOutput)
}

func (o LookupPrivateServiceResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateServiceResultOutput) BuildFilter() BuildFilterPtrOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) *BuildFilter { return v.BuildFilter }).(BuildFilterPtrOutput)
}

func (o LookupPrivateServiceResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The URL to view the service in the Render Dashboard
func (o LookupPrivateServiceResultOutput) DashboardUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) string { return v.DashboardUrl }).(pulumi.StringOutput)
}

func (o LookupPrivateServiceResultOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) *string { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateServiceResultOutput) ImagePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) *string { return v.ImagePath }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateServiceResultOutput) NotifyOnFail() ServiceNotifyOnFailOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) ServiceNotifyOnFail { return v.NotifyOnFail }).(ServiceNotifyOnFailOutput)
}

func (o LookupPrivateServiceResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

func (o LookupPrivateServiceResultOutput) RegistryCredential() RegistryCredentialSummaryPtrOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) *RegistryCredentialSummary { return v.RegistryCredential }).(RegistryCredentialSummaryPtrOutput)
}

func (o LookupPrivateServiceResultOutput) Repo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) *string { return v.Repo }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateServiceResultOutput) RootDir() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) string { return v.RootDir }).(pulumi.StringOutput)
}

func (o LookupPrivateServiceResultOutput) ServiceDetails() PrivateServiceDetailsOutputPtrOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) *PrivateServiceDetailsOutput { return v.ServiceDetails }).(PrivateServiceDetailsOutputPtrOutput)
}

func (o LookupPrivateServiceResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) string { return v.Slug }).(pulumi.StringOutput)
}

func (o LookupPrivateServiceResultOutput) Suspended() ServiceSuspendedOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) ServiceSuspended { return v.Suspended }).(ServiceSuspendedOutput)
}

func (o LookupPrivateServiceResultOutput) Suspenders() ServiceSuspendersItemArrayOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) []ServiceSuspendersItem { return v.Suspenders }).(ServiceSuspendersItemArrayOutput)
}

func (o LookupPrivateServiceResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateServiceResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateServiceResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateServiceResultOutput{})
}
