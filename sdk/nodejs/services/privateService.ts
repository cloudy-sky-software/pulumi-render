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

    public /*out*/ readonly autoDeploy!: pulumi.Output<enums.services.PrivateServiceAutoDeploy>;
    public /*out*/ readonly branch!: pulumi.Output<string | undefined>;
    public /*out*/ readonly buildFilter!: pulumi.Output<outputs.services.BuildFilter | undefined>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public /*out*/ readonly imagePath!: pulumi.Output<string | undefined>;
    public /*out*/ readonly name!: pulumi.Output<string>;
    public /*out*/ readonly notifyOnFail!: pulumi.Output<enums.services.PrivateServiceNotifyOnFail>;
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    public /*out*/ readonly repo!: pulumi.Output<string | undefined>;
    public /*out*/ readonly rootDir!: pulumi.Output<string>;
    public /*out*/ readonly serviceDetails!: pulumi.Output<outputs.services.StaticSiteDetails | outputs.services.WebServiceDetails | outputs.services.PrivateServiceDetails | outputs.services.BackgroundWorkerDetails | outputs.services.CronJobDetails>;
    public /*out*/ readonly slug!: pulumi.Output<string>;
    public /*out*/ readonly suspended!: pulumi.Output<enums.services.PrivateServiceSuspended>;
    public /*out*/ readonly suspenders!: pulumi.Output<enums.services.PrivateServiceSuspendersItem[]>;
    public readonly type!: pulumi.Output<string>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a PrivateService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PrivateServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["type"] = (args ? args.type : undefined) ?? "private_service";
            resourceInputs["autoDeploy"] = undefined /*out*/;
            resourceInputs["branch"] = undefined /*out*/;
            resourceInputs["buildFilter"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["imagePath"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["notifyOnFail"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["repo"] = undefined /*out*/;
            resourceInputs["rootDir"] = undefined /*out*/;
            resourceInputs["serviceDetails"] = undefined /*out*/;
            resourceInputs["slug"] = undefined /*out*/;
            resourceInputs["suspended"] = undefined /*out*/;
            resourceInputs["suspenders"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        } else {
            resourceInputs["autoDeploy"] = undefined /*out*/;
            resourceInputs["branch"] = undefined /*out*/;
            resourceInputs["buildFilter"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["imagePath"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["notifyOnFail"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["repo"] = undefined /*out*/;
            resourceInputs["rootDir"] = undefined /*out*/;
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
    type?: pulumi.Input<string>;
}
