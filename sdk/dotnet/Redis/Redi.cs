// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Redis
{
    /// <summary>
    /// Input type for creating a Redis instance
    /// </summary>
    [RenderResourceType("render:redis:Redi")]
    public partial class Redi : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The creation time of the Redis instance
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The ID of the environment the Redis instance is associated with
        /// </summary>
        [Output("environmentId")]
        public Output<string?> EnvironmentId { get; private set; } = null!;

        /// <summary>
        /// The IP allow list for the Redis instance
        /// </summary>
        [Output("ipAllowList")]
        public Output<ImmutableArray<Outputs.CidrBlockAndDescription>> IpAllowList { get; private set; } = null!;

        [Output("maintenance")]
        public Output<Outputs.MaintenanceProperties?> Maintenance { get; private set; } = null!;

        /// <summary>
        /// The eviction policy for the Redis instance
        /// </summary>
        [Output("maxmemoryPolicy")]
        public Output<Pulumi.Render.Redis.MaxmemoryPolicy?> MaxmemoryPolicy { get; private set; } = null!;

        /// <summary>
        /// The name of the Redis instance
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Options for a Redis instance
        /// </summary>
        [Output("options")]
        public Output<Outputs.RedisOptions> Options { get; private set; } = null!;

        [Output("owner")]
        public Output<Outputs.Owner> Owner { get; private set; } = null!;

        /// <summary>
        /// The ID of the owner of the Redis instance
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        [Output("plan")]
        public Output<Pulumi.Render.Redis.Plan> Plan { get; private set; } = null!;

        /// <summary>
        /// Defaults to "oregon"
        /// </summary>
        [Output("region")]
        public Output<Pulumi.Render.Redis.Region> Region { get; private set; } = null!;

        [Output("status")]
        public Output<Pulumi.Render.Redis.Status> Status { get; private set; } = null!;

        /// <summary>
        /// The last updated time of the Redis instance
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The version of Redis
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Redi resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Redi(string name, RediArgs args, CustomResourceOptions? options = null)
            : base("render:redis:Redi", name, args ?? new RediArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Redi(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("render:redis:Redi", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Redi resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Redi Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Redi(name, id, options);
        }
    }

    public sealed class RediArgs : global::Pulumi.ResourceArgs
    {
        [Input("environmentId")]
        public Input<string>? EnvironmentId { get; set; }

        [Input("ipAllowList")]
        private InputList<Inputs.CidrBlockAndDescriptionArgs>? _ipAllowList;
        public InputList<Inputs.CidrBlockAndDescriptionArgs> IpAllowList
        {
            get => _ipAllowList ?? (_ipAllowList = new InputList<Inputs.CidrBlockAndDescriptionArgs>());
            set => _ipAllowList = value;
        }

        /// <summary>
        /// The eviction policy for the Redis instance
        /// </summary>
        [Input("maxmemoryPolicy")]
        public Input<Pulumi.Render.Redis.MaxmemoryPolicy>? MaxmemoryPolicy { get; set; }

        /// <summary>
        /// The name of the Redis instance
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the owner of the Redis instance
        /// </summary>
        [Input("ownerId", required: true)]
        public Input<string> OwnerId { get; set; } = null!;

        [Input("plan", required: true)]
        public Input<Pulumi.Render.Redis.Plan> Plan { get; set; } = null!;

        /// <summary>
        /// The region where the Redis instance is located
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public RediArgs()
        {
        }
        public static new RediArgs Empty => new RediArgs();
    }
}
