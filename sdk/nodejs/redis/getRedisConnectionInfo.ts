// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getRedisConnectionInfo(args: GetRedisConnectionInfoArgs, opts?: pulumi.InvokeOptions): Promise<GetRedisConnectionInfoResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:redis:getRedisConnectionInfo", {
        "redisId": args.redisId,
    }, opts);
}

export interface GetRedisConnectionInfoArgs {
    redisId: string;
}

/**
 * A Redis instance
 */
export interface GetRedisConnectionInfoResult {
    /**
     * The connection string to use from outside Render
     */
    readonly externalConnectionString: string;
    /**
     * The connection string to use from within Render
     */
    readonly internalConnectionString: string;
    /**
     * The Redis CLI command to connect to the Redis instance
     */
    readonly redisCLICommand: string;
}
export function getRedisConnectionInfoOutput(args: GetRedisConnectionInfoOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRedisConnectionInfoResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("render:redis:getRedisConnectionInfo", {
        "redisId": args.redisId,
    }, opts);
}

export interface GetRedisConnectionInfoOutputArgs {
    redisId: pulumi.Input<string>;
}
