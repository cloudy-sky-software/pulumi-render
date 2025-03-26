// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listWebhookEvents(args: ListWebhookEventsArgs, opts?: pulumi.InvokeOptions): Promise<ListWebhookEventsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:webhooks:listWebhookEvents", {
        "webhookId": args.webhookId,
    }, opts);
}

export interface ListWebhookEventsArgs {
    /**
     * Unique identifier for the webhook
     */
    webhookId: string;
}

export interface ListWebhookEventsResult {
    readonly items: outputs.webhooks.WebhookEventWithCursor[];
}
export function listWebhookEventsOutput(args: ListWebhookEventsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListWebhookEventsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("render:webhooks:listWebhookEvents", {
        "webhookId": args.webhookId,
    }, opts);
}

export interface ListWebhookEventsOutputArgs {
    /**
     * Unique identifier for the webhook
     */
    webhookId: pulumi.Input<string>;
}
