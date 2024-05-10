// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class GetJob
    {
        public static Task<GetJobResult> InvokeAsync(GetJobArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetJobResult>("render:services:getJob", args ?? new GetJobArgs(), options.WithDefaults());

        public static Output<GetJobResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetJobResult>("render:services:getJob", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetJobArgs : global::Pulumi.InvokeArgs
    {
        public GetJobArgs()
        {
        }
        public static new GetJobArgs Empty => new GetJobArgs();
    }


    [OutputType]
    public sealed class GetJobResult
    {
        public readonly Outputs.Job Items;

        [OutputConstructor]
        private GetJobResult(Outputs.Job items)
        {
            Items = items;
        }
    }
}
