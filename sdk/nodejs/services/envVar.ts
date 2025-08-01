// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class EnvVar extends pulumi.CustomResource {
    /**
     * Get an existing EnvVar resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EnvVar {
        return new EnvVar(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'render:services:EnvVar';

    /**
     * Returns true if the given object is an instance of EnvVar.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EnvVar {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EnvVar.__pulumiType;
    }

    public readonly generateValue!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly key!: pulumi.Output<string>;
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a EnvVar resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EnvVarArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["envVarKey"] = args ? args.envVarKey : undefined;
            resourceInputs["generateValue"] = args ? args.generateValue : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
            resourceInputs["key"] = undefined /*out*/;
        } else {
            resourceInputs["generateValue"] = undefined /*out*/;
            resourceInputs["key"] = undefined /*out*/;
            resourceInputs["value"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EnvVar.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EnvVar resource.
 */
export interface EnvVarArgs {
    /**
     * The name of the environment variable
     */
    envVarKey?: pulumi.Input<string>;
    generateValue?: pulumi.Input<boolean>;
    /**
     * The ID of the service
     */
    serviceId?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}
