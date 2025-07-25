// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listWebhooks(args?: ListWebhooksArgs, opts?: pulumi.InvokeOptions): Promise<ListWebhooksResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:webhooks:listWebhooks", {
    }, opts);
}

export interface ListWebhooksArgs {
}

export interface ListWebhooksResult {
    readonly items: outputs.webhooks.WebhookWithCursor[];
}
export function listWebhooksOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListWebhooksResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("render:webhooks:listWebhooks", {
    }, opts);
}

