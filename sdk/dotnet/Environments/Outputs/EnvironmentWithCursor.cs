// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Environments.Outputs
{

    /// <summary>
    /// An environment with a cursor
    /// </summary>
    [OutputType]
    public sealed class EnvironmentWithCursor
    {
        public readonly string Cursor;
        public readonly Outputs.Environment Environment;

        [OutputConstructor]
        private EnvironmentWithCursor(
            string cursor,

            Outputs.Environment environment)
        {
            Cursor = cursor;
            Environment = environment;
        }
    }
}