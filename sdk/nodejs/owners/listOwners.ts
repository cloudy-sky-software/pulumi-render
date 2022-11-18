// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listOwners(args?: ListOwnersArgs, opts?: pulumi.InvokeOptions): Promise<ListOwnersResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:owners:listOwners", {
    }, opts);
}

export interface ListOwnersArgs {
}

export interface ListOwnersResult {
    readonly items: outputs.owners.ListOwnersResponse[];
}
