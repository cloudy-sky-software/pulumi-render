// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class ListServices
    {
        public static Task<ListServicesResult> InvokeAsync(ListServicesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListServicesResult>("render:services:listServices", args ?? new ListServicesArgs(), options.WithDefaults());

        public static Output<ListServicesResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListServicesResult>("render:services:listServices", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListServicesArgs : global::Pulumi.InvokeArgs
    {
        public ListServicesArgs()
        {
        }
        public static new ListServicesArgs Empty => new ListServicesArgs();
    }


    [OutputType]
    public sealed class ListServicesResult
    {
        public readonly ImmutableArray<Outputs.ListServiceResponse> Items;

        [OutputConstructor]
        private ListServicesResult(ImmutableArray<Outputs.ListServiceResponse> items)
        {
            Items = items;
        }
    }
}
