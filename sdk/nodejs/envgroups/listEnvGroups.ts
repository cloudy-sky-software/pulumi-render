// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listEnvGroups(args?: ListEnvGroupsArgs, opts?: pulumi.InvokeOptions): Promise<ListEnvGroupsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:env-groups:listEnvGroups", {
    }, opts);
}

export interface ListEnvGroupsArgs {
}

export interface ListEnvGroupsResult {
    readonly items: outputs.envgroups.EnvGroupMeta[];
}
export function listEnvGroupsOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListEnvGroupsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("render:env-groups:listEnvGroups", {
    }, opts);
}

