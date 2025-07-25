// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class GetWebService
    {
        public static Task<GetWebServiceResult> InvokeAsync(GetWebServiceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebServiceResult>("render:services:getWebService", args ?? new GetWebServiceArgs(), options.WithDefaults());

        public static Output<GetWebServiceResult> Invoke(GetWebServiceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebServiceResult>("render:services:getWebService", args ?? new GetWebServiceInvokeArgs(), options.WithDefaults());

        public static Output<GetWebServiceResult> Invoke(GetWebServiceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebServiceResult>("render:services:getWebService", args ?? new GetWebServiceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebServiceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetWebServiceArgs()
        {
        }
        public static new GetWebServiceArgs Empty => new GetWebServiceArgs();
    }

    public sealed class GetWebServiceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetWebServiceInvokeArgs()
        {
        }
        public static new GetWebServiceInvokeArgs Empty => new GetWebServiceInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebServiceResult
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
        public readonly Outputs.WebServiceDetailsOutput? ServiceDetails;
        public readonly string Slug;
        public readonly Pulumi.Render.Services.ServiceSuspended Suspended;
        public readonly ImmutableArray<Pulumi.Render.Services.ServiceSuspendersItem> Suspenders;
        public readonly string? Type;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetWebServiceResult(
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

            Outputs.WebServiceDetailsOutput? serviceDetails,

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
