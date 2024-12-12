// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.NotificationSettings
{
    public static class GetServiceNotificationOverride
    {
        public static Task<Outputs.GetServiceNotificationOverrideProperties> InvokeAsync(GetServiceNotificationOverrideArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetServiceNotificationOverrideProperties>("render:notification-settings:getServiceNotificationOverride", args ?? new GetServiceNotificationOverrideArgs(), options.WithDefaults());

        public static Output<Outputs.GetServiceNotificationOverrideProperties> Invoke(GetServiceNotificationOverrideInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetServiceNotificationOverrideProperties>("render:notification-settings:getServiceNotificationOverride", args ?? new GetServiceNotificationOverrideInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.GetServiceNotificationOverrideProperties> Invoke(GetServiceNotificationOverrideInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetServiceNotificationOverrideProperties>("render:notification-settings:getServiceNotificationOverride", args ?? new GetServiceNotificationOverrideInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceNotificationOverrideArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetServiceNotificationOverrideArgs()
        {
        }
        public static new GetServiceNotificationOverrideArgs Empty => new GetServiceNotificationOverrideArgs();
    }

    public sealed class GetServiceNotificationOverrideInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetServiceNotificationOverrideInvokeArgs()
        {
        }
        public static new GetServiceNotificationOverrideInvokeArgs Empty => new GetServiceNotificationOverrideInvokeArgs();
    }
}
