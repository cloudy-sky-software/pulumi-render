// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetRedisArgs, GetRedisOutputArgs } from "./getRedis";
export const getRedis: typeof import("./getRedis").getRedis = null as any;
export const getRedisOutput: typeof import("./getRedis").getRedisOutput = null as any;
utilities.lazyLoad(exports, ["getRedis","getRedisOutput"], () => require("./getRedis"));

export { GetRedisConnectionInfoArgs, GetRedisConnectionInfoOutputArgs } from "./getRedisConnectionInfo";
export const getRedisConnectionInfo: typeof import("./getRedisConnectionInfo").getRedisConnectionInfo = null as any;
export const getRedisConnectionInfoOutput: typeof import("./getRedisConnectionInfo").getRedisConnectionInfoOutput = null as any;
utilities.lazyLoad(exports, ["getRedisConnectionInfo","getRedisConnectionInfoOutput"], () => require("./getRedisConnectionInfo"));

export { ListRedisArgs, ListRedisResult } from "./listRedis";
export const listRedis: typeof import("./listRedis").listRedis = null as any;
export const listRedisOutput: typeof import("./listRedis").listRedisOutput = null as any;
utilities.lazyLoad(exports, ["listRedis","listRedisOutput"], () => require("./listRedis"));

export { RedisArgs } from "./redis";
export type Redis = import("./redis").Redis;
export const Redis: typeof import("./redis").Redis = null as any;
utilities.lazyLoad(exports, ["Redis"], () => require("./redis"));


// Export enums:
export * from "../types/enums/redis";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "render:redis:Redis":
                return new Redis(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("render", "redis", _module)
