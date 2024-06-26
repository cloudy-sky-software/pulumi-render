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
    public partial class Deploy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Defaults to "do_not_clear"
        /// </summary>
        [Output("clearCache")]
        public Output<Pulumi.Render.Services.ClearCache?> ClearCache { get; private set; } = null!;

        [Output("commit")]
        public Output<Outputs.CommitProperties?> Commit { get; private set; } = null!;

        /// <summary>
        /// Specific ID of commit to deploy for a web service, defaults to latest commit. Not supported for Cron Job deploys.
        /// </summary>
        [Output("commitId")]
        public Output<string?> CommitId { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string?> CreatedAt { get; private set; } = null!;

        [Output("finishedAt")]
        public Output<string?> FinishedAt { get; private set; } = null!;

        /// <summary>
        /// Image information used when creating the deploy. Not present for Git-backed deploys
        /// </summary>
        [Output("image")]
        public Output<Outputs.ImageProperties?> Image { get; private set; } = null!;

        /// <summary>
        /// URL of the image to deploy for an image-backed service. The host, repository, and image name must match the currently configured image for the service.
        /// </summary>
        [Output("imageUrl")]
        public Output<string?> ImageUrl { get; private set; } = null!;

        [Output("status")]
        public Output<Pulumi.Render.Services.Status?> Status { get; private set; } = null!;

        [Output("trigger")]
        public Output<Pulumi.Render.Services.Trigger?> Trigger { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string?> UpdatedAt { get; private set; } = null!;


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
                PluginDownloadURL = "github://api.github.com/cloudy-sky-software/pulumi-render",
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

    public sealed class DeployArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defaults to "do_not_clear"
        /// </summary>
        [Input("clearCache")]
        public Input<Pulumi.Render.Services.ClearCache>? ClearCache { get; set; }

        /// <summary>
        /// Specific ID of commit to deploy for a web service, defaults to latest commit. Not supported for Cron Job deploys.
        /// </summary>
        [Input("commitId")]
        public Input<string>? CommitId { get; set; }

        /// <summary>
        /// URL of the image to deploy for an image-backed service. The host, repository, and image name must match the currently configured image for the service.
        /// </summary>
        [Input("imageUrl")]
        public Input<string>? ImageUrl { get; set; }

        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        public DeployArgs()
        {
            ClearCache = Pulumi.Render.Services.ClearCache.DoNotClear;
        }
        public static new DeployArgs Empty => new DeployArgs();
    }
}
