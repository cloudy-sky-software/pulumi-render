// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const BackgroundWorkerServiceDetailsEnv = {
    Docker: "docker",
    Elixir: "elixir",
    Go: "go",
    Node: "node",
    Python: "python",
    Ruby: "ruby",
    Rust: "rust",
} as const;

export type BackgroundWorkerServiceDetailsEnv = (typeof BackgroundWorkerServiceDetailsEnv)[keyof typeof BackgroundWorkerServiceDetailsEnv];

export const BackgroundWorkerServiceDetailsPlan = {
    Starter: "starter",
    StarterPlus: "starter_plus",
    Standard: "standard",
    StandardPlus: "standard_plus",
    Pro: "pro",
    ProPlus: "pro_plus",
    ProMax: "pro_max",
    ProUltra: "pro_ultra",
} as const;

export type BackgroundWorkerServiceDetailsPlan = (typeof BackgroundWorkerServiceDetailsPlan)[keyof typeof BackgroundWorkerServiceDetailsPlan];

export const BackgroundWorkerServiceDetailsPullRequestPreviewsEnabled = {
    Yes: "yes",
    No: "no",
} as const;

export type BackgroundWorkerServiceDetailsPullRequestPreviewsEnabled = (typeof BackgroundWorkerServiceDetailsPullRequestPreviewsEnabled)[keyof typeof BackgroundWorkerServiceDetailsPullRequestPreviewsEnabled];

export const BackgroundWorkerServiceDetailsRegion = {
    Oregon: "oregon",
    Frankfurt: "frankfurt",
} as const;

export type BackgroundWorkerServiceDetailsRegion = (typeof BackgroundWorkerServiceDetailsRegion)[keyof typeof BackgroundWorkerServiceDetailsRegion];

export const ClearCache = {
    DoNotClear: "do_not_clear",
    Clear: "clear",
} as const;

export type ClearCache = (typeof ClearCache)[keyof typeof ClearCache];

export const CronJobServiceDetailsEnv = {
    Docker: "docker",
    Elixir: "elixir",
    Go: "go",
    Node: "node",
    Python: "python",
    Ruby: "ruby",
    Rust: "rust",
} as const;

export type CronJobServiceDetailsEnv = (typeof CronJobServiceDetailsEnv)[keyof typeof CronJobServiceDetailsEnv];

export const CronJobServiceDetailsPlan = {
    Starter: "starter",
    StarterPlus: "starter_plus",
    Standard: "standard",
    StandardPlus: "standard_plus",
    Pro: "pro",
    ProPlus: "pro_plus",
    ProMax: "pro_max",
    ProUltra: "pro_ultra",
} as const;

export type CronJobServiceDetailsPlan = (typeof CronJobServiceDetailsPlan)[keyof typeof CronJobServiceDetailsPlan];

export const CronJobServiceDetailsRegion = {
    Oregon: "oregon",
    Frankfurt: "frankfurt",
} as const;

export type CronJobServiceDetailsRegion = (typeof CronJobServiceDetailsRegion)[keyof typeof CronJobServiceDetailsRegion];

export const CustomDomainDomainType = {
    Apex: "apex",
    Subdomain: "subdomain",
} as const;

export type CustomDomainDomainType = (typeof CustomDomainDomainType)[keyof typeof CustomDomainDomainType];

export const CustomDomainVerificationStatus = {
    Verified: "verified",
    Unverified: "unverified",
} as const;

export type CustomDomainVerificationStatus = (typeof CustomDomainVerificationStatus)[keyof typeof CustomDomainVerificationStatus];

export const DeployClearCache = {
    DoNotClear: "do_not_clear",
    Clear: "clear",
} as const;

export type DeployClearCache = (typeof DeployClearCache)[keyof typeof DeployClearCache];

export const DomainType = {
    Apex: "apex",
    Subdomain: "subdomain",
} as const;

export type DomainType = (typeof DomainType)[keyof typeof DomainType];

export const EnvVarKeyValueGenerateValue = {
    Yes: "yes",
} as const;

export type EnvVarKeyValueGenerateValue = (typeof EnvVarKeyValueGenerateValue)[keyof typeof EnvVarKeyValueGenerateValue];

export const OpenPortsProtocol = {
    Tcp: "TCP",
    Udp: "UDP",
} as const;

export type OpenPortsProtocol = (typeof OpenPortsProtocol)[keyof typeof OpenPortsProtocol];

export const PrivateServiceDetailsEnv = {
    Docker: "docker",
    Elixir: "elixir",
    Go: "go",
    Node: "node",
    Python: "python",
    Ruby: "ruby",
    Rust: "rust",
} as const;

export type PrivateServiceDetailsEnv = (typeof PrivateServiceDetailsEnv)[keyof typeof PrivateServiceDetailsEnv];

export const PrivateServiceDetailsPlan = {
    Starter: "starter",
    StarterPlus: "starter_plus",
    Standard: "standard",
    StandardPlus: "standard_plus",
    Pro: "pro",
    ProPlus: "pro_plus",
    ProMax: "pro_max",
    ProUltra: "pro_ultra",
} as const;

export type PrivateServiceDetailsPlan = (typeof PrivateServiceDetailsPlan)[keyof typeof PrivateServiceDetailsPlan];

export const PrivateServiceDetailsPullRequestPreviewsEnabled = {
    Yes: "yes",
    No: "no",
} as const;

export type PrivateServiceDetailsPullRequestPreviewsEnabled = (typeof PrivateServiceDetailsPullRequestPreviewsEnabled)[keyof typeof PrivateServiceDetailsPullRequestPreviewsEnabled];

export const PrivateServiceDetailsRegion = {
    Oregon: "oregon",
    Frankfurt: "frankfurt",
} as const;

export type PrivateServiceDetailsRegion = (typeof PrivateServiceDetailsRegion)[keyof typeof PrivateServiceDetailsRegion];

export const ServiceAutoDeploy = {
    Yes: "yes",
    No: "no",
} as const;

/**
 * Whether to auto deploy the service or not upon git push.
 */
export type ServiceAutoDeploy = (typeof ServiceAutoDeploy)[keyof typeof ServiceAutoDeploy];

export const ServiceNotifyOnFail = {
    Default: "default",
    Notify: "notify",
    Ignore: "ignore",
} as const;

/**
 * The notification setting for this service upon deployment failure.
 */
export type ServiceNotifyOnFail = (typeof ServiceNotifyOnFail)[keyof typeof ServiceNotifyOnFail];

export const ServiceSuspended = {
    Suspended: "suspended",
    NotSuspended: "not_suspended",
} as const;

export type ServiceSuspended = (typeof ServiceSuspended)[keyof typeof ServiceSuspended];

export const StaticSiteRouteType = {
    Redirect: "redirect",
    Rewrite: "rewrite",
} as const;

export type StaticSiteRouteType = (typeof StaticSiteRouteType)[keyof typeof StaticSiteRouteType];

export const StaticSiteServiceDetailsPullRequestPreviewsEnabled = {
    Yes: "yes",
    No: "no",
} as const;

export type StaticSiteServiceDetailsPullRequestPreviewsEnabled = (typeof StaticSiteServiceDetailsPullRequestPreviewsEnabled)[keyof typeof StaticSiteServiceDetailsPullRequestPreviewsEnabled];

export const VerificationStatus = {
    Verified: "verified",
    Unverified: "unverified",
} as const;

export type VerificationStatus = (typeof VerificationStatus)[keyof typeof VerificationStatus];

export const WebServiceServiceDetailsEnv = {
    Docker: "docker",
    Elixir: "elixir",
    Go: "go",
    Node: "node",
    Python: "python",
    Ruby: "ruby",
    Rust: "rust",
} as const;

export type WebServiceServiceDetailsEnv = (typeof WebServiceServiceDetailsEnv)[keyof typeof WebServiceServiceDetailsEnv];

export const WebServiceServiceDetailsPlan = {
    Starter: "starter",
    StarterPlus: "starter_plus",
    Standard: "standard",
    StandardPlus: "standard_plus",
    Pro: "pro",
    ProPlus: "pro_plus",
    ProMax: "pro_max",
    ProUltra: "pro_ultra",
} as const;

export type WebServiceServiceDetailsPlan = (typeof WebServiceServiceDetailsPlan)[keyof typeof WebServiceServiceDetailsPlan];

export const WebServiceServiceDetailsPullRequestPreviewsEnabled = {
    Yes: "yes",
    No: "no",
} as const;

export type WebServiceServiceDetailsPullRequestPreviewsEnabled = (typeof WebServiceServiceDetailsPullRequestPreviewsEnabled)[keyof typeof WebServiceServiceDetailsPullRequestPreviewsEnabled];

export const WebServiceServiceDetailsRegion = {
    Oregon: "oregon",
    Frankfurt: "frankfurt",
} as const;

export type WebServiceServiceDetailsRegion = (typeof WebServiceServiceDetailsRegion)[keyof typeof WebServiceServiceDetailsRegion];
