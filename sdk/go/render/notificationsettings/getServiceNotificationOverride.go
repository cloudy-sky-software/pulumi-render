// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package notificationsettings

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetServiceNotificationOverride(ctx *pulumi.Context, args *GetServiceNotificationOverrideArgs, opts ...pulumi.InvokeOption) (*GetServiceNotificationOverrideResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServiceNotificationOverrideResult
	err := ctx.Invoke("render:notification-settings:getServiceNotificationOverride", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetServiceNotificationOverrideArgs struct {
	// The ID of the service
	ServiceId string `pulumi:"serviceId"`
}

type GetServiceNotificationOverrideResult struct {
	NotificationsToSend         GetServiceNotificationOverridePropertiesNotificationsToSend         `pulumi:"notificationsToSend"`
	PreviewNotificationsEnabled GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled `pulumi:"previewNotificationsEnabled"`
	ServiceId                   string                                                              `pulumi:"serviceId"`
}

func GetServiceNotificationOverrideOutput(ctx *pulumi.Context, args GetServiceNotificationOverrideOutputArgs, opts ...pulumi.InvokeOption) GetServiceNotificationOverrideResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServiceNotificationOverrideResult, error) {
			args := v.(GetServiceNotificationOverrideArgs)
			r, err := GetServiceNotificationOverride(ctx, &args, opts...)
			var s GetServiceNotificationOverrideResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServiceNotificationOverrideResultOutput)
}

type GetServiceNotificationOverrideOutputArgs struct {
	// The ID of the service
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (GetServiceNotificationOverrideOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceNotificationOverrideArgs)(nil)).Elem()
}

type GetServiceNotificationOverrideResultOutput struct{ *pulumi.OutputState }

func (GetServiceNotificationOverrideResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceNotificationOverrideResult)(nil)).Elem()
}

func (o GetServiceNotificationOverrideResultOutput) ToGetServiceNotificationOverrideResultOutput() GetServiceNotificationOverrideResultOutput {
	return o
}

func (o GetServiceNotificationOverrideResultOutput) ToGetServiceNotificationOverrideResultOutputWithContext(ctx context.Context) GetServiceNotificationOverrideResultOutput {
	return o
}

func (o GetServiceNotificationOverrideResultOutput) NotificationsToSend() GetServiceNotificationOverridePropertiesNotificationsToSendOutput {
	return o.ApplyT(func(v GetServiceNotificationOverrideResult) GetServiceNotificationOverridePropertiesNotificationsToSend {
		return v.NotificationsToSend
	}).(GetServiceNotificationOverridePropertiesNotificationsToSendOutput)
}

func (o GetServiceNotificationOverrideResultOutput) PreviewNotificationsEnabled() GetServiceNotificationOverridePropertiesPreviewNotificationsEnabledOutput {
	return o.ApplyT(func(v GetServiceNotificationOverrideResult) GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled {
		return v.PreviewNotificationsEnabled
	}).(GetServiceNotificationOverridePropertiesPreviewNotificationsEnabledOutput)
}

func (o GetServiceNotificationOverrideResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceNotificationOverrideResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceNotificationOverrideResultOutput{})
}
