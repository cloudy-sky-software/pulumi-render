// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package webhooks

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-render/sdk/go/render/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebhookEvents(ctx *pulumi.Context, args *ListWebhookEventsArgs, opts ...pulumi.InvokeOption) (*ListWebhookEventsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListWebhookEventsResult
	err := ctx.Invoke("render:webhooks:listWebhookEvents", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebhookEventsArgs struct {
	// Unique identifier for the webhook
	WebhookId string `pulumi:"webhookId"`
}

type ListWebhookEventsResult struct {
	Items []WebhookEventWithCursor `pulumi:"items"`
}

func ListWebhookEventsOutput(ctx *pulumi.Context, args ListWebhookEventsOutputArgs, opts ...pulumi.InvokeOption) ListWebhookEventsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListWebhookEventsResultOutput, error) {
			args := v.(ListWebhookEventsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("render:webhooks:listWebhookEvents", args, ListWebhookEventsResultOutput{}, options).(ListWebhookEventsResultOutput), nil
		}).(ListWebhookEventsResultOutput)
}

type ListWebhookEventsOutputArgs struct {
	// Unique identifier for the webhook
	WebhookId pulumi.StringInput `pulumi:"webhookId"`
}

func (ListWebhookEventsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebhookEventsArgs)(nil)).Elem()
}

type ListWebhookEventsResultOutput struct{ *pulumi.OutputState }

func (ListWebhookEventsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebhookEventsResult)(nil)).Elem()
}

func (o ListWebhookEventsResultOutput) ToListWebhookEventsResultOutput() ListWebhookEventsResultOutput {
	return o
}

func (o ListWebhookEventsResultOutput) ToListWebhookEventsResultOutputWithContext(ctx context.Context) ListWebhookEventsResultOutput {
	return o
}

func (o ListWebhookEventsResultOutput) Items() WebhookEventWithCursorArrayOutput {
	return o.ApplyT(func(v ListWebhookEventsResult) []WebhookEventWithCursor { return v.Items }).(WebhookEventWithCursorArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebhookEventsResultOutput{})
}
