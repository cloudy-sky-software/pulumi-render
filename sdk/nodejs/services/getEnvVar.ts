// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getEnvVar(args: GetEnvVarArgs, opts?: pulumi.InvokeOptions): Promise<outputs.services.EnvVar> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:services:getEnvVar", {
        "envVarKey": args.envVarKey,
        "serviceId": args.serviceId,
    }, opts);
}

export interface GetEnvVarArgs {
    /**
     * The name of the environment variable
     */
    envVarKey: string;
    /**
     * The ID of the service
     */
    serviceId: string;
}
export function getEnvVarOutput(args: GetEnvVarOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.services.EnvVar> {
    return pulumi.output(args).apply((a: any) => getEnvVar(a, opts))
}

export interface GetEnvVarOutputArgs {
    /**
     * The name of the environment variable
     */
    envVarKey: pulumi.Input<string>;
    /**
     * The ID of the service
     */
    serviceId: pulumi.Input<string>;
}