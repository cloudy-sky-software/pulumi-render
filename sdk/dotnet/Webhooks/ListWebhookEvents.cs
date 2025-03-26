// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Webhooks
{
    public static class ListWebhookEvents
    {
        public static Task<ListWebhookEventsResult> InvokeAsync(ListWebhookEventsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListWebhookEventsResult>("render:webhooks:listWebhookEvents", args ?? new ListWebhookEventsArgs(), options.WithDefaults());

        public static Output<ListWebhookEventsResult> Invoke(ListWebhookEventsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListWebhookEventsResult>("render:webhooks:listWebhookEvents", args ?? new ListWebhookEventsInvokeArgs(), options.WithDefaults());

        public static Output<ListWebhookEventsResult> Invoke(ListWebhookEventsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListWebhookEventsResult>("render:webhooks:listWebhookEvents", args ?? new ListWebhookEventsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListWebhookEventsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for the webhook
        /// </summary>
        [Input("webhookId", required: true)]
        public string WebhookId { get; set; } = null!;

        public ListWebhookEventsArgs()
        {
        }
        public static new ListWebhookEventsArgs Empty => new ListWebhookEventsArgs();
    }

    public sealed class ListWebhookEventsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for the webhook
        /// </summary>
        [Input("webhookId", required: true)]
        public Input<string> WebhookId { get; set; } = null!;

        public ListWebhookEventsInvokeArgs()
        {
        }
        public static new ListWebhookEventsInvokeArgs Empty => new ListWebhookEventsInvokeArgs();
    }


    [OutputType]
    public sealed class ListWebhookEventsResult
    {
        public readonly ImmutableArray<Outputs.WebhookEventWithCursor> Items;

        [OutputConstructor]
        private ListWebhookEventsResult(ImmutableArray<Outputs.WebhookEventWithCursor> items)
        {
            Items = items;
        }
    }
}
