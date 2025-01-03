// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Postgres
{
    public static class GetPostgres
    {
        public static Task<Outputs.PostgresDetail> InvokeAsync(GetPostgresArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.PostgresDetail>("render:postgres:getPostgres", args ?? new GetPostgresArgs(), options.WithDefaults());

        public static Output<Outputs.PostgresDetail> Invoke(GetPostgresInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.PostgresDetail>("render:postgres:getPostgres", args ?? new GetPostgresInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.PostgresDetail> Invoke(GetPostgresInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.PostgresDetail>("render:postgres:getPostgres", args ?? new GetPostgresInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPostgresArgs : global::Pulumi.InvokeArgs
    {
        [Input("postgresId", required: true)]
        public string PostgresId { get; set; } = null!;

        public GetPostgresArgs()
        {
        }
        public static new GetPostgresArgs Empty => new GetPostgresArgs();
    }

    public sealed class GetPostgresInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("postgresId", required: true)]
        public Input<string> PostgresId { get; set; } = null!;

        public GetPostgresInvokeArgs()
        {
        }
        public static new GetPostgresInvokeArgs Empty => new GetPostgresInvokeArgs();
    }
}
