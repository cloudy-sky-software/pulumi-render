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
        public static Task<Outputs.Deploy> InvokeAsync(GetDeployArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.Deploy>("render:services:getDeploy", args ?? new GetDeployArgs(), options.WithDefaults());

        public static Output<Outputs.Deploy> Invoke(GetDeployInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.Deploy>("render:services:getDeploy", args ?? new GetDeployInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.Deploy> Invoke(GetDeployInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.Deploy>("render:services:getDeploy", args ?? new GetDeployInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeployArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the deploy
        /// </summary>
        [Input("deployId", required: true)]
        public string DeployId { get; set; } = null!;

        /// <summary>
        /// The ID of the service
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
        /// The ID of the deploy
        /// </summary>
        [Input("deployId", required: true)]
        public Input<string> DeployId { get; set; } = null!;

        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetDeployInvokeArgs()
        {
        }
        public static new GetDeployInvokeArgs Empty => new GetDeployInvokeArgs();
    }
}
