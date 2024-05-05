// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    [RenderResourceType("render:services:AutoScaling")]
    public partial class AutoScaling : global::Pulumi.CustomResource
    {
        [Output("criteria")]
        public Output<Outputs.AutoScalingCriteria?> Criteria { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The maximum number of instances for the service
        /// </summary>
        [Output("max")]
        public Output<double?> Max { get; private set; } = null!;

        /// <summary>
        /// The minimum number of instances for the service
        /// </summary>
        [Output("min")]
        public Output<double?> Min { get; private set; } = null!;


        /// <summary>
        /// Create a AutoScaling resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AutoScaling(string name, AutoScalingArgs? args = null, CustomResourceOptions? options = null)
            : base("render:services:AutoScaling", name, args ?? new AutoScalingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AutoScaling(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("render:services:AutoScaling", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/cloudy-sky-software/pulumi-render",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AutoScaling resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AutoScaling Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AutoScaling(name, id, options);
        }
    }

    public sealed class AutoScalingArgs : global::Pulumi.ResourceArgs
    {
        [Input("criteria")]
        public Input<Inputs.AutoScalingCriteriaArgs>? Criteria { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The maximum number of instances for the service
        /// </summary>
        [Input("max")]
        public Input<double>? Max { get; set; }

        /// <summary>
        /// The minimum number of instances for the service
        /// </summary>
        [Input("min")]
        public Input<double>? Min { get; set; }

        /// <summary>
        /// (Required) The ID of the service
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        public AutoScalingArgs()
        {
        }
        public static new AutoScalingArgs Empty => new AutoScalingArgs();
    }
}
