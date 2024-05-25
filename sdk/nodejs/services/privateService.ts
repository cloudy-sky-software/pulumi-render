// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class PrivateService extends pulumi.CustomResource {
    /**
     * Get an existing PrivateService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PrivateService {
        return new PrivateService(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'render:services:PrivateService';

    /**
     * Returns true if the given object is an instance of PrivateService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateService.__pulumiType;
    }

    public readonly autoDeploy!: pulumi.Output<enums.services.ServiceAutoDeploy | undefined>;
    public readonly branch!: pulumi.Output<string | undefined>;
    public readonly buildFilter!: pulumi.Output<outputs.services.BuildFilter | undefined>;
    public /*out*/ readonly createdAt!: pulumi.Output<string | undefined>;
    public readonly envVars!: pulumi.Output<(outputs.services.EnvVarKeyValue | outputs.services.EnvVarKeyGenerateValue)[] | undefined>;
    public readonly image!: pulumi.Output<outputs.services.Image | undefined>;
    public /*out*/ readonly imagePath!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string | undefined>;
    public /*out*/ readonly notifyOnFail!: pulumi.Output<enums.services.ServiceNotifyOnFail | undefined>;
    public readonly ownerId!: pulumi.Output<string | undefined>;
    public readonly repo!: pulumi.Output<string | undefined>;
    public readonly rootDir!: pulumi.Output<string | undefined>;
    public readonly secretFiles!: pulumi.Output<outputs.services.SecretFile[] | undefined>;
    public readonly serviceDetails!: pulumi.Output<outputs.services.PrivateServiceDetailsOutput | undefined>;
    public /*out*/ readonly slug!: pulumi.Output<string | undefined>;
    public /*out*/ readonly suspended!: pulumi.Output<enums.services.ServiceSuspended | undefined>;
    public /*out*/ readonly suspenders!: pulumi.Output<enums.services.ServiceSuspendersItem[] | undefined>;
    public readonly type!: pulumi.Output<string | undefined>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string | undefined>;

    /**
     * Create a PrivateService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.ownerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ownerId'");
            }
            resourceInputs["autoDeploy"] = (args ? args.autoDeploy : undefined) ?? "yes";
            resourceInputs["branch"] = args ? args.branch : undefined;
            resourceInputs["buildFilter"] = args ? args.buildFilter : undefined;
            resourceInputs["envVars"] = args ? args.envVars : undefined;
            resourceInputs["image"] = args ? args.image : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ownerId"] = args ? args.ownerId : undefined;
            resourceInputs["repo"] = args ? args.repo : undefined;
            resourceInputs["rootDir"] = args ? args.rootDir : undefined;
            resourceInputs["secretFiles"] = args ? args.secretFiles : undefined;
            resourceInputs["serviceDetails"] = args ? (args.serviceDetails ? pulumi.output(args.serviceDetails).apply(inputs.services.privateServiceDetailsCreateArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["type"] = (args ? args.type : undefined) ?? "private_service";
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["imagePath"] = undefined /*out*/;
            resourceInputs["notifyOnFail"] = undefined /*out*/;
            resourceInputs["slug"] = undefined /*out*/;
            resourceInputs["suspended"] = undefined /*out*/;
            resourceInputs["suspenders"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        } else {
            resourceInputs["autoDeploy"] = undefined /*out*/;
            resourceInputs["branch"] = undefined /*out*/;
            resourceInputs["buildFilter"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["envVars"] = undefined /*out*/;
            resourceInputs["image"] = undefined /*out*/;
            resourceInputs["imagePath"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["notifyOnFail"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["repo"] = undefined /*out*/;
            resourceInputs["rootDir"] = undefined /*out*/;
            resourceInputs["secretFiles"] = undefined /*out*/;
            resourceInputs["serviceDetails"] = undefined /*out*/;
            resourceInputs["slug"] = undefined /*out*/;
            resourceInputs["suspended"] = undefined /*out*/;
            resourceInputs["suspenders"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrivateService.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PrivateService resource.
 */
export interface PrivateServiceArgs {
    /**
     * Defaults to "yes"
     */
    autoDeploy?: pulumi.Input<enums.services.ServiceCreateAutoDeploy>;
    /**
     * If left empty, this will fall back to the default branch of the repository
     */
    branch?: pulumi.Input<string>;
    buildFilter?: pulumi.Input<inputs.services.BuildFilterArgs>;
    envVars?: pulumi.Input<pulumi.Input<inputs.services.EnvVarKeyValueArgs | inputs.services.EnvVarKeyGenerateValueArgs>[]>;
    image?: pulumi.Input<inputs.services.ImageArgs>;
    name: pulumi.Input<string>;
    ownerId: pulumi.Input<string>;
    /**
     * Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
     */
    repo?: pulumi.Input<string>;
    rootDir?: pulumi.Input<string>;
    secretFiles?: pulumi.Input<pulumi.Input<inputs.services.SecretFileArgs>[]>;
    serviceDetails?: pulumi.Input<inputs.services.PrivateServiceDetailsCreateArgs>;
    type?: pulumi.Input<string>;
}
