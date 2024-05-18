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
    public sealed class GetBackgroundWorker
    {
        public readonly Pulumi.Render.Services.PreviewServiceServiceAutoDeploy AutoDeploy;
        public readonly string? Branch;
        public readonly Outputs.BuildFilter? BuildFilter;
        public readonly string CreatedAt;
        public readonly string Id;
        public readonly string? ImagePath;
        public readonly string Name;
        public readonly Pulumi.Render.Services.PreviewServiceServiceNotifyOnFail NotifyOnFail;
        public readonly string OwnerId;
        public readonly string? Repo;
        public readonly string RootDir;
        public readonly Outputs.BackgroundWorkerDetailsOutput? ServiceDetails;
        public readonly string Slug;
        public readonly Pulumi.Render.Services.PreviewServiceServiceSuspended Suspended;
        public readonly ImmutableArray<Pulumi.Render.Services.PreviewServiceServiceSuspendersItem> Suspenders;
        public readonly string? Type;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetBackgroundWorker(
            Pulumi.Render.Services.PreviewServiceServiceAutoDeploy autoDeploy,

            string? branch,

            Outputs.BuildFilter? buildFilter,

            string createdAt,

            string id,

            string? imagePath,

            string name,

            Pulumi.Render.Services.PreviewServiceServiceNotifyOnFail notifyOnFail,

            string ownerId,

            string? repo,

            string rootDir,

            Outputs.BackgroundWorkerDetailsOutput? serviceDetails,

            string slug,

            Pulumi.Render.Services.PreviewServiceServiceSuspended suspended,

            ImmutableArray<Pulumi.Render.Services.PreviewServiceServiceSuspendersItem> suspenders,

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
