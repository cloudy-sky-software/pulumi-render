// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Postgres
{
    public static class ListPostgresBackup
    {
        public static Task<ListPostgresBackupResult> InvokeAsync(ListPostgresBackupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListPostgresBackupResult>("render:postgres:listPostgresBackup", args ?? new ListPostgresBackupArgs(), options.WithDefaults());

        public static Output<ListPostgresBackupResult> Invoke(ListPostgresBackupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListPostgresBackupResult>("render:postgres:listPostgresBackup", args ?? new ListPostgresBackupInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListPostgresBackupArgs : global::Pulumi.InvokeArgs
    {
        [Input("postgresId", required: true)]
        public string PostgresId { get; set; } = null!;

        public ListPostgresBackupArgs()
        {
        }
        public static new ListPostgresBackupArgs Empty => new ListPostgresBackupArgs();
    }

    public sealed class ListPostgresBackupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("postgresId", required: true)]
        public Input<string> PostgresId { get; set; } = null!;

        public ListPostgresBackupInvokeArgs()
        {
        }
        public static new ListPostgresBackupInvokeArgs Empty => new ListPostgresBackupInvokeArgs();
    }


    [OutputType]
    public sealed class ListPostgresBackupResult
    {
        public readonly ImmutableArray<Outputs.ListPostgresBackupItemProperties> Items;

        [OutputConstructor]
        private ListPostgresBackupResult(ImmutableArray<Outputs.ListPostgresBackupItemProperties> items)
        {
            Items = items;
        }
    }
}