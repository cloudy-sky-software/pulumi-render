// Code generated by pulumi-language-go DO NOT EDIT.
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
	AutoDeploy  ServiceAutoDeploy `pulumi:"autoDeploy"`
	Branch      *string           `pulumi:"branch"`
	BuildFilter *BuildFilter      `pulumi:"buildFilter"`
	CreatedAt   string            `pulumi:"createdAt"`
	// The URL to view the service in the Render Dashboard
	DashboardUrl       string                     `pulumi:"dashboardUrl"`
	EnvironmentId      *string                    `pulumi:"environmentId"`
	Id                 string                     `pulumi:"id"`
	ImagePath          *string                    `pulumi:"imagePath"`
	Name               string                     `pulumi:"name"`
	NotifyOnFail       ServiceNotifyOnFail        `pulumi:"notifyOnFail"`
	OwnerId            string                     `pulumi:"ownerId"`
	RegistryCredential *RegistryCredentialSummary `pulumi:"registryCredential"`
	Repo               *string                    `pulumi:"repo"`
	RootDir            string                     `pulumi:"rootDir"`
	ServiceDetails     *StaticSiteDetailsOutput   `pulumi:"serviceDetails"`
	Slug               string                     `pulumi:"slug"`
	Suspended          ServiceSuspended           `pulumi:"suspended"`
	Suspenders         []ServiceSuspendersItem    `pulumi:"suspenders"`
	Type               *string                    `pulumi:"type"`
	UpdatedAt          string                     `pulumi:"updatedAt"`
}

// Defaults sets the appropriate defaults for LookupStaticSiteResult
func (val *LookupStaticSiteResult) Defaults() *LookupStaticSiteResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if internal.IsZero(tmp.AutoDeploy) {
		tmp.AutoDeploy = ServiceAutoDeploy("yes")
	}
	tmp.ServiceDetails = tmp.ServiceDetails.Defaults()

	if tmp.Type == nil {
		type_ := "static_site"
		tmp.Type = &type_
	}
	return &tmp
}
func LookupStaticSiteOutput(ctx *pulumi.Context, args LookupStaticSiteOutputArgs, opts ...pulumi.InvokeOption) LookupStaticSiteResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupStaticSiteResultOutput, error) {
			args := v.(LookupStaticSiteArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("render:services:getStaticSite", args, LookupStaticSiteResultOutput{}, options).(LookupStaticSiteResultOutput), nil
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

func (o LookupStaticSiteResultOutput) AutoDeploy() ServiceAutoDeployOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) ServiceAutoDeploy { return v.AutoDeploy }).(ServiceAutoDeployOutput)
}

func (o LookupStaticSiteResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteResultOutput) BuildFilter() BuildFilterPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *BuildFilter { return v.BuildFilter }).(BuildFilterPtrOutput)
}

func (o LookupStaticSiteResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The URL to view the service in the Render Dashboard
func (o LookupStaticSiteResultOutput) DashboardUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.DashboardUrl }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) ImagePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.ImagePath }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) NotifyOnFail() ServiceNotifyOnFailOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) ServiceNotifyOnFail { return v.NotifyOnFail }).(ServiceNotifyOnFailOutput)
}

func (o LookupStaticSiteResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) RegistryCredential() RegistryCredentialSummaryPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *RegistryCredentialSummary { return v.RegistryCredential }).(RegistryCredentialSummaryPtrOutput)
}

func (o LookupStaticSiteResultOutput) Repo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.Repo }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteResultOutput) RootDir() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.RootDir }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) ServiceDetails() StaticSiteDetailsOutputPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *StaticSiteDetailsOutput { return v.ServiceDetails }).(StaticSiteDetailsOutputPtrOutput)
}

func (o LookupStaticSiteResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.Slug }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) Suspended() ServiceSuspendedOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) ServiceSuspended { return v.Suspended }).(ServiceSuspendedOutput)
}

func (o LookupStaticSiteResultOutput) Suspenders() ServiceSuspendersItemArrayOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) []ServiceSuspendersItem { return v.Suspenders }).(ServiceSuspendersItemArrayOutput)
}

func (o LookupStaticSiteResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticSiteResultOutput{})
}
