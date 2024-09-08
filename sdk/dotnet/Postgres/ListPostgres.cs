// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Postgres
{
    public static class ListPostgres
    {
        public static Task<ListPostgresResult> InvokeAsync(ListPostgresArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListPostgresResult>("render:postgres:listPostgres", args ?? new ListPostgresArgs(), options.WithDefaults());

        public static Output<ListPostgresResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListPostgresResult>("render:postgres:listPostgres", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListPostgresArgs : global::Pulumi.InvokeArgs
    {
        public ListPostgresArgs()
        {
        }
        public static new ListPostgresArgs Empty => new ListPostgresArgs();
    }


    [OutputType]
    public sealed class ListPostgresResult
    {
        public readonly ImmutableArray<Outputs.ListPostgresItemProperties> Items;

        [OutputConstructor]
        private ListPostgresResult(ImmutableArray<Outputs.ListPostgresItemProperties> items)
        {
            Items = items;
        }
    }
}