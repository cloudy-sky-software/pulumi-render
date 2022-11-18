// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export interface ListOwnersResponse {
    cursor?: string;
    /**
     * The owner object represents an authorized user or team. The `type` property indicates if the owner is a user or team.
     */
    owner?: outputs.owners.Owner;
}

/**
 * The owner object represents an authorized user or team. The `type` property indicates if the owner is a user or team.
 */
export interface Owner {
    /**
     * The email of the owner.
     */
    email?: string;
    /**
     * The owner ID.
     */
    id?: string;
    /**
     * The name of the owner.
     */
    name?: string;
    /**
     * The type of the authorized user.
     */
    type?: enums.owners.OwnerType;
}

