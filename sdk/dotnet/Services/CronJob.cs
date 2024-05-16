// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    [RenderResourceType("render:services:CronJob")]
    public partial class CronJob : global::Pulumi.CustomResource
    {
        [Output("autoDeploy")]
        public Output<Pulumi.Render.Services.CronJobAutoDeploy> AutoDeploy { get; private set; } = null!;

        [Output("branch")]
        public Output<string?> Branch { get; private set; } = null!;

        [Output("buildFilter")]
        public Output<Outputs.BuildFilter?> BuildFilter { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("imagePath")]
        public Output<string?> ImagePath { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("notifyOnFail")]
        public Output<Pulumi.Render.Services.CronJobNotifyOnFail> NotifyOnFail { get; private set; } = null!;

        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        [Output("repo")]
        public Output<string?> Repo { get; private set; } = null!;

        [Output("rootDir")]
        public Output<string> RootDir { get; private set; } = null!;

        [Output("serviceDetails")]
        public Output<Outputs.CronJobDetailsCreate> ServiceDetails { get; private set; } = null!;

        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

        [Output("suspended")]
        public Output<Pulumi.Render.Services.CronJobSuspended> Suspended { get; private set; } = null!;

        [Output("suspenders")]
        public Output<ImmutableArray<Pulumi.Render.Services.CronJobSuspendersItem>> Suspenders { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a CronJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CronJob(string name, CronJobArgs? args = null, CustomResourceOptions? options = null)
            : base("render:services:CronJob", name, args ?? new CronJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CronJob(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("render:services:CronJob", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing CronJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CronJob Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CronJob(name, id, options);
        }
    }

    public sealed class CronJobArgs : global::Pulumi.ResourceArgs
    {
        [Input("serviceDetails")]
        public Input<Inputs.CronJobDetailsCreateArgs>? ServiceDetails { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public CronJobArgs()
        {
            Type = "cron_job";
        }
        public static new CronJobArgs Empty => new CronJobArgs();
    }
}
