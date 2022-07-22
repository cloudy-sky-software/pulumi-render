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

    /// <summary>
    /// A service header object
    /// </summary>
    public sealed class ServiceServiceHeaderArgs : Pulumi.ResourceArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ServiceServiceHeaderArgs()
        {
        }
    }
}