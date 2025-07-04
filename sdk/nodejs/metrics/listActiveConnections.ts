// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listActiveConnections(args?: ListActiveConnectionsArgs, opts?: pulumi.InvokeOptions): Promise<ListActiveConnectionsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:metrics:listActiveConnections", {
    }, opts);
}

export interface ListActiveConnectionsArgs {
}

export interface ListActiveConnectionsResult {
    readonly items: outputs.metrics.ListActiveConnectionsItemProperties[];
}
export function listActiveConnectionsOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListActiveConnectionsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("render:metrics:listActiveConnections", {
    }, opts);
}

