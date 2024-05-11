// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class RestartServer extends pulumi.CustomResource {
    /**
     * Get an existing RestartServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RestartServer {
        return new RestartServer(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'render:services:RestartServer';

    /**
     * Returns true if the given object is an instance of RestartServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RestartServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RestartServer.__pulumiType;
    }


    /**
     * Create a RestartServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RestartServerArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
        } else {
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RestartServer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RestartServer resource.
 */
export interface RestartServerArgs {
    /**
     * The ID of the service
     */
    serviceId?: pulumi.Input<string>;
}
