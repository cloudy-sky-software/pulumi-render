// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listDeploys(args: ListDeploysArgs, opts?: pulumi.InvokeOptions): Promise<ListDeploysResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:services:listDeploys", {
        "serviceId": args.serviceId,
    }, opts);
}

export interface ListDeploysArgs {
    /**
     * (Required) The ID of the service
     */
    serviceId: string;
}

export interface ListDeploysResult {
    readonly items: outputs.services.ListDeploysResponse[];
}
export function listDeploysOutput(args: ListDeploysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListDeploysResult> {
    return pulumi.output(args).apply((a: any) => listDeploys(a, opts))
}

export interface ListDeploysOutputArgs {
    /**
     * (Required) The ID of the service
     */
    serviceId: pulumi.Input<string>;
}
