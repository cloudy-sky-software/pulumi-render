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
    public sealed class SecretFile
    {
        public readonly string Name;

        [OutputConstructor]
        private SecretFile(string name)
        {
            Name = name;
        }
    }
}
