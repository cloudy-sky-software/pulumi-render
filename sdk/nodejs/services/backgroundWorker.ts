// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class BackgroundWorker extends pulumi.CustomResource {
    /**
     * Get an existing BackgroundWorker resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BackgroundWorker {
        return new BackgroundWorker(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'render:services:BackgroundWorker';

    /**
     * Returns true if the given object is an instance of BackgroundWorker.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackgroundWorker {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackgroundWorker.__pulumiType;
    }

    public readonly autoDeploy!: pulumi.Output<enums.services.ServiceAutoDeploy | undefined>;
    public readonly branch!: pulumi.Output<string | undefined>;
    public readonly buildFilter!: pulumi.Output<outputs.services.BuildFilter | undefined>;
    public /*out*/ readonly createdAt!: pulumi.Output<string | undefined>;
    /**
     * The URL to view the service in the Render Dashboard
     */
    public /*out*/ readonly dashboardUrl!: pulumi.Output<string | undefined>;
    public readonly envVars!: pulumi.Output<outputs.services.EnvVarInput[] | undefined>;
    public /*out*/ readonly environmentId!: pulumi.Output<string | undefined>;
    public readonly image!: pulumi.Output<outputs.services.Image | undefined>;
    public /*out*/ readonly imagePath!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string | undefined>;
    public /*out*/ readonly notifyOnFail!: pulumi.Output<enums.services.ServiceNotifyOnFail | undefined>;
    public readonly ownerId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly registryCredential!: pulumi.Output<outputs.services.RegistryCredentialSummary | undefined>;
    public readonly repo!: pulumi.Output<string | undefined>;
    public readonly rootDir!: pulumi.Output<string | undefined>;
    public readonly secretFiles!: pulumi.Output<outputs.services.SecretFileInput[] | undefined>;
    public readonly serviceDetails!: pulumi.Output<outputs.services.BackgroundWorkerDetailsOutput | undefined>;
    public /*out*/ readonly slug!: pulumi.Output<string | undefined>;
    public /*out*/ readonly suspended!: pulumi.Output<enums.services.ServiceSuspended | undefined>;
    public /*out*/ readonly suspenders!: pulumi.Output<enums.services.ServiceSuspendersItem[] | undefined>;
    public readonly type!: pulumi.Output<string | undefined>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string | undefined>;

    /**
     * Create a BackgroundWorker resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackgroundWorkerArgs, opts?: pulumi.CustomResourceOptions) {
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
            resourceInputs["serviceDetails"] = args ? (args.serviceDetails ? pulumi.output(args.serviceDetails).apply(inputs.services.backgroundWorkerDetailsCreateArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["type"] = (args ? args.type : undefined) ?? "background_worker";
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["dashboardUrl"] = undefined /*out*/;
            resourceInputs["environmentId"] = undefined /*out*/;
            resourceInputs["imagePath"] = undefined /*out*/;
            resourceInputs["notifyOnFail"] = undefined /*out*/;
            resourceInputs["registryCredential"] = undefined /*out*/;
            resourceInputs["slug"] = undefined /*out*/;
            resourceInputs["suspended"] = undefined /*out*/;
            resourceInputs["suspenders"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        } else {
            resourceInputs["autoDeploy"] = undefined /*out*/;
            resourceInputs["branch"] = undefined /*out*/;
            resourceInputs["buildFilter"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["dashboardUrl"] = undefined /*out*/;
            resourceInputs["envVars"] = undefined /*out*/;
            resourceInputs["environmentId"] = undefined /*out*/;
            resourceInputs["image"] = undefined /*out*/;
            resourceInputs["imagePath"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["notifyOnFail"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["registryCredential"] = undefined /*out*/;
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
        super(BackgroundWorker.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a BackgroundWorker resource.
 */
export interface BackgroundWorkerArgs {
    autoDeploy?: pulumi.Input<enums.services.ServiceCreateAutoDeploy>;
    /**
     * If left empty, this will fall back to the default branch of the repository
     */
    branch?: pulumi.Input<string>;
    buildFilter?: pulumi.Input<inputs.services.BuildFilterArgs>;
    envVars?: pulumi.Input<pulumi.Input<inputs.services.EnvVarInputArgs>[]>;
    image?: pulumi.Input<inputs.services.ImageArgs>;
    name: pulumi.Input<string>;
    ownerId: pulumi.Input<string>;
    /**
     * Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
     */
    repo?: pulumi.Input<string>;
    rootDir?: pulumi.Input<string>;
    secretFiles?: pulumi.Input<pulumi.Input<inputs.services.SecretFileInputArgs>[]>;
    serviceDetails?: pulumi.Input<inputs.services.BackgroundWorkerDetailsCreateArgs>;
    type?: pulumi.Input<string>;
}
