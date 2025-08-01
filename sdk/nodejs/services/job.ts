// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'render:services:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public /*out*/ readonly finishedAt!: pulumi.Output<string | undefined>;
    public readonly planId!: pulumi.Output<string>;
    public readonly serviceId!: pulumi.Output<string>;
    public readonly startCommand!: pulumi.Output<string>;
    public /*out*/ readonly startedAt!: pulumi.Output<string | undefined>;
    public /*out*/ readonly status!: pulumi.Output<enums.services.JobStatus | undefined>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.startCommand === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startCommand'");
            }
            resourceInputs["planId"] = args ? args.planId : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["startCommand"] = args ? args.startCommand : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["finishedAt"] = undefined /*out*/;
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
        super(Job.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    planId?: pulumi.Input<string>;
    /**
     * The ID of the service
     */
    serviceId?: pulumi.Input<string>;
    startCommand: pulumi.Input<string>;
}
