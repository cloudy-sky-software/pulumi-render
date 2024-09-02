// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Inputs
{

    public sealed class BackgroundWorkerDetailsCreateArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoscaling")]
        public Input<Inputs.WebServiceDetailspropertiesautoscalingArgs>? Autoscaling { get; set; }

        [Input("disk")]
        public Input<Inputs.ServiceDiskArgs>? Disk { get; set; }

        /// <summary>
        /// This field has been deprecated, runtime should be used in its place.
        /// </summary>
        [Input("env")]
        public Input<Pulumi.Render.Services.BackgroundWorkerDetailsCreateEnv>? Env { get; set; }

        [Input("envSpecificDetails")]
        public Input<Inputs.EnvSpecificDetailsCreateArgs>? EnvSpecificDetails { get; set; }

        /// <summary>
        /// The maximum amount of time (in seconds) that Render waits for your application process to exit gracefully after sending it a SIGTERM signal.
        /// </summary>
        [Input("maxShutdownDelaySeconds")]
        public Input<int>? MaxShutdownDelaySeconds { get; set; }

        /// <summary>
        /// Defaults to 1
        /// </summary>
        [Input("numInstances")]
        public Input<int>? NumInstances { get; set; }

        /// <summary>
        /// Defaults to "starter"
        /// </summary>
        [Input("plan")]
        public Input<Pulumi.Render.Services.BackgroundWorkerDetailsCreatePlan>? Plan { get; set; }

        [Input("preDeployCommand")]
        public Input<string>? PreDeployCommand { get; set; }

        /// <summary>
        /// Defaults to "no"
        /// </summary>
        [Input("pullRequestPreviewsEnabled")]
        public Input<Pulumi.Render.Services.BackgroundWorkerDetailsCreatePullRequestPreviewsEnabled>? PullRequestPreviewsEnabled { get; set; }

        /// <summary>
        /// Defaults to "oregon"
        /// </summary>
        [Input("region")]
        public Input<Pulumi.Render.Services.BackgroundWorkerDetailsCreateRegion>? Region { get; set; }

        /// <summary>
        /// Runtime
        /// </summary>
        [Input("runtime", required: true)]
        public Input<Pulumi.Render.Services.BackgroundWorkerDetailsCreateRuntime> Runtime { get; set; } = null!;

        public BackgroundWorkerDetailsCreateArgs()
        {
            MaxShutdownDelaySeconds = 30;
            NumInstances = 1;
            Plan = Pulumi.Render.Services.BackgroundWorkerDetailsCreatePlan.Starter;
            PullRequestPreviewsEnabled = Pulumi.Render.Services.BackgroundWorkerDetailsCreatePullRequestPreviewsEnabled.No;
            Region = Pulumi.Render.Services.BackgroundWorkerDetailsCreateRegion.Oregon;
        }
        public static new BackgroundWorkerDetailsCreateArgs Empty => new BackgroundWorkerDetailsCreateArgs();
    }
}
