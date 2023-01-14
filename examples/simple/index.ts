import * as pulumi from "@pulumi/pulumi";
import * as render from "@cloudyskysoftware/pulumi-render";

import { services } from "@cloudyskysoftware/pulumi-render/types/input";

const ownerId = pulumi
    .output(render.owners.listOwners())
    .apply(
        (result) =>
            result.items.filter(
                (i) => i.owner?.email === "pl@cloudysky.software"
            )[0].owner?.id || ""
    );

const staticSiteDetails: services.StaticSiteServiceDetailsArgs = {
    publishPath: "public",
};

const staticSite = new render.services.StaticSite("staticsite", {
    name: "My custom static site",
    ownerId,
    repo: "https://github.com/cloudy-sky-software/test-static-site",
    autoDeploy: "no",
    branch: "main",
    serviceDetails: staticSiteDetails,
});

const port = 8080;

const webServiceDetails: services.WebServiceServiceDetailsArgs = {
    env: "node",
    plan: "starter",
    region: "oregon",
    openPorts: [
        {
            port,
            protocol: "TCP",
        },
    ],
    envSpecificDetails: {
        buildCommand: "yarn",
        startCommand: "node app.js",
    },
};

const webService = new render.services.WebService("webservice", {
    name: "An Express.js web service",
    ownerId,
    repo: "https://github.com/render-examples/express-hello-world",
    autoDeploy: "yes",
    envVars: [
        {
            key: "PORT",
            value: `${port}`,
        },
    ],
    branch: "master",
    serviceDetails: webServiceDetails,
});

export const url = (
    staticSite.serviceDetails as services.StaticSiteServiceDetailsArgs
).url;

export const webServiceUrl = (
    webService.serviceDetails as services.WebServiceServiceDetailsArgs
).url;
