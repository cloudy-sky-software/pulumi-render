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
    public sealed class Route
    {
        public readonly string Destination;
        public readonly string Id;
        /// <summary>
        /// Redirect and Rewrite Rules are applied in priority order starting at 0
        /// </summary>
        public readonly int Priority;
        public readonly string Source;
        public readonly Pulumi.Render.Services.ServiceRouteType Type;

        [OutputConstructor]
        private Route(
            string destination,

            string id,

            int priority,

            string source,

            Pulumi.Render.Services.ServiceRouteType type)
        {
            Destination = destination;
            Id = id;
            Priority = priority;
            Source = source;
            Type = type;
        }
    }
}
