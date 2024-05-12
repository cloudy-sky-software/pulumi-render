// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Inputs
{

    public sealed class WebServiceDetailsCreateArgs : global::Pulumi.ResourceArgs
    {
        [Input("disk")]
        public Input<Inputs.WebServiceDetailsCreateDiskPropertiesArgs>? Disk { get; set; }

        /// <summary>
        /// Environment (runtime)
        /// </summary>
        [Input("env", required: true)]
        public Input<Pulumi.Render.Services.WebServiceDetailsCreateEnv> Env { get; set; } = null!;

        [Input("envSpecificDetails")]
        public InputUnion<Inputs.DockerDetailsArgs, Inputs.NativeEnvironmentDetailsArgs>? EnvSpecificDetails { get; set; }

        [Input("healthCheckPath")]
        public Input<string>? HealthCheckPath { get; set; }

        /// <summary>
        /// Defaults to 1
        /// </summary>
        [Input("numInstances")]
        public Input<int>? NumInstances { get; set; }

        [Input("plan")]
        public Input<Pulumi.Render.Services.WebServiceDetailsCreatePlan>? Plan { get; set; }

        /// <summary>
        /// Defaults to "no"
        /// </summary>
        [Input("pullRequestPreviewsEnabled")]
        public Input<Pulumi.Render.Services.WebServiceDetailsCreatePullRequestPreviewsEnabled>? PullRequestPreviewsEnabled { get; set; }

        [Input("region")]
        public Input<Pulumi.Render.Services.WebServiceDetailsCreateRegion>? Region { get; set; }

        public WebServiceDetailsCreateArgs()
        {
            PullRequestPreviewsEnabled = Pulumi.Render.Services.WebServiceDetailsCreatePullRequestPreviewsEnabled.No;
        }
        public static new WebServiceDetailsCreateArgs Empty => new WebServiceDetailsCreateArgs();
    }
}