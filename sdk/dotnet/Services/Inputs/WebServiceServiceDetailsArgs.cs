// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Inputs
{

    public sealed class WebServiceServiceDetailsArgs : global::Pulumi.ResourceArgs
    {
        [Input("disk")]
        public Input<Inputs.DiskArgs>? Disk { get; set; }

        [Input("env", required: true)]
        public Input<Pulumi.Render.Services.WebServiceServiceDetailsEnv> Env { get; set; } = null!;

        [Input("envSpecificDetails")]
        public InputUnion<Inputs.DockerDetailsArgs, Inputs.NativeEnvironmentDetailsArgs>? EnvSpecificDetails { get; set; }

        [Input("healthCheckPath")]
        public Input<string>? HealthCheckPath { get; set; }

        [Input("numInstances")]
        public Input<double>? NumInstances { get; set; }

        [Input("openPorts")]
        private InputList<Inputs.OpenPortsArgs>? _openPorts;
        public InputList<Inputs.OpenPortsArgs> OpenPorts
        {
            get => _openPorts ?? (_openPorts = new InputList<Inputs.OpenPortsArgs>());
            set => _openPorts = value;
        }

        [Input("parentServer")]
        public Input<Inputs.WebServiceServiceDetailsParentServerPropertiesArgs>? ParentServer { get; set; }

        [Input("plan")]
        public Input<Pulumi.Render.Services.WebServiceServiceDetailsPlan>? Plan { get; set; }

        [Input("pullRequestPreviewsEnabled")]
        public Input<Pulumi.Render.Services.WebServiceServiceDetailsPullRequestPreviewsEnabled>? PullRequestPreviewsEnabled { get; set; }

        [Input("region")]
        public Input<Pulumi.Render.Services.WebServiceServiceDetailsRegion>? Region { get; set; }

        [Input("url")]
        public Input<string>? Url { get; set; }

        public WebServiceServiceDetailsArgs()
        {
            NumInstances = 1;
            Plan = Pulumi.Render.Services.WebServiceServiceDetailsPlan.Starter;
            PullRequestPreviewsEnabled = Pulumi.Render.Services.WebServiceServiceDetailsPullRequestPreviewsEnabled.No;
            Region = Pulumi.Render.Services.WebServiceServiceDetailsRegion.Oregon;
        }
        public static new WebServiceServiceDetailsArgs Empty => new WebServiceServiceDetailsArgs();
    }
}
