// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Render.Services.Inputs
{

    public sealed class ServiceWebServiceArgs : Pulumi.ResourceArgs
    {
        [Input("disk")]
        public Input<Inputs.ServiceDiskArgs>? Disk { get; set; }

        [Input("env", required: true)]
        public Input<CloudySkySoftware.Render.Services.ServiceWebServiceEnv> Env { get; set; } = null!;

        [Input("envSpecificDetails")]
        public InputUnion<Inputs.ServiceDockerDetailsArgs, Inputs.ServiceNativeEnvironmentDetailsArgs>? EnvSpecificDetails { get; set; }

        [Input("healthCheckPath")]
        public Input<string>? HealthCheckPath { get; set; }

        [Input("numInstances")]
        public Input<double>? NumInstances { get; set; }

        [Input("plan")]
        public Input<CloudySkySoftware.Render.Services.ServiceWebServicePlan>? Plan { get; set; }

        [Input("pullRequestPreviewsEnabled")]
        public Input<CloudySkySoftware.Render.Services.ServiceWebServicePullRequestPreviewsEnabled>? PullRequestPreviewsEnabled { get; set; }

        [Input("region")]
        public Input<CloudySkySoftware.Render.Services.ServiceWebServiceRegion>? Region { get; set; }

        public ServiceWebServiceArgs()
        {
            NumInstances = 1;
            Plan = CloudySkySoftware.Render.Services.ServiceWebServicePlan.Starter;
            PullRequestPreviewsEnabled = CloudySkySoftware.Render.Services.ServiceWebServicePullRequestPreviewsEnabled.No;
            Region = CloudySkySoftware.Render.Services.ServiceWebServiceRegion.Oregon;
        }
    }
}