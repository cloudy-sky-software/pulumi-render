// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getCustomDomain(args: GetCustomDomainArgs, opts?: pulumi.InvokeOptions): Promise<GetCustomDomainResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:services:getCustomDomain", {
        "id": args.id,
        "serviceId": args.serviceId,
    }, opts);
}

export interface GetCustomDomainArgs {
    /**
     * (Required) The ID or name of the custom domain
     */
    id: string;
    /**
     * (Required) The ID of the service
     */
    serviceId: string;
}

export interface GetCustomDomainResult {
    readonly items: outputs.services.CustomDomain;
}

export function getCustomDomainOutput(args: GetCustomDomainOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCustomDomainResult> {
    return pulumi.output(args).apply(a => getCustomDomain(a, opts))
}

export interface GetCustomDomainOutputArgs {
    /**
     * (Required) The ID or name of the custom domain
     */
    id: pulumi.Input<string>;
    /**
     * (Required) The ID of the service
     */
    serviceId: pulumi.Input<string>;
}
