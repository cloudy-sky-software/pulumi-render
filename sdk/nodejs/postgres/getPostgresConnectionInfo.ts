// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getPostgresConnectionInfo(args: GetPostgresConnectionInfoArgs, opts?: pulumi.InvokeOptions): Promise<outputs.postgres.PostgresConnectionInfo> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("render:postgres:getPostgresConnectionInfo", {
        "postgresId": args.postgresId,
    }, opts);
}

export interface GetPostgresConnectionInfoArgs {
    postgresId: string;
}
export function getPostgresConnectionInfoOutput(args: GetPostgresConnectionInfoOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.postgres.PostgresConnectionInfo> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("render:postgres:getPostgresConnectionInfo", {
        "postgresId": args.postgresId,
    }, opts);
}

export interface GetPostgresConnectionInfoOutputArgs {
    postgresId: pulumi.Input<string>;
}
