// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Registrycredentials.Outputs
{

    [OutputType]
    public sealed class ListRegistrycredentialsItemProperties
    {
        public readonly string? Id;
        public readonly string? Name;
        public readonly string? Registry;
        public readonly string? Username;

        [OutputConstructor]
        private ListRegistrycredentialsItemProperties(
            string? id,

            string? name,

            string? registry,

            string? username)
        {
            Id = id;
            Name = name;
            Registry = registry;
            Username = username;
        }
    }
}
