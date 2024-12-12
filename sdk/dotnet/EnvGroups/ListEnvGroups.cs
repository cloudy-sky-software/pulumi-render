// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.EnvGroups
{
    public static class ListEnvGroups
    {
        public static Task<ListEnvGroupsResult> InvokeAsync(ListEnvGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListEnvGroupsResult>("render:env-groups:listEnvGroups", args ?? new ListEnvGroupsArgs(), options.WithDefaults());

        public static Output<ListEnvGroupsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListEnvGroupsResult>("render:env-groups:listEnvGroups", InvokeArgs.Empty, options.WithDefaults());

        public static Output<ListEnvGroupsResult> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListEnvGroupsResult>("render:env-groups:listEnvGroups", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class ListEnvGroupsArgs : global::Pulumi.InvokeArgs
    {
        public ListEnvGroupsArgs()
        {
        }
        public static new ListEnvGroupsArgs Empty => new ListEnvGroupsArgs();
    }


    [OutputType]
    public sealed class ListEnvGroupsResult
    {
        public readonly ImmutableArray<Outputs.EnvGroupMeta> Items;

        [OutputConstructor]
        private ListEnvGroupsResult(ImmutableArray<Outputs.EnvGroupMeta> items)
        {
            Items = items;
        }
    }
}
