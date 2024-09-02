// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { EnvironmentArgs } from "./environment";
export type Environment = import("./environment").Environment;
export const Environment: typeof import("./environment").Environment = null as any;
utilities.lazyLoad(exports, ["Environment"], () => require("./environment"));

export { GetEnvironmentArgs, GetEnvironmentResult, GetEnvironmentOutputArgs } from "./getEnvironment";
export const getEnvironment: typeof import("./getEnvironment").getEnvironment = null as any;
export const getEnvironmentOutput: typeof import("./getEnvironment").getEnvironmentOutput = null as any;
utilities.lazyLoad(exports, ["getEnvironment","getEnvironmentOutput"], () => require("./getEnvironment"));

export { ListEnvironmentsArgs, ListEnvironmentsResult } from "./listEnvironments";
export const listEnvironments: typeof import("./listEnvironments").listEnvironments = null as any;
export const listEnvironmentsOutput: typeof import("./listEnvironments").listEnvironmentsOutput = null as any;
utilities.lazyLoad(exports, ["listEnvironments","listEnvironmentsOutput"], () => require("./listEnvironments"));

export { ResourcesToEnvironmentArgs } from "./resourcesToEnvironment";
export type ResourcesToEnvironment = import("./resourcesToEnvironment").ResourcesToEnvironment;
export const ResourcesToEnvironment: typeof import("./resourcesToEnvironment").ResourcesToEnvironment = null as any;
utilities.lazyLoad(exports, ["ResourcesToEnvironment"], () => require("./resourcesToEnvironment"));


// Export enums:
export * from "../types/enums/environments";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "render:environments:Environment":
                return new Environment(name, <any>undefined, { urn })
            case "render:environments:ResourcesToEnvironment":
                return new ResourcesToEnvironment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("render", "environments", _module)
