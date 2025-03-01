// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class GetStaticSite
    {
        public static Task<Outputs.GetStaticSite> InvokeAsync(GetStaticSiteArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetStaticSite>("render:services:getStaticSite", args ?? new GetStaticSiteArgs(), options.WithDefaults());

        public static Output<Outputs.GetStaticSite> Invoke(GetStaticSiteInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetStaticSite>("render:services:getStaticSite", args ?? new GetStaticSiteInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.GetStaticSite> Invoke(GetStaticSiteInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetStaticSite>("render:services:getStaticSite", args ?? new GetStaticSiteInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStaticSiteArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetStaticSiteArgs()
        {
        }
        public static new GetStaticSiteArgs Empty => new GetStaticSiteArgs();
    }

    public sealed class GetStaticSiteInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetStaticSiteInvokeArgs()
        {
        }
        public static new GetStaticSiteInvokeArgs Empty => new GetStaticSiteInvokeArgs();
    }
}
