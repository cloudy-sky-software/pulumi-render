// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Logs
{
    public static class GetResourceLogStream
    {
        public static Task<Outputs.GetResourceLogStreamProperties> InvokeAsync(GetResourceLogStreamArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetResourceLogStreamProperties>("render:logs:getResourceLogStream", args ?? new GetResourceLogStreamArgs(), options.WithDefaults());

        public static Output<Outputs.GetResourceLogStreamProperties> Invoke(GetResourceLogStreamInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetResourceLogStreamProperties>("render:logs:getResourceLogStream", args ?? new GetResourceLogStreamInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourceLogStreamArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the resource (server, cron job, postgres, or redis) whose log streams should be returned
        /// </summary>
        [Input("resourceId", required: true)]
        public string ResourceId { get; set; } = null!;

        public GetResourceLogStreamArgs()
        {
        }
        public static new GetResourceLogStreamArgs Empty => new GetResourceLogStreamArgs();
    }

    public sealed class GetResourceLogStreamInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the resource (server, cron job, postgres, or redis) whose log streams should be returned
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        public GetResourceLogStreamInvokeArgs()
        {
        }
        public static new GetResourceLogStreamInvokeArgs Empty => new GetResourceLogStreamInvokeArgs();
    }
}
