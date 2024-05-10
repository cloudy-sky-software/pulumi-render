// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class ListEnvVarsForService
    {
        public static Task<ListEnvVarsForServiceResult> InvokeAsync(ListEnvVarsForServiceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListEnvVarsForServiceResult>("render:services:listEnvVarsForService", args ?? new ListEnvVarsForServiceArgs(), options.WithDefaults());

        public static Output<ListEnvVarsForServiceResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListEnvVarsForServiceResult>("render:services:listEnvVarsForService", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListEnvVarsForServiceArgs : global::Pulumi.InvokeArgs
    {
        public ListEnvVarsForServiceArgs()
        {
        }
        public static new ListEnvVarsForServiceArgs Empty => new ListEnvVarsForServiceArgs();
    }


    [OutputType]
    public sealed class ListEnvVarsForServiceResult
    {
        public readonly ImmutableArray<Outputs.EnvVarWithCursor> Items;

        [OutputConstructor]
        private ListEnvVarsForServiceResult(ImmutableArray<Outputs.EnvVarWithCursor> items)
        {
            Items = items;
        }
    }
}
