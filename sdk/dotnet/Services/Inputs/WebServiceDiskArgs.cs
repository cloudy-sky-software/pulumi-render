// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Render.Services.Inputs
{

    public sealed class WebServiceDiskArgs : Pulumi.ResourceArgs
    {
        [Input("mountPath", required: true)]
        public Input<string> MountPath { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("sizeGB")]
        public Input<double>? SizeGB { get; set; }

        public WebServiceDiskArgs()
        {
            SizeGB = 1;
        }
    }
}
