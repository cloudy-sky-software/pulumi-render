// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Metrics.Outputs
{

    /// <summary>
    /// A time series data point
    /// </summary>
    [OutputType]
    public sealed class ListActiveConnectionsItemProperties
    {
        /// <summary>
        /// List of labels describing the time series
        /// </summary>
        public readonly ImmutableArray<Outputs.ListActiveConnectionsItemPropertiesLabelsItemProperties> Labels;
        public readonly string Unit;
        /// <summary>
        /// The values of the time series
        /// </summary>
        public readonly ImmutableArray<Outputs.ListActiveConnectionsItemPropertiesValuesItemProperties> Values;

        [OutputConstructor]
        private ListActiveConnectionsItemProperties(
            ImmutableArray<Outputs.ListActiveConnectionsItemPropertiesLabelsItemProperties> labels,

            string unit,

            ImmutableArray<Outputs.ListActiveConnectionsItemPropertiesValuesItemProperties> values)
        {
            Labels = labels;
            Unit = unit;
            Values = values;
        }
    }
}
