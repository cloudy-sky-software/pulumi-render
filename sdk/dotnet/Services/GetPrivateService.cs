// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class GetPrivateService
    {
        public static Task<Outputs.GetPrivateService> InvokeAsync(GetPrivateServiceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.GetPrivateService>("render:services:getPrivateService", args ?? new GetPrivateServiceArgs(), options.WithDefaults());

        public static Output<Outputs.GetPrivateService> Invoke(GetPrivateServiceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.GetPrivateService>("render:services:getPrivateService", args ?? new GetPrivateServiceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrivateServiceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetPrivateServiceArgs()
        {
        }
        public static new GetPrivateServiceArgs Empty => new GetPrivateServiceArgs();
    }

    public sealed class GetPrivateServiceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetPrivateServiceInvokeArgs()
        {
        }
        public static new GetPrivateServiceInvokeArgs Empty => new GetPrivateServiceInvokeArgs();
    }
}
