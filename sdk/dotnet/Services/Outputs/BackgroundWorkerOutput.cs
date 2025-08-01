// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Outputs
{

    [OutputType]
    public sealed class BackgroundWorkerOutput
    {
        public readonly Pulumi.Render.Services.ServiceAutoDeploy AutoDeploy;
        public readonly string? Branch;
        public readonly Outputs.BuildFilter? BuildFilter;
        public readonly string CreatedAt;
        /// <summary>
        /// The URL to view the service in the Render Dashboard
        /// </summary>
        public readonly string DashboardUrl;
        public readonly string? EnvironmentId;
        public readonly string Id;
        public readonly string? ImagePath;
        public readonly string Name;
        public readonly Pulumi.Render.Services.ServiceNotifyOnFail NotifyOnFail;
        public readonly string OwnerId;
        public readonly Outputs.RegistryCredentialSummary? RegistryCredential;
        public readonly string? Repo;
        public readonly string RootDir;
        public readonly Outputs.BackgroundWorkerDetailsOutput? ServiceDetails;
        public readonly string Slug;
        public readonly Pulumi.Render.Services.ServiceSuspended Suspended;
        public readonly ImmutableArray<Pulumi.Render.Services.ServiceSuspendersItem> Suspenders;
        public readonly string? Type;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private BackgroundWorkerOutput(
            Pulumi.Render.Services.ServiceAutoDeploy autoDeploy,

            string? branch,

            Outputs.BuildFilter? buildFilter,

            string createdAt,

            string dashboardUrl,

            string? environmentId,

            string id,

            string? imagePath,

            string name,

            Pulumi.Render.Services.ServiceNotifyOnFail notifyOnFail,

            string ownerId,

            Outputs.RegistryCredentialSummary? registryCredential,

            string? repo,

            string rootDir,

            Outputs.BackgroundWorkerDetailsOutput? serviceDetails,

            string slug,

            Pulumi.Render.Services.ServiceSuspended suspended,

            ImmutableArray<Pulumi.Render.Services.ServiceSuspendersItem> suspenders,

            string? type,

            string updatedAt)
        {
            AutoDeploy = autoDeploy;
            Branch = branch;
            BuildFilter = buildFilter;
            CreatedAt = createdAt;
            DashboardUrl = dashboardUrl;
            EnvironmentId = environmentId;
            Id = id;
            ImagePath = imagePath;
            Name = name;
            NotifyOnFail = notifyOnFail;
            OwnerId = ownerId;
            RegistryCredential = registryCredential;
            Repo = repo;
            RootDir = rootDir;
            ServiceDetails = serviceDetails;
            Slug = slug;
            Suspended = suspended;
            Suspenders = suspenders;
            Type = type;
            UpdatedAt = updatedAt;
        }
    }
}
