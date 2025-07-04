// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getStaticSite(args: GetStaticSiteArgs, opts?: pulumi.InvokeOptions): Promise<GetStaticSiteResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:services:getStaticSite", {
        "serviceId": args.serviceId,
    }, opts);
}

export interface GetStaticSiteArgs {
    /**
     * The ID of the service
     */
    serviceId: string;
}

export interface GetStaticSiteResult {
    readonly autoDeploy: enums.services.ServiceAutoDeploy;
    readonly branch?: string;
    readonly buildFilter?: outputs.services.BuildFilter;
    readonly createdAt: string;
    /**
     * The URL to view the service in the Render Dashboard
     */
    readonly dashboardUrl: string;
    readonly environmentId?: string;
    readonly id: string;
    readonly imagePath?: string;
    readonly name: string;
    readonly notifyOnFail: enums.services.ServiceNotifyOnFail;
    readonly ownerId: string;
    readonly registryCredential?: outputs.services.RegistryCredentialSummary;
    readonly repo?: string;
    readonly rootDir: string;
    readonly serviceDetails?: outputs.services.StaticSiteDetailsOutput;
    readonly slug: string;
    readonly suspended: enums.services.ServiceSuspended;
    readonly suspenders: enums.services.ServiceSuspendersItem[];
    readonly type?: string;
    readonly updatedAt: string;
}
export function getStaticSiteOutput(args: GetStaticSiteOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetStaticSiteResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("render:services:getStaticSite", {
        "serviceId": args.serviceId,
    }, opts);
}

export interface GetStaticSiteOutputArgs {
    /**
     * The ID of the service
     */
    serviceId: pulumi.Input<string>;
}
