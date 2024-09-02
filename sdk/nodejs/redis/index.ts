// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetRediArgs, GetRediResult, GetRediOutputArgs } from "./getRedi";
export const getRedi: typeof import("./getRedi").getRedi = null as any;
export const getRediOutput: typeof import("./getRedi").getRediOutput = null as any;
utilities.lazyLoad(exports, ["getRedi","getRediOutput"], () => require("./getRedi"));

export { GetRedisConnectionInfoArgs, GetRedisConnectionInfoResult, GetRedisConnectionInfoOutputArgs } from "./getRedisConnectionInfo";
export const getRedisConnectionInfo: typeof import("./getRedisConnectionInfo").getRedisConnectionInfo = null as any;
export const getRedisConnectionInfoOutput: typeof import("./getRedisConnectionInfo").getRedisConnectionInfoOutput = null as any;
utilities.lazyLoad(exports, ["getRedisConnectionInfo","getRedisConnectionInfoOutput"], () => require("./getRedisConnectionInfo"));

export { ListRedisArgs, ListRedisResult } from "./listRedis";
export const listRedis: typeof import("./listRedis").listRedis = null as any;
export const listRedisOutput: typeof import("./listRedis").listRedisOutput = null as any;
utilities.lazyLoad(exports, ["listRedis","listRedisOutput"], () => require("./listRedis"));

export { RediArgs } from "./redi";
export type Redi = import("./redi").Redi;
export const Redi: typeof import("./redi").Redi = null as any;
utilities.lazyLoad(exports, ["Redi"], () => require("./redi"));


// Export enums:
export * from "../types/enums/redis";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "render:redis:Redi":
                return new Redi(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("render", "redis", _module)
