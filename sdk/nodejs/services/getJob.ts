// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getJob(args: GetJobArgs, opts?: pulumi.InvokeOptions): Promise<GetJobResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:services:getJob", {
        "id": args.id,
        "serviceId": args.serviceId,
    }, opts);
}

export interface GetJobArgs {
    /**
     * (Required) The ID of the job
     */
    id: string;
    /**
     * (Required) The ID of the service
     */
    serviceId: string;
}

export interface GetJobResult {
    readonly items: outputs.services.Job;
}
export function getJobOutput(args: GetJobOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetJobResult> {
    return pulumi.output(args).apply((a: any) => getJob(a, opts))
}

export interface GetJobOutputArgs {
    /**
     * (Required) The ID of the job
     */
    id: pulumi.Input<string>;
    /**
     * (Required) The ID of the service
     */
    serviceId: pulumi.Input<string>;
}
