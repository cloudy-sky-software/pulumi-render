// *** WARNING: this file was generated by pulumigen. ***
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
            => Pulumi.Deployment.Instance.InvokeAsync<ListCustomDomainsResult>("render:services:listCustomDomains", args ?? new ListCustomDomainsArgs(), options.WithDefaults());

        public static Output<ListCustomDomainsResult> Invoke(ListCustomDomainsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<ListCustomDomainsResult>("render:services:listCustomDomains", args ?? new ListCustomDomainsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListCustomDomainsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// (Required) The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public ListCustomDomainsArgs()
        {
        }
    }

    public sealed class ListCustomDomainsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// (Required) The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public ListCustomDomainsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListCustomDomainsResult
    {
        [OutputConstructor]
        private ListCustomDomainsResult()
        {
        }
    }
}
