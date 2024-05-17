// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getStaticSiteOutput(args: GetStaticSiteOutputArgs, opts?: pulumi.InvokeOptions): Promise<GetStaticSiteOutputResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:services:getStaticSiteOutput", {
        "serviceId": args.serviceId,
    }, opts);
}

export interface GetStaticSiteOutputArgs {
    /**
     * The ID of the service
     */
    serviceId: string;
}

export interface GetStaticSiteOutputResult {
    readonly items: outputs.services.GetStaticSiteOutput;
}
export function getStaticSiteOutputOutput(args: GetStaticSiteOutputOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStaticSiteOutputResult> {
    return pulumi.output(args).apply((a: any) => getStaticSiteOutput(a, opts))
}

export interface GetStaticSiteOutputOutputArgs {
    /**
     * The ID of the service
     */
    serviceId: pulumi.Input<string>;
}
