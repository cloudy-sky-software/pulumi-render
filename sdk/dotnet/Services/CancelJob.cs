// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    [RenderResourceType("render:services:CancelJob")]
    public partial class CancelJob : global::Pulumi.CustomResource
    {
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("finishedAt")]
        public Output<string?> FinishedAt { get; private set; } = null!;

        [Output("planId")]
        public Output<string> PlanId { get; private set; } = null!;

        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;

        [Output("startCommand")]
        public Output<string> StartCommand { get; private set; } = null!;

        [Output("startedAt")]
        public Output<string?> StartedAt { get; private set; } = null!;

        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;


        /// <summary>
        /// Create a CancelJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CancelJob(string name, CancelJobArgs? args = null, CustomResourceOptions? options = null)
            : base("render:services:CancelJob", name, args ?? new CancelJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CancelJob(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("render:services:CancelJob", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing CancelJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CancelJob Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CancelJob(name, id, options);
        }
    }

    public sealed class CancelJobArgs : global::Pulumi.ResourceArgs
    {
        public CancelJobArgs()
        {
        }
        public static new CancelJobArgs Empty => new CancelJobArgs();
    }
}
