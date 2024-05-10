// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listCustomDomains(args?: ListCustomDomainsArgs, opts?: pulumi.InvokeOptions): Promise<ListCustomDomainsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:services:listCustomDomains", {
    }, opts);
}

export interface ListCustomDomainsArgs {
}

export interface ListCustomDomainsResult {
    readonly items: outputs.services.ListCustomDomainsItemProperties[];
}
export function listCustomDomainsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<ListCustomDomainsResult> {
    return pulumi.output(listCustomDomains(opts))
}
