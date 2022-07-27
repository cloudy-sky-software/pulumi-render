// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    [RenderResourceType("render:services:Deploy")]
    public partial class Deploy : Pulumi.CustomResource
    {
        [Output("clearCache")]
        public Output<Pulumi.Render.Services.DeployClearCache?> ClearCache { get; private set; } = null!;

        [Output("commit")]
        public Output<Outputs.DeployCommit?> Commit { get; private set; } = null!;


        /// <summary>
        /// Create a Deploy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Deploy(string name, DeployArgs? args = null, CustomResourceOptions? options = null)
            : base("render:services:Deploy", name, args ?? new DeployArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Deploy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("render:services:Deploy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Deploy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Deploy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Deploy(name, id, options);
        }
    }

    public sealed class DeployArgs : Pulumi.ResourceArgs
    {
        [Input("clearCache")]
        public Input<Pulumi.Render.Services.DeployClearCache>? ClearCache { get; set; }

        /// <summary>
        /// (Required) The ID of the service
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public DeployArgs()
        {
            ClearCache = Pulumi.Render.Services.DeployClearCache.DoNotClear;
        }
    }
}
