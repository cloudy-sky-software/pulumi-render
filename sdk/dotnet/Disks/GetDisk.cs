// *** WARNING: this file was generated by pulumigen. ***
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
        public static Task<Outputs.DiskWithCursorpropertiesdisk> InvokeAsync(GetDiskArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.DiskWithCursorpropertiesdisk>("render:disks:getDisk", args ?? new GetDiskArgs(), options.WithDefaults());

        public static Output<Outputs.DiskWithCursorpropertiesdisk> Invoke(GetDiskInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.DiskWithCursorpropertiesdisk>("render:disks:getDisk", args ?? new GetDiskInvokeArgs(), options.WithDefaults());
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
}
