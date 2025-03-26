// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.KeyValue.Outputs
{

    [OutputType]
    public sealed class KeyValueWithCursor
    {
        public readonly string Cursor;
        /// <summary>
        /// A Key Value instance
        /// </summary>
        public readonly Outputs.KeyValue KeyValue;

        [OutputConstructor]
        private KeyValueWithCursor(
            string cursor,

            Outputs.KeyValue keyValue)
        {
            Cursor = cursor;
            KeyValue = keyValue;
        }
    }
}
