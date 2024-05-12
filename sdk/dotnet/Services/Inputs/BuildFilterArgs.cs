// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Inputs
{

    public sealed class BuildFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("ignoredPaths", required: true)]
        private InputList<string>? _ignoredPaths;
        public InputList<string> IgnoredPaths
        {
            get => _ignoredPaths ?? (_ignoredPaths = new InputList<string>());
            set => _ignoredPaths = value;
        }

        [Input("paths", required: true)]
        private InputList<string>? _paths;
        public InputList<string> Paths
        {
            get => _paths ?? (_paths = new InputList<string>());
            set => _paths = value;
        }

        public BuildFilterArgs()
        {
        }
        public static new BuildFilterArgs Empty => new BuildFilterArgs();
    }
}