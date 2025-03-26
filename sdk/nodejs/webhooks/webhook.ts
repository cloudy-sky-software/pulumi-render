// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class Webhook extends pulumi.CustomResource {
    /**
     * Get an existing Webhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Webhook {
        return new Webhook(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'render:webhooks:Webhook';

    /**
     * Returns true if the given object is an instance of Webhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Webhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Webhook.__pulumiType;
    }

    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * The event types that will trigger the webhook. An empty list means all event types will trigger the webhook.
     */
    public readonly eventFilter!: pulumi.Output<enums.webhooks.EventFilterItem[]>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the owner (team or personal user) whose resources should be returned
     */
    public readonly ownerId!: pulumi.Output<string>;
    public /*out*/ readonly secret!: pulumi.Output<string>;
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a Webhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebhookArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.eventFilter === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventFilter'");
            }
            if ((!args || args.ownerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ownerId'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["eventFilter"] = args ? args.eventFilter : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ownerId"] = args ? args.ownerId : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["secret"] = undefined /*out*/;
        } else {
            resourceInputs["enabled"] = undefined /*out*/;
            resourceInputs["eventFilter"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["secret"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Webhook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Webhook resource.
 */
export interface WebhookArgs {
    enabled: pulumi.Input<boolean>;
    /**
     * The event types that will trigger the webhook. An empty list means all event types will trigger the webhook.
     */
    eventFilter: pulumi.Input<pulumi.Input<enums.webhooks.EventFilterItem>[]>;
    name?: pulumi.Input<string>;
    /**
     * The ID of the owner (team or personal user) whose resources should be returned
     */
    ownerId: pulumi.Input<string>;
    url: pulumi.Input<string>;
}
