// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Owners.Outputs
{

    [OutputType]
    public sealed class ListOwnersResponse
    {
        public readonly string? Cursor;
        /// <summary>
        /// The owner object represents an authorized user or team. The `type` property indicates if the owner is a user or team.
        /// </summary>
        public readonly Outputs.Owner? Owner;

        [OutputConstructor]
        private ListOwnersResponse(
            string? cursor,

            Outputs.Owner? owner)
        {
            Cursor = cursor;
            Owner = owner;
        }
    }
}
