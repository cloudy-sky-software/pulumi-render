// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class PreviewService extends pulumi.CustomResource {
    /**
     * Get an existing PreviewService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PreviewService {
        return new PreviewService(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'render:services:PreviewService';

    /**
     * Returns true if the given object is an instance of PreviewService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PreviewService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PreviewService.__pulumiType;
    }

    public /*out*/ readonly deployId!: pulumi.Output<string | undefined>;
    /**
     * Must be either a full URL or the relative path to an image. If a relative path, Render uses the base service's image URL as its root. For example, if the base service's image URL is `docker.io/library/nginx:latest`, then valid values are: `docker.io/library/nginx:<any tag or SHA>`, `library/nginx:<any tag or SHA>`, or `nginx:<any tag or SHA>`. Note that the path must match (only the tag or SHA can vary).
     */
    public readonly imagePath!: pulumi.Output<string>;
    /**
     * A name for the service preview instance. If not specified, Render generates the name using the base service's name and the specified tag or SHA.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
     */
    public readonly plan!: pulumi.Output<enums.services.Plan | undefined>;
    public /*out*/ readonly service!: pulumi.Output<outputs.services.Service | undefined>;

    /**
     * Create a PreviewService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PreviewServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.imagePath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imagePath'");
            }
            resourceInputs["imagePath"] = args ? args.imagePath : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["plan"] = args ? args.plan : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["deployId"] = undefined /*out*/;
            resourceInputs["service"] = undefined /*out*/;
        } else {
            resourceInputs["deployId"] = undefined /*out*/;
            resourceInputs["imagePath"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["plan"] = undefined /*out*/;
            resourceInputs["service"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PreviewService.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PreviewService resource.
 */
export interface PreviewServiceArgs {
    /**
     * Must be either a full URL or the relative path to an image. If a relative path, Render uses the base service's image URL as its root. For example, if the base service's image URL is `docker.io/library/nginx:latest`, then valid values are: `docker.io/library/nginx:<any tag or SHA>`, `library/nginx:<any tag or SHA>`, or `nginx:<any tag or SHA>`. Note that the path must match (only the tag or SHA can vary).
     */
    imagePath: pulumi.Input<string>;
    /**
     * A name for the service preview instance. If not specified, Render generates the name using the base service's name and the specified tag or SHA.
     */
    name?: pulumi.Input<string>;
    /**
     * The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
     */
    plan?: pulumi.Input<enums.services.Plan>;
    /**
     * The ID of the service
     */
    serviceId?: pulumi.Input<string>;
}