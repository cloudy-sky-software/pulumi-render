// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'render:projects:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The environments associated with the project
     */
    public /*out*/ readonly environmentIds!: pulumi.Output<string[]>;
    /**
     * The environments to create when creating the project
     */
    public readonly environments!: pulumi.Output<outputs.projects.ProjectCreateEnvironmentInput[]>;
    /**
     * The name of the project
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly owner!: pulumi.Output<outputs.projects.Owner>;
    /**
     * The ID of the owner that the project belongs to
     */
    public readonly ownerId!: pulumi.Output<string>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.environments === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environments'");
            }
            if ((!args || args.ownerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ownerId'");
            }
            resourceInputs["environments"] = args ? args.environments : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ownerId"] = args ? args.ownerId : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["environmentIds"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        } else {
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["environmentIds"] = undefined /*out*/;
            resourceInputs["environments"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Project.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * The environments to create when creating the project
     */
    environments: pulumi.Input<pulumi.Input<inputs.projects.ProjectCreateEnvironmentInputArgs>[]>;
    /**
     * The name of the project
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the owner that the project belongs to
     */
    ownerId: pulumi.Input<string>;
}
