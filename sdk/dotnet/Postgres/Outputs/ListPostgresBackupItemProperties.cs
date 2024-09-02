// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Postgres.Outputs
{

    [OutputType]
    public sealed class ListPostgresBackupItemProperties
    {
        public readonly string CreatedAt;
        public readonly string Id;
        /// <summary>
        /// URL to download the Postgres backup
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private ListPostgresBackupItemProperties(
            string createdAt,

            string id,

            string? url)
        {
            CreatedAt = createdAt;
            Id = id;
            Url = url;
        }
    }
}
