// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class ListHeaders
    {
        public static Task<ListHeadersResult> InvokeAsync(ListHeadersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListHeadersResult>("render:services:listHeaders", args ?? new ListHeadersArgs(), options.WithDefaults());

        public static Output<ListHeadersResult> Invoke(ListHeadersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListHeadersResult>("render:services:listHeaders", args ?? new ListHeadersInvokeArgs(), options.WithDefaults());

        public static Output<ListHeadersResult> Invoke(ListHeadersInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListHeadersResult>("render:services:listHeaders", args ?? new ListHeadersInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListHeadersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public ListHeadersArgs()
        {
        }
        public static new ListHeadersArgs Empty => new ListHeadersArgs();
    }

    public sealed class ListHeadersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public ListHeadersInvokeArgs()
        {
        }
        public static new ListHeadersInvokeArgs Empty => new ListHeadersInvokeArgs();
    }


    [OutputType]
    public sealed class ListHeadersResult
    {
        public readonly ImmutableArray<Outputs.HeaderWithCursor> Items;

        [OutputConstructor]
        private ListHeadersResult(ImmutableArray<Outputs.HeaderWithCursor> items)
        {
            Items = items;
        }
    }
}
