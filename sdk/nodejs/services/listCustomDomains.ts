// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listCustomDomains(args: ListCustomDomainsArgs, opts?: pulumi.InvokeOptions): Promise<ListCustomDomainsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:services:listCustomDomains", {
        "serviceId": args.serviceId,
    }, opts);
}

export interface ListCustomDomainsArgs {
    /**
     * The ID of the service
     */
    serviceId: string;
}

export interface ListCustomDomainsResult {
    readonly items: outputs.services.CustomDomainWithCursor[];
}
export function listCustomDomainsOutput(args: ListCustomDomainsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListCustomDomainsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("render:services:listCustomDomains", {
        "serviceId": args.serviceId,
    }, opts);
}

export interface ListCustomDomainsOutputArgs {
    /**
     * The ID of the service
     */
    serviceId: pulumi.Input<string>;
}
