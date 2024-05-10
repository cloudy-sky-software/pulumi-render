// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Outputs
{

    /// <summary>
    /// Image information used when creating the deploy. Not present for Git-backed deploys
    /// </summary>
    [OutputType]
    public sealed class DeployImageProperties
    {
        /// <summary>
        /// Image reference used when creating the deploy
        /// </summary>
        public readonly string? Ref;
        /// <summary>
        /// Name of credential used to pull the image, if provided
        /// </summary>
        public readonly string? RegistryCredential;
        /// <summary>
        /// SHA that the image reference was resolved to when creating the deploy
        /// </summary>
        public readonly string? Sha;

        [OutputConstructor]
        private DeployImageProperties(
            string? @ref,

            string? registryCredential,

            string? sha)
        {
            Ref = @ref;
            RegistryCredential = registryCredential;
            Sha = sha;
        }
    }
}
