// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { EnvGroupArgs } from "./envGroup";
export type EnvGroup = import("./envGroup").EnvGroup;
export const EnvGroup: typeof import("./envGroup").EnvGroup = null as any;
utilities.lazyLoad(exports, ["EnvGroup"], () => require("./envGroup"));

export { GetEnvGroupArgs, GetEnvGroupOutputArgs } from "./getEnvGroup";
export const getEnvGroup: typeof import("./getEnvGroup").getEnvGroup = null as any;
export const getEnvGroupOutput: typeof import("./getEnvGroup").getEnvGroupOutput = null as any;
utilities.lazyLoad(exports, ["getEnvGroup","getEnvGroupOutput"], () => require("./getEnvGroup"));

export { GetEnvGroupEnvVarArgs, GetEnvGroupEnvVarOutputArgs } from "./getEnvGroupEnvVar";
export const getEnvGroupEnvVar: typeof import("./getEnvGroupEnvVar").getEnvGroupEnvVar = null as any;
export const getEnvGroupEnvVarOutput: typeof import("./getEnvGroupEnvVar").getEnvGroupEnvVarOutput = null as any;
utilities.lazyLoad(exports, ["getEnvGroupEnvVar","getEnvGroupEnvVarOutput"], () => require("./getEnvGroupEnvVar"));

export { GetEnvGroupSecretFileArgs, GetEnvGroupSecretFileOutputArgs } from "./getEnvGroupSecretFile";
export const getEnvGroupSecretFile: typeof import("./getEnvGroupSecretFile").getEnvGroupSecretFile = null as any;
export const getEnvGroupSecretFileOutput: typeof import("./getEnvGroupSecretFile").getEnvGroupSecretFileOutput = null as any;
utilities.lazyLoad(exports, ["getEnvGroupSecretFile","getEnvGroupSecretFileOutput"], () => require("./getEnvGroupSecretFile"));

export { LinkServiceToEnvGroupArgs } from "./linkServiceToEnvGroup";
export type LinkServiceToEnvGroup = import("./linkServiceToEnvGroup").LinkServiceToEnvGroup;
export const LinkServiceToEnvGroup: typeof import("./linkServiceToEnvGroup").LinkServiceToEnvGroup = null as any;
utilities.lazyLoad(exports, ["LinkServiceToEnvGroup"], () => require("./linkServiceToEnvGroup"));

export { ListEnvGroupsArgs, ListEnvGroupsResult } from "./listEnvGroups";
export const listEnvGroups: typeof import("./listEnvGroups").listEnvGroups = null as any;
export const listEnvGroupsOutput: typeof import("./listEnvGroups").listEnvGroupsOutput = null as any;
utilities.lazyLoad(exports, ["listEnvGroups","listEnvGroupsOutput"], () => require("./listEnvGroups"));


// Export enums:
export * from "../types/enums/envgroups";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "render:env-groups:EnvGroup":
                return new EnvGroup(name, <any>undefined, { urn })
            case "render:env-groups:LinkServiceToEnvGroup":
                return new LinkServiceToEnvGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("render", "env-groups", _module)
