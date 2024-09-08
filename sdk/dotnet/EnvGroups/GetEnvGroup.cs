// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.EnvGroups
{
    public static class GetEnvGroup
    {
        public static Task<Outputs.EnvGroup> InvokeAsync(GetEnvGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.EnvGroup>("render:env-groups:getEnvGroup", args ?? new GetEnvGroupArgs(), options.WithDefaults());

        public static Output<Outputs.EnvGroup> Invoke(GetEnvGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.EnvGroup>("render:env-groups:getEnvGroup", args ?? new GetEnvGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnvGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter for resources that belong to an environment group
        /// </summary>
        [Input("envGroupId", required: true)]
        public string EnvGroupId { get; set; } = null!;

        public GetEnvGroupArgs()
        {
        }
        public static new GetEnvGroupArgs Empty => new GetEnvGroupArgs();
    }

    public sealed class GetEnvGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter for resources that belong to an environment group
        /// </summary>
        [Input("envGroupId", required: true)]
        public Input<string> EnvGroupId { get; set; } = null!;

        public GetEnvGroupInvokeArgs()
        {
        }
        public static new GetEnvGroupInvokeArgs Empty => new GetEnvGroupInvokeArgs();
    }
}