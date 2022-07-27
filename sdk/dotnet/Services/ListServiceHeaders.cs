// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class ListServiceHeaders
    {
        public static Task<ListServiceHeadersResult> InvokeAsync(ListServiceHeadersArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListServiceHeadersResult>("render:services:listServiceHeaders", args ?? new ListServiceHeadersArgs(), options.WithDefaults());

        public static Output<ListServiceHeadersResult> Invoke(ListServiceHeadersInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<ListServiceHeadersResult>("render:services:listServiceHeaders", args ?? new ListServiceHeadersInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListServiceHeadersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// (Required) The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public ListServiceHeadersArgs()
        {
        }
    }

    public sealed class ListServiceHeadersInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// (Required) The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public ListServiceHeadersInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListServiceHeadersResult
    {
        [OutputConstructor]
        private ListServiceHeadersResult()
        {
        }
    }
}
