// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class RestoreSnapshot extends pulumi.CustomResource {
    /**
     * Get an existing RestoreSnapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RestoreSnapshot {
        return new RestoreSnapshot(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'render:disks:RestoreSnapshot';

    /**
     * Returns true if the given object is an instance of RestoreSnapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RestoreSnapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RestoreSnapshot.__pulumiType;
    }

    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * When a service with a disk is scaled, the instanceId is used to identify the instance that the disk is attached to. Each instance's disks get their own snapshots, and can be restored separately.
     */
    public readonly instanceId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly mountPath!: pulumi.Output<string>;
    public /*out*/ readonly name!: pulumi.Output<string>;
    public /*out*/ readonly serviceId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly sizeGB!: pulumi.Output<number>;
    public readonly snapshotKey!: pulumi.Output<string>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a RestoreSnapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RestoreSnapshotArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.snapshotKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'snapshotKey'");
            }
            resourceInputs["diskId"] = args ? args.diskId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["snapshotKey"] = args ? args.snapshotKey : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["mountPath"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serviceId"] = undefined /*out*/;
            resourceInputs["sizeGB"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        } else {
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["instanceId"] = undefined /*out*/;
            resourceInputs["mountPath"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serviceId"] = undefined /*out*/;
            resourceInputs["sizeGB"] = undefined /*out*/;
            resourceInputs["snapshotKey"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RestoreSnapshot.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RestoreSnapshot resource.
 */
export interface RestoreSnapshotArgs {
    /**
     * The ID of the disk
     */
    diskId?: pulumi.Input<string>;
    /**
     * When a service with a disk is scaled, the instanceId is used to identify the instance that the disk is attached to. Each instance's disks get their own snapshots, and can be restored separately.
     */
    instanceId?: pulumi.Input<string>;
    snapshotKey: pulumi.Input<string>;
}
