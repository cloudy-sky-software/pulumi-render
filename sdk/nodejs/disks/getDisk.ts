// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getDisk(args: GetDiskArgs, opts?: pulumi.InvokeOptions): Promise<outputs.disks.DiskWithCursorpropertiesdisk> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:disks:getDisk", {
        "diskId": args.diskId,
    }, opts);
}

export interface GetDiskArgs {
    /**
     * The ID of the disk
     */
    diskId: string;
}
export function getDiskOutput(args: GetDiskOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.disks.DiskWithCursorpropertiesdisk> {
    return pulumi.output(args).apply((a: any) => getDisk(a, opts))
}

export interface GetDiskOutputArgs {
    /**
     * The ID of the disk
     */
    diskId: pulumi.Input<string>;
}
