import * as pulumi from "@pulumi/pulumi";
import * as render from "@cloudyskysoftware/pulumi-render";

import type { services as servicesInputs } from "@cloudyskysoftware/pulumi-render/types/input";

const ownerId = pulumi
    .output(render.owners.listOwners())
    .apply(
        (result) =>
            result.items.filter(
                (i) => i.owner?.email === "pl@cloudysky.software"
            )[0].owner?.id || ""
    );

const staticSiteDetails: servicesInputs.StaticSiteDetailsCreateArgs = {
    publishPath: "public",
};

const staticSite = new render.services.StaticSite("staticsite", {
    name: "My custom static site",
    ownerId,
    repo: "https://github.com/cloudy-sky-software/test-static-site",
    autoDeploy: "no",
    branch: "main",
    serviceDetails: staticSiteDetails,
    type: "static_site",
});

const port = "8080";

const webServiceDetails: servicesInputs.WebServiceDetailsCreateArgs = {
    env: "node",
    plan: "starter",
    region: "oregon",
    envSpecificDetails: {
        buildCommand: "yarn",
        startCommand: "node app.js",
    },
    runtime: "node",
};

const db = new render.postgres.Postgres("db", {
    ownerId,
    version: "16",
    databaseUser: "myuser",
    databaseName: "test",
    plan: "basic_256mb",
    diskSizeGB: 1,
});

export const connectionInfo = render.postgres.getPostgresConnectionInfoOutput({
    postgresId: db.id,
});

const webService = new render.services.WebService("webservice", {
    name: "An Express.js web service",
    ownerId,
    repo: "https://github.com/render-examples/express-hello-world",
    autoDeploy: "yes",
    branch: "master",
    serviceDetails: webServiceDetails,
    type: "web_service",
});

new render.services.EnvVarsForService("webServiceEnvVars", {
    serviceId: webService.id,
    envVars: [
        {
            key: "PORT",
            value: port,
        },
        {
            key: "DB_URL",
            value: connectionInfo.internalConnectionString,
        },
    ],
});

export const url = staticSite.serviceDetails.apply((s) => s?.url);

export const webServiceUrl = webService.serviceDetails.apply((s) => s?.url);
