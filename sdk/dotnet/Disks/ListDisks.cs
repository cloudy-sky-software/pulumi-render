// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Disks
{
    public static class ListDisks
    {
        public static Task<ListDisksResult> InvokeAsync(ListDisksArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListDisksResult>("render:disks:listDisks", args ?? new ListDisksArgs(), options.WithDefaults());

        public static Output<ListDisksResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListDisksResult>("render:disks:listDisks", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListDisksArgs : global::Pulumi.InvokeArgs
    {
        public ListDisksArgs()
        {
        }
        public static new ListDisksArgs Empty => new ListDisksArgs();
    }


    [OutputType]
    public sealed class ListDisksResult
    {
        public readonly ImmutableArray<Outputs.DiskWithCursor> Items;

        [OutputConstructor]
        private ListDisksResult(ImmutableArray<Outputs.DiskWithCursor> items)
        {
            Items = items;
        }
    }
}