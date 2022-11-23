// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * A background worker service
 */
export interface BackgroundWorker {
    /**
     * Whether to auto deploy the service or not upon git push.
     */
    autoDeploy?: enums.services.ServiceAutoDeploy;
    /**
     * If left empty, this will fall back to the default branch of the repository.
     */
    branch?: string;
    createdAt?: string;
    envVars?: outputs.services.EnvVarKeyValue[];
    name: string;
    /**
     * The notification setting for this service upon deployment failure.
     */
    notifyOnFail?: enums.services.ServiceNotifyOnFail;
    /**
     * The id of the owner (user/team).
     */
    ownerId: string;
    /**
     * Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
     */
    repo: string;
    secretFiles?: outputs.services.SecretFile[];
    serviceDetails?: outputs.services.BackgroundWorkerServiceDetails;
    slug?: string;
    suspended?: enums.services.ServiceSuspended;
    suspenders?: string[];
    type?: string;
    updatedAt?: string;
}
/**
 * backgroundWorkerProvideDefaults sets the appropriate defaults for BackgroundWorker
 */
export function backgroundWorkerProvideDefaults(val: BackgroundWorker): BackgroundWorker {
    return {
        ...val,
        autoDeploy: (val.autoDeploy) ?? "no",
        serviceDetails: (val.serviceDetails ? outputs.services.backgroundWorkerServiceDetailsProvideDefaults(val.serviceDetails) : undefined),
        type: (val.type) ?? "background_worker",
    };
}

export interface BackgroundWorkerServiceDetails {
    disk?: outputs.services.Disk;
    env: enums.services.BackgroundWorkerServiceDetailsEnv;
    envSpecificDetails?: outputs.services.DockerDetails | outputs.services.NativeEnvironmentDetails;
    numInstances?: number;
    parentServer?: outputs.services.BackgroundWorkerServiceDetailsParentServerProperties;
    plan?: enums.services.BackgroundWorkerServiceDetailsPlan;
    pullRequestPreviewsEnabled?: enums.services.BackgroundWorkerServiceDetailsPullRequestPreviewsEnabled;
    region?: enums.services.BackgroundWorkerServiceDetailsRegion;
    url?: string;
}
/**
 * backgroundWorkerServiceDetailsProvideDefaults sets the appropriate defaults for BackgroundWorkerServiceDetails
 */
export function backgroundWorkerServiceDetailsProvideDefaults(val: BackgroundWorkerServiceDetails): BackgroundWorkerServiceDetails {
    return {
        ...val,
        disk: (val.disk ? outputs.services.diskProvideDefaults(val.disk) : undefined),
        numInstances: (val.numInstances) ?? 1,
        plan: (val.plan) ?? "starter",
        pullRequestPreviewsEnabled: (val.pullRequestPreviewsEnabled) ?? "no",
        region: (val.region) ?? "oregon",
    };
}

export interface BackgroundWorkerServiceDetailsParentServerProperties {
    id?: string;
    name?: string;
}

export interface Commit {
    createdAt?: string;
    id?: string;
    message?: string;
}

/**
 * A cron job
 */
export interface CronJob {
    /**
     * Whether to auto deploy the service or not upon git push.
     */
    autoDeploy?: enums.services.ServiceAutoDeploy;
    /**
     * If left empty, this will fall back to the default branch of the repository.
     */
    branch?: string;
    createdAt?: string;
    envVars?: outputs.services.EnvVarKeyValue[];
    name: string;
    /**
     * The notification setting for this service upon deployment failure.
     */
    notifyOnFail?: enums.services.ServiceNotifyOnFail;
    /**
     * The id of the owner (user/team).
     */
    ownerId: string;
    /**
     * Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
     */
    repo: string;
    secretFiles?: outputs.services.SecretFile[];
    serviceDetails?: outputs.services.CronJobServiceDetails;
    slug?: string;
    suspended?: enums.services.ServiceSuspended;
    suspenders?: string[];
    type?: string;
    updatedAt?: string;
}
/**
 * cronJobProvideDefaults sets the appropriate defaults for CronJob
 */
export function cronJobProvideDefaults(val: CronJob): CronJob {
    return {
        ...val,
        autoDeploy: (val.autoDeploy) ?? "no",
        serviceDetails: (val.serviceDetails ? outputs.services.cronJobServiceDetailsProvideDefaults(val.serviceDetails) : undefined),
        type: (val.type) ?? "cron_job",
    };
}

export interface CronJobServiceDetails {
    env: enums.services.CronJobServiceDetailsEnv;
    envSpecificDetails?: outputs.services.DockerDetails | outputs.services.NativeEnvironmentDetails;
    lastSuccessfulRunAt?: string;
    plan?: enums.services.CronJobServiceDetailsPlan;
    region?: enums.services.CronJobServiceDetailsRegion;
    schedule: string;
}
/**
 * cronJobServiceDetailsProvideDefaults sets the appropriate defaults for CronJobServiceDetails
 */
export function cronJobServiceDetailsProvideDefaults(val: CronJobServiceDetails): CronJobServiceDetails {
    return {
        ...val,
        plan: (val.plan) ?? "starter",
        region: (val.region) ?? "oregon",
    };
}

export interface CustomDomain {
    createdAt?: string;
    domainType: enums.services.CustomDomainDomainType;
    name: string;
    publicSuffix?: string;
    redirectForName: string;
    server: outputs.services.CustomDomainServerProperties;
    verificationStatus: enums.services.CustomDomainVerificationStatus;
}

export interface CustomDomainServerProperties {
    id?: string;
    name?: string;
}

export interface Deploy {
    clearCache?: enums.services.DeployClearCache;
    commit?: outputs.services.Commit;
}
/**
 * deployProvideDefaults sets the appropriate defaults for Deploy
 */
export function deployProvideDefaults(val: Deploy): Deploy {
    return {
        ...val,
        clearCache: (val.clearCache) ?? "do_not_clear",
    };
}

export interface Disk {
    mountPath: string;
    name: string;
    sizeGB?: number;
}
/**
 * diskProvideDefaults sets the appropriate defaults for Disk
 */
export function diskProvideDefaults(val: Disk): Disk {
    return {
        ...val,
        sizeGB: (val.sizeGB) ?? 1,
    };
}

export interface DockerDetails {
    dockerCommand: string;
    dockerContext: string;
    dockerfilePath?: string;
}

export interface EnvVarKeyValue {
    generateValue?: enums.services.EnvVarKeyValueGenerateValue;
    key: string;
    value?: string;
}

/**
 * A background worker service
 */
export interface GetBackgroundWorker {
    /**
     * Whether to auto deploy the service or not upon git push.
     */
    autoDeploy?: enums.services.ServiceAutoDeploy;
    /**
     * If left empty, this will fall back to the default branch of the repository.
     */
    branch?: string;
    createdAt?: string;
    envVars?: outputs.services.EnvVarKeyValue[];
    name: string;
    /**
     * The notification setting for this service upon deployment failure.
     */
    notifyOnFail?: enums.services.ServiceNotifyOnFail;
    /**
     * The id of the owner (user/team).
     */
    ownerId: string;
    /**
     * Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
     */
    repo: string;
    secretFiles?: outputs.services.SecretFile[];
    serviceDetails?: outputs.services.BackgroundWorkerServiceDetails;
    slug?: string;
    suspended?: enums.services.ServiceSuspended;
    suspenders?: string[];
    type?: string;
    updatedAt?: string;
}
/**
 * getBackgroundWorkerProvideDefaults sets the appropriate defaults for GetBackgroundWorker
 */
export function getBackgroundWorkerProvideDefaults(val: GetBackgroundWorker): GetBackgroundWorker {
    return {
        ...val,
        autoDeploy: (val.autoDeploy) ?? "no",
        serviceDetails: (val.serviceDetails ? outputs.services.backgroundWorkerServiceDetailsProvideDefaults(val.serviceDetails) : undefined),
        type: (val.type) ?? "background_worker",
    };
}

/**
 * A cron job
 */
export interface GetCronJob {
    /**
     * Whether to auto deploy the service or not upon git push.
     */
    autoDeploy?: enums.services.ServiceAutoDeploy;
    /**
     * If left empty, this will fall back to the default branch of the repository.
     */
    branch?: string;
    createdAt?: string;
    envVars?: outputs.services.EnvVarKeyValue[];
    name: string;
    /**
     * The notification setting for this service upon deployment failure.
     */
    notifyOnFail?: enums.services.ServiceNotifyOnFail;
    /**
     * The id of the owner (user/team).
     */
    ownerId: string;
    /**
     * Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
     */
    repo: string;
    secretFiles?: outputs.services.SecretFile[];
    serviceDetails?: outputs.services.CronJobServiceDetails;
    slug?: string;
    suspended?: enums.services.ServiceSuspended;
    suspenders?: string[];
    type?: string;
    updatedAt?: string;
}
/**
 * getCronJobProvideDefaults sets the appropriate defaults for GetCronJob
 */
export function getCronJobProvideDefaults(val: GetCronJob): GetCronJob {
    return {
        ...val,
        autoDeploy: (val.autoDeploy) ?? "no",
        serviceDetails: (val.serviceDetails ? outputs.services.cronJobServiceDetailsProvideDefaults(val.serviceDetails) : undefined),
        type: (val.type) ?? "cron_job",
    };
}

/**
 * A private service
 */
export interface GetPrivateService {
    /**
     * Whether to auto deploy the service or not upon git push.
     */
    autoDeploy?: enums.services.ServiceAutoDeploy;
    /**
     * If left empty, this will fall back to the default branch of the repository.
     */
    branch?: string;
    createdAt?: string;
    envVars?: outputs.services.EnvVarKeyValue[];
    name: string;
    /**
     * The notification setting for this service upon deployment failure.
     */
    notifyOnFail?: enums.services.ServiceNotifyOnFail;
    /**
     * The id of the owner (user/team).
     */
    ownerId: string;
    /**
     * Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
     */
    repo: string;
    secretFiles?: outputs.services.SecretFile[];
    serviceDetails?: outputs.services.PrivateServiceDetails;
    slug?: string;
    suspended?: enums.services.ServiceSuspended;
    suspenders?: string[];
    type?: string;
    updatedAt?: string;
}
/**
 * getPrivateServiceProvideDefaults sets the appropriate defaults for GetPrivateService
 */
export function getPrivateServiceProvideDefaults(val: GetPrivateService): GetPrivateService {
    return {
        ...val,
        autoDeploy: (val.autoDeploy) ?? "no",
        serviceDetails: (val.serviceDetails ? outputs.services.privateServiceDetailsProvideDefaults(val.serviceDetails) : undefined),
        type: (val.type) ?? "private_service",
    };
}

/**
 * A static website service
 */
export interface GetStaticSite {
    /**
     * Whether to auto deploy the service or not upon git push.
     */
    autoDeploy?: enums.services.ServiceAutoDeploy;
    /**
     * If left empty, this will fall back to the default branch of the repository.
     */
    branch?: string;
    createdAt?: string;
    envVars?: outputs.services.EnvVarKeyValue[];
    name: string;
    /**
     * The notification setting for this service upon deployment failure.
     */
    notifyOnFail?: enums.services.ServiceNotifyOnFail;
    /**
     * The id of the owner (user/team).
     */
    ownerId: string;
    /**
     * Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
     */
    repo: string;
    secretFiles?: outputs.services.SecretFile[];
    serviceDetails?: outputs.services.StaticSiteServiceDetails;
    slug?: string;
    suspended?: enums.services.ServiceSuspended;
    suspenders?: string[];
    type?: string;
    updatedAt?: string;
}
/**
 * getStaticSiteProvideDefaults sets the appropriate defaults for GetStaticSite
 */
export function getStaticSiteProvideDefaults(val: GetStaticSite): GetStaticSite {
    return {
        ...val,
        autoDeploy: (val.autoDeploy) ?? "no",
        serviceDetails: (val.serviceDetails ? outputs.services.staticSiteServiceDetailsProvideDefaults(val.serviceDetails) : undefined),
        type: (val.type) ?? "static_site",
    };
}

/**
 * A web service
 */
export interface GetWebService {
    /**
     * Whether to auto deploy the service or not upon git push.
     */
    autoDeploy?: enums.services.ServiceAutoDeploy;
    /**
     * If left empty, this will fall back to the default branch of the repository.
     */
    branch?: string;
    createdAt?: string;
    envVars?: outputs.services.EnvVarKeyValue[];
    name: string;
    /**
     * The notification setting for this service upon deployment failure.
     */
    notifyOnFail?: enums.services.ServiceNotifyOnFail;
    /**
     * The id of the owner (user/team).
     */
    ownerId: string;
    /**
     * Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
     */
    repo: string;
    secretFiles?: outputs.services.SecretFile[];
    serviceDetails?: outputs.services.WebServiceServiceDetails;
    slug?: string;
    suspended?: enums.services.ServiceSuspended;
    suspenders?: string[];
    type?: string;
    updatedAt?: string;
}
/**
 * getWebServiceProvideDefaults sets the appropriate defaults for GetWebService
 */
export function getWebServiceProvideDefaults(val: GetWebService): GetWebService {
    return {
        ...val,
        autoDeploy: (val.autoDeploy) ?? "no",
        serviceDetails: (val.serviceDetails ? outputs.services.webServiceServiceDetailsProvideDefaults(val.serviceDetails) : undefined),
        type: (val.type) ?? "web_service",
    };
}

export interface Job {
    createdAt?: string;
    finishedAt?: string;
    planId: string;
    startCommand: string;
    startedAt?: string;
    status?: string;
}

export interface ListCustomDomainsResponse {
    cursor?: string;
    customDomain?: outputs.services.CustomDomain;
}

export interface ListDeploysResponse {
    cursor?: string;
    customDomain?: outputs.services.Deploy;
}
/**
 * listDeploysResponseProvideDefaults sets the appropriate defaults for ListDeploysResponse
 */
export function listDeploysResponseProvideDefaults(val: ListDeploysResponse): ListDeploysResponse {
    return {
        ...val,
        customDomain: (val.customDomain ? outputs.services.deployProvideDefaults(val.customDomain) : undefined),
    };
}

export interface ListEnvVarsResponse {
    cursor?: string;
    envVar?: outputs.services.EnvVarKeyValue;
}

export interface ListJobsResponse {
    cursor?: string;
    job?: outputs.services.Job;
}

export interface ListServiceHeadersResponse {
    cursor?: string;
    /**
     * A service header object
     */
    header?: outputs.services.ServiceHeader;
}

export interface ListServiceResponse {
    cursor?: string;
    service?: outputs.services.StaticSite | outputs.services.WebService | outputs.services.PrivateService | outputs.services.BackgroundWorker | outputs.services.CronJob;
}

export interface ListStaticSiteRoutesResponse {
    cursor?: string;
    /**
     * A route object for a static site
     */
    route?: outputs.services.StaticSiteRoute;
}

export interface NativeEnvironmentDetails {
    buildCommand: string;
    startCommand: string;
}

export interface OpenPorts {
    port?: number;
    protocol?: enums.services.OpenPortsProtocol;
}

/**
 * A private service
 */
export interface PrivateService {
    /**
     * Whether to auto deploy the service or not upon git push.
     */
    autoDeploy?: enums.services.ServiceAutoDeploy;
    /**
     * If left empty, this will fall back to the default branch of the repository.
     */
    branch?: string;
    createdAt?: string;
    envVars?: outputs.services.EnvVarKeyValue[];
    name: string;
    /**
     * The notification setting for this service upon deployment failure.
     */
    notifyOnFail?: enums.services.ServiceNotifyOnFail;
    /**
     * The id of the owner (user/team).
     */
    ownerId: string;
    /**
     * Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
     */
    repo: string;
    secretFiles?: outputs.services.SecretFile[];
    serviceDetails?: outputs.services.PrivateServiceDetails;
    slug?: string;
    suspended?: enums.services.ServiceSuspended;
    suspenders?: string[];
    type?: string;
    updatedAt?: string;
}
/**
 * privateServiceProvideDefaults sets the appropriate defaults for PrivateService
 */
export function privateServiceProvideDefaults(val: PrivateService): PrivateService {
    return {
        ...val,
        autoDeploy: (val.autoDeploy) ?? "no",
        serviceDetails: (val.serviceDetails ? outputs.services.privateServiceDetailsProvideDefaults(val.serviceDetails) : undefined),
        type: (val.type) ?? "private_service",
    };
}

export interface PrivateServiceDetails {
    disk?: outputs.services.Disk;
    env: enums.services.PrivateServiceDetailsEnv;
    envSpecificDetails?: outputs.services.DockerDetails | outputs.services.NativeEnvironmentDetails;
    numInstances?: number;
    openPorts?: outputs.services.OpenPorts[];
    parentServer?: outputs.services.PrivateServiceDetailsParentServerProperties;
    plan?: enums.services.PrivateServiceDetailsPlan;
    pullRequestPreviewsEnabled?: enums.services.PrivateServiceDetailsPullRequestPreviewsEnabled;
    region?: enums.services.PrivateServiceDetailsRegion;
    url?: string;
}
/**
 * privateServiceDetailsProvideDefaults sets the appropriate defaults for PrivateServiceDetails
 */
export function privateServiceDetailsProvideDefaults(val: PrivateServiceDetails): PrivateServiceDetails {
    return {
        ...val,
        disk: (val.disk ? outputs.services.diskProvideDefaults(val.disk) : undefined),
        numInstances: (val.numInstances) ?? 1,
        plan: (val.plan) ?? "starter",
        pullRequestPreviewsEnabled: (val.pullRequestPreviewsEnabled) ?? "no",
        region: (val.region) ?? "oregon",
    };
}

export interface PrivateServiceDetailsParentServerProperties {
    id?: string;
    name?: string;
}

export interface SecretFile {
    contents: string;
    name: string;
}

export interface ServerProperties {
    id?: string;
    name?: string;
}

/**
 * A service header object
 */
export interface ServiceHeader {
    name: string;
    path: string;
    value: string;
}

/**
 * A static website service
 */
export interface StaticSite {
    /**
     * Whether to auto deploy the service or not upon git push.
     */
    autoDeploy?: enums.services.ServiceAutoDeploy;
    /**
     * If left empty, this will fall back to the default branch of the repository.
     */
    branch?: string;
    createdAt?: string;
    envVars?: outputs.services.EnvVarKeyValue[];
    name: string;
    /**
     * The notification setting for this service upon deployment failure.
     */
    notifyOnFail?: enums.services.ServiceNotifyOnFail;
    /**
     * The id of the owner (user/team).
     */
    ownerId: string;
    /**
     * Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
     */
    repo: string;
    secretFiles?: outputs.services.SecretFile[];
    serviceDetails?: outputs.services.StaticSiteServiceDetails;
    slug?: string;
    suspended?: enums.services.ServiceSuspended;
    suspenders?: string[];
    type?: string;
    updatedAt?: string;
}
/**
 * staticSiteProvideDefaults sets the appropriate defaults for StaticSite
 */
export function staticSiteProvideDefaults(val: StaticSite): StaticSite {
    return {
        ...val,
        autoDeploy: (val.autoDeploy) ?? "no",
        serviceDetails: (val.serviceDetails ? outputs.services.staticSiteServiceDetailsProvideDefaults(val.serviceDetails) : undefined),
        type: (val.type) ?? "static_site",
    };
}

/**
 * A route object for a static site
 */
export interface StaticSiteRoute {
    destination: string;
    source: string;
    type: enums.services.StaticSiteRouteType;
}

export interface StaticSiteServiceDetails {
    buildCommand?: string;
    headers?: outputs.services.ServiceHeader[];
    parentServer?: outputs.services.StaticSiteServiceDetailsParentServerProperties;
    publishPath?: string;
    pullRequestPreviewsEnabled?: enums.services.StaticSiteServiceDetailsPullRequestPreviewsEnabled;
    routes?: outputs.services.StaticSiteRoute[];
    /**
     * The HTTPS service URL. A subdomain of onrender.com, by default.
     */
    url?: string;
}
/**
 * staticSiteServiceDetailsProvideDefaults sets the appropriate defaults for StaticSiteServiceDetails
 */
export function staticSiteServiceDetailsProvideDefaults(val: StaticSiteServiceDetails): StaticSiteServiceDetails {
    return {
        ...val,
        publishPath: (val.publishPath) ?? "public",
        pullRequestPreviewsEnabled: (val.pullRequestPreviewsEnabled) ?? "no",
    };
}

export interface StaticSiteServiceDetailsParentServerProperties {
    id?: string;
    name?: string;
}

/**
 * A web service
 */
export interface WebService {
    /**
     * Whether to auto deploy the service or not upon git push.
     */
    autoDeploy?: enums.services.ServiceAutoDeploy;
    /**
     * If left empty, this will fall back to the default branch of the repository.
     */
    branch?: string;
    createdAt?: string;
    envVars?: outputs.services.EnvVarKeyValue[];
    name: string;
    /**
     * The notification setting for this service upon deployment failure.
     */
    notifyOnFail?: enums.services.ServiceNotifyOnFail;
    /**
     * The id of the owner (user/team).
     */
    ownerId: string;
    /**
     * Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
     */
    repo: string;
    secretFiles?: outputs.services.SecretFile[];
    serviceDetails?: outputs.services.WebServiceServiceDetails;
    slug?: string;
    suspended?: enums.services.ServiceSuspended;
    suspenders?: string[];
    type?: string;
    updatedAt?: string;
}
/**
 * webServiceProvideDefaults sets the appropriate defaults for WebService
 */
export function webServiceProvideDefaults(val: WebService): WebService {
    return {
        ...val,
        autoDeploy: (val.autoDeploy) ?? "no",
        serviceDetails: (val.serviceDetails ? outputs.services.webServiceServiceDetailsProvideDefaults(val.serviceDetails) : undefined),
        type: (val.type) ?? "web_service",
    };
}

export interface WebServiceServiceDetails {
    disk?: outputs.services.Disk;
    env: enums.services.WebServiceServiceDetailsEnv;
    envSpecificDetails?: outputs.services.DockerDetails | outputs.services.NativeEnvironmentDetails;
    healthCheckPath?: string;
    numInstances?: number;
    openPorts?: outputs.services.OpenPorts[];
    parentServer?: outputs.services.WebServiceServiceDetailsParentServerProperties;
    plan?: enums.services.WebServiceServiceDetailsPlan;
    pullRequestPreviewsEnabled?: enums.services.WebServiceServiceDetailsPullRequestPreviewsEnabled;
    region?: enums.services.WebServiceServiceDetailsRegion;
    url?: string;
}
/**
 * webServiceServiceDetailsProvideDefaults sets the appropriate defaults for WebServiceServiceDetails
 */
export function webServiceServiceDetailsProvideDefaults(val: WebServiceServiceDetails): WebServiceServiceDetails {
    return {
        ...val,
        disk: (val.disk ? outputs.services.diskProvideDefaults(val.disk) : undefined),
        numInstances: (val.numInstances) ?? 1,
        plan: (val.plan) ?? "starter",
        pullRequestPreviewsEnabled: (val.pullRequestPreviewsEnabled) ?? "no",
        region: (val.region) ?? "oregon",
    };
}

export interface WebServiceServiceDetailsParentServerProperties {
    id?: string;
    name?: string;
}
