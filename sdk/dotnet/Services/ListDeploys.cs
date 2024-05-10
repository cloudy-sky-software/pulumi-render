// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class ListDeploys
    {
        public static Task<ListDeploysResult> InvokeAsync(ListDeploysArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListDeploysResult>("render:services:listDeploys", args ?? new ListDeploysArgs(), options.WithDefaults());

        public static Output<ListDeploysResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListDeploysResult>("render:services:listDeploys", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListDeploysArgs : global::Pulumi.InvokeArgs
    {
        public ListDeploysArgs()
        {
        }
        public static new ListDeploysArgs Empty => new ListDeploysArgs();
    }


    [OutputType]
    public sealed class ListDeploysResult
    {
        public readonly ImmutableArray<Outputs.ListDeploysItemProperties> Items;

        [OutputConstructor]
        private ListDeploysResult(ImmutableArray<Outputs.ListDeploysItemProperties> items)
        {
            Items = items;
        }
    }
}
