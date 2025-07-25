// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Blueprints
{
    public static class ListBlueprintSyncs
    {
        public static Task<ListBlueprintSyncsResult> InvokeAsync(ListBlueprintSyncsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListBlueprintSyncsResult>("render:blueprints:listBlueprintSyncs", args ?? new ListBlueprintSyncsArgs(), options.WithDefaults());

        public static Output<ListBlueprintSyncsResult> Invoke(ListBlueprintSyncsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListBlueprintSyncsResult>("render:blueprints:listBlueprintSyncs", args ?? new ListBlueprintSyncsInvokeArgs(), options.WithDefaults());

        public static Output<ListBlueprintSyncsResult> Invoke(ListBlueprintSyncsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListBlueprintSyncsResult>("render:blueprints:listBlueprintSyncs", args ?? new ListBlueprintSyncsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListBlueprintSyncsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the blueprint
        /// </summary>
        [Input("blueprintId", required: true)]
        public string BlueprintId { get; set; } = null!;

        public ListBlueprintSyncsArgs()
        {
        }
        public static new ListBlueprintSyncsArgs Empty => new ListBlueprintSyncsArgs();
    }

    public sealed class ListBlueprintSyncsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the blueprint
        /// </summary>
        [Input("blueprintId", required: true)]
        public Input<string> BlueprintId { get; set; } = null!;

        public ListBlueprintSyncsInvokeArgs()
        {
        }
        public static new ListBlueprintSyncsInvokeArgs Empty => new ListBlueprintSyncsInvokeArgs();
    }


    [OutputType]
    public sealed class ListBlueprintSyncsResult
    {
        public readonly ImmutableArray<Outputs.SyncWithCursor> Items;

        [OutputConstructor]
        private ListBlueprintSyncsResult(ImmutableArray<Outputs.SyncWithCursor> items)
        {
            Items = items;
        }
    }
}
