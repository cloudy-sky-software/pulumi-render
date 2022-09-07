// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { BackgroundWorkerArgs } from "./backgroundWorker";
export type BackgroundWorker = import("./backgroundWorker").BackgroundWorker;
export const BackgroundWorker: typeof import("./backgroundWorker").BackgroundWorker = null as any;

export { CronJobArgs } from "./cronJob";
export type CronJob = import("./cronJob").CronJob;
export const CronJob: typeof import("./cronJob").CronJob = null as any;

export { CustomDomainArgs } from "./customDomain";
export type CustomDomain = import("./customDomain").CustomDomain;
export const CustomDomain: typeof import("./customDomain").CustomDomain = null as any;

export { DeployArgs } from "./deploy";
export type Deploy = import("./deploy").Deploy;
export const Deploy: typeof import("./deploy").Deploy = null as any;

export { GetBackgroundWorkerArgs, GetBackgroundWorkerResult, GetBackgroundWorkerOutputArgs } from "./getBackgroundWorker";
export const getBackgroundWorker: typeof import("./getBackgroundWorker").getBackgroundWorker = null as any;
export const getBackgroundWorkerOutput: typeof import("./getBackgroundWorker").getBackgroundWorkerOutput = null as any;

export { GetCronJobArgs, GetCronJobResult, GetCronJobOutputArgs } from "./getCronJob";
export const getCronJob: typeof import("./getCronJob").getCronJob = null as any;
export const getCronJobOutput: typeof import("./getCronJob").getCronJobOutput = null as any;

export { GetCustomDomainArgs, GetCustomDomainResult, GetCustomDomainOutputArgs } from "./getCustomDomain";
export const getCustomDomain: typeof import("./getCustomDomain").getCustomDomain = null as any;
export const getCustomDomainOutput: typeof import("./getCustomDomain").getCustomDomainOutput = null as any;

export { GetDeployArgs, GetDeployResult, GetDeployOutputArgs } from "./getDeploy";
export const getDeploy: typeof import("./getDeploy").getDeploy = null as any;
export const getDeployOutput: typeof import("./getDeploy").getDeployOutput = null as any;

export { GetPrivateServiceArgs, GetPrivateServiceResult, GetPrivateServiceOutputArgs } from "./getPrivateService";
export const getPrivateService: typeof import("./getPrivateService").getPrivateService = null as any;
export const getPrivateServiceOutput: typeof import("./getPrivateService").getPrivateServiceOutput = null as any;

export { GetStaticSiteArgs, GetStaticSiteResult, GetStaticSiteOutputArgs } from "./getStaticSite";
export const getStaticSite: typeof import("./getStaticSite").getStaticSite = null as any;
export const getStaticSiteOutput: typeof import("./getStaticSite").getStaticSiteOutput = null as any;

export { GetWebServiceArgs, GetWebServiceResult, GetWebServiceOutputArgs } from "./getWebService";
export const getWebService: typeof import("./getWebService").getWebService = null as any;
export const getWebServiceOutput: typeof import("./getWebService").getWebServiceOutput = null as any;

export { ListCustomDomainsArgs, ListCustomDomainsResult, ListCustomDomainsOutputArgs } from "./listCustomDomains";
export const listCustomDomains: typeof import("./listCustomDomains").listCustomDomains = null as any;
export const listCustomDomainsOutput: typeof import("./listCustomDomains").listCustomDomainsOutput = null as any;

export { ListDeploysArgs, ListDeploysResult, ListDeploysOutputArgs } from "./listDeploys";
export const listDeploys: typeof import("./listDeploys").listDeploys = null as any;
export const listDeploysOutput: typeof import("./listDeploys").listDeploysOutput = null as any;

export { ListEnvVarsArgs, ListEnvVarsResult, ListEnvVarsOutputArgs } from "./listEnvVars";
export const listEnvVars: typeof import("./listEnvVars").listEnvVars = null as any;
export const listEnvVarsOutput: typeof import("./listEnvVars").listEnvVarsOutput = null as any;

export { ListServiceHeadersArgs, ListServiceHeadersResult, ListServiceHeadersOutputArgs } from "./listServiceHeaders";
export const listServiceHeaders: typeof import("./listServiceHeaders").listServiceHeaders = null as any;
export const listServiceHeadersOutput: typeof import("./listServiceHeaders").listServiceHeadersOutput = null as any;

export { ListServicesArgs, ListServicesResult } from "./listServices";
export const listServices: typeof import("./listServices").listServices = null as any;

export { ListStaticSiteRoutesArgs, ListStaticSiteRoutesResult, ListStaticSiteRoutesOutputArgs } from "./listStaticSiteRoutes";
export const listStaticSiteRoutes: typeof import("./listStaticSiteRoutes").listStaticSiteRoutes = null as any;
export const listStaticSiteRoutesOutput: typeof import("./listStaticSiteRoutes").listStaticSiteRoutesOutput = null as any;

export { PrivateServiceArgs } from "./privateService";
export type PrivateService = import("./privateService").PrivateService;
export const PrivateService: typeof import("./privateService").PrivateService = null as any;

export { ScaleArgs } from "./scale";
export type Scale = import("./scale").Scale;
export const Scale: typeof import("./scale").Scale = null as any;

export { StaticSiteArgs } from "./staticSite";
export type StaticSite = import("./staticSite").StaticSite;
export const StaticSite: typeof import("./staticSite").StaticSite = null as any;

export { SuspendArgs } from "./suspend";
export type Suspend = import("./suspend").Suspend;
export const Suspend: typeof import("./suspend").Suspend = null as any;

export { WebServiceArgs } from "./webService";
export type WebService = import("./webService").WebService;
export const WebService: typeof import("./webService").WebService = null as any;

utilities.lazyLoad(exports, ["BackgroundWorker"], () => require("./backgroundWorker"));
utilities.lazyLoad(exports, ["CronJob"], () => require("./cronJob"));
utilities.lazyLoad(exports, ["CustomDomain"], () => require("./customDomain"));
utilities.lazyLoad(exports, ["Deploy"], () => require("./deploy"));
utilities.lazyLoad(exports, ["getBackgroundWorker","getBackgroundWorkerOutput"], () => require("./getBackgroundWorker"));
utilities.lazyLoad(exports, ["getCronJob","getCronJobOutput"], () => require("./getCronJob"));
utilities.lazyLoad(exports, ["getCustomDomain","getCustomDomainOutput"], () => require("./getCustomDomain"));
utilities.lazyLoad(exports, ["getDeploy","getDeployOutput"], () => require("./getDeploy"));
utilities.lazyLoad(exports, ["getPrivateService","getPrivateServiceOutput"], () => require("./getPrivateService"));
utilities.lazyLoad(exports, ["getStaticSite","getStaticSiteOutput"], () => require("./getStaticSite"));
utilities.lazyLoad(exports, ["getWebService","getWebServiceOutput"], () => require("./getWebService"));
utilities.lazyLoad(exports, ["listCustomDomains","listCustomDomainsOutput"], () => require("./listCustomDomains"));
utilities.lazyLoad(exports, ["listDeploys","listDeploysOutput"], () => require("./listDeploys"));
utilities.lazyLoad(exports, ["listEnvVars","listEnvVarsOutput"], () => require("./listEnvVars"));
utilities.lazyLoad(exports, ["listServiceHeaders","listServiceHeadersOutput"], () => require("./listServiceHeaders"));
utilities.lazyLoad(exports, ["listServices"], () => require("./listServices"));
utilities.lazyLoad(exports, ["listStaticSiteRoutes","listStaticSiteRoutesOutput"], () => require("./listStaticSiteRoutes"));
utilities.lazyLoad(exports, ["PrivateService"], () => require("./privateService"));
utilities.lazyLoad(exports, ["Scale"], () => require("./scale"));
utilities.lazyLoad(exports, ["StaticSite"], () => require("./staticSite"));
utilities.lazyLoad(exports, ["Suspend"], () => require("./suspend"));
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
            case "render:services:CustomDomain":
                return new CustomDomain(name, <any>undefined, { urn })
            case "render:services:Deploy":
                return new Deploy(name, <any>undefined, { urn })
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
