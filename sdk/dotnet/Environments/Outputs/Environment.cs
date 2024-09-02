// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Environments.Outputs
{

    [OutputType]
    public sealed class Environment
    {
        public readonly ImmutableArray<string> DatabasesIds;
        public readonly ImmutableArray<string> EnvGroupIds;
        public readonly string Id;
        public readonly string Name;
        public readonly string ProjectId;
        /// <summary>
        /// Indicates whether an environment is `unprotected` or `protected`. Only admin users can perform destructive actions in `protected` environments.
        /// </summary>
        public readonly Pulumi.Render.Environments.EnvironmentProtectedStatus ProtectedStatus;
        public readonly ImmutableArray<string> RedisIds;
        public readonly ImmutableArray<string> ServiceIds;

        [OutputConstructor]
        private Environment(
            ImmutableArray<string> databasesIds,

            ImmutableArray<string> envGroupIds,

            string id,

            string name,

            string projectId,

            Pulumi.Render.Environments.EnvironmentProtectedStatus protectedStatus,

            ImmutableArray<string> redisIds,

            ImmutableArray<string> serviceIds)
        {
            DatabasesIds = databasesIds;
            EnvGroupIds = envGroupIds;
            Id = id;
            Name = name;
            ProjectId = projectId;
            ProtectedStatus = protectedStatus;
            RedisIds = redisIds;
            ServiceIds = serviceIds;
        }
    }
}
