// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Postgres
{
    public static class GetPostgresRecoveryInfo
    {
        public static Task<GetPostgresRecoveryInfoResult> InvokeAsync(GetPostgresRecoveryInfoArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPostgresRecoveryInfoResult>("render:postgres:getPostgresRecoveryInfo", args ?? new GetPostgresRecoveryInfoArgs(), options.WithDefaults());

        public static Output<GetPostgresRecoveryInfoResult> Invoke(GetPostgresRecoveryInfoInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPostgresRecoveryInfoResult>("render:postgres:getPostgresRecoveryInfo", args ?? new GetPostgresRecoveryInfoInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPostgresRecoveryInfoArgs : global::Pulumi.InvokeArgs
    {
        [Input("postgresId", required: true)]
        public string PostgresId { get; set; } = null!;

        public GetPostgresRecoveryInfoArgs()
        {
        }
        public static new GetPostgresRecoveryInfoArgs Empty => new GetPostgresRecoveryInfoArgs();
    }

    public sealed class GetPostgresRecoveryInfoInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("postgresId", required: true)]
        public Input<string> PostgresId { get; set; } = null!;

        public GetPostgresRecoveryInfoInvokeArgs()
        {
        }
        public static new GetPostgresRecoveryInfoInvokeArgs Empty => new GetPostgresRecoveryInfoInvokeArgs();
    }


    [OutputType]
    public sealed class GetPostgresRecoveryInfoResult
    {
        public readonly Outputs.GetPostgresRecoveryInfoProperties Items;

        [OutputConstructor]
        private GetPostgresRecoveryInfoResult(Outputs.GetPostgresRecoveryInfoProperties items)
        {
            Items = items;
        }
    }
}
