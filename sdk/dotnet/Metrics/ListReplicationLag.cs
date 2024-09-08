// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Metrics
{
    public static class ListReplicationLag
    {
        public static Task<ListReplicationLagResult> InvokeAsync(ListReplicationLagArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListReplicationLagResult>("render:metrics:listReplicationLag", args ?? new ListReplicationLagArgs(), options.WithDefaults());

        public static Output<ListReplicationLagResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListReplicationLagResult>("render:metrics:listReplicationLag", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListReplicationLagArgs : global::Pulumi.InvokeArgs
    {
        public ListReplicationLagArgs()
        {
        }
        public static new ListReplicationLagArgs Empty => new ListReplicationLagArgs();
    }


    [OutputType]
    public sealed class ListReplicationLagResult
    {
        public readonly ImmutableArray<Outputs.ListReplicationLagItemProperties> Items;

        [OutputConstructor]
        private ListReplicationLagResult(ImmutableArray<Outputs.ListReplicationLagItemProperties> items)
        {
            Items = items;
        }
    }
}