// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Outputs
{

    [OutputType]
    public sealed class StaticSiteOutput
    {
        public readonly Pulumi.Render.Services.ServiceAutoDeploy AutoDeploy;
        public readonly string? Branch;
        public readonly Outputs.BuildFilter? BuildFilter;
        public readonly string CreatedAt;
        public readonly string Id;
        public readonly string? ImagePath;
        public readonly string Name;
        public readonly Pulumi.Render.Services.ServiceNotifyOnFail NotifyOnFail;
        public readonly string OwnerId;
        public readonly string? Repo;
        public readonly string RootDir;
        public readonly Outputs.StaticSiteDetailsOutput? ServiceDetails;
        public readonly string Slug;
        public readonly Pulumi.Render.Services.ServiceSuspended Suspended;
        public readonly ImmutableArray<Pulumi.Render.Services.ServiceSuspendersItem> Suspenders;
        public readonly string? Type;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private StaticSiteOutput(
            Pulumi.Render.Services.ServiceAutoDeploy autoDeploy,

            string? branch,

            Outputs.BuildFilter? buildFilter,

            string createdAt,

            string id,

            string? imagePath,

            string name,

            Pulumi.Render.Services.ServiceNotifyOnFail notifyOnFail,

            string ownerId,

            string? repo,

            string rootDir,

            Outputs.StaticSiteDetailsOutput? serviceDetails,

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
            Id = id;
            ImagePath = imagePath;
            Name = name;
            NotifyOnFail = notifyOnFail;
            OwnerId = ownerId;
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
