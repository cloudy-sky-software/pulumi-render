// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class ListRoutes
    {
        public static Task<ListRoutesResult> InvokeAsync(ListRoutesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListRoutesResult>("render:services:listRoutes", args ?? new ListRoutesArgs(), options.WithDefaults());

        public static Output<ListRoutesResult> Invoke(ListRoutesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListRoutesResult>("render:services:listRoutes", args ?? new ListRoutesInvokeArgs(), options.WithDefaults());

        public static Output<ListRoutesResult> Invoke(ListRoutesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListRoutesResult>("render:services:listRoutes", args ?? new ListRoutesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListRoutesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public ListRoutesArgs()
        {
        }
        public static new ListRoutesArgs Empty => new ListRoutesArgs();
    }

    public sealed class ListRoutesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public ListRoutesInvokeArgs()
        {
        }
        public static new ListRoutesInvokeArgs Empty => new ListRoutesInvokeArgs();
    }


    [OutputType]
    public sealed class ListRoutesResult
    {
        public readonly ImmutableArray<Outputs.RouteWithCursor> Items;

        [OutputConstructor]
        private ListRoutesResult(ImmutableArray<Outputs.RouteWithCursor> items)
        {
            Items = items;
        }
    }
}
