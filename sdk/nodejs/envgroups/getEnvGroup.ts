// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getEnvGroup(args: GetEnvGroupArgs, opts?: pulumi.InvokeOptions): Promise<outputs.envgroups.EnvGroup> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:env-groups:getEnvGroup", {
        "envGroupId": args.envGroupId,
    }, opts);
}

export interface GetEnvGroupArgs {
    /**
     * Filter for resources that belong to an environment group
     */
    envGroupId: string;
}
export function getEnvGroupOutput(args: GetEnvGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.envgroups.EnvGroup> {
    return pulumi.output(args).apply((a: any) => getEnvGroup(a, opts))
}

export interface GetEnvGroupOutputArgs {
    /**
     * Filter for resources that belong to an environment group
     */
    envGroupId: pulumi.Input<string>;
}