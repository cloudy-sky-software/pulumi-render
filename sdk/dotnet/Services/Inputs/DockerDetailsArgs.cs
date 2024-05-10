// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Inputs
{

    public sealed class DockerDetailsArgs : global::Pulumi.ResourceArgs
    {
        [Input("dockerCommand", required: true)]
        public Input<string> DockerCommand { get; set; } = null!;

        [Input("dockerContext", required: true)]
        public Input<string> DockerContext { get; set; } = null!;

        [Input("dockerfilePath", required: true)]
        public Input<string> DockerfilePath { get; set; } = null!;

        [Input("preDeployCommand")]
        public Input<string>? PreDeployCommand { get; set; }

        [Input("registryCredential")]
        public Input<Inputs.RegistryCredentialArgs>? RegistryCredential { get; set; }

        public DockerDetailsArgs()
        {
        }
        public static new DockerDetailsArgs Empty => new DockerDetailsArgs();
    }
}
