// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    [RenderResourceType("render:services:CustomDomain")]
    public partial class CustomDomain : global::Pulumi.CustomResource
    {
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a CustomDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomDomain(string name, CustomDomainArgs? args = null, CustomResourceOptions? options = null)
            : base("render:services:CustomDomain", name, args ?? new CustomDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomDomain(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("render:services:CustomDomain", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing CustomDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomDomain Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CustomDomain(name, id, options);
        }
    }

    public sealed class CustomDomainArgs : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the service
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        public CustomDomainArgs()
        {
        }
        public static new CustomDomainArgs Empty => new CustomDomainArgs();
    }
}
