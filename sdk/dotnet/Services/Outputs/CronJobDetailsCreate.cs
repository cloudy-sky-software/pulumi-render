// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Outputs
{

    [OutputType]
    public sealed class CronJobDetailsCreate
    {
        /// <summary>
        /// Environment (runtime)
        /// </summary>
        public readonly Pulumi.Render.Services.CronJobDetailsCreateEnv Env;
        public readonly Union<Outputs.DockerDetails, Outputs.NativeEnvironmentDetails>? EnvSpecificDetails;
        public readonly Pulumi.Render.Services.CronJobDetailsCreatePlan? Plan;
        public readonly Pulumi.Render.Services.CronJobDetailsCreateRegion? Region;
        public readonly string Schedule;

        [OutputConstructor]
        private CronJobDetailsCreate(
            Pulumi.Render.Services.CronJobDetailsCreateEnv env,

            Union<Outputs.DockerDetails, Outputs.NativeEnvironmentDetails>? envSpecificDetails,

            Pulumi.Render.Services.CronJobDetailsCreatePlan? plan,

            Pulumi.Render.Services.CronJobDetailsCreateRegion? region,

            string schedule)
        {
            Env = env;
            EnvSpecificDetails = envSpecificDetails;
            Plan = plan;
            Region = region;
            Schedule = schedule;
        }
    }
}
