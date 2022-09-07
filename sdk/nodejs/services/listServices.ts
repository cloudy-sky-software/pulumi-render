// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listServices(args?: ListServicesArgs, opts?: pulumi.InvokeOptions): Promise<ListServicesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("render:services:listServices", {
    }, opts);
}

export interface ListServicesArgs {
}

export interface ListServicesResult {
    readonly items: outputs.services.ListServiceResponse[];
}
