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
    public sealed class StaticSiteDetailsOutput
    {
        public readonly string BuildCommand;
        public readonly string BuildPlan;
        public readonly Outputs.Resource? ParentServer;
        public readonly string PublishPath;
        public readonly Pulumi.Render.Services.StaticSiteDetailsOutputPullRequestPreviewsEnabled PullRequestPreviewsEnabled;
        public readonly string Url;

        [OutputConstructor]
        private StaticSiteDetailsOutput(
            string buildCommand,

            string buildPlan,

            Outputs.Resource? parentServer,

            string publishPath,

            Pulumi.Render.Services.StaticSiteDetailsOutputPullRequestPreviewsEnabled pullRequestPreviewsEnabled,

            string url)
        {
            BuildCommand = buildCommand;
            BuildPlan = buildPlan;
            ParentServer = parentServer;
            PublishPath = publishPath;
            PullRequestPreviewsEnabled = pullRequestPreviewsEnabled;
            Url = url;
        }
    }
}
