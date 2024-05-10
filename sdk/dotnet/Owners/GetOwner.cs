// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Owners
{
    public static class GetOwner
    {
        public static Task<GetOwnerResult> InvokeAsync(GetOwnerArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOwnerResult>("render:owners:getOwner", args ?? new GetOwnerArgs(), options.WithDefaults());

        public static Output<GetOwnerResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOwnerResult>("render:owners:getOwner", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetOwnerArgs : global::Pulumi.InvokeArgs
    {
        public GetOwnerArgs()
        {
        }
        public static new GetOwnerArgs Empty => new GetOwnerArgs();
    }


    [OutputType]
    public sealed class GetOwnerResult
    {
        public readonly Outputs.Owner Items;

        [OutputConstructor]
        private GetOwnerResult(Outputs.Owner items)
        {
            Items = items;
        }
    }
}
