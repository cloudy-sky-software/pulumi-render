// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Outputs
{

    [OutputType]
    public sealed class Deploy
    {
        public readonly Pulumi.Render.Services.DeployClearCache? ClearCache;
        public readonly Outputs.Commit? Commit;

        [OutputConstructor]
        private Deploy(
            Pulumi.Render.Services.DeployClearCache? clearCache,

            Outputs.Commit? commit)
        {
            ClearCache = clearCache;
            Commit = commit;
        }
    }
}
