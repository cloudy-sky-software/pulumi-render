// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Redis
{
    public static class GetRedis
    {
        public static Task<GetRedisResult> InvokeAsync(GetRedisArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRedisResult>("render:redis:getRedis", args ?? new GetRedisArgs(), options.WithDefaults());

        public static Output<GetRedisResult> Invoke(GetRedisInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRedisResult>("render:redis:getRedis", args ?? new GetRedisInvokeArgs(), options.WithDefaults());

        public static Output<GetRedisResult> Invoke(GetRedisInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRedisResult>("render:redis:getRedis", args ?? new GetRedisInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRedisArgs : global::Pulumi.InvokeArgs
    {
        [Input("redisId", required: true)]
        public string RedisId { get; set; } = null!;

        public GetRedisArgs()
        {
        }
        public static new GetRedisArgs Empty => new GetRedisArgs();
    }

    public sealed class GetRedisInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("redisId", required: true)]
        public Input<string> RedisId { get; set; } = null!;

        public GetRedisInvokeArgs()
        {
        }
        public static new GetRedisInvokeArgs Empty => new GetRedisInvokeArgs();
    }


    [OutputType]
    public sealed class GetRedisResult
    {
        /// <summary>
        /// The creation time of the Redis instance
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The ID of the environment the Redis instance is associated with
        /// </summary>
        public readonly string? EnvironmentId;
        /// <summary>
        /// The ID of the Redis instance
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The IP allow list for the Redis instance
        /// </summary>
        public readonly ImmutableArray<Outputs.CidrBlockAndDescription> IpAllowList;
        public readonly Outputs.RedisDetailMaintenanceProperties? Maintenance;
        /// <summary>
        /// The name of the Redis instance
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Options for a Redis instance
        /// </summary>
        public readonly Outputs.RedisOptions Options;
        public readonly Outputs.Owner Owner;
        public readonly Pulumi.Render.Redis.RedisDetailPlan Plan;
        /// <summary>
        /// Defaults to "oregon"
        /// </summary>
        public readonly Pulumi.Render.Redis.RedisDetailRegion Region;
        public readonly Pulumi.Render.Redis.RedisDetailStatus Status;
        /// <summary>
        /// The last updated time of the Redis instance
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// The version of Redis
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetRedisResult(
            string createdAt,

            string? environmentId,

            string id,

            ImmutableArray<Outputs.CidrBlockAndDescription> ipAllowList,

            Outputs.RedisDetailMaintenanceProperties? maintenance,

            string name,

            Outputs.RedisOptions options,

            Outputs.Owner owner,

            Pulumi.Render.Redis.RedisDetailPlan plan,

            Pulumi.Render.Redis.RedisDetailRegion region,

            Pulumi.Render.Redis.RedisDetailStatus status,

            string updatedAt,

            string version)
        {
            CreatedAt = createdAt;
            EnvironmentId = environmentId;
            Id = id;
            IpAllowList = ipAllowList;
            Maintenance = maintenance;
            Name = name;
            Options = options;
            Owner = owner;
            Plan = plan;
            Region = region;
            Status = status;
            UpdatedAt = updatedAt;
            Version = version;
        }
    }
}
