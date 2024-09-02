// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const Registry = {
    Github: "GITHUB",
    Gitlab: "GITLAB",
    Docker: "DOCKER",
    GoogleArtifact: "GOOGLE_ARTIFACT",
} as const;

/**
 * The registry to use this credential with
 */
export type Registry = (typeof Registry)[keyof typeof Registry];

export const RegistryCredentialRegistry = {
    Github: "GITHUB",
    Gitlab: "GITLAB",
    Docker: "DOCKER",
    GoogleArtifact: "GOOGLE_ARTIFACT",
} as const;

/**
 * The registry to use this credential with
 */
export type RegistryCredentialRegistry = (typeof RegistryCredentialRegistry)[keyof typeof RegistryCredentialRegistry];
