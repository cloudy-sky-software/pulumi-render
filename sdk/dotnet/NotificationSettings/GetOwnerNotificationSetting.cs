// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.NotificationSettings
{
    public static class GetOwnerNotificationSetting
    {
        public static Task<Outputs.GetOwnerNotificationSettingProperties> InvokeAsync(GetOwnerNotificationSettingArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetOwnerNotificationSettingProperties>("render:notification-settings:getOwnerNotificationSetting", args ?? new GetOwnerNotificationSettingArgs(), options.WithDefaults());

        public static Output<Outputs.GetOwnerNotificationSettingProperties> Invoke(GetOwnerNotificationSettingInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetOwnerNotificationSettingProperties>("render:notification-settings:getOwnerNotificationSetting", args ?? new GetOwnerNotificationSettingInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.GetOwnerNotificationSettingProperties> Invoke(GetOwnerNotificationSettingInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetOwnerNotificationSettingProperties>("render:notification-settings:getOwnerNotificationSetting", args ?? new GetOwnerNotificationSettingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOwnerNotificationSettingArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the owner (team or personal user) whose resources should be returned
        /// </summary>
        [Input("ownerId", required: true)]
        public string OwnerId { get; set; } = null!;

        public GetOwnerNotificationSettingArgs()
        {
        }
        public static new GetOwnerNotificationSettingArgs Empty => new GetOwnerNotificationSettingArgs();
    }

    public sealed class GetOwnerNotificationSettingInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the owner (team or personal user) whose resources should be returned
        /// </summary>
        [Input("ownerId", required: true)]
        public Input<string> OwnerId { get; set; } = null!;

        public GetOwnerNotificationSettingInvokeArgs()
        {
        }
        public static new GetOwnerNotificationSettingInvokeArgs Empty => new GetOwnerNotificationSettingInvokeArgs();
    }
}
