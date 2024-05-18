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
    public sealed class PrivateServiceDetailsOutput
    {
        public readonly Outputs.AutoscalingConfig? Autoscaling;
        public readonly string BuildPlan;
        public readonly Outputs.Disk? Disk;
        /// <summary>
        /// Environment (runtime)
        /// </summary>
        public readonly Pulumi.Render.Services.PrivateServiceDetailsOutputEnv Env;
        public readonly Union<Outputs.DockerDetails, Outputs.NativeEnvironmentDetails> EnvSpecificDetails;
        /// <summary>
        /// For a *manually* scaled service, this is the number of instances the service is scaled to. DOES NOT indicate the number of running instances for an *autoscaled* service.
        /// </summary>
        public readonly int NumInstances;
        public readonly ImmutableArray<Outputs.ServerPort> OpenPorts;
        public readonly Outputs.Resource? ParentServer;
        /// <summary>
        /// The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
        /// </summary>
        public readonly Pulumi.Render.Services.PrivateServiceDetailsOutputPlan Plan;
        public readonly Pulumi.Render.Services.PrivateServiceDetailsOutputPullRequestPreviewsEnabled PullRequestPreviewsEnabled;
        public readonly Pulumi.Render.Services.PrivateServiceDetailsOutputRegion Region;
        public readonly string Url;

        [OutputConstructor]
        private PrivateServiceDetailsOutput(
            Outputs.AutoscalingConfig? autoscaling,

            string buildPlan,

            Outputs.Disk? disk,

            Pulumi.Render.Services.PrivateServiceDetailsOutputEnv env,

            Union<Outputs.DockerDetails, Outputs.NativeEnvironmentDetails> envSpecificDetails,

            int numInstances,

            ImmutableArray<Outputs.ServerPort> openPorts,

            Outputs.Resource? parentServer,

            Pulumi.Render.Services.PrivateServiceDetailsOutputPlan plan,

            Pulumi.Render.Services.PrivateServiceDetailsOutputPullRequestPreviewsEnabled pullRequestPreviewsEnabled,

            Pulumi.Render.Services.PrivateServiceDetailsOutputRegion region,

            string url)
        {
            Autoscaling = autoscaling;
            BuildPlan = buildPlan;
            Disk = disk;
            Env = env;
            EnvSpecificDetails = envSpecificDetails;
            NumInstances = numInstances;
            OpenPorts = openPorts;
            ParentServer = parentServer;
            Plan = plan;
            PullRequestPreviewsEnabled = pullRequestPreviewsEnabled;
            Region = region;
            Url = url;
        }
    }
}
