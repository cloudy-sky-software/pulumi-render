// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.EnvGroups
{
    public static class GetEnvGroupSecretFile
    {
        public static Task<GetEnvGroupSecretFileResult> InvokeAsync(GetEnvGroupSecretFileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEnvGroupSecretFileResult>("render:env-groups:getEnvGroupSecretFile", args ?? new GetEnvGroupSecretFileArgs(), options.WithDefaults());

        public static Output<GetEnvGroupSecretFileResult> Invoke(GetEnvGroupSecretFileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnvGroupSecretFileResult>("render:env-groups:getEnvGroupSecretFile", args ?? new GetEnvGroupSecretFileInvokeArgs(), options.WithDefaults());

        public static Output<GetEnvGroupSecretFileResult> Invoke(GetEnvGroupSecretFileInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnvGroupSecretFileResult>("render:env-groups:getEnvGroupSecretFile", args ?? new GetEnvGroupSecretFileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnvGroupSecretFileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter for resources that belong to an environment group
        /// </summary>
        [Input("envGroupId", required: true)]
        public string EnvGroupId { get; set; } = null!;

        /// <summary>
        /// The name of the secret file
        /// </summary>
        [Input("secretFileName", required: true)]
        public string SecretFileName { get; set; } = null!;

        public GetEnvGroupSecretFileArgs()
        {
        }
        public static new GetEnvGroupSecretFileArgs Empty => new GetEnvGroupSecretFileArgs();
    }

    public sealed class GetEnvGroupSecretFileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter for resources that belong to an environment group
        /// </summary>
        [Input("envGroupId", required: true)]
        public Input<string> EnvGroupId { get; set; } = null!;

        /// <summary>
        /// The name of the secret file
        /// </summary>
        [Input("secretFileName", required: true)]
        public Input<string> SecretFileName { get; set; } = null!;

        public GetEnvGroupSecretFileInvokeArgs()
        {
        }
        public static new GetEnvGroupSecretFileInvokeArgs Empty => new GetEnvGroupSecretFileInvokeArgs();
    }


    [OutputType]
    public sealed class GetEnvGroupSecretFileResult
    {
        public readonly string Content;
        public readonly string Name;

        [OutputConstructor]
        private GetEnvGroupSecretFileResult(
            string content,

            string name)
        {
            Content = content;
            Name = name;
        }
    }
}
