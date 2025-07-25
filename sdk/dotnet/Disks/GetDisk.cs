// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Disks
{
    public static class GetDisk
    {
        public static Task<GetDiskResult> InvokeAsync(GetDiskArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDiskResult>("render:disks:getDisk", args ?? new GetDiskArgs(), options.WithDefaults());

        public static Output<GetDiskResult> Invoke(GetDiskInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDiskResult>("render:disks:getDisk", args ?? new GetDiskInvokeArgs(), options.WithDefaults());

        public static Output<GetDiskResult> Invoke(GetDiskInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDiskResult>("render:disks:getDisk", args ?? new GetDiskInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDiskArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the disk
        /// </summary>
        [Input("diskId", required: true)]
        public string DiskId { get; set; } = null!;

        public GetDiskArgs()
        {
        }
        public static new GetDiskArgs Empty => new GetDiskArgs();
    }

    public sealed class GetDiskInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the disk
        /// </summary>
        [Input("diskId", required: true)]
        public Input<string> DiskId { get; set; } = null!;

        public GetDiskInvokeArgs()
        {
        }
        public static new GetDiskInvokeArgs Empty => new GetDiskInvokeArgs();
    }


    [OutputType]
    public sealed class GetDiskResult
    {
        public readonly string CreatedAt;
        public readonly string Id;
        public readonly string MountPath;
        public readonly string Name;
        public readonly string? ServiceId;
        public readonly int SizeGB;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetDiskResult(
            string createdAt,

            string id,

            string mountPath,

            string name,

            string? serviceId,

            int sizeGB,

            string updatedAt)
        {
            CreatedAt = createdAt;
            Id = id;
            MountPath = mountPath;
            Name = name;
            ServiceId = serviceId;
            SizeGB = sizeGB;
            UpdatedAt = updatedAt;
        }
    }
}
