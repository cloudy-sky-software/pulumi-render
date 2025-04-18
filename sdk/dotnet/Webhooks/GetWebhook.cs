// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Webhooks
{
    public static class GetWebhook
    {
        public static Task<Outputs.WebhookWithCursorpropertieswebhook> InvokeAsync(GetWebhookArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.WebhookWithCursorpropertieswebhook>("render:webhooks:getWebhook", args ?? new GetWebhookArgs(), options.WithDefaults());

        public static Output<Outputs.WebhookWithCursorpropertieswebhook> Invoke(GetWebhookInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.WebhookWithCursorpropertieswebhook>("render:webhooks:getWebhook", args ?? new GetWebhookInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.WebhookWithCursorpropertieswebhook> Invoke(GetWebhookInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.WebhookWithCursorpropertieswebhook>("render:webhooks:getWebhook", args ?? new GetWebhookInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebhookArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for the webhook
        /// </summary>
        [Input("webhookId", required: true)]
        public string WebhookId { get; set; } = null!;

        public GetWebhookArgs()
        {
        }
        public static new GetWebhookArgs Empty => new GetWebhookArgs();
    }

    public sealed class GetWebhookInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for the webhook
        /// </summary>
        [Input("webhookId", required: true)]
        public Input<string> WebhookId { get; set; } = null!;

        public GetWebhookInvokeArgs()
        {
        }
        public static new GetWebhookInvokeArgs Empty => new GetWebhookInvokeArgs();
    }
}
