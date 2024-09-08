// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class Header extends pulumi.CustomResource {
    /**
     * Get an existing Header resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Header {
        return new Header(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'render:services:Header';

    /**
     * Returns true if the given object is an instance of Header.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Header {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Header.__pulumiType;
    }

    public /*out*/ readonly headers!: pulumi.Output<outputs.services.Header | undefined>;
    /**
     * Header name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The request path to add the header to. Wildcards will cause headers to be applied to all matching paths.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * Header value
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a Header resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HeaderArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
            resourceInputs["headers"] = undefined /*out*/;
        } else {
            resourceInputs["headers"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["path"] = undefined /*out*/;
            resourceInputs["value"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Header.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Header resource.
 */
export interface HeaderArgs {
    /**
     * Header name
     */
    name?: pulumi.Input<string>;
    /**
     * The request path to add the header to. Wildcards will cause headers to be applied to all matching paths.
     */
    path: pulumi.Input<string>;
    /**
     * The ID of the service
     */
    serviceId?: pulumi.Input<string>;
    /**
     * Header value
     */
    value: pulumi.Input<string>;
}