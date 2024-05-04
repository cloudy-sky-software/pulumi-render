// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { BackgroundWorkerArgs } from "./backgroundWorker";
export type BackgroundWorker = import("./backgroundWorker").BackgroundWorker;
export const BackgroundWorker: typeof import("./backgroundWorker").BackgroundWorker = null as any;
utilities.lazyLoad(exports, ["BackgroundWorker"], () => require("./backgroundWorker"));

export { CronJobArgs } from "./cronJob";
export type CronJob = import("./cronJob").CronJob;
export const CronJob: typeof import("./cronJob").CronJob = null as any;
utilities.lazyLoad(exports, ["CronJob"], () => require("./cronJob"));

export { CustomDomainsArgs } from "./customDomains";
export type CustomDomains = import("./customDomains").CustomDomains;
export const CustomDomains: typeof import("./customDomains").CustomDomains = null as any;
utilities.lazyLoad(exports, ["CustomDomains"], () => require("./customDomains"));

export { DeploysArgs } from "./deploys";
export type Deploys = import("./deploys").Deploys;
export const Deploys: typeof import("./deploys").Deploys = null as any;
utilities.lazyLoad(exports, ["Deploys"], () => require("./deploys"));

export { EnvVarsArgs } from "./envVars";
export type EnvVars = import("./envVars").EnvVars;
export const EnvVars: typeof import("./envVars").EnvVars = null as any;
utilities.lazyLoad(exports, ["EnvVars"], () => require("./envVars"));

export { GetBackgroundWorkerArgs, GetBackgroundWorkerResult, GetBackgroundWorkerOutputArgs } from "./getBackgroundWorker";
export const getBackgroundWorker: typeof import("./getBackgroundWorker").getBackgroundWorker = null as any;
export const getBackgroundWorkerOutput: typeof import("./getBackgroundWorker").getBackgroundWorkerOutput = null as any;
utilities.lazyLoad(exports, ["getBackgroundWorker","getBackgroundWorkerOutput"], () => require("./getBackgroundWorker"));

export { GetCronJobArgs, GetCronJobResult, GetCronJobOutputArgs } from "./getCronJob";
export const getCronJob: typeof import("./getCronJob").getCronJob = null as any;
export const getCronJobOutput: typeof import("./getCronJob").getCronJobOutput = null as any;
utilities.lazyLoad(exports, ["getCronJob","getCronJobOutput"], () => require("./getCronJob"));

export { GetCustomDomainArgs, GetCustomDomainResult, GetCustomDomainOutputArgs } from "./getCustomDomain";
export const getCustomDomain: typeof import("./getCustomDomain").getCustomDomain = null as any;
export const getCustomDomainOutput: typeof import("./getCustomDomain").getCustomDomainOutput = null as any;
utilities.lazyLoad(exports, ["getCustomDomain","getCustomDomainOutput"], () => require("./getCustomDomain"));

export { GetDeployArgs, GetDeployResult, GetDeployOutputArgs } from "./getDeploy";
export const getDeploy: typeof import("./getDeploy").getDeploy = null as any;
export const getDeployOutput: typeof import("./getDeploy").getDeployOutput = null as any;
utilities.lazyLoad(exports, ["getDeploy","getDeployOutput"], () => require("./getDeploy"));

export { GetJobArgs, GetJobResult, GetJobOutputArgs } from "./getJob";
export const getJob: typeof import("./getJob").getJob = null as any;
export const getJobOutput: typeof import("./getJob").getJobOutput = null as any;
utilities.lazyLoad(exports, ["getJob","getJobOutput"], () => require("./getJob"));

export { GetPrivateServiceArgs, GetPrivateServiceResult, GetPrivateServiceOutputArgs } from "./getPrivateService";
export const getPrivateService: typeof import("./getPrivateService").getPrivateService = null as any;
export const getPrivateServiceOutput: typeof import("./getPrivateService").getPrivateServiceOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateService","getPrivateServiceOutput"], () => require("./getPrivateService"));

export { GetStaticSiteArgs, GetStaticSiteResult, GetStaticSiteOutputArgs } from "./getStaticSite";
export const getStaticSite: typeof import("./getStaticSite").getStaticSite = null as any;
export const getStaticSiteOutput: typeof import("./getStaticSite").getStaticSiteOutput = null as any;
utilities.lazyLoad(exports, ["getStaticSite","getStaticSiteOutput"], () => require("./getStaticSite"));

export { GetWebServiceArgs, GetWebServiceResult, GetWebServiceOutputArgs } from "./getWebService";
export const getWebService: typeof import("./getWebService").getWebService = null as any;
export const getWebServiceOutput: typeof import("./getWebService").getWebServiceOutput = null as any;
utilities.lazyLoad(exports, ["getWebService","getWebServiceOutput"], () => require("./getWebService"));

export { JobsArgs } from "./jobs";
export type Jobs = import("./jobs").Jobs;
export const Jobs: typeof import("./jobs").Jobs = null as any;
utilities.lazyLoad(exports, ["Jobs"], () => require("./jobs"));

export { ListCustomDomainsArgs, ListCustomDomainsResult, ListCustomDomainsOutputArgs } from "./listCustomDomains";
export const listCustomDomains: typeof import("./listCustomDomains").listCustomDomains = null as any;
export const listCustomDomainsOutput: typeof import("./listCustomDomains").listCustomDomainsOutput = null as any;
utilities.lazyLoad(exports, ["listCustomDomains","listCustomDomainsOutput"], () => require("./listCustomDomains"));

export { ListDeploysArgs, ListDeploysResult, ListDeploysOutputArgs } from "./listDeploys";
export const listDeploys: typeof import("./listDeploys").listDeploys = null as any;
export const listDeploysOutput: typeof import("./listDeploys").listDeploysOutput = null as any;
utilities.lazyLoad(exports, ["listDeploys","listDeploysOutput"], () => require("./listDeploys"));

export { ListEnvVarsArgs, ListEnvVarsResult, ListEnvVarsOutputArgs } from "./listEnvVars";
export const listEnvVars: typeof import("./listEnvVars").listEnvVars = null as any;
export const listEnvVarsOutput: typeof import("./listEnvVars").listEnvVarsOutput = null as any;
utilities.lazyLoad(exports, ["listEnvVars","listEnvVarsOutput"], () => require("./listEnvVars"));

export { ListHeadersArgs, ListHeadersResult, ListHeadersOutputArgs } from "./listHeaders";
export const listHeaders: typeof import("./listHeaders").listHeaders = null as any;
export const listHeadersOutput: typeof import("./listHeaders").listHeadersOutput = null as any;
utilities.lazyLoad(exports, ["listHeaders","listHeadersOutput"], () => require("./listHeaders"));

export { ListJobsArgs, ListJobsResult, ListJobsOutputArgs } from "./listJobs";
export const listJobs: typeof import("./listJobs").listJobs = null as any;
export const listJobsOutput: typeof import("./listJobs").listJobsOutput = null as any;
utilities.lazyLoad(exports, ["listJobs","listJobsOutput"], () => require("./listJobs"));

export { ListRoutesArgs, ListRoutesResult, ListRoutesOutputArgs } from "./listRoutes";
export const listRoutes: typeof import("./listRoutes").listRoutes = null as any;
export const listRoutesOutput: typeof import("./listRoutes").listRoutesOutput = null as any;
utilities.lazyLoad(exports, ["listRoutes","listRoutesOutput"], () => require("./listRoutes"));

export { ListServicesArgs, ListServicesResult } from "./listServices";
export const listServices: typeof import("./listServices").listServices = null as any;
export const listServicesOutput: typeof import("./listServices").listServicesOutput = null as any;
utilities.lazyLoad(exports, ["listServices","listServicesOutput"], () => require("./listServices"));

export { PrivateServiceArgs } from "./privateService";
export type PrivateService = import("./privateService").PrivateService;
export const PrivateService: typeof import("./privateService").PrivateService = null as any;
utilities.lazyLoad(exports, ["PrivateService"], () => require("./privateService"));

export { ScaleArgs } from "./scale";
export type Scale = import("./scale").Scale;
export const Scale: typeof import("./scale").Scale = null as any;
utilities.lazyLoad(exports, ["Scale"], () => require("./scale"));

export { StaticSiteArgs } from "./staticSite";
export type StaticSite = import("./staticSite").StaticSite;
export const StaticSite: typeof import("./staticSite").StaticSite = null as any;
utilities.lazyLoad(exports, ["StaticSite"], () => require("./staticSite"));

export { SuspendArgs } from "./suspend";
export type Suspend = import("./suspend").Suspend;
export const Suspend: typeof import("./suspend").Suspend = null as any;
utilities.lazyLoad(exports, ["Suspend"], () => require("./suspend"));

export { WebServiceArgs } from "./webService";
export type WebService = import("./webService").WebService;
export const WebService: typeof import("./webService").WebService = null as any;
utilities.lazyLoad(exports, ["WebService"], () => require("./webService"));


// Export enums:
export * from "../types/enums/services";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "render:services:BackgroundWorker":
                return new BackgroundWorker(name, <any>undefined, { urn })
            case "render:services:CronJob":
                return new CronJob(name, <any>undefined, { urn })
            case "render:services:CustomDomains":
                return new CustomDomains(name, <any>undefined, { urn })
            case "render:services:Deploys":
                return new Deploys(name, <any>undefined, { urn })
            case "render:services:EnvVars":
                return new EnvVars(name, <any>undefined, { urn })
            case "render:services:Jobs":
                return new Jobs(name, <any>undefined, { urn })
            case "render:services:PrivateService":
                return new PrivateService(name, <any>undefined, { urn })
            case "render:services:Scale":
                return new Scale(name, <any>undefined, { urn })
            case "render:services:StaticSite":
                return new StaticSite(name, <any>undefined, { urn })
            case "render:services:Suspend":
                return new Suspend(name, <any>undefined, { urn })
            case "render:services:WebService":
                return new WebService(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("render", "services", _module)
