// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    [RenderResourceType("render:services:RestartServer")]
    public partial class RestartServer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Create a RestartServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RestartServer(string name, RestartServerArgs? args = null, CustomResourceOptions? options = null)
            : base("render:services:RestartServer", name, args ?? new RestartServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RestartServer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("render:services:RestartServer", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing RestartServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RestartServer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RestartServer(name, id, options);
        }
    }

    public sealed class RestartServerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        public RestartServerArgs()
        {
        }
        public static new RestartServerArgs Empty => new RestartServerArgs();
    }
}
