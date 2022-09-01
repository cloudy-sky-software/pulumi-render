// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class GetDeploy
    {
        public static Task<GetDeployResult> InvokeAsync(GetDeployArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDeployResult>("render:services:getDeploy", args ?? new GetDeployArgs(), options.WithDefaults());

        public static Output<GetDeployResult> Invoke(GetDeployInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDeployResult>("render:services:getDeploy", args ?? new GetDeployInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeployArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// (Required) The ID of the deploy
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// (Required) The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetDeployArgs()
        {
        }
        public static new GetDeployArgs Empty => new GetDeployArgs();
    }

    public sealed class GetDeployInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// (Required) The ID of the deploy
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// (Required) The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetDeployInvokeArgs()
        {
        }
        public static new GetDeployInvokeArgs Empty => new GetDeployInvokeArgs();
    }


    [OutputType]
    public sealed class GetDeployResult
    {
        public readonly Outputs.Deploy Items;

        [OutputConstructor]
        private GetDeployResult(Outputs.Deploy items)
        {
            Items = items;
        }
    }
}
