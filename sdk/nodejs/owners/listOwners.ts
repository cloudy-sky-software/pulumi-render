// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function listOwners(args?: ListOwnersArgs, opts?: pulumi.InvokeOptions): Promise<ListOwnersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("render:owners:listOwners", {
    }, opts);
}

export interface ListOwnersArgs {
}

export interface ListOwnersResult {
}