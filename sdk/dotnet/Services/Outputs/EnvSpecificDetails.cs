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
    public sealed class EnvSpecificDetails
    {
        public readonly string? BuildCommand;
        public readonly string? DockerCommand;
        public readonly string? DockerContext;
        public readonly string? DockerfilePath;
        public readonly string? PreDeployCommand;
        public readonly Outputs.RegistryCredential? RegistryCredential;
        public readonly string? StartCommand;

        [OutputConstructor]
        private EnvSpecificDetails(
            string? buildCommand,

            string? dockerCommand,

            string? dockerContext,

            string? dockerfilePath,

            string? preDeployCommand,

            Outputs.RegistryCredential? registryCredential,

            string? startCommand)
        {
            BuildCommand = buildCommand;
            DockerCommand = dockerCommand;
            DockerContext = dockerContext;
            DockerfilePath = dockerfilePath;
            PreDeployCommand = preDeployCommand;
            RegistryCredential = registryCredential;
            StartCommand = startCommand;
        }
    }
}