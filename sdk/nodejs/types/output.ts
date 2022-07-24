// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";

import * as utilities from "../utilities";

export namespace services {
    export interface DeployCommit {
        createdAt?: string;
        id?: string;
        message?: string;
    }

    export interface ServerProperties {
        id?: string;
        name?: string;
    }

    export interface ServiceDisk {
        mountPath: string;
        name: string;
        sizeGB?: number;
    }
    /**
     * serviceDiskProvideDefaults sets the appropriate defaults for ServiceDisk
     */
    export function serviceDiskProvideDefaults(val: ServiceDisk): ServiceDisk {
        return {
            ...val,
            sizeGB: (val.sizeGB) ?? 1,
        };
    }

    export interface ServiceDockerDetails {
        dockerCommand?: string;
        dockerContext?: string;
        dockerfilePath?: string;
    }
    /**
     * serviceDockerDetailsProvideDefaults sets the appropriate defaults for ServiceDockerDetails
     */
    export function serviceDockerDetailsProvideDefaults(val: ServiceDockerDetails): ServiceDockerDetails {
        return {
            ...val,
            dockerfilePath: (val.dockerfilePath) ?? "./Dockerfile",
        };
    }

    export interface ServiceNativeEnvironmentDetails {
        buildCommand: string;
        startCommand: string;
    }

    /**
     * A service header object
     */
    export interface ServiceServiceHeader {
        name: string;
        path: string;
        value: string;
    }

    /**
     * A static website service
     */
    export interface ServiceStaticSite {
        buildCommand?: string;
        headers?: outputs.services.ServiceServiceHeader[];
        parentServer?: outputs.services.ServiceStaticSiteParentServerProperties;
        publishPath?: string;
        pullRequestPreviewsEnabled?: enums.services.ServiceStaticSitePullRequestPreviewsEnabled;
        routes?: outputs.services.ServiceStaticSiteRoute[];
        type?: string;
        /**
         * The HTTPS service URL. A subdomain of onrender.com, by default.
         */
        url?: string;
    }
    /**
     * serviceStaticSiteProvideDefaults sets the appropriate defaults for ServiceStaticSite
     */
    export function serviceStaticSiteProvideDefaults(val: ServiceStaticSite): ServiceStaticSite {
        return {
            ...val,
            publishPath: (val.publishPath) ?? "public",
            pullRequestPreviewsEnabled: (val.pullRequestPreviewsEnabled) ?? "no",
            type: (val.type) ?? "static_site",
        };
    }

    export interface ServiceStaticSiteParentServerProperties {
        id?: string;
        name?: string;
    }

    /**
     * A route object for a static site
     */
    export interface ServiceStaticSiteRoute {
        destination: string;
        source: string;
        type: enums.services.ServiceStaticSiteRouteType;
    }

    export interface ServiceWebService {
        disk?: outputs.services.ServiceDisk;
        env: enums.services.ServiceWebServiceEnv;
        envSpecificDetails?: outputs.services.ServiceDockerDetails | outputs.services.ServiceNativeEnvironmentDetails;
        healthCheckPath?: string;
        numInstances?: number;
        plan?: enums.services.ServiceWebServicePlan;
        pullRequestPreviewsEnabled?: enums.services.ServiceWebServicePullRequestPreviewsEnabled;
        region?: enums.services.ServiceWebServiceRegion;
        type?: string;
    }
    /**
     * serviceWebServiceProvideDefaults sets the appropriate defaults for ServiceWebService
     */
    export function serviceWebServiceProvideDefaults(val: ServiceWebService): ServiceWebService {
        return {
            ...val,
            disk: (val.disk ? outputs.services.serviceDiskProvideDefaults(val.disk) : undefined),
            numInstances: (val.numInstances) ?? 1,
            plan: (val.plan) ?? "starter",
            pullRequestPreviewsEnabled: (val.pullRequestPreviewsEnabled) ?? "no",
            region: (val.region) ?? "oregon",
            type: (val.type) ?? "web_service",
        };
    }

}
