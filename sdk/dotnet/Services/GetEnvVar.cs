// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class GetEnvVar
    {
        public static Task<GetEnvVarResult> InvokeAsync(GetEnvVarArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEnvVarResult>("render:services:getEnvVar", args ?? new GetEnvVarArgs(), options.WithDefaults());

        public static Output<GetEnvVarResult> Invoke(GetEnvVarInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnvVarResult>("render:services:getEnvVar", args ?? new GetEnvVarInvokeArgs(), options.WithDefaults());

        public static Output<GetEnvVarResult> Invoke(GetEnvVarInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnvVarResult>("render:services:getEnvVar", args ?? new GetEnvVarInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnvVarArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the environment variable
        /// </summary>
        [Input("envVarKey", required: true)]
        public string EnvVarKey { get; set; } = null!;

        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetEnvVarArgs()
        {
        }
        public static new GetEnvVarArgs Empty => new GetEnvVarArgs();
    }

    public sealed class GetEnvVarInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the environment variable
        /// </summary>
        [Input("envVarKey", required: true)]
        public Input<string> EnvVarKey { get; set; } = null!;

        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetEnvVarInvokeArgs()
        {
        }
        public static new GetEnvVarInvokeArgs Empty => new GetEnvVarInvokeArgs();
    }


    [OutputType]
    public sealed class GetEnvVarResult
    {
        public readonly string Key;
        public readonly string Value;

        [OutputConstructor]
        private GetEnvVarResult(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
