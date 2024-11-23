// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getOwnerLogStream(args: GetOwnerLogStreamArgs, opts?: pulumi.InvokeOptions): Promise<outputs.logs.GetOwnerLogStreamProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:logs:getOwnerLogStream", {
        "ownerId": args.ownerId,
    }, opts);
}

export interface GetOwnerLogStreamArgs {
    /**
     * The ID of the owner (team or personal user) whose log streams should be returned
     */
    ownerId: string;
}
export function getOwnerLogStreamOutput(args: GetOwnerLogStreamOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.logs.GetOwnerLogStreamProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("render:logs:getOwnerLogStream", {
        "ownerId": args.ownerId,
    }, opts);
}

export interface GetOwnerLogStreamOutputArgs {
    /**
     * The ID of the owner (team or personal user) whose log streams should be returned
     */
    ownerId: pulumi.Input<string>;
}
