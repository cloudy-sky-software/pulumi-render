// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";

import * as utilities from "../utilities";

export namespace blueprints {
}

export namespace disks {
}

export namespace envgroups {
    export interface EnvVarInputArgs {
        generateValue?: pulumi.Input<boolean>;
        key?: pulumi.Input<string>;
        value?: pulumi.Input<string>;
    }

    export interface SecretFileInputArgs {
        content: pulumi.Input<string>;
        name: pulumi.Input<string>;
    }
}

export namespace environments {
}

export namespace keyvalue {
    export interface CidrBlockAndDescriptionArgs {
        cidrBlock: pulumi.Input<string>;
        /**
         * User-provided description of the CIDR block
         */
        description: pulumi.Input<string>;
    }

}

export namespace logs {
}

export namespace metrics {
}

export namespace notificationsettings {
}

export namespace owners {
}

export namespace postgres {
    export interface CidrBlockAndDescriptionArgs {
        cidrBlock: pulumi.Input<string>;
        /**
         * User-provided description of the CIDR block
         */
        description: pulumi.Input<string>;
    }

    export interface ReadReplicaInputArgs {
        /**
         * The display name of the replica instance.
         */
        name: pulumi.Input<string>;
    }

}

export namespace projects {
    export interface ProjectCreateEnvironmentInputArgs {
        name: pulumi.Input<string>;
        /**
         * Indicates whether network connections across environments are allowed.
         */
        networkIsolationEnabled?: pulumi.Input<boolean>;
        /**
         * Indicates whether an environment is `unprotected` or `protected`. Only admin users can perform destructive actions in `protected` environments.
         */
        protectedStatus?: pulumi.Input<enums.projects.ProjectCreateEnvironmentInputProtectedStatus>;
    }

}

export namespace redis {
    export interface CidrBlockAndDescriptionArgs {
        cidrBlock: pulumi.Input<string>;
        /**
         * User-provided description of the CIDR block
         */
        description: pulumi.Input<string>;
    }

}

export namespace registrycredentials {
}

export namespace services {
    export interface BackgroundWorkerDetailsCreateArgs {
        autoscaling?: pulumi.Input<inputs.services.WebServiceDetailspropertiesautoscalingArgs>;
        disk?: pulumi.Input<inputs.services.ServiceDiskArgs>;
        /**
         * This field has been deprecated, runtime should be used in its place.
         */
        env?: pulumi.Input<enums.services.BackgroundWorkerDetailsCreateEnv>;
        envSpecificDetails?: pulumi.Input<inputs.services.EnvSpecificDetailsCreateArgs>;
        /**
         * The maximum amount of time (in seconds) that Render waits for your application process to exit gracefully after sending it a SIGTERM signal.
         */
        maxShutdownDelaySeconds?: pulumi.Input<number>;
        /**
         * Defaults to 1
         */
        numInstances?: pulumi.Input<number>;
        /**
         * Defaults to "starter"
         */
        plan?: pulumi.Input<enums.services.BackgroundWorkerDetailsCreatePlan>;
        preDeployCommand?: pulumi.Input<string>;
        previews?: pulumi.Input<inputs.services.PreviewsArgs>;
        /**
         * This field has been deprecated. previews.generation should be used in its place.
         */
        pullRequestPreviewsEnabled?: pulumi.Input<enums.services.BackgroundWorkerDetailsCreatePullRequestPreviewsEnabled>;
        /**
         * Defaults to "oregon"
         */
        region?: pulumi.Input<enums.services.BackgroundWorkerDetailsCreateRegion>;
        /**
         * Runtime
         */
        runtime: pulumi.Input<enums.services.BackgroundWorkerDetailsCreateRuntime>;
    }
    /**
     * backgroundWorkerDetailsCreateArgsProvideDefaults sets the appropriate defaults for BackgroundWorkerDetailsCreateArgs
     */
    export function backgroundWorkerDetailsCreateArgsProvideDefaults(val: BackgroundWorkerDetailsCreateArgs): BackgroundWorkerDetailsCreateArgs {
        return {
            ...val,
            autoscaling: (val.autoscaling ? pulumi.output(val.autoscaling).apply(inputs.services.webServiceDetailspropertiesautoscalingArgsProvideDefaults) : undefined),
            maxShutdownDelaySeconds: (val.maxShutdownDelaySeconds) ?? 30,
            numInstances: (val.numInstances) ?? 1,
            plan: (val.plan) ?? "starter",
            previews: (val.previews ? pulumi.output(val.previews).apply(inputs.services.previewsArgsProvideDefaults) : undefined),
            pullRequestPreviewsEnabled: (val.pullRequestPreviewsEnabled) ?? "no",
            region: (val.region) ?? "oregon",
        };
    }

    export interface BuildFilterArgs {
        ignoredPaths: pulumi.Input<pulumi.Input<string>[]>;
        paths: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface CriteriaPropertiesArgs {
        cpu: pulumi.Input<inputs.services.CriteriaPropertiesCpuPropertiesArgs>;
        memory: pulumi.Input<inputs.services.WebServiceDetailspropertiesautoscalingpropertiescriteriapropertiescpuArgs>;
    }
    /**
     * criteriaPropertiesArgsProvideDefaults sets the appropriate defaults for CriteriaPropertiesArgs
     */
    export function criteriaPropertiesArgsProvideDefaults(val: CriteriaPropertiesArgs): CriteriaPropertiesArgs {
        return {
            ...val,
            cpu: pulumi.output(val.cpu).apply(inputs.services.criteriaPropertiesCpuPropertiesArgsProvideDefaults),
            memory: pulumi.output(val.memory).apply(inputs.services.webServiceDetailspropertiesautoscalingpropertiescriteriapropertiescpuArgsProvideDefaults),
        };
    }

    export interface CriteriaPropertiesCpuPropertiesArgs {
        enabled: pulumi.Input<boolean>;
        /**
         * Determines when your service will be scaled. If the average resource utilization is significantly above/below the target, we will increase/decrease the number of instances.
         */
        percentage: pulumi.Input<number>;
    }
    /**
     * criteriaPropertiesCpuPropertiesArgsProvideDefaults sets the appropriate defaults for CriteriaPropertiesCpuPropertiesArgs
     */
    export function criteriaPropertiesCpuPropertiesArgsProvideDefaults(val: CriteriaPropertiesCpuPropertiesArgs): CriteriaPropertiesCpuPropertiesArgs {
        return {
            ...val,
            enabled: (val.enabled) ?? false,
        };
    }

    export interface CronJobDetailsCreateArgs {
        /**
         * This field has been deprecated, runtime should be used in its place.
         */
        env?: pulumi.Input<enums.services.CronJobDetailsCreateEnv>;
        envSpecificDetails?: pulumi.Input<inputs.services.EnvSpecificDetailsArgs>;
        /**
         * Defaults to "starter"
         */
        plan?: pulumi.Input<enums.services.CronJobDetailsCreatePlan>;
        /**
         * Defaults to "oregon"
         */
        region?: pulumi.Input<enums.services.CronJobDetailsCreateRegion>;
        /**
         * Runtime
         */
        runtime: pulumi.Input<enums.services.CronJobDetailsCreateRuntime>;
        schedule: pulumi.Input<string>;
    }
    /**
     * cronJobDetailsCreateArgsProvideDefaults sets the appropriate defaults for CronJobDetailsCreateArgs
     */
    export function cronJobDetailsCreateArgsProvideDefaults(val: CronJobDetailsCreateArgs): CronJobDetailsCreateArgs {
        return {
            ...val,
            plan: (val.plan) ?? "starter",
            region: (val.region) ?? "oregon",
        };
    }

    export interface EnvSpecificDetailsArgs {
        buildCommand?: pulumi.Input<string>;
        dockerCommand?: pulumi.Input<string>;
        dockerContext?: pulumi.Input<string>;
        dockerfilePath?: pulumi.Input<string>;
        preDeployCommand?: pulumi.Input<string>;
        registryCredential?: pulumi.Input<inputs.services.RegistryCredentialArgs>;
        startCommand?: pulumi.Input<string>;
    }

    export interface EnvSpecificDetailsCreateArgs {
        buildCommand?: pulumi.Input<string>;
        dockerCommand?: pulumi.Input<string>;
        dockerContext?: pulumi.Input<string>;
        /**
         * Defaults to "./Dockerfile"
         */
        dockerfilePath?: pulumi.Input<string>;
        registryCredentialId?: pulumi.Input<string>;
        startCommand?: pulumi.Input<string>;
    }

    export interface EnvVarInputArgs {
        generateValue?: pulumi.Input<boolean>;
        key?: pulumi.Input<string>;
        value?: pulumi.Input<string>;
    }

    export interface HeaderInputArgs {
        /**
         * Header name
         */
        name: pulumi.Input<string>;
        /**
         * The request path to add the header to. Wildcards will cause headers to be applied to all matching paths.
         */
        path: pulumi.Input<string>;
        /**
         * Header value
         */
        value: pulumi.Input<string>;
    }

    export interface ImageArgs {
        /**
         * Path to the image used for this server (e.g docker.io/library/nginx:latest).
         */
        imagePath: pulumi.Input<string>;
        /**
         * The ID of the owner for this image. This should match the owner of the service as well as the owner of any specified registry credential.
         */
        ownerId: pulumi.Input<string>;
        /**
         * Optional reference to the registry credential passed to the image repository to retrieve this image.
         */
        registryCredentialId?: pulumi.Input<string>;
    }

    export interface MaintenanceModeArgs {
        enabled: pulumi.Input<boolean>;
        /**
         * The page to be served when [maintenance mode](https://docs.render.com/maintenance-mode) is enabled. When empty, the default maintenance mode page is served.
         */
        uri: pulumi.Input<string>;
    }

    export interface PreviewsArgs {
        /**
         * Defaults to "off"
         */
        generation?: pulumi.Input<enums.services.PreviewsGeneration>;
    }
    /**
     * previewsArgsProvideDefaults sets the appropriate defaults for PreviewsArgs
     */
    export function previewsArgsProvideDefaults(val: PreviewsArgs): PreviewsArgs {
        return {
            ...val,
            generation: (val.generation) ?? "off",
        };
    }

    export interface PrivateServiceDetailsCreateArgs {
        autoscaling?: pulumi.Input<inputs.services.WebServiceDetailspropertiesautoscalingArgs>;
        disk?: pulumi.Input<inputs.services.ServiceDiskArgs>;
        /**
         * This field has been deprecated, runtime should be used in its place.
         */
        env?: pulumi.Input<enums.services.PrivateServiceDetailsCreateEnv>;
        envSpecificDetails?: pulumi.Input<inputs.services.EnvSpecificDetailsCreateArgs>;
        /**
         * The maximum amount of time (in seconds) that Render waits for your application process to exit gracefully after sending it a SIGTERM signal.
         */
        maxShutdownDelaySeconds?: pulumi.Input<number>;
        /**
         * Defaults to 1
         */
        numInstances?: pulumi.Input<number>;
        /**
         * Defaults to "starter"
         */
        plan?: pulumi.Input<enums.services.PrivateServiceDetailsCreatePlan>;
        preDeployCommand?: pulumi.Input<string>;
        previews?: pulumi.Input<inputs.services.PreviewsArgs>;
        /**
         * This field has been deprecated. previews.generation should be used in its place.
         */
        pullRequestPreviewsEnabled?: pulumi.Input<enums.services.PrivateServiceDetailsCreatePullRequestPreviewsEnabled>;
        /**
         * Defaults to "oregon"
         */
        region?: pulumi.Input<enums.services.PrivateServiceDetailsCreateRegion>;
        /**
         * Runtime
         */
        runtime: pulumi.Input<enums.services.PrivateServiceDetailsCreateRuntime>;
    }
    /**
     * privateServiceDetailsCreateArgsProvideDefaults sets the appropriate defaults for PrivateServiceDetailsCreateArgs
     */
    export function privateServiceDetailsCreateArgsProvideDefaults(val: PrivateServiceDetailsCreateArgs): PrivateServiceDetailsCreateArgs {
        return {
            ...val,
            autoscaling: (val.autoscaling ? pulumi.output(val.autoscaling).apply(inputs.services.webServiceDetailspropertiesautoscalingArgsProvideDefaults) : undefined),
            maxShutdownDelaySeconds: (val.maxShutdownDelaySeconds) ?? 30,
            numInstances: (val.numInstances) ?? 1,
            plan: (val.plan) ?? "starter",
            previews: (val.previews ? pulumi.output(val.previews).apply(inputs.services.previewsArgsProvideDefaults) : undefined),
            pullRequestPreviewsEnabled: (val.pullRequestPreviewsEnabled) ?? "no",
            region: (val.region) ?? "oregon",
        };
    }

    export interface RegistryCredentialArgs {
        /**
         * Unique identifier for this credential
         */
        id: pulumi.Input<string>;
        /**
         * Descriptive name for this credential
         */
        name: pulumi.Input<string>;
        /**
         * The registry to use this credential with
         */
        registry: pulumi.Input<enums.services.RegistryCredentialRegistry>;
        /**
         * Last updated time for the credential
         */
        updatedAt: pulumi.Input<string>;
        /**
         * The username associated with the credential
         */
        username: pulumi.Input<string>;
    }

    export interface RouteCreateArgs {
        destination: pulumi.Input<string>;
        /**
         * Redirect and Rewrite Rules are applied in priority order starting at 0. Defaults to last in the priority list.
         */
        priority?: pulumi.Input<number>;
        source: pulumi.Input<string>;
        type: pulumi.Input<enums.services.RouteCreateType>;
    }

    export interface SecretFileInputArgs {
        content: pulumi.Input<string>;
        name: pulumi.Input<string>;
    }

    export interface ServiceDiskArgs {
        mountPath: pulumi.Input<string>;
        name: pulumi.Input<string>;
        /**
         * Defaults to 1
         */
        sizeGB?: pulumi.Input<number>;
    }

    export interface StaticSiteDetailsCreateArgs {
        buildCommand?: pulumi.Input<string>;
        headers?: pulumi.Input<pulumi.Input<inputs.services.HeaderInputArgs>[]>;
        previews?: pulumi.Input<inputs.services.PreviewsArgs>;
        /**
         * Defaults to "public"
         */
        publishPath?: pulumi.Input<string>;
        /**
         * This field has been deprecated. previews.generation should be used in its place.
         */
        pullRequestPreviewsEnabled?: pulumi.Input<enums.services.StaticSiteDetailsCreatePullRequestPreviewsEnabled>;
        routes?: pulumi.Input<pulumi.Input<inputs.services.RouteCreateArgs>[]>;
    }
    /**
     * staticSiteDetailsCreateArgsProvideDefaults sets the appropriate defaults for StaticSiteDetailsCreateArgs
     */
    export function staticSiteDetailsCreateArgsProvideDefaults(val: StaticSiteDetailsCreateArgs): StaticSiteDetailsCreateArgs {
        return {
            ...val,
            previews: (val.previews ? pulumi.output(val.previews).apply(inputs.services.previewsArgsProvideDefaults) : undefined),
            pullRequestPreviewsEnabled: (val.pullRequestPreviewsEnabled) ?? "no",
        };
    }

    export interface WebServiceDetailsCreateArgs {
        autoscaling?: pulumi.Input<inputs.services.WebServiceDetailspropertiesautoscalingArgs>;
        disk?: pulumi.Input<inputs.services.ServiceDiskArgs>;
        /**
         * This field has been deprecated, runtime should be used in its place.
         */
        env?: pulumi.Input<enums.services.WebServiceDetailsCreateEnv>;
        envSpecificDetails?: pulumi.Input<inputs.services.EnvSpecificDetailsCreateArgs>;
        healthCheckPath?: pulumi.Input<string>;
        maintenanceMode?: pulumi.Input<inputs.services.MaintenanceModeArgs>;
        /**
         * The maximum amount of time (in seconds) that Render waits for your application process to exit gracefully after sending it a SIGTERM signal.
         */
        maxShutdownDelaySeconds?: pulumi.Input<number>;
        /**
         * Defaults to 1
         */
        numInstances?: pulumi.Input<number>;
        /**
         * Defaults to "starter"
         */
        plan?: pulumi.Input<enums.services.WebServiceDetailsCreatePlan>;
        preDeployCommand?: pulumi.Input<string>;
        previews?: pulumi.Input<inputs.services.PreviewsArgs>;
        /**
         * This field has been deprecated. previews.generation should be used in its place.
         */
        pullRequestPreviewsEnabled?: pulumi.Input<enums.services.WebServiceDetailsCreatePullRequestPreviewsEnabled>;
        /**
         * Defaults to "oregon"
         */
        region?: pulumi.Input<enums.services.WebServiceDetailsCreateRegion>;
        /**
         * Runtime
         */
        runtime: pulumi.Input<enums.services.WebServiceDetailsCreateRuntime>;
    }
    /**
     * webServiceDetailsCreateArgsProvideDefaults sets the appropriate defaults for WebServiceDetailsCreateArgs
     */
    export function webServiceDetailsCreateArgsProvideDefaults(val: WebServiceDetailsCreateArgs): WebServiceDetailsCreateArgs {
        return {
            ...val,
            autoscaling: (val.autoscaling ? pulumi.output(val.autoscaling).apply(inputs.services.webServiceDetailspropertiesautoscalingArgsProvideDefaults) : undefined),
            maxShutdownDelaySeconds: (val.maxShutdownDelaySeconds) ?? 30,
            plan: (val.plan) ?? "starter",
            previews: (val.previews ? pulumi.output(val.previews).apply(inputs.services.previewsArgsProvideDefaults) : undefined),
            pullRequestPreviewsEnabled: (val.pullRequestPreviewsEnabled) ?? "no",
            region: (val.region) ?? "oregon",
        };
    }

    export interface WebServiceDetailspropertiesautoscalingArgs {
        criteria: pulumi.Input<inputs.services.WebServiceDetailspropertiesautoscalingCriteriaPropertiesArgs>;
        enabled: pulumi.Input<boolean>;
        /**
         * The maximum number of instances for the service
         */
        max: pulumi.Input<number>;
        /**
         * The minimum number of instances for the service
         */
        min: pulumi.Input<number>;
    }
    /**
     * webServiceDetailspropertiesautoscalingArgsProvideDefaults sets the appropriate defaults for WebServiceDetailspropertiesautoscalingArgs
     */
    export function webServiceDetailspropertiesautoscalingArgsProvideDefaults(val: WebServiceDetailspropertiesautoscalingArgs): WebServiceDetailspropertiesautoscalingArgs {
        return {
            ...val,
            criteria: pulumi.output(val.criteria).apply(inputs.services.webServiceDetailspropertiesautoscalingCriteriaPropertiesArgsProvideDefaults),
            enabled: (val.enabled) ?? false,
        };
    }

    export interface WebServiceDetailspropertiesautoscalingCriteriaPropertiesArgs {
        cpu: pulumi.Input<inputs.services.WebServiceDetailspropertiesautoscalingCriteriaPropertiesCpuPropertiesArgs>;
        memory: pulumi.Input<inputs.services.WebServiceDetailspropertiesautoscalingpropertiescriteriapropertiescpuArgs>;
    }
    /**
     * webServiceDetailspropertiesautoscalingCriteriaPropertiesArgsProvideDefaults sets the appropriate defaults for WebServiceDetailspropertiesautoscalingCriteriaPropertiesArgs
     */
    export function webServiceDetailspropertiesautoscalingCriteriaPropertiesArgsProvideDefaults(val: WebServiceDetailspropertiesautoscalingCriteriaPropertiesArgs): WebServiceDetailspropertiesautoscalingCriteriaPropertiesArgs {
        return {
            ...val,
            cpu: pulumi.output(val.cpu).apply(inputs.services.webServiceDetailspropertiesautoscalingCriteriaPropertiesCpuPropertiesArgsProvideDefaults),
            memory: pulumi.output(val.memory).apply(inputs.services.webServiceDetailspropertiesautoscalingpropertiescriteriapropertiescpuArgsProvideDefaults),
        };
    }

    export interface WebServiceDetailspropertiesautoscalingCriteriaPropertiesCpuPropertiesArgs {
        enabled: pulumi.Input<boolean>;
        /**
         * Determines when your service will be scaled. If the average resource utilization is significantly above/below the target, we will increase/decrease the number of instances.
         */
        percentage: pulumi.Input<number>;
    }
    /**
     * webServiceDetailspropertiesautoscalingCriteriaPropertiesCpuPropertiesArgsProvideDefaults sets the appropriate defaults for WebServiceDetailspropertiesautoscalingCriteriaPropertiesCpuPropertiesArgs
     */
    export function webServiceDetailspropertiesautoscalingCriteriaPropertiesCpuPropertiesArgsProvideDefaults(val: WebServiceDetailspropertiesautoscalingCriteriaPropertiesCpuPropertiesArgs): WebServiceDetailspropertiesautoscalingCriteriaPropertiesCpuPropertiesArgs {
        return {
            ...val,
            enabled: (val.enabled) ?? false,
        };
    }

    export interface WebServiceDetailspropertiesautoscalingpropertiescriteriapropertiescpuArgs {
        enabled: pulumi.Input<boolean>;
        /**
         * Determines when your service will be scaled. If the average resource utilization is significantly above/below the target, we will increase/decrease the number of instances.
         */
        percentage: pulumi.Input<number>;
    }
    /**
     * webServiceDetailspropertiesautoscalingpropertiescriteriapropertiescpuArgsProvideDefaults sets the appropriate defaults for WebServiceDetailspropertiesautoscalingpropertiescriteriapropertiescpuArgs
     */
    export function webServiceDetailspropertiesautoscalingpropertiescriteriapropertiescpuArgsProvideDefaults(val: WebServiceDetailspropertiesautoscalingpropertiescriteriapropertiescpuArgs): WebServiceDetailspropertiesautoscalingpropertiescriteriapropertiescpuArgs {
        return {
            ...val,
            enabled: (val.enabled) ?? false,
        };
    }

}

export namespace users {
}

export namespace webhooks {
}
