// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.RegistryCredentials.Outputs
{

    [OutputType]
    public sealed class RegistryCredential
    {
        /// <summary>
        /// Unique identifier for this credential
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Descriptive name for this credential
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The registry to use this credential with
        /// </summary>
        public readonly Pulumi.Render.RegistryCredentials.RegistryCredentialRegistry Registry;
        /// <summary>
        /// Last updated time for the credential
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// The username associated with the credential
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private RegistryCredential(
            string id,

            string name,

            Pulumi.Render.RegistryCredentials.RegistryCredentialRegistry registry,

            string updatedAt,

            string username)
        {
            Id = id;
            Name = name;
            Registry = registry;
            UpdatedAt = updatedAt;
            Username = username;
        }
    }
}
