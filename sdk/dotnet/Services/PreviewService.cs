// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    [RenderResourceType("render:services:PreviewService")]
    public partial class PreviewService : global::Pulumi.CustomResource
    {
        [Output("deployId")]
        public Output<string?> DeployId { get; private set; } = null!;

        /// <summary>
        /// Must be either a full URL or the relative path to an image. If a relative path, Render uses the base service's image URL as its root. For example, if the base service's image URL is `docker.io/library/nginx:latest`, then valid values are: `docker.io/library/nginx:&lt;any tag or SHA&gt;`, `library/nginx:&lt;any tag or SHA&gt;`, or `nginx:&lt;any tag or SHA&gt;`. Note that the path must match (only the tag or SHA can vary).
        /// </summary>
        [Output("imagePath")]
        public Output<string> ImagePath { get; private set; } = null!;

        /// <summary>
        /// A name for the service preview instance. If not specified, Render generates the name using the base service's name and the specified tag or SHA.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
        /// </summary>
        [Output("plan")]
        public Output<Pulumi.Render.Services.Plan?> Plan { get; private set; } = null!;

        [Output("service")]
        public Output<Outputs.Service?> Service { get; private set; } = null!;


        /// <summary>
        /// Create a PreviewService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PreviewService(string name, PreviewServiceArgs args, CustomResourceOptions? options = null)
            : base("render:services:PreviewService", name, args ?? new PreviewServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PreviewService(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("render:services:PreviewService", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing PreviewService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PreviewService Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PreviewService(name, id, options);
        }
    }

    public sealed class PreviewServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Must be either a full URL or the relative path to an image. If a relative path, Render uses the base service's image URL as its root. For example, if the base service's image URL is `docker.io/library/nginx:latest`, then valid values are: `docker.io/library/nginx:&lt;any tag or SHA&gt;`, `library/nginx:&lt;any tag or SHA&gt;`, or `nginx:&lt;any tag or SHA&gt;`. Note that the path must match (only the tag or SHA can vary).
        /// </summary>
        [Input("imagePath", required: true)]
        public Input<string> ImagePath { get; set; } = null!;

        /// <summary>
        /// A name for the service preview instance. If not specified, Render generates the name using the base service's name and the specified tag or SHA.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
        /// </summary>
        [Input("plan")]
        public Input<Pulumi.Render.Services.Plan>? Plan { get; set; }

        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        public PreviewServiceArgs()
        {
        }
        public static new PreviewServiceArgs Empty => new PreviewServiceArgs();
    }
}
