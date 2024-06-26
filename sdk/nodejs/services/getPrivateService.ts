// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getPrivateService(args: GetPrivateServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateServiceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:services:getPrivateService", {
        "serviceId": args.serviceId,
    }, opts);
}

export interface GetPrivateServiceArgs {
    /**
     * The ID of the service
     */
    serviceId: string;
}

export interface GetPrivateServiceResult {
    readonly items: outputs.services.GetPrivateService;
}
export function getPrivateServiceOutput(args: GetPrivateServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPrivateServiceResult> {
    return pulumi.output(args).apply((a: any) => getPrivateService(a, opts))
}

export interface GetPrivateServiceOutputArgs {
    /**
     * The ID of the service
     */
    serviceId: pulumi.Input<string>;
}
