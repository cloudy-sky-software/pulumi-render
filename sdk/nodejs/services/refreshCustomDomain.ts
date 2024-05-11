// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class RefreshCustomDomain extends pulumi.CustomResource {
    /**
     * Get an existing RefreshCustomDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RefreshCustomDomain {
        return new RefreshCustomDomain(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'render:services:RefreshCustomDomain';

    /**
     * Returns true if the given object is an instance of RefreshCustomDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RefreshCustomDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RefreshCustomDomain.__pulumiType;
    }


    /**
     * Create a RefreshCustomDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RefreshCustomDomainArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["customDomainIdOrName"] = args ? args.customDomainIdOrName : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
        } else {
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RefreshCustomDomain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RefreshCustomDomain resource.
 */
export interface RefreshCustomDomainArgs {
    /**
     * The ID or name of the custom domain
     */
    customDomainIdOrName?: pulumi.Input<string>;
    /**
     * The ID of the service
     */
    serviceId?: pulumi.Input<string>;
}
