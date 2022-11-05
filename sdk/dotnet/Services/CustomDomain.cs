// *** WARNING: this file was generated by pulumigen. ***
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
        [Output("createdAt")]
        public Output<string?> CreatedAt { get; private set; } = null!;

        [Output("domainType")]
        public Output<Pulumi.Render.Services.DomainType> DomainType { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("publicSuffix")]
        public Output<string?> PublicSuffix { get; private set; } = null!;

        [Output("redirectForName")]
        public Output<string> RedirectForName { get; private set; } = null!;

        [Output("server")]
        public Output<Outputs.ServerProperties> Server { get; private set; } = null!;

        [Output("verificationStatus")]
        public Output<Pulumi.Render.Services.VerificationStatus> VerificationStatus { get; private set; } = null!;


        /// <summary>
        /// Create a CustomDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomDomain(string name, CustomDomainArgs args, CustomResourceOptions? options = null)
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
                PluginDownloadURL = "github://cloudy-sky-software/pulumi-render",
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
        /// <summary>
        /// (Required) The ID of the service
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("server", required: true)]
        public Input<Inputs.ServerPropertiesArgs> Server { get; set; } = null!;

        public CustomDomainArgs()
        {
        }
        public static new CustomDomainArgs Empty => new CustomDomainArgs();
    }
}
