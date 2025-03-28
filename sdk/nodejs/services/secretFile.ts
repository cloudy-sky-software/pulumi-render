// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class SecretFile extends pulumi.CustomResource {
    /**
     * Get an existing SecretFile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SecretFile {
        return new SecretFile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'render:services:SecretFile';

    /**
     * Returns true if the given object is an instance of SecretFile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretFile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretFile.__pulumiType;
    }

    public readonly content!: pulumi.Output<string>;
    public /*out*/ readonly name!: pulumi.Output<string>;

    /**
     * Create a SecretFile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecretFileArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["secretFileName"] = args ? args.secretFileName : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["content"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretFile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SecretFile resource.
 */
export interface SecretFileArgs {
    content?: pulumi.Input<string>;
    /**
     * The file name of the secret file
     */
    secretFileName?: pulumi.Input<string>;
    /**
     * The ID of the service
     */
    serviceId?: pulumi.Input<string>;
}
