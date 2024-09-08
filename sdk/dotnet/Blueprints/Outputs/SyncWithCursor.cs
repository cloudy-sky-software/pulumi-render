// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Blueprints.Outputs
{

    /// <summary>
    /// A blueprint sync with a cursor
    /// </summary>
    [OutputType]
    public sealed class SyncWithCursor
    {
        public readonly string Cursor;
        public readonly Outputs.SyncWithCursorSyncProperties Sync;

        [OutputConstructor]
        private SyncWithCursor(
            string cursor,

            Outputs.SyncWithCursorSyncProperties sync)
        {
            Cursor = cursor;
            Sync = sync;
        }
    }
}