// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Postgres
{
    [RenderResourceType("render:postgres:RecoverPostgres")]
    public partial class RecoverPostgres : global::Pulumi.CustomResource
    {
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The URL to view the Postgres instance in the Render Dashboard
        /// </summary>
        [Output("dashboardUrl")]
        public Output<string> DashboardUrl { get; private set; } = null!;

        [Output("databaseName")]
        public Output<string> DatabaseName { get; private set; } = null!;

        [Output("databaseUser")]
        public Output<string> DatabaseUser { get; private set; } = null!;

        /// <summary>
        /// Datadog API key to use for monitoring the new database. Defaults to the API key of the original database. Use an empty string to prevent copying of the API key to the new database.
        /// </summary>
        [Output("datadogApiKey")]
        public Output<string?> DatadogApiKey { get; private set; } = null!;

        [Output("diskSizeGB")]
        public Output<int?> DiskSizeGB { get; private set; } = null!;

        [Output("environmentId")]
        public Output<string?> EnvironmentId { get; private set; } = null!;

        /// <summary>
        /// The time at which the database will be expire. Applies to free tier databases only.
        /// </summary>
        [Output("expiresAt")]
        public Output<string?> ExpiresAt { get; private set; } = null!;

        [Output("highAvailabilityEnabled")]
        public Output<bool> HighAvailabilityEnabled { get; private set; } = null!;

        [Output("ipAllowList")]
        public Output<ImmutableArray<Outputs.CidrBlockAndDescription>> IpAllowList { get; private set; } = null!;

        [Output("maintenance")]
        public Output<Outputs.RedisDetailpropertiesmaintenance?> Maintenance { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("owner")]
        public Output<Outputs.Owner> Owner { get; private set; } = null!;

        [Output("plan")]
        public Output<Pulumi.Render.Postgres.Plan> Plan { get; private set; } = null!;

        [Output("primaryPostgresID")]
        public Output<string?> PrimaryPostgresID { get; private set; } = null!;

        [Output("readReplicas")]
        public Output<ImmutableArray<Outputs.ReadReplica>> ReadReplicas { get; private set; } = null!;

        /// <summary>
        /// Defaults to "oregon"
        /// </summary>
        [Output("region")]
        public Output<Pulumi.Render.Postgres.Region> Region { get; private set; } = null!;

        /// <summary>
        /// Name of the new database.
        /// </summary>
        [Output("restoreName")]
        public Output<string?> RestoreName { get; private set; } = null!;

        /// <summary>
        /// The point in time to restore the database to. See `/recovery-info` for restore availability
        /// </summary>
        [Output("restoreTime")]
        public Output<string> RestoreTime { get; private set; } = null!;

        [Output("role")]
        public Output<Pulumi.Render.Postgres.Role> Role { get; private set; } = null!;

        [Output("status")]
        public Output<Pulumi.Render.Postgres.Status> Status { get; private set; } = null!;

        [Output("suspended")]
        public Output<Pulumi.Render.Postgres.RecoverPostgresSuspended> Suspended { get; private set; } = null!;

        [Output("suspenders")]
        public Output<ImmutableArray<Pulumi.Render.Postgres.SuspendersItem>> Suspenders { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The PostgreSQL version
        /// </summary>
        [Output("version")]
        public Output<Pulumi.Render.Postgres.Version> Version { get; private set; } = null!;


        /// <summary>
        /// Create a RecoverPostgres resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RecoverPostgres(string name, RecoverPostgresArgs args, CustomResourceOptions? options = null)
            : base("render:postgres:RecoverPostgres", name, args ?? new RecoverPostgresArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RecoverPostgres(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("render:postgres:RecoverPostgres", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing RecoverPostgres resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RecoverPostgres Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RecoverPostgres(name, id, options);
        }
    }

    public sealed class RecoverPostgresArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Datadog API key to use for monitoring the new database. Defaults to the API key of the original database. Use an empty string to prevent copying of the API key to the new database.
        /// </summary>
        [Input("datadogApiKey")]
        public Input<string>? DatadogApiKey { get; set; }

        /// <summary>
        /// The environment to create the new database in. Defaults to the environment of the original database.
        /// </summary>
        [Input("environmentId")]
        public Input<string>? EnvironmentId { get; set; }

        /// <summary>
        /// The plan to use for the new database. Defaults to the same plan as the original database. Cannot be a lower tier plan than the original database.
        /// </summary>
        [Input("plan")]
        public Input<string>? Plan { get; set; }

        [Input("postgresId")]
        public Input<string>? PostgresId { get; set; }

        /// <summary>
        /// Name of the new database.
        /// </summary>
        [Input("restoreName")]
        public Input<string>? RestoreName { get; set; }

        /// <summary>
        /// The point in time to restore the database to. See `/recovery-info` for restore availability
        /// </summary>
        [Input("restoreTime", required: true)]
        public Input<string> RestoreTime { get; set; } = null!;

        public RecoverPostgresArgs()
        {
        }
        public static new RecoverPostgresArgs Empty => new RecoverPostgresArgs();
    }
}
