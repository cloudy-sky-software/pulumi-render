// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getKeyValue(args: GetKeyValueArgs, opts?: pulumi.InvokeOptions): Promise<GetKeyValueResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:key-value:getKeyValue", {
        "keyValueId": args.keyValueId,
    }, opts);
}

export interface GetKeyValueArgs {
    keyValueId: string;
}

/**
 * A Key Value instance
 */
export interface GetKeyValueResult {
    /**
     * The creation time of the Key Value instance
     */
    readonly createdAt: string;
    /**
     * The ID of the environment the Key Value instance is associated with
     */
    readonly environmentId?: string;
    /**
     * The ID of the Key Value instance
     */
    readonly id: string;
    /**
     * The IP allow list for the Key Value instance
     */
    readonly ipAllowList: outputs.keyvalue.CidrBlockAndDescription[];
    readonly maintenance?: outputs.keyvalue.RedisDetailpropertiesmaintenance;
    /**
     * The name of the Key Value instance
     */
    readonly name: string;
    /**
     * Options for a Key Value instance
     */
    readonly options: outputs.keyvalue.KeyValueOptions;
    readonly owner: outputs.keyvalue.Owner;
    readonly plan: enums.keyvalue.KeyValueDetailPlan;
    /**
     * Defaults to "oregon"
     */
    readonly region: enums.keyvalue.KeyValueDetailRegion;
    readonly status: enums.keyvalue.KeyValueDetailStatus;
    /**
     * The last updated time of the Key Value instance
     */
    readonly updatedAt: string;
    /**
     * The version of Key Value
     */
    readonly version: string;
}
export function getKeyValueOutput(args: GetKeyValueOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetKeyValueResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("render:key-value:getKeyValue", {
        "keyValueId": args.keyValueId,
    }, opts);
}

export interface GetKeyValueOutputArgs {
    keyValueId: pulumi.Input<string>;
}
