// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class GetSecretFile
    {
        public static Task<Outputs.SecretFile> InvokeAsync(GetSecretFileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.SecretFile>("render:services:getSecretFile", args ?? new GetSecretFileArgs(), options.WithDefaults());

        public static Output<Outputs.SecretFile> Invoke(GetSecretFileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.SecretFile>("render:services:getSecretFile", args ?? new GetSecretFileInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.SecretFile> Invoke(GetSecretFileInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.SecretFile>("render:services:getSecretFile", args ?? new GetSecretFileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecretFileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The file name of the secret file
        /// </summary>
        [Input("secretFileName", required: true)]
        public string SecretFileName { get; set; } = null!;

        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetSecretFileArgs()
        {
        }
        public static new GetSecretFileArgs Empty => new GetSecretFileArgs();
    }

    public sealed class GetSecretFileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The file name of the secret file
        /// </summary>
        [Input("secretFileName", required: true)]
        public Input<string> SecretFileName { get; set; } = null!;

        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetSecretFileInvokeArgs()
        {
        }
        public static new GetSecretFileInvokeArgs Empty => new GetSecretFileInvokeArgs();
    }
}
