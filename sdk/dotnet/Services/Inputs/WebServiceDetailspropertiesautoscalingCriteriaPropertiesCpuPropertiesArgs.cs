// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Inputs
{

    public sealed class WebServiceDetailspropertiesautoscalingCriteriaPropertiesCpuPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Determines when your service will be scaled. If the average resource utilization is significantly above/below the target, we will increase/decrease the number of instances.
        /// </summary>
        [Input("percentage", required: true)]
        public Input<int> Percentage { get; set; } = null!;

        public WebServiceDetailspropertiesautoscalingCriteriaPropertiesCpuPropertiesArgs()
        {
            Enabled = false;
        }
        public static new WebServiceDetailspropertiesautoscalingCriteriaPropertiesCpuPropertiesArgs Empty => new WebServiceDetailspropertiesautoscalingCriteriaPropertiesCpuPropertiesArgs();
    }
}
