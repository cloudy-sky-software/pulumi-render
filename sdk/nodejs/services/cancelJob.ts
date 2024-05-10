// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class CancelJob extends pulumi.CustomResource {
    /**
     * Get an existing CancelJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CancelJob {
        return new CancelJob(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'render:services:CancelJob';

    /**
     * Returns true if the given object is an instance of CancelJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CancelJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CancelJob.__pulumiType;
    }

    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public /*out*/ readonly finishedAt!: pulumi.Output<string | undefined>;
    public /*out*/ readonly planId!: pulumi.Output<string>;
    public /*out*/ readonly serviceId!: pulumi.Output<string>;
    public /*out*/ readonly startCommand!: pulumi.Output<string>;
    public /*out*/ readonly startedAt!: pulumi.Output<string | undefined>;
    public /*out*/ readonly status!: pulumi.Output<string | undefined>;

    /**
     * Create a CancelJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CancelJobArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["finishedAt"] = undefined /*out*/;
            resourceInputs["planId"] = undefined /*out*/;
            resourceInputs["serviceId"] = undefined /*out*/;
            resourceInputs["startCommand"] = undefined /*out*/;
            resourceInputs["startedAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["finishedAt"] = undefined /*out*/;
            resourceInputs["planId"] = undefined /*out*/;
            resourceInputs["serviceId"] = undefined /*out*/;
            resourceInputs["startCommand"] = undefined /*out*/;
            resourceInputs["startedAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CancelJob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CancelJob resource.
 */
export interface CancelJobArgs {
}
