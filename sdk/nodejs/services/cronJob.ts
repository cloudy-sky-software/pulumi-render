// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * A cron job
 */
export class CronJob extends pulumi.CustomResource {
    /**
     * Get an existing CronJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CronJob {
        return new CronJob(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'render:services:CronJob';

    /**
     * Returns true if the given object is an instance of CronJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CronJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CronJob.__pulumiType;
    }

    /**
     * Whether to auto deploy the service or not upon git push.
     */
    public readonly autoDeploy!: pulumi.Output<enums.services.ServiceAutoDeploy | undefined>;
    /**
     * If left empty, this will fall back to the default branch of the repository.
     */
    public readonly branch!: pulumi.Output<string | undefined>;
    public readonly createdAt!: pulumi.Output<string | undefined>;
    public readonly envVars!: pulumi.Output<outputs.services.EnvVarKeyValue[] | undefined>;
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The notification setting for this service upon deployment failure.
     */
    public readonly notifyOnFail!: pulumi.Output<enums.services.ServiceNotifyOnFail | undefined>;
    /**
     * The id of the owner (user/team).
     */
    public readonly ownerId!: pulumi.Output<string | undefined>;
    /**
     * Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
     */
    public readonly repo!: pulumi.Output<string | undefined>;
    public readonly secretFiles!: pulumi.Output<outputs.services.SecretFile[] | undefined>;
    public readonly serviceDetails!: pulumi.Output<outputs.services.CronJobServiceDetails | undefined>;
    public readonly slug!: pulumi.Output<string | undefined>;
    public readonly suspended!: pulumi.Output<enums.services.ServiceSuspended | undefined>;
    public readonly suspenders!: pulumi.Output<string[] | undefined>;
    public readonly type!: pulumi.Output<string | undefined>;
    public readonly updatedAt!: pulumi.Output<string | undefined>;

    /**
     * Create a CronJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CronJobArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.ownerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ownerId'");
            }
            if ((!args || args.repo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repo'");
            }
            resourceInputs["autoDeploy"] = (args ? args.autoDeploy : undefined) ?? "no";
            resourceInputs["branch"] = args ? args.branch : undefined;
            resourceInputs["createdAt"] = args ? args.createdAt : undefined;
            resourceInputs["envVars"] = args ? args.envVars : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notifyOnFail"] = args ? args.notifyOnFail : undefined;
            resourceInputs["ownerId"] = args ? args.ownerId : undefined;
            resourceInputs["repo"] = args ? args.repo : undefined;
            resourceInputs["secretFiles"] = args ? args.secretFiles : undefined;
            resourceInputs["serviceDetails"] = args ? (args.serviceDetails ? pulumi.output(args.serviceDetails).apply(inputs.services.cronJobServiceDetailsArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
            resourceInputs["suspended"] = args ? args.suspended : undefined;
            resourceInputs["suspenders"] = args ? args.suspenders : undefined;
            resourceInputs["type"] = (args ? args.type : undefined) ?? "cron_job";
            resourceInputs["updatedAt"] = args ? args.updatedAt : undefined;
        } else {
            resourceInputs["autoDeploy"] = undefined /*out*/;
            resourceInputs["branch"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["envVars"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["notifyOnFail"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["repo"] = undefined /*out*/;
            resourceInputs["secretFiles"] = undefined /*out*/;
            resourceInputs["serviceDetails"] = undefined /*out*/;
            resourceInputs["slug"] = undefined /*out*/;
            resourceInputs["suspended"] = undefined /*out*/;
            resourceInputs["suspenders"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CronJob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CronJob resource.
 */
export interface CronJobArgs {
    /**
     * Whether to auto deploy the service or not upon git push.
     */
    autoDeploy?: pulumi.Input<enums.services.ServiceAutoDeploy>;
    /**
     * If left empty, this will fall back to the default branch of the repository.
     */
    branch?: pulumi.Input<string>;
    createdAt?: pulumi.Input<string>;
    envVars?: pulumi.Input<pulumi.Input<inputs.services.EnvVarKeyValueArgs>[]>;
    name: pulumi.Input<string>;
    /**
     * The notification setting for this service upon deployment failure.
     */
    notifyOnFail?: pulumi.Input<enums.services.ServiceNotifyOnFail>;
    /**
     * The id of the owner (user/team).
     */
    ownerId: pulumi.Input<string>;
    /**
     * Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
     */
    repo: pulumi.Input<string>;
    secretFiles?: pulumi.Input<pulumi.Input<inputs.services.SecretFileArgs>[]>;
    serviceDetails?: pulumi.Input<inputs.services.CronJobServiceDetailsArgs>;
    slug?: pulumi.Input<string>;
    suspended?: pulumi.Input<enums.services.ServiceSuspended>;
    suspenders?: pulumi.Input<pulumi.Input<string>[]>;
    type?: pulumi.Input<string>;
    updatedAt?: pulumi.Input<string>;
}
