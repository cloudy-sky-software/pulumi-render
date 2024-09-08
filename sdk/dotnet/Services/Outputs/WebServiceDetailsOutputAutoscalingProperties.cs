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
    public sealed class WebServiceDetailsOutputAutoscalingProperties
    {
        public readonly Outputs.WebServiceDetailsOutputAutoscalingPropertiesCriteriaProperties Criteria;
        public readonly bool Enabled;
        /// <summary>
        /// The maximum number of instances for the service
        /// </summary>
        public readonly int Max;
        /// <summary>
        /// The minimum number of instances for the service
        /// </summary>
        public readonly int Min;

        [OutputConstructor]
        private WebServiceDetailsOutputAutoscalingProperties(
            Outputs.WebServiceDetailsOutputAutoscalingPropertiesCriteriaProperties criteria,

            bool enabled,

            int max,

            int min)
        {
            Criteria = criteria;
            Enabled = enabled;
            Max = max;
            Min = min;
        }
    }
}