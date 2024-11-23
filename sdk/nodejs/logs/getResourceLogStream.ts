// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getResourceLogStream(args: GetResourceLogStreamArgs, opts?: pulumi.InvokeOptions): Promise<outputs.logs.GetResourceLogStreamProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:logs:getResourceLogStream", {
        "resourceId": args.resourceId,
    }, opts);
}

export interface GetResourceLogStreamArgs {
    /**
     * The ID of the resource (server, cron job, postgres, or redis) whose log streams should be returned
     */
    resourceId: string;
}
export function getResourceLogStreamOutput(args: GetResourceLogStreamOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.logs.GetResourceLogStreamProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("render:logs:getResourceLogStream", {
        "resourceId": args.resourceId,
    }, opts);
}

export interface GetResourceLogStreamOutputArgs {
    /**
     * The ID of the resource (server, cron job, postgres, or redis) whose log streams should be returned
     */
    resourceId: pulumi.Input<string>;
}
