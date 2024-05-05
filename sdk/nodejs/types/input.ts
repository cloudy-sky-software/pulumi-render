// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";

import * as utilities from "../utilities";

export namespace owners {
}

export namespace registrycredentials {
}

export namespace services {
    export interface AutoScalingCriteriaArgs {
        cpu?: pulumi.Input<inputs.services.AutoScalingCriteriaSpecArgs>;
        memory?: pulumi.Input<inputs.services.AutoScalingCriteriaSpecArgs>;
    }

    export interface AutoScalingCriteriaSpecArgs {
        enabled?: pulumi.Input<boolean>;
        /**
         * Determines when your service will be scaled. If the average resource utilization is significantly above/below the target, we will increase/decrease the number of instances.
         */
        percentage?: pulumi.Input<number>;
    }

    export interface BackgroundWorkerServiceDetailsArgs {
        disk?: pulumi.Input<inputs.services.DiskArgs>;
        env: pulumi.Input<enums.services.BackgroundWorkerServiceDetailsEnv>;
        envSpecificDetails?: pulumi.Input<inputs.services.DockerDetailsArgs | inputs.services.NativeEnvironmentDetailsArgs>;
        numInstances?: pulumi.Input<number>;
        parentServer?: pulumi.Input<inputs.services.BackgroundWorkerServiceDetailsParentServerPropertiesArgs>;
        plan?: pulumi.Input<enums.services.BackgroundWorkerServiceDetailsPlan>;
        pullRequestPreviewsEnabled?: pulumi.Input<enums.services.BackgroundWorkerServiceDetailsPullRequestPreviewsEnabled>;
        region?: pulumi.Input<enums.services.BackgroundWorkerServiceDetailsRegion>;
        url?: pulumi.Input<string>;
    }
    /**
     * backgroundWorkerServiceDetailsArgsProvideDefaults sets the appropriate defaults for BackgroundWorkerServiceDetailsArgs
     */
    export function backgroundWorkerServiceDetailsArgsProvideDefaults(val: BackgroundWorkerServiceDetailsArgs): BackgroundWorkerServiceDetailsArgs {
        return {
            ...val,
            disk: (val.disk ? pulumi.output(val.disk).apply(inputs.services.diskArgsProvideDefaults) : undefined),
            numInstances: (val.numInstances) ?? 1,
            plan: (val.plan) ?? "starter",
            pullRequestPreviewsEnabled: (val.pullRequestPreviewsEnabled) ?? "no",
            region: (val.region) ?? "oregon",
        };
    }

    export interface BackgroundWorkerServiceDetailsParentServerPropertiesArgs {
        id?: pulumi.Input<string>;
        name?: pulumi.Input<string>;
    }

    export interface CronJobServiceDetailsArgs {
        env: pulumi.Input<enums.services.CronJobServiceDetailsEnv>;
        envSpecificDetails?: pulumi.Input<inputs.services.DockerDetailsArgs | inputs.services.NativeEnvironmentDetailsArgs>;
        lastSuccessfulRunAt?: pulumi.Input<string>;
        plan?: pulumi.Input<enums.services.CronJobServiceDetailsPlan>;
        region?: pulumi.Input<enums.services.CronJobServiceDetailsRegion>;
        schedule: pulumi.Input<string>;
    }
    /**
     * cronJobServiceDetailsArgsProvideDefaults sets the appropriate defaults for CronJobServiceDetailsArgs
     */
    export function cronJobServiceDetailsArgsProvideDefaults(val: CronJobServiceDetailsArgs): CronJobServiceDetailsArgs {
        return {
            ...val,
            plan: (val.plan) ?? "starter",
            region: (val.region) ?? "oregon",
        };
    }

    export interface DiskArgs {
        mountPath: pulumi.Input<string>;
        name: pulumi.Input<string>;
        sizeGB?: pulumi.Input<number>;
    }
    /**
     * diskArgsProvideDefaults sets the appropriate defaults for DiskArgs
     */
    export function diskArgsProvideDefaults(val: DiskArgs): DiskArgs {
        return {
            ...val,
            sizeGB: (val.sizeGB) ?? 1,
        };
    }

    export interface DockerDetailsArgs {
        dockerCommand: pulumi.Input<string>;
        dockerContext: pulumi.Input<string>;
        dockerfilePath?: pulumi.Input<string>;
    }

    export interface EnvVarKeyValueArgs {
        generateValue?: pulumi.Input<enums.services.EnvVarKeyValueGenerateValue>;
        key: pulumi.Input<string>;
        value?: pulumi.Input<string>;
    }

    export interface NativeEnvironmentDetailsArgs {
        buildCommand: pulumi.Input<string>;
        startCommand: pulumi.Input<string>;
    }

    export interface OpenPortsArgs {
        port?: pulumi.Input<number>;
        protocol?: pulumi.Input<enums.services.OpenPortsProtocol>;
    }

    export interface PrivateServiceDetailsArgs {
        disk?: pulumi.Input<inputs.services.DiskArgs>;
        env: pulumi.Input<enums.services.PrivateServiceDetailsEnv>;
        envSpecificDetails?: pulumi.Input<inputs.services.DockerDetailsArgs | inputs.services.NativeEnvironmentDetailsArgs>;
        numInstances?: pulumi.Input<number>;
        openPorts?: pulumi.Input<pulumi.Input<inputs.services.OpenPortsArgs>[]>;
        parentServer?: pulumi.Input<inputs.services.PrivateServiceDetailsParentServerPropertiesArgs>;
        plan?: pulumi.Input<enums.services.PrivateServiceDetailsPlan>;
        pullRequestPreviewsEnabled?: pulumi.Input<enums.services.PrivateServiceDetailsPullRequestPreviewsEnabled>;
        region?: pulumi.Input<enums.services.PrivateServiceDetailsRegion>;
        url?: pulumi.Input<string>;
    }
    /**
     * privateServiceDetailsArgsProvideDefaults sets the appropriate defaults for PrivateServiceDetailsArgs
     */
    export function privateServiceDetailsArgsProvideDefaults(val: PrivateServiceDetailsArgs): PrivateServiceDetailsArgs {
        return {
            ...val,
            disk: (val.disk ? pulumi.output(val.disk).apply(inputs.services.diskArgsProvideDefaults) : undefined),
            numInstances: (val.numInstances) ?? 1,
            plan: (val.plan) ?? "starter",
            pullRequestPreviewsEnabled: (val.pullRequestPreviewsEnabled) ?? "no",
            region: (val.region) ?? "oregon",
        };
    }

    export interface PrivateServiceDetailsParentServerPropertiesArgs {
        id?: pulumi.Input<string>;
        name?: pulumi.Input<string>;
    }

    export interface SecretFileArgs {
        contents: pulumi.Input<string>;
        name: pulumi.Input<string>;
    }

    /**
     * A service header object
     */
    export interface ServiceHeaderArgs {
        name: pulumi.Input<string>;
        path: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }

    /**
     * A route object for a static site
     */
    export interface StaticSiteRouteArgs {
        destination: pulumi.Input<string>;
        source: pulumi.Input<string>;
        type: pulumi.Input<enums.services.StaticSiteRouteType>;
    }

    export interface StaticSiteServiceDetailsArgs {
        buildCommand?: pulumi.Input<string>;
        headers?: pulumi.Input<pulumi.Input<inputs.services.ServiceHeaderArgs>[]>;
        parentServer?: pulumi.Input<inputs.services.StaticSiteServiceDetailsParentServerPropertiesArgs>;
        publishPath?: pulumi.Input<string>;
        pullRequestPreviewsEnabled?: pulumi.Input<enums.services.StaticSiteServiceDetailsPullRequestPreviewsEnabled>;
        routes?: pulumi.Input<pulumi.Input<inputs.services.StaticSiteRouteArgs>[]>;
        /**
         * The HTTPS service URL. A subdomain of onrender.com, by default.
         */
        url?: pulumi.Input<string>;
    }
    /**
     * staticSiteServiceDetailsArgsProvideDefaults sets the appropriate defaults for StaticSiteServiceDetailsArgs
     */
    export function staticSiteServiceDetailsArgsProvideDefaults(val: StaticSiteServiceDetailsArgs): StaticSiteServiceDetailsArgs {
        return {
            ...val,
            publishPath: (val.publishPath) ?? "public",
            pullRequestPreviewsEnabled: (val.pullRequestPreviewsEnabled) ?? "no",
        };
    }

    export interface StaticSiteServiceDetailsParentServerPropertiesArgs {
        id?: pulumi.Input<string>;
        name?: pulumi.Input<string>;
    }

    export interface WebServiceServiceDetailsArgs {
        disk?: pulumi.Input<inputs.services.DiskArgs>;
        env: pulumi.Input<enums.services.WebServiceServiceDetailsEnv>;
        envSpecificDetails?: pulumi.Input<inputs.services.DockerDetailsArgs | inputs.services.NativeEnvironmentDetailsArgs>;
        healthCheckPath?: pulumi.Input<string>;
        numInstances?: pulumi.Input<number>;
        openPorts?: pulumi.Input<pulumi.Input<inputs.services.OpenPortsArgs>[]>;
        parentServer?: pulumi.Input<inputs.services.WebServiceServiceDetailsParentServerPropertiesArgs>;
        plan?: pulumi.Input<enums.services.WebServiceServiceDetailsPlan>;
        pullRequestPreviewsEnabled?: pulumi.Input<enums.services.WebServiceServiceDetailsPullRequestPreviewsEnabled>;
        region?: pulumi.Input<enums.services.WebServiceServiceDetailsRegion>;
        url?: pulumi.Input<string>;
    }
    /**
     * webServiceServiceDetailsArgsProvideDefaults sets the appropriate defaults for WebServiceServiceDetailsArgs
     */
    export function webServiceServiceDetailsArgsProvideDefaults(val: WebServiceServiceDetailsArgs): WebServiceServiceDetailsArgs {
        return {
            ...val,
            disk: (val.disk ? pulumi.output(val.disk).apply(inputs.services.diskArgsProvideDefaults) : undefined),
            numInstances: (val.numInstances) ?? 1,
            plan: (val.plan) ?? "starter",
            pullRequestPreviewsEnabled: (val.pullRequestPreviewsEnabled) ?? "no",
            region: (val.region) ?? "oregon",
        };
    }

    export interface WebServiceServiceDetailsParentServerPropertiesArgs {
        id?: pulumi.Input<string>;
        name?: pulumi.Input<string>;
    }
}
