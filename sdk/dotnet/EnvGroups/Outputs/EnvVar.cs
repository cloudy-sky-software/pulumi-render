// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.EnvGroups.Outputs
{

    [OutputType]
    public sealed class EnvVar
    {
        public readonly string Key;
        public readonly string Value;

        [OutputConstructor]
        private EnvVar(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
