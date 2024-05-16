// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Inputs
{

    public sealed class WebServiceDetailsCreateDiskPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("mountPath", required: true)]
        public Input<string> MountPath { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Defaults to 1
        /// </summary>
        [Input("sizeGB")]
        public Input<int>? SizeGB { get; set; }

        public WebServiceDetailsCreateDiskPropertiesArgs()
        {
        }
        public static new WebServiceDetailsCreateDiskPropertiesArgs Empty => new WebServiceDetailsCreateDiskPropertiesArgs();
    }
}
