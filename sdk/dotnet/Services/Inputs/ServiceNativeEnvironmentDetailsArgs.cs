// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Inputs
{

    public sealed class ServiceNativeEnvironmentDetailsArgs : Pulumi.ResourceArgs
    {
        [Input("buildCommand", required: true)]
        public Input<string> BuildCommand { get; set; } = null!;

        [Input("startCommand", required: true)]
        public Input<string> StartCommand { get; set; } = null!;

        public ServiceNativeEnvironmentDetailsArgs()
        {
        }
    }
}
