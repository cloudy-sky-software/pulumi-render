// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetRegistryCredentialArgs, GetRegistryCredentialOutputArgs } from "./getRegistryCredential";
export const getRegistryCredential: typeof import("./getRegistryCredential").getRegistryCredential = null as any;
export const getRegistryCredentialOutput: typeof import("./getRegistryCredential").getRegistryCredentialOutput = null as any;
utilities.lazyLoad(exports, ["getRegistryCredential","getRegistryCredentialOutput"], () => require("./getRegistryCredential"));

export { ListRegistryCredentialsArgs, ListRegistryCredentialsResult } from "./listRegistryCredentials";
export const listRegistryCredentials: typeof import("./listRegistryCredentials").listRegistryCredentials = null as any;
export const listRegistryCredentialsOutput: typeof import("./listRegistryCredentials").listRegistryCredentialsOutput = null as any;
utilities.lazyLoad(exports, ["listRegistryCredentials","listRegistryCredentialsOutput"], () => require("./listRegistryCredentials"));

export { RegistryCredentialArgs } from "./registryCredential";
export type RegistryCredential = import("./registryCredential").RegistryCredential;
export const RegistryCredential: typeof import("./registryCredential").RegistryCredential = null as any;
utilities.lazyLoad(exports, ["RegistryCredential"], () => require("./registryCredential"));


// Export enums:
export * from "../types/enums/registrycredentials";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "render:registrycredentials:RegistryCredential":
                return new RegistryCredential(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("render", "registrycredentials", _module)
