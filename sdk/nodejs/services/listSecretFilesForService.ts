// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listSecretFilesForService(args: ListSecretFilesForServiceArgs, opts?: pulumi.InvokeOptions): Promise<ListSecretFilesForServiceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:services:listSecretFilesForService", {
        "serviceId": args.serviceId,
    }, opts);
}

export interface ListSecretFilesForServiceArgs {
    /**
     * The ID of the service
     */
    serviceId: string;
}

export interface ListSecretFilesForServiceResult {
    readonly items: outputs.services.SecretFileWithCursor[];
}
export function listSecretFilesForServiceOutput(args: ListSecretFilesForServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListSecretFilesForServiceResult> {
    return pulumi.output(args).apply((a: any) => listSecretFilesForService(a, opts))
}

export interface ListSecretFilesForServiceOutputArgs {
    /**
     * The ID of the service
     */
    serviceId: pulumi.Input<string>;
}