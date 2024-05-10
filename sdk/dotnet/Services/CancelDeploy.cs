// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    [RenderResourceType("render:services:CancelDeploy")]
    public partial class CancelDeploy : global::Pulumi.CustomResource
    {
        [Output("commit")]
        public Output<Outputs.CommitProperties?> Commit { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string?> CreatedAt { get; private set; } = null!;

        [Output("finishedAt")]
        public Output<string?> FinishedAt { get; private set; } = null!;

        /// <summary>
        /// Image information used when creating the deploy. Not present for Git-backed deploys
        /// </summary>
        [Output("image")]
        public Output<Outputs.ImageProperties?> Image { get; private set; } = null!;

        [Output("status")]
        public Output<Pulumi.Render.Services.Status?> Status { get; private set; } = null!;

        [Output("trigger")]
        public Output<Pulumi.Render.Services.Trigger?> Trigger { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string?> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a CancelDeploy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CancelDeploy(string name, CancelDeployArgs? args = null, CustomResourceOptions? options = null)
            : base("render:services:CancelDeploy", name, args ?? new CancelDeployArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CancelDeploy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("render:services:CancelDeploy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing CancelDeploy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CancelDeploy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CancelDeploy(name, id, options);
        }
    }

    public sealed class CancelDeployArgs : global::Pulumi.ResourceArgs
    {
        public CancelDeployArgs()
        {
        }
        public static new CancelDeployArgs Empty => new CancelDeployArgs();
    }
}
