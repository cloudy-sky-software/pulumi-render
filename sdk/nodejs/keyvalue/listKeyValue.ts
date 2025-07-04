// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listKeyValue(args?: ListKeyValueArgs, opts?: pulumi.InvokeOptions): Promise<ListKeyValueResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:key-value:listKeyValue", {
    }, opts);
}

export interface ListKeyValueArgs {
}

export interface ListKeyValueResult {
    readonly items: outputs.keyvalue.KeyValueWithCursor[];
}
export function listKeyValueOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListKeyValueResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("render:key-value:listKeyValue", {
    }, opts);
}

