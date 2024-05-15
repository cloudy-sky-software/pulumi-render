// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Inputs
{

    public sealed class CronJobDetailsCreateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Environment (runtime)
        /// </summary>
        [Input("env", required: true)]
        public Input<Pulumi.Render.Services.ServiceCronJobDetailsCreateEnv> Env { get; set; } = null!;

        [Input("envSpecificDetails")]
        public InputUnion<Inputs.DockerDetailsArgs, Inputs.NativeEnvironmentDetailsArgs>? EnvSpecificDetails { get; set; }

        [Input("plan")]
        public Input<Pulumi.Render.Services.ServiceCronJobDetailsCreatePlan>? Plan { get; set; }

        [Input("region")]
        public Input<Pulumi.Render.Services.ServiceCronJobDetailsCreateRegion>? Region { get; set; }

        [Input("schedule", required: true)]
        public Input<string> Schedule { get; set; } = null!;

        [Input("type")]
        public Input<string>? Type { get; set; }

        public CronJobDetailsCreateArgs()
        {
            Type = "cron_job";
        }
        public static new CronJobDetailsCreateArgs Empty => new CronJobDetailsCreateArgs();
    }
}
