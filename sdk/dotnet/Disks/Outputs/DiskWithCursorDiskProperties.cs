// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Disks.Outputs
{

    [OutputType]
    public sealed class DiskWithCursorDiskProperties
    {
        public readonly string CreatedAt;
        public readonly string Id;
        public readonly string MountPath;
        public readonly string Name;
        public readonly string? ServiceId;
        public readonly int SizeGB;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private DiskWithCursorDiskProperties(
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