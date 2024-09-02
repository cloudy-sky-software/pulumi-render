// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Environments
{
    [RenderResourceType("render:environments:ResourcesToEnvironment")]
    public partial class ResourcesToEnvironment : global::Pulumi.CustomResource
    {
        [Output("databasesIds")]
        public Output<ImmutableArray<string>> DatabasesIds { get; private set; } = null!;

        [Output("envGroupIds")]
        public Output<ImmutableArray<string>> EnvGroupIds { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether an environment is `unprotected` or `protected`. Only admin users can perform destructive actions in `protected` environments.
        /// </summary>
        [Output("protectedStatus")]
        public Output<Pulumi.Render.Environments.ProtectedStatus> ProtectedStatus { get; private set; } = null!;

        [Output("redisIds")]
        public Output<ImmutableArray<string>> RedisIds { get; private set; } = null!;

        [Output("resourceIds")]
        public Output<ImmutableArray<string>> ResourceIds { get; private set; } = null!;

        [Output("serviceIds")]
        public Output<ImmutableArray<string>> ServiceIds { get; private set; } = null!;


        /// <summary>
        /// Create a ResourcesToEnvironment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourcesToEnvironment(string name, ResourcesToEnvironmentArgs args, CustomResourceOptions? options = null)
            : base("render:environments:ResourcesToEnvironment", name, args ?? new ResourcesToEnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourcesToEnvironment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("render:environments:ResourcesToEnvironment", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ResourcesToEnvironment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourcesToEnvironment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ResourcesToEnvironment(name, id, options);
        }
    }

    public sealed class ResourcesToEnvironmentArgs : global::Pulumi.ResourceArgs
    {
        [Input("environmentId")]
        public Input<string>? EnvironmentId { get; set; }

        [Input("resourceIds", required: true)]
        private InputList<string>? _resourceIds;
        public InputList<string> ResourceIds
        {
            get => _resourceIds ?? (_resourceIds = new InputList<string>());
            set => _resourceIds = value;
        }

        public ResourcesToEnvironmentArgs()
        {
        }
        public static new ResourcesToEnvironmentArgs Empty => new ResourcesToEnvironmentArgs();
    }
}
