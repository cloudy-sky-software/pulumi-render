// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Inputs
{

    public sealed class EnvVarInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("generateValue")]
        public Input<bool>? GenerateValue { get; set; }

        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public EnvVarInputArgs()
        {
        }
        public static new EnvVarInputArgs Empty => new EnvVarInputArgs();
    }
}
