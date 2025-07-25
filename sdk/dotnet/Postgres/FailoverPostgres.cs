// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Postgres
{
    [RenderResourceType("render:postgres:FailoverPostgres")]
    public partial class FailoverPostgres : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Create a FailoverPostgres resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FailoverPostgres(string name, FailoverPostgresArgs? args = null, CustomResourceOptions? options = null)
            : base("render:postgres:FailoverPostgres", name, args ?? new FailoverPostgresArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FailoverPostgres(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("render:postgres:FailoverPostgres", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing FailoverPostgres resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FailoverPostgres Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FailoverPostgres(name, id, options);
        }
    }

    public sealed class FailoverPostgresArgs : global::Pulumi.ResourceArgs
    {
        [Input("postgresId")]
        public Input<string>? PostgresId { get; set; }

        public FailoverPostgresArgs()
        {
        }
        public static new FailoverPostgresArgs Empty => new FailoverPostgresArgs();
    }
}
