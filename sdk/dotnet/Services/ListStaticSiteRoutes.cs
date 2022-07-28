// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class ListStaticSiteRoutes
    {
        public static Task<ListStaticSiteRoutesResult> InvokeAsync(ListStaticSiteRoutesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListStaticSiteRoutesResult>("render:services:listStaticSiteRoutes", args ?? new ListStaticSiteRoutesArgs(), options.WithDefaults());

        public static Output<ListStaticSiteRoutesResult> Invoke(ListStaticSiteRoutesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<ListStaticSiteRoutesResult>("render:services:listStaticSiteRoutes", args ?? new ListStaticSiteRoutesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListStaticSiteRoutesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// (Required) The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public ListStaticSiteRoutesArgs()
        {
        }
    }

    public sealed class ListStaticSiteRoutesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// (Required) The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public ListStaticSiteRoutesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListStaticSiteRoutesResult
    {
        public readonly ImmutableArray<Outputs.ListStaticSiteRoutesResponse> Value;

        [OutputConstructor]
        private ListStaticSiteRoutesResult(ImmutableArray<Outputs.ListStaticSiteRoutesResponse> value)
        {
            Value = value;
        }
    }
}
