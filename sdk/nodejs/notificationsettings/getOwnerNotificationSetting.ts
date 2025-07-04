// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getOwnerNotificationSetting(args: GetOwnerNotificationSettingArgs, opts?: pulumi.InvokeOptions): Promise<GetOwnerNotificationSettingResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:notification-settings:getOwnerNotificationSetting", {
        "ownerId": args.ownerId,
    }, opts);
}

export interface GetOwnerNotificationSettingArgs {
    /**
     * The ID of the owner (team or personal user) whose resources should be returned
     */
    ownerId: string;
}

export interface GetOwnerNotificationSettingResult {
    readonly emailEnabled: boolean;
    readonly notificationsToSend: enums.notificationsettings.GetOwnerNotificationSettingPropertiesNotificationsToSend;
    readonly ownerId: string;
    readonly previewNotificationsEnabled: boolean;
    readonly slackEnabled: boolean;
}
export function getOwnerNotificationSettingOutput(args: GetOwnerNotificationSettingOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetOwnerNotificationSettingResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("render:notification-settings:getOwnerNotificationSetting", {
        "ownerId": args.ownerId,
    }, opts);
}

export interface GetOwnerNotificationSettingOutputArgs {
    /**
     * The ID of the owner (team or personal user) whose resources should be returned
     */
    ownerId: pulumi.Input<string>;
}
