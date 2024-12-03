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
    public sealed class JobWithCursorpropertiesjob
    {
        public readonly string CreatedAt;
        public readonly string? FinishedAt;
        public readonly string Id;
        public readonly string PlanId;
        public readonly string ServiceId;
        public readonly string StartCommand;
        public readonly string? StartedAt;
        public readonly Pulumi.Render.Services.JobWithCursorpropertiesjobStatus? Status;

        [OutputConstructor]
        private JobWithCursorpropertiesjob(
            string createdAt,

            string? finishedAt,

            string id,

            string planId,

            string serviceId,

            string startCommand,

            string? startedAt,

            Pulumi.Render.Services.JobWithCursorpropertiesjobStatus? status)
        {
            CreatedAt = createdAt;
            FinishedAt = finishedAt;
            Id = id;
            PlanId = planId;
            ServiceId = serviceId;
            StartCommand = startCommand;
            StartedAt = startedAt;
            Status = status;
        }
    }
}