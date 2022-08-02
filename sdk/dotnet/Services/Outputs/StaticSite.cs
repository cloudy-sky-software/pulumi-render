// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Outputs
{

    /// <summary>
    /// A static website service
    /// </summary>
    [OutputType]
    public sealed class StaticSite
    {
        /// <summary>
        /// Whether to auto deploy the service or not upon git push.
        /// </summary>
        public readonly Pulumi.Render.Services.ServiceAutoDeploy? AutoDeploy;
        /// <summary>
        /// If left empty, this will fall back to the default branch of the repository.
        /// </summary>
        public readonly string? Branch;
        public readonly string? CreatedAt;
        public readonly ImmutableArray<Outputs.EnvVarKeyValueOrGenerateValue> EnvVars;
        public readonly string Name;
        /// <summary>
        /// The notification setting for this service upon deployment failure.
        /// </summary>
        public readonly Pulumi.Render.Services.ServiceNotifyOnFail? NotifyOnFail;
        /// <summary>
        /// The id of the owner (user/team).
        /// </summary>
        public readonly string OwnerId;
        /// <summary>
        /// Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
        /// </summary>
        public readonly string Repo;
        public readonly ImmutableArray<Outputs.SecretFile> SecretFiles;
        public readonly Outputs.StaticSiteServiceDetails? ServiceDetails;
        public readonly string? Slug;
        public readonly Pulumi.Render.Services.ServiceSuspended? Suspended;
        public readonly ImmutableArray<string> Suspenders;
        public readonly string? Type;
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private StaticSite(
            Pulumi.Render.Services.ServiceAutoDeploy? autoDeploy,

            string? branch,

            string? createdAt,

            ImmutableArray<Outputs.EnvVarKeyValueOrGenerateValue> envVars,

            string name,

            Pulumi.Render.Services.ServiceNotifyOnFail? notifyOnFail,

            string ownerId,

            string repo,

            ImmutableArray<Outputs.SecretFile> secretFiles,

            Outputs.StaticSiteServiceDetails? serviceDetails,

            string? slug,

            Pulumi.Render.Services.ServiceSuspended? suspended,

            ImmutableArray<string> suspenders,

            string? type,

            string? updatedAt)
        {
            AutoDeploy = autoDeploy;
            Branch = branch;
            CreatedAt = createdAt;
            EnvVars = envVars;
            Name = name;
            NotifyOnFail = notifyOnFail;
            OwnerId = ownerId;
            Repo = repo;
            SecretFiles = secretFiles;
            ServiceDetails = serviceDetails;
            Slug = slug;
            Suspended = suspended;
            Suspenders = suspenders;
            Type = type;
            UpdatedAt = updatedAt;
        }
    }
}
