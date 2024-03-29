// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Inputs
{

    public sealed class EnvVarKeyValueArgs : global::Pulumi.ResourceArgs
    {
        [Input("generateValue")]
        public Input<Pulumi.Render.Services.EnvVarKeyValueGenerateValue>? GenerateValue { get; set; }

        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("value")]
        public Input<string>? Value { get; set; }

        public EnvVarKeyValueArgs()
        {
        }
        public static new EnvVarKeyValueArgs Empty => new EnvVarKeyValueArgs();
    }
}
