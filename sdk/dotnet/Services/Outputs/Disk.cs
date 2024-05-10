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
    public sealed class Disk
    {
        public readonly string MountPath;
        public readonly string Name;
        public readonly int SizeGB;

        [OutputConstructor]
        private Disk(
            string mountPath,

            string name,

            int sizeGB)
        {
            MountPath = mountPath;
            Name = name;
            SizeGB = sizeGB;
        }
    }
}
