// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

export function getBackgroundWorker(args: GetBackgroundWorkerArgs, opts?: pulumi.InvokeOptions): Promise<GetBackgroundWorkerResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("render:services:getBackgroundWorker", {
        "id": args.id,
    }, opts);
}

export interface GetBackgroundWorkerArgs {
    /**
     * (Required) The ID of the service
     */
    id: string;
}

export interface GetBackgroundWorkerResult {
    readonly items: outputs.services.GetBackgroundWorker;
}

export function getBackgroundWorkerOutput(args: GetBackgroundWorkerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBackgroundWorkerResult> {
    return pulumi.output(args).apply(a => getBackgroundWorker(a, opts))
}

export interface GetBackgroundWorkerOutputArgs {
    /**
     * (Required) The ID of the service
     */
    id: pulumi.Input<string>;
}