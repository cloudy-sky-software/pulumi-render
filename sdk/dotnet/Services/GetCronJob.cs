// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class GetCronJob
    {
        public static Task<Outputs.GetCronJob> InvokeAsync(GetCronJobArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetCronJob>("render:services:getCronJob", args ?? new GetCronJobArgs(), options.WithDefaults());

        public static Output<Outputs.GetCronJob> Invoke(GetCronJobInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetCronJob>("render:services:getCronJob", args ?? new GetCronJobInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.GetCronJob> Invoke(GetCronJobInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetCronJob>("render:services:getCronJob", args ?? new GetCronJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCronJobArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetCronJobArgs()
        {
        }
        public static new GetCronJobArgs Empty => new GetCronJobArgs();
    }

    public sealed class GetCronJobInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetCronJobInvokeArgs()
        {
        }
        public static new GetCronJobInvokeArgs Empty => new GetCronJobInvokeArgs();
    }
}
