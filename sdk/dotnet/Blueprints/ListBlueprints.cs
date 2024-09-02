// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Blueprints
{
    public static class ListBlueprints
    {
        public static Task<ListBlueprintsResult> InvokeAsync(ListBlueprintsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListBlueprintsResult>("render:blueprints:listBlueprints", args ?? new ListBlueprintsArgs(), options.WithDefaults());

        public static Output<ListBlueprintsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListBlueprintsResult>("render:blueprints:listBlueprints", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListBlueprintsArgs : global::Pulumi.InvokeArgs
    {
        public ListBlueprintsArgs()
        {
        }
        public static new ListBlueprintsArgs Empty => new ListBlueprintsArgs();
    }


    [OutputType]
    public sealed class ListBlueprintsResult
    {
        public readonly ImmutableArray<Outputs.BlueprintWithCursor> Items;

        [OutputConstructor]
        private ListBlueprintsResult(ImmutableArray<Outputs.BlueprintWithCursor> items)
        {
            Items = items;
        }
    }
}
