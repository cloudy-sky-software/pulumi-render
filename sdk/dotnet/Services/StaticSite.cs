// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    /// <summary>
    /// A static website service
    /// 
    /// ## Example Usage
    /// </summary>
    [RenderResourceType("render:services:StaticSite")]
    public partial class StaticSite : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to auto deploy the service or not upon git push.
        /// </summary>
        [Output("autoDeploy")]
        public Output<Pulumi.Render.Services.ServiceAutoDeploy?> AutoDeploy { get; private set; } = null!;

        /// <summary>
        /// If left empty, this will fall back to the default branch of the repository.
        /// </summary>
        [Output("branch")]
        public Output<string?> Branch { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string?> CreatedAt { get; private set; } = null!;

        [Output("envVars")]
        public Output<ImmutableArray<Outputs.EnvVarKeyValue>> EnvVars { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The notification setting for this service upon deployment failure.
        /// </summary>
        [Output("notifyOnFail")]
        public Output<Pulumi.Render.Services.ServiceNotifyOnFail?> NotifyOnFail { get; private set; } = null!;

        /// <summary>
        /// The id of the owner (user/team).
        /// </summary>
        [Output("ownerId")]
        public Output<string?> OwnerId { get; private set; } = null!;

        /// <summary>
        /// Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
        /// </summary>
        [Output("repo")]
        public Output<string?> Repo { get; private set; } = null!;

        [Output("secretFiles")]
        public Output<ImmutableArray<Outputs.SecretFile>> SecretFiles { get; private set; } = null!;

        [Output("serviceDetails")]
        public Output<Outputs.StaticSiteServiceDetails?> ServiceDetails { get; private set; } = null!;

        [Output("slug")]
        public Output<string?> Slug { get; private set; } = null!;

        [Output("suspended")]
        public Output<Pulumi.Render.Services.ServiceSuspended?> Suspended { get; private set; } = null!;

        [Output("suspenders")]
        public Output<ImmutableArray<string>> Suspenders { get; private set; } = null!;

        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string?> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a StaticSite resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StaticSite(string name, StaticSiteArgs args, CustomResourceOptions? options = null)
            : base("render:services:StaticSite", name, args ?? new StaticSiteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StaticSite(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("render:services:StaticSite", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://cloudy-sky-software/pulumi-render",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StaticSite resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StaticSite Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StaticSite(name, id, options);
        }
    }

    public sealed class StaticSiteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to auto deploy the service or not upon git push.
        /// </summary>
        [Input("autoDeploy")]
        public Input<Pulumi.Render.Services.ServiceAutoDeploy>? AutoDeploy { get; set; }

        /// <summary>
        /// If left empty, this will fall back to the default branch of the repository.
        /// </summary>
        [Input("branch")]
        public Input<string>? Branch { get; set; }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("envVars")]
        private InputList<Inputs.EnvVarKeyValueArgs>? _envVars;
        public InputList<Inputs.EnvVarKeyValueArgs> EnvVars
        {
            get => _envVars ?? (_envVars = new InputList<Inputs.EnvVarKeyValueArgs>());
            set => _envVars = value;
        }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The notification setting for this service upon deployment failure.
        /// </summary>
        [Input("notifyOnFail")]
        public Input<Pulumi.Render.Services.ServiceNotifyOnFail>? NotifyOnFail { get; set; }

        /// <summary>
        /// The id of the owner (user/team).
        /// </summary>
        [Input("ownerId", required: true)]
        public Input<string> OwnerId { get; set; } = null!;

        /// <summary>
        /// Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
        /// </summary>
        [Input("repo", required: true)]
        public Input<string> Repo { get; set; } = null!;

        [Input("secretFiles")]
        private InputList<Inputs.SecretFileArgs>? _secretFiles;
        public InputList<Inputs.SecretFileArgs> SecretFiles
        {
            get => _secretFiles ?? (_secretFiles = new InputList<Inputs.SecretFileArgs>());
            set => _secretFiles = value;
        }

        [Input("serviceDetails")]
        public Input<Inputs.StaticSiteServiceDetailsArgs>? ServiceDetails { get; set; }

        [Input("slug")]
        public Input<string>? Slug { get; set; }

        [Input("suspended")]
        public Input<Pulumi.Render.Services.ServiceSuspended>? Suspended { get; set; }

        [Input("suspenders")]
        private InputList<string>? _suspenders;
        public InputList<string> Suspenders
        {
            get => _suspenders ?? (_suspenders = new InputList<string>());
            set => _suspenders = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public StaticSiteArgs()
        {
            AutoDeploy = Pulumi.Render.Services.ServiceAutoDeploy.No;
            Type = "static_site";
        }
        public static new StaticSiteArgs Empty => new StaticSiteArgs();
    }
}
