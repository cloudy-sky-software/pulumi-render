// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    [RenderResourceType("render:services:EnvVarsForService")]
    public partial class EnvVarsForService : global::Pulumi.CustomResource
    {
        [Output("envVars")]
        public Output<ImmutableArray<Union<Outputs.EnvVarKeyValue, Outputs.EnvVarKeyGenerateValue>>> EnvVars { get; private set; } = null!;


        /// <summary>
        /// Create a EnvVarsForService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EnvVarsForService(string name, EnvVarsForServiceArgs? args = null, CustomResourceOptions? options = null)
            : base("render:services:EnvVarsForService", name, args ?? new EnvVarsForServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EnvVarsForService(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("render:services:EnvVarsForService", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing EnvVarsForService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EnvVarsForService Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EnvVarsForService(name, id, options);
        }
    }

    public sealed class EnvVarsForServiceArgs : global::Pulumi.ResourceArgs
    {
        [Input("envVars")]
        private InputList<Union<Inputs.EnvVarKeyValueArgs, Inputs.EnvVarKeyGenerateValueArgs>>? _envVars;
        public InputList<Union<Inputs.EnvVarKeyValueArgs, Inputs.EnvVarKeyGenerateValueArgs>> EnvVars
        {
            get => _envVars ?? (_envVars = new InputList<Union<Inputs.EnvVarKeyValueArgs, Inputs.EnvVarKeyGenerateValueArgs>>());
            set => _envVars = value;
        }

        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        public EnvVarsForServiceArgs()
        {
        }
        public static new EnvVarsForServiceArgs Empty => new EnvVarsForServiceArgs();
    }
}
