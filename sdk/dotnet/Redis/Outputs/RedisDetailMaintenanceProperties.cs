// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Redis.Outputs
{

    [OutputType]
    public sealed class RedisDetailMaintenanceProperties
    {
        public readonly string Id;
        /// <summary>
        /// If present, the maintenance run cannot be scheduled for later than this date-time.
        /// </summary>
        public readonly string? PendingMaintenanceBy;
        public readonly string ScheduledAt;
        public readonly Pulumi.Render.Redis.RedisDetailMaintenancePropertiesState State;
        public readonly string Type;

        [OutputConstructor]
        private RedisDetailMaintenanceProperties(
            string id,

            string? pendingMaintenanceBy,

            string scheduledAt,

            Pulumi.Render.Redis.RedisDetailMaintenancePropertiesState state,

            string type)
        {
            Id = id;
            PendingMaintenanceBy = pendingMaintenanceBy;
            ScheduledAt = scheduledAt;
            State = state;
            Type = type;
        }
    }
}
