// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Projects.Inputs
{

    public sealed class ProjectCreateEnvironmentInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Indicates whether an environment is `unprotected` or `protected`. Only admin users can perform destructive actions in `protected` environments.
        /// </summary>
        [Input("protectedStatus")]
        public Input<Pulumi.Render.Projects.ProjectCreateEnvironmentInputProtectedStatus>? ProtectedStatus { get; set; }

        public ProjectCreateEnvironmentInputArgs()
        {
        }
        public static new ProjectCreateEnvironmentInputArgs Empty => new ProjectCreateEnvironmentInputArgs();
    }
}
