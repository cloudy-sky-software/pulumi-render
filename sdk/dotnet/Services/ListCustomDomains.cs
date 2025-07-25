// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class ListCustomDomains
    {
        public static Task<ListCustomDomainsResult> InvokeAsync(ListCustomDomainsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListCustomDomainsResult>("render:services:listCustomDomains", args ?? new ListCustomDomainsArgs(), options.WithDefaults());

        public static Output<ListCustomDomainsResult> Invoke(ListCustomDomainsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListCustomDomainsResult>("render:services:listCustomDomains", args ?? new ListCustomDomainsInvokeArgs(), options.WithDefaults());

        public static Output<ListCustomDomainsResult> Invoke(ListCustomDomainsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListCustomDomainsResult>("render:services:listCustomDomains", args ?? new ListCustomDomainsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListCustomDomainsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public ListCustomDomainsArgs()
        {
        }
        public static new ListCustomDomainsArgs Empty => new ListCustomDomainsArgs();
    }

    public sealed class ListCustomDomainsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public ListCustomDomainsInvokeArgs()
        {
        }
        public static new ListCustomDomainsInvokeArgs Empty => new ListCustomDomainsInvokeArgs();
    }


    [OutputType]
    public sealed class ListCustomDomainsResult
    {
        public readonly ImmutableArray<Outputs.CustomDomainWithCursor> Items;

        [OutputConstructor]
        private ListCustomDomainsResult(ImmutableArray<Outputs.CustomDomainWithCursor> items)
        {
            Items = items;
        }
    }
}
