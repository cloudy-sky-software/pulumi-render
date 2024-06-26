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
    public sealed class DockerDetails
    {
        public readonly string DockerCommand;
        public readonly string DockerContext;
        public readonly string DockerfilePath;
        public readonly string? PreDeployCommand;
        public readonly Outputs.RegistryCredential? RegistryCredential;

        [OutputConstructor]
        private DockerDetails(
            string dockerCommand,

            string dockerContext,

            string dockerfilePath,

            string? preDeployCommand,

            Outputs.RegistryCredential? registryCredential)
        {
            DockerCommand = dockerCommand;
            DockerContext = dockerContext;
            DockerfilePath = dockerfilePath;
            PreDeployCommand = preDeployCommand;
            RegistryCredential = registryCredential;
        }
    }
}
