// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Inputs
{

    /// <summary>
    /// A route object for a static site
    /// </summary>
    public sealed class ServiceStaticSiteRouteArgs : Pulumi.ResourceArgs
    {
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        [Input("type", required: true)]
        public Input<Pulumi.Render.Services.ServiceStaticSiteRouteType> Type { get; set; } = null!;

        public ServiceStaticSiteRouteArgs()
        {
        }
    }
}
