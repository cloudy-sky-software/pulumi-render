// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace CloudySkySoftware.Render.Services
{
    [RenderResourceType("render:services:Scale")]
    public partial class Scale : Pulumi.CustomResource
    {
        [Output("numInstances")]
        public Output<double> NumInstances { get; private set; } = null!;


        /// <summary>
        /// Create a Scale resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Scale(string name, ScaleArgs args, CustomResourceOptions? options = null)
            : base("render:services:Scale", name, args ?? new ScaleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Scale(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("render:services:Scale", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/cloudy-sky-software/pulumi-render/releases/download/${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Scale resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Scale Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Scale(name, id, options);
        }
    }

    public sealed class ScaleArgs : Pulumi.ResourceArgs
    {
        [Input("numInstances", required: true)]
        public Input<double> NumInstances { get; set; } = null!;

        public ScaleArgs()
        {
        }
    }
}
