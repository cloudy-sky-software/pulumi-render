// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Inputs
{

    public sealed class CronJobServiceDetailsArgs : Pulumi.ResourceArgs
    {
        [Input("env", required: true)]
        public Input<Pulumi.Render.Services.CronJobServiceDetailsEnv> Env { get; set; } = null!;

        [Input("envSpecificDetails")]
        public InputUnion<Inputs.DockerDetailsArgs, Inputs.NativeEnvironmentDetailsArgs>? EnvSpecificDetails { get; set; }

        [Input("lastSuccessfulRunAt")]
        public Input<string>? LastSuccessfulRunAt { get; set; }

        [Input("plan")]
        public Input<Pulumi.Render.Services.CronJobServiceDetailsPlan>? Plan { get; set; }

        [Input("region")]
        public Input<Pulumi.Render.Services.CronJobServiceDetailsRegion>? Region { get; set; }

        [Input("schedule", required: true)]
        public Input<string> Schedule { get; set; } = null!;

        public CronJobServiceDetailsArgs()
        {
            Plan = Pulumi.Render.Services.CronJobServiceDetailsPlan.Starter;
            Region = Pulumi.Render.Services.CronJobServiceDetailsRegion.Oregon;
        }
    }
}
