// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetWebhookArgs, GetWebhookResult, GetWebhookOutputArgs } from "./getWebhook";
export const getWebhook: typeof import("./getWebhook").getWebhook = null as any;
export const getWebhookOutput: typeof import("./getWebhook").getWebhookOutput = null as any;
utilities.lazyLoad(exports, ["getWebhook","getWebhookOutput"], () => require("./getWebhook"));

export { ListWebhooksArgs, ListWebhooksResult } from "./listWebhooks";
export const listWebhooks: typeof import("./listWebhooks").listWebhooks = null as any;
export const listWebhooksOutput: typeof import("./listWebhooks").listWebhooksOutput = null as any;
utilities.lazyLoad(exports, ["listWebhooks","listWebhooksOutput"], () => require("./listWebhooks"));

export { WebhookArgs } from "./webhook";
export type Webhook = import("./webhook").Webhook;
export const Webhook: typeof import("./webhook").Webhook = null as any;
utilities.lazyLoad(exports, ["Webhook"], () => require("./webhook"));


// Export enums:
export * from "../types/enums/webhooks";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "render:webhooks:Webhook":
                return new Webhook(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("render", "webhooks", _module)
