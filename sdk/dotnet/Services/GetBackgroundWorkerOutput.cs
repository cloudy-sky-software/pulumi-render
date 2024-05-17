// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class GetBackgroundWorkerOutput
    {
        public static Task<GetBackgroundWorkerOutputResult> InvokeAsync(GetBackgroundWorkerOutputArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBackgroundWorkerOutputResult>("render:services:getBackgroundWorkerOutput", args ?? new GetBackgroundWorkerOutputArgs(), options.WithDefaults());

        public static Output<GetBackgroundWorkerOutputResult> Invoke(GetBackgroundWorkerOutputInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBackgroundWorkerOutputResult>("render:services:getBackgroundWorkerOutput", args ?? new GetBackgroundWorkerOutputInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackgroundWorkerOutputArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetBackgroundWorkerOutputArgs()
        {
        }
        public static new GetBackgroundWorkerOutputArgs Empty => new GetBackgroundWorkerOutputArgs();
    }

    public sealed class GetBackgroundWorkerOutputInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetBackgroundWorkerOutputInvokeArgs()
        {
        }
        public static new GetBackgroundWorkerOutputInvokeArgs Empty => new GetBackgroundWorkerOutputInvokeArgs();
    }


    [OutputType]
    public sealed class GetBackgroundWorkerOutputResult
    {
        public readonly Outputs.GetBackgroundWorkerOutput Items;

        [OutputConstructor]
        private GetBackgroundWorkerOutputResult(Outputs.GetBackgroundWorkerOutput items)
        {
            Items = items;
        }
    }
}
