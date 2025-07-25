// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    [RenderResourceType("render:services:Header")]
    public partial class Header : global::Pulumi.CustomResource
    {
        [Output("headers")]
        public Output<Outputs.Header?> Headers { get; private set; } = null!;

        /// <summary>
        /// Header name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The request path to add the header to. Wildcards will cause headers to be applied to all matching paths.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// Header value
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a Header resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Header(string name, HeaderArgs args, CustomResourceOptions? options = null)
            : base("render:services:Header", name, args ?? new HeaderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Header(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("render:services:Header", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Header resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Header Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Header(name, id, options);
        }
    }

    public sealed class HeaderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Header name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The request path to add the header to. Wildcards will cause headers to be applied to all matching paths.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// Header value
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public HeaderArgs()
        {
        }
        public static new HeaderArgs Empty => new HeaderArgs();
    }
}
