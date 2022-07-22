// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./customDomain";
export * from "./deploy";
export * from "./scale";
export * from "./service";
export * from "./suspend";

// Export enums:
export * from "../types/enums/services";

// Import resources to register:
import { CustomDomain } from "./customDomain";
import { Deploy } from "./deploy";
import { Scale } from "./scale";
import { Service } from "./service";
import { Suspend } from "./suspend";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "render:services:CustomDomain":
                return new CustomDomain(name, <any>undefined, { urn })
            case "render:services:Deploy":
                return new Deploy(name, <any>undefined, { urn })
            case "render:services:Scale":
                return new Scale(name, <any>undefined, { urn })
            case "render:services:Service":
                return new Service(name, <any>undefined, { urn })
            case "render:services:Suspend":
                return new Suspend(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("render", "services", _module)