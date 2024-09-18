// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getEnvGroupEnvVar(args: GetEnvGroupEnvVarArgs, opts?: pulumi.InvokeOptions): Promise<outputs.envgroups.EnvVar> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:env-groups:getEnvGroupEnvVar", {
        "envGroupId": args.envGroupId,
        "envVarKey": args.envVarKey,
    }, opts);
}

export interface GetEnvGroupEnvVarArgs {
    /**
     * Filter for resources that belong to an environment group
     */
    envGroupId: string;
    /**
     * The name of the environment variable
     */
    envVarKey: string;
}
export function getEnvGroupEnvVarOutput(args: GetEnvGroupEnvVarOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.envgroups.EnvVar> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("render:env-groups:getEnvGroupEnvVar", {
        "envGroupId": args.envGroupId,
        "envVarKey": args.envVarKey,
    }, opts);
}

export interface GetEnvGroupEnvVarOutputArgs {
    /**
     * Filter for resources that belong to an environment group
     */
    envGroupId: pulumi.Input<string>;
    /**
     * The name of the environment variable
     */
    envVarKey: pulumi.Input<string>;
}
