// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const BackgroundWorkerDetailsCreateEnv = {
    Docker: "docker",
    Elixir: "elixir",
    Go: "go",
    Node: "node",
    Python: "python",
    Ruby: "ruby",
    Rust: "rust",
    Image: "image",
} as const;

/**
 * Environment (runtime)
 */
export type BackgroundWorkerDetailsCreateEnv = (typeof BackgroundWorkerDetailsCreateEnv)[keyof typeof BackgroundWorkerDetailsCreateEnv];

export const BackgroundWorkerDetailsCreatePlan = {
    Starter: "starter",
    Standard: "standard",
    Pro: "pro",
    ProPlus: "pro_plus",
    ProMax: "pro_max",
    ProUltra: "pro_ultra",
} as const;

export type BackgroundWorkerDetailsCreatePlan = (typeof BackgroundWorkerDetailsCreatePlan)[keyof typeof BackgroundWorkerDetailsCreatePlan];

export const BackgroundWorkerDetailsCreatePullRequestPreviewsEnabled = {
    Yes: "yes",
    No: "no",
} as const;

/**
 * Defaults to "no"
 */
export type BackgroundWorkerDetailsCreatePullRequestPreviewsEnabled = (typeof BackgroundWorkerDetailsCreatePullRequestPreviewsEnabled)[keyof typeof BackgroundWorkerDetailsCreatePullRequestPreviewsEnabled];

export const BackgroundWorkerDetailsCreateRegion = {
    Frankfurt: "frankfurt",
    Oregon: "oregon",
    Ohio: "ohio",
    Singapore: "singapore",
} as const;

export type BackgroundWorkerDetailsCreateRegion = (typeof BackgroundWorkerDetailsCreateRegion)[keyof typeof BackgroundWorkerDetailsCreateRegion];

export const BackgroundWorkerDetailsEnv = {
    Docker: "docker",
    Elixir: "elixir",
    Go: "go",
    Node: "node",
    Python: "python",
    Ruby: "ruby",
    Rust: "rust",
    Image: "image",
} as const;

/**
 * Environment (runtime)
 */
export type BackgroundWorkerDetailsEnv = (typeof BackgroundWorkerDetailsEnv)[keyof typeof BackgroundWorkerDetailsEnv];

export const BackgroundWorkerDetailsPlan = {
    Starter: "starter",
    StarterPlus: "starter_plus",
    Standard: "standard",
    StandardPlus: "standard_plus",
    Pro: "pro",
    ProPlus: "pro_plus",
    ProMax: "pro_max",
    ProUltra: "pro_ultra",
    Free: "free",
    Custom: "custom",
} as const;

/**
 * The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
 */
export type BackgroundWorkerDetailsPlan = (typeof BackgroundWorkerDetailsPlan)[keyof typeof BackgroundWorkerDetailsPlan];

export const BackgroundWorkerDetailsPullRequestPreviewsEnabled = {
    Yes: "yes",
    No: "no",
} as const;

export type BackgroundWorkerDetailsPullRequestPreviewsEnabled = (typeof BackgroundWorkerDetailsPullRequestPreviewsEnabled)[keyof typeof BackgroundWorkerDetailsPullRequestPreviewsEnabled];

export const BackgroundWorkerDetailsRegion = {
    Frankfurt: "frankfurt",
    Oregon: "oregon",
    Ohio: "ohio",
    Singapore: "singapore",
} as const;

export type BackgroundWorkerDetailsRegion = (typeof BackgroundWorkerDetailsRegion)[keyof typeof BackgroundWorkerDetailsRegion];

export const BackgroundWorkerServiceAutoDeploy = {
    Yes: "yes",
    No: "no",
} as const;

export type BackgroundWorkerServiceAutoDeploy = (typeof BackgroundWorkerServiceAutoDeploy)[keyof typeof BackgroundWorkerServiceAutoDeploy];

export const BackgroundWorkerServiceNotifyOnFail = {
    Default: "default",
    Notify: "notify",
    Ignore: "ignore",
} as const;

export type BackgroundWorkerServiceNotifyOnFail = (typeof BackgroundWorkerServiceNotifyOnFail)[keyof typeof BackgroundWorkerServiceNotifyOnFail];

export const BackgroundWorkerServiceSuspended = {
    Suspended: "suspended",
    NotSuspended: "not_suspended",
} as const;

export type BackgroundWorkerServiceSuspended = (typeof BackgroundWorkerServiceSuspended)[keyof typeof BackgroundWorkerServiceSuspended];

export const BackgroundWorkerServiceSuspendersItem = {
    Admin: "admin",
    Billing: "billing",
    User: "user",
    ParentService: "parent_service",
    Unknown: "unknown",
} as const;

export type BackgroundWorkerServiceSuspendersItem = (typeof BackgroundWorkerServiceSuspendersItem)[keyof typeof BackgroundWorkerServiceSuspendersItem];

export const CancelDeployStatus = {
    Created: "created",
    BuildInProgress: "build_in_progress",
    UpdateInProgress: "update_in_progress",
    Live: "live",
    Deactivated: "deactivated",
    BuildFailed: "build_failed",
    UpdateFailed: "update_failed",
    Canceled: "canceled",
    PreDeployInProgress: "pre_deploy_in_progress",
    PreDeployFailed: "pre_deploy_failed",
} as const;

export type CancelDeployStatus = (typeof CancelDeployStatus)[keyof typeof CancelDeployStatus];

export const CancelDeployTrigger = {
    Api: "api",
    BlueprintSync: "blueprint_sync",
    DeployHook: "deploy_hook",
    DeployedByRender: "deployed_by_render",
    Manual: "manual",
    Other: "other",
    NewCommit: "new_commit",
    Rollback: "rollback",
    ServiceResumed: "service_resumed",
    ServiceUpdated: "service_updated",
} as const;

export type CancelDeployTrigger = (typeof CancelDeployTrigger)[keyof typeof CancelDeployTrigger];

export const CronJobDetailsCreateEnv = {
    Docker: "docker",
    Elixir: "elixir",
    Go: "go",
    Node: "node",
    Python: "python",
    Ruby: "ruby",
    Rust: "rust",
    Image: "image",
} as const;

/**
 * Environment (runtime)
 */
export type CronJobDetailsCreateEnv = (typeof CronJobDetailsCreateEnv)[keyof typeof CronJobDetailsCreateEnv];

export const CronJobDetailsCreatePlan = {
    Starter: "starter",
    Standard: "standard",
    Pro: "pro",
    ProPlus: "pro_plus",
    ProMax: "pro_max",
    ProUltra: "pro_ultra",
} as const;

export type CronJobDetailsCreatePlan = (typeof CronJobDetailsCreatePlan)[keyof typeof CronJobDetailsCreatePlan];

export const CronJobDetailsCreateRegion = {
    Frankfurt: "frankfurt",
    Oregon: "oregon",
    Ohio: "ohio",
    Singapore: "singapore",
} as const;

export type CronJobDetailsCreateRegion = (typeof CronJobDetailsCreateRegion)[keyof typeof CronJobDetailsCreateRegion];

export const CronJobDetailsEnv = {
    Docker: "docker",
    Elixir: "elixir",
    Go: "go",
    Node: "node",
    Python: "python",
    Ruby: "ruby",
    Rust: "rust",
    Image: "image",
} as const;

/**
 * Environment (runtime)
 */
export type CronJobDetailsEnv = (typeof CronJobDetailsEnv)[keyof typeof CronJobDetailsEnv];

export const CronJobDetailsPlan = {
    Starter: "starter",
    StarterPlus: "starter_plus",
    Standard: "standard",
    StandardPlus: "standard_plus",
    Pro: "pro",
    ProPlus: "pro_plus",
    ProMax: "pro_max",
    ProUltra: "pro_ultra",
    Free: "free",
    Custom: "custom",
} as const;

/**
 * The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
 */
export type CronJobDetailsPlan = (typeof CronJobDetailsPlan)[keyof typeof CronJobDetailsPlan];

export const CronJobDetailsRegion = {
    Frankfurt: "frankfurt",
    Oregon: "oregon",
    Ohio: "ohio",
    Singapore: "singapore",
} as const;

export type CronJobDetailsRegion = (typeof CronJobDetailsRegion)[keyof typeof CronJobDetailsRegion];

export const CronJobServiceAutoDeploy = {
    Yes: "yes",
    No: "no",
} as const;

export type CronJobServiceAutoDeploy = (typeof CronJobServiceAutoDeploy)[keyof typeof CronJobServiceAutoDeploy];

export const CronJobServiceNotifyOnFail = {
    Default: "default",
    Notify: "notify",
    Ignore: "ignore",
} as const;

export type CronJobServiceNotifyOnFail = (typeof CronJobServiceNotifyOnFail)[keyof typeof CronJobServiceNotifyOnFail];

export const CronJobServiceSuspended = {
    Suspended: "suspended",
    NotSuspended: "not_suspended",
} as const;

export type CronJobServiceSuspended = (typeof CronJobServiceSuspended)[keyof typeof CronJobServiceSuspended];

export const CronJobServiceSuspendersItem = {
    Admin: "admin",
    Billing: "billing",
    User: "user",
    ParentService: "parent_service",
    Unknown: "unknown",
} as const;

export type CronJobServiceSuspendersItem = (typeof CronJobServiceSuspendersItem)[keyof typeof CronJobServiceSuspendersItem];

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
    Clear: "clear",
    DoNotClear: "do_not_clear",
} as const;

/**
 * Defaults to "do_not_clear"
 */
export type DeployClearCache = (typeof DeployClearCache)[keyof typeof DeployClearCache];

export const DeployStatus = {
    Created: "created",
    BuildInProgress: "build_in_progress",
    UpdateInProgress: "update_in_progress",
    Live: "live",
    Deactivated: "deactivated",
    BuildFailed: "build_failed",
    UpdateFailed: "update_failed",
    Canceled: "canceled",
    PreDeployInProgress: "pre_deploy_in_progress",
    PreDeployFailed: "pre_deploy_failed",
} as const;

export type DeployStatus = (typeof DeployStatus)[keyof typeof DeployStatus];

export const DeployTrigger = {
    Api: "api",
    BlueprintSync: "blueprint_sync",
    DeployHook: "deploy_hook",
    DeployedByRender: "deployed_by_render",
    Manual: "manual",
    Other: "other",
    NewCommit: "new_commit",
    Rollback: "rollback",
    ServiceResumed: "service_resumed",
    ServiceUpdated: "service_updated",
} as const;

export type DeployTrigger = (typeof DeployTrigger)[keyof typeof DeployTrigger];

export const PreviewServicePlan = {
    Starter: "starter",
    StarterPlus: "starter_plus",
    Standard: "standard",
    StandardPlus: "standard_plus",
    Pro: "pro",
    ProPlus: "pro_plus",
    ProMax: "pro_max",
    ProUltra: "pro_ultra",
    Free: "free",
    Custom: "custom",
} as const;

/**
 * The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
 */
export type PreviewServicePlan = (typeof PreviewServicePlan)[keyof typeof PreviewServicePlan];

export const PreviewServiceServiceAutoDeploy = {
    Yes: "yes",
    No: "no",
} as const;

export type PreviewServiceServiceAutoDeploy = (typeof PreviewServiceServiceAutoDeploy)[keyof typeof PreviewServiceServiceAutoDeploy];

export const PreviewServiceServiceNotifyOnFail = {
    Default: "default",
    Notify: "notify",
    Ignore: "ignore",
} as const;

export type PreviewServiceServiceNotifyOnFail = (typeof PreviewServiceServiceNotifyOnFail)[keyof typeof PreviewServiceServiceNotifyOnFail];

export const PreviewServiceServiceSuspended = {
    Suspended: "suspended",
    NotSuspended: "not_suspended",
} as const;

export type PreviewServiceServiceSuspended = (typeof PreviewServiceServiceSuspended)[keyof typeof PreviewServiceServiceSuspended];

export const PreviewServiceServiceSuspendersItem = {
    Admin: "admin",
    Billing: "billing",
    User: "user",
    ParentService: "parent_service",
    Unknown: "unknown",
} as const;

export type PreviewServiceServiceSuspendersItem = (typeof PreviewServiceServiceSuspendersItem)[keyof typeof PreviewServiceServiceSuspendersItem];

export const PrivateServiceDetailsCreateEnv = {
    Docker: "docker",
    Elixir: "elixir",
    Go: "go",
    Node: "node",
    Python: "python",
    Ruby: "ruby",
    Rust: "rust",
    Image: "image",
} as const;

/**
 * Environment (runtime)
 */
export type PrivateServiceDetailsCreateEnv = (typeof PrivateServiceDetailsCreateEnv)[keyof typeof PrivateServiceDetailsCreateEnv];

export const PrivateServiceDetailsCreatePlan = {
    Starter: "starter",
    Standard: "standard",
    Pro: "pro",
    ProPlus: "pro_plus",
    ProMax: "pro_max",
    ProUltra: "pro_ultra",
} as const;

export type PrivateServiceDetailsCreatePlan = (typeof PrivateServiceDetailsCreatePlan)[keyof typeof PrivateServiceDetailsCreatePlan];

export const PrivateServiceDetailsCreatePullRequestPreviewsEnabled = {
    Yes: "yes",
    No: "no",
} as const;

/**
 * Defaults to "no"
 */
export type PrivateServiceDetailsCreatePullRequestPreviewsEnabled = (typeof PrivateServiceDetailsCreatePullRequestPreviewsEnabled)[keyof typeof PrivateServiceDetailsCreatePullRequestPreviewsEnabled];

export const PrivateServiceDetailsCreateRegion = {
    Frankfurt: "frankfurt",
    Oregon: "oregon",
    Ohio: "ohio",
    Singapore: "singapore",
} as const;

export type PrivateServiceDetailsCreateRegion = (typeof PrivateServiceDetailsCreateRegion)[keyof typeof PrivateServiceDetailsCreateRegion];

export const PrivateServiceDetailsEnv = {
    Docker: "docker",
    Elixir: "elixir",
    Go: "go",
    Node: "node",
    Python: "python",
    Ruby: "ruby",
    Rust: "rust",
    Image: "image",
} as const;

/**
 * Environment (runtime)
 */
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
    Free: "free",
    Custom: "custom",
} as const;

/**
 * The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
 */
export type PrivateServiceDetailsPlan = (typeof PrivateServiceDetailsPlan)[keyof typeof PrivateServiceDetailsPlan];

export const PrivateServiceDetailsPullRequestPreviewsEnabled = {
    Yes: "yes",
    No: "no",
} as const;

export type PrivateServiceDetailsPullRequestPreviewsEnabled = (typeof PrivateServiceDetailsPullRequestPreviewsEnabled)[keyof typeof PrivateServiceDetailsPullRequestPreviewsEnabled];

export const PrivateServiceDetailsRegion = {
    Frankfurt: "frankfurt",
    Oregon: "oregon",
    Ohio: "ohio",
    Singapore: "singapore",
} as const;

export type PrivateServiceDetailsRegion = (typeof PrivateServiceDetailsRegion)[keyof typeof PrivateServiceDetailsRegion];

export const PrivateServiceServiceAutoDeploy = {
    Yes: "yes",
    No: "no",
} as const;

export type PrivateServiceServiceAutoDeploy = (typeof PrivateServiceServiceAutoDeploy)[keyof typeof PrivateServiceServiceAutoDeploy];

export const PrivateServiceServiceNotifyOnFail = {
    Default: "default",
    Notify: "notify",
    Ignore: "ignore",
} as const;

export type PrivateServiceServiceNotifyOnFail = (typeof PrivateServiceServiceNotifyOnFail)[keyof typeof PrivateServiceServiceNotifyOnFail];

export const PrivateServiceServiceSuspended = {
    Suspended: "suspended",
    NotSuspended: "not_suspended",
} as const;

export type PrivateServiceServiceSuspended = (typeof PrivateServiceServiceSuspended)[keyof typeof PrivateServiceServiceSuspended];

export const PrivateServiceServiceSuspendersItem = {
    Admin: "admin",
    Billing: "billing",
    User: "user",
    ParentService: "parent_service",
    Unknown: "unknown",
} as const;

export type PrivateServiceServiceSuspendersItem = (typeof PrivateServiceServiceSuspendersItem)[keyof typeof PrivateServiceServiceSuspendersItem];

export const RegistryCredentialRegistry = {
    Github: "GITHUB",
    Gitlab: "GITLAB",
    Docker: "DOCKER",
} as const;

/**
 * The registry to use this credential with
 */
export type RegistryCredentialRegistry = (typeof RegistryCredentialRegistry)[keyof typeof RegistryCredentialRegistry];

export const RollbackDeployStatus = {
    Created: "created",
    BuildInProgress: "build_in_progress",
    UpdateInProgress: "update_in_progress",
    Live: "live",
    Deactivated: "deactivated",
    BuildFailed: "build_failed",
    UpdateFailed: "update_failed",
    Canceled: "canceled",
    PreDeployInProgress: "pre_deploy_in_progress",
    PreDeployFailed: "pre_deploy_failed",
} as const;

export type RollbackDeployStatus = (typeof RollbackDeployStatus)[keyof typeof RollbackDeployStatus];

export const RollbackDeployTrigger = {
    Api: "api",
    BlueprintSync: "blueprint_sync",
    DeployHook: "deploy_hook",
    DeployedByRender: "deployed_by_render",
    Manual: "manual",
    Other: "other",
    NewCommit: "new_commit",
    Rollback: "rollback",
    ServiceResumed: "service_resumed",
    ServiceUpdated: "service_updated",
} as const;

export type RollbackDeployTrigger = (typeof RollbackDeployTrigger)[keyof typeof RollbackDeployTrigger];

export const ServerPortProtocol = {
    Tcp: "TCP",
    Udp: "UDP",
} as const;

export type ServerPortProtocol = (typeof ServerPortProtocol)[keyof typeof ServerPortProtocol];

export const ServiceAutoDeploy = {
    Yes: "yes",
    No: "no",
} as const;

export type ServiceAutoDeploy = (typeof ServiceAutoDeploy)[keyof typeof ServiceAutoDeploy];

export const ServiceNotifyOnFail = {
    Default: "default",
    Notify: "notify",
    Ignore: "ignore",
} as const;

export type ServiceNotifyOnFail = (typeof ServiceNotifyOnFail)[keyof typeof ServiceNotifyOnFail];

export const ServiceSuspended = {
    Suspended: "suspended",
    NotSuspended: "not_suspended",
} as const;

export type ServiceSuspended = (typeof ServiceSuspended)[keyof typeof ServiceSuspended];

export const ServiceSuspendersItem = {
    Admin: "admin",
    Billing: "billing",
    User: "user",
    ParentService: "parent_service",
    Unknown: "unknown",
} as const;

export type ServiceSuspendersItem = (typeof ServiceSuspendersItem)[keyof typeof ServiceSuspendersItem];

export const StaticSiteDetailsCreatePullRequestPreviewsEnabled = {
    Yes: "yes",
    No: "no",
} as const;

/**
 * Defaults to "no"
 */
export type StaticSiteDetailsCreatePullRequestPreviewsEnabled = (typeof StaticSiteDetailsCreatePullRequestPreviewsEnabled)[keyof typeof StaticSiteDetailsCreatePullRequestPreviewsEnabled];

export const StaticSiteDetailsPullRequestPreviewsEnabled = {
    Yes: "yes",
    No: "no",
} as const;

export type StaticSiteDetailsPullRequestPreviewsEnabled = (typeof StaticSiteDetailsPullRequestPreviewsEnabled)[keyof typeof StaticSiteDetailsPullRequestPreviewsEnabled];

export const StaticSiteRouteType = {
    Redirect: "redirect",
    Rewrite: "rewrite",
} as const;

export type StaticSiteRouteType = (typeof StaticSiteRouteType)[keyof typeof StaticSiteRouteType];

export const StaticSiteServiceAutoDeploy = {
    Yes: "yes",
    No: "no",
} as const;

export type StaticSiteServiceAutoDeploy = (typeof StaticSiteServiceAutoDeploy)[keyof typeof StaticSiteServiceAutoDeploy];

export const StaticSiteServiceNotifyOnFail = {
    Default: "default",
    Notify: "notify",
    Ignore: "ignore",
} as const;

export type StaticSiteServiceNotifyOnFail = (typeof StaticSiteServiceNotifyOnFail)[keyof typeof StaticSiteServiceNotifyOnFail];

export const StaticSiteServiceSuspended = {
    Suspended: "suspended",
    NotSuspended: "not_suspended",
} as const;

export type StaticSiteServiceSuspended = (typeof StaticSiteServiceSuspended)[keyof typeof StaticSiteServiceSuspended];

export const StaticSiteServiceSuspendersItem = {
    Admin: "admin",
    Billing: "billing",
    User: "user",
    ParentService: "parent_service",
    Unknown: "unknown",
} as const;

export type StaticSiteServiceSuspendersItem = (typeof StaticSiteServiceSuspendersItem)[keyof typeof StaticSiteServiceSuspendersItem];

export const WebServiceDetailsCreateEnv = {
    Docker: "docker",
    Elixir: "elixir",
    Go: "go",
    Node: "node",
    Python: "python",
    Ruby: "ruby",
    Rust: "rust",
    Image: "image",
} as const;

/**
 * Environment (runtime)
 */
export type WebServiceDetailsCreateEnv = (typeof WebServiceDetailsCreateEnv)[keyof typeof WebServiceDetailsCreateEnv];

export const WebServiceDetailsCreatePlan = {
    Starter: "starter",
    Standard: "standard",
    Pro: "pro",
    ProPlus: "pro_plus",
    ProMax: "pro_max",
    ProUltra: "pro_ultra",
} as const;

export type WebServiceDetailsCreatePlan = (typeof WebServiceDetailsCreatePlan)[keyof typeof WebServiceDetailsCreatePlan];

export const WebServiceDetailsCreatePullRequestPreviewsEnabled = {
    Yes: "yes",
    No: "no",
} as const;

/**
 * Defaults to "no"
 */
export type WebServiceDetailsCreatePullRequestPreviewsEnabled = (typeof WebServiceDetailsCreatePullRequestPreviewsEnabled)[keyof typeof WebServiceDetailsCreatePullRequestPreviewsEnabled];

export const WebServiceDetailsCreateRegion = {
    Frankfurt: "frankfurt",
    Oregon: "oregon",
    Ohio: "ohio",
    Singapore: "singapore",
} as const;

export type WebServiceDetailsCreateRegion = (typeof WebServiceDetailsCreateRegion)[keyof typeof WebServiceDetailsCreateRegion];

export const WebServiceDetailsEnv = {
    Docker: "docker",
    Elixir: "elixir",
    Go: "go",
    Node: "node",
    Python: "python",
    Ruby: "ruby",
    Rust: "rust",
    Image: "image",
} as const;

/**
 * Environment (runtime)
 */
export type WebServiceDetailsEnv = (typeof WebServiceDetailsEnv)[keyof typeof WebServiceDetailsEnv];

export const WebServiceDetailsPlan = {
    Starter: "starter",
    StarterPlus: "starter_plus",
    Standard: "standard",
    StandardPlus: "standard_plus",
    Pro: "pro",
    ProPlus: "pro_plus",
    ProMax: "pro_max",
    ProUltra: "pro_ultra",
    Free: "free",
    Custom: "custom",
} as const;

/**
 * The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
 */
export type WebServiceDetailsPlan = (typeof WebServiceDetailsPlan)[keyof typeof WebServiceDetailsPlan];

export const WebServiceDetailsPullRequestPreviewsEnabled = {
    Yes: "yes",
    No: "no",
} as const;

export type WebServiceDetailsPullRequestPreviewsEnabled = (typeof WebServiceDetailsPullRequestPreviewsEnabled)[keyof typeof WebServiceDetailsPullRequestPreviewsEnabled];

export const WebServiceDetailsRegion = {
    Frankfurt: "frankfurt",
    Oregon: "oregon",
    Ohio: "ohio",
    Singapore: "singapore",
} as const;

export type WebServiceDetailsRegion = (typeof WebServiceDetailsRegion)[keyof typeof WebServiceDetailsRegion];

export const WebServiceServiceAutoDeploy = {
    Yes: "yes",
    No: "no",
} as const;

export type WebServiceServiceAutoDeploy = (typeof WebServiceServiceAutoDeploy)[keyof typeof WebServiceServiceAutoDeploy];

export const WebServiceServiceNotifyOnFail = {
    Default: "default",
    Notify: "notify",
    Ignore: "ignore",
} as const;

export type WebServiceServiceNotifyOnFail = (typeof WebServiceServiceNotifyOnFail)[keyof typeof WebServiceServiceNotifyOnFail];

export const WebServiceServiceSuspended = {
    Suspended: "suspended",
    NotSuspended: "not_suspended",
} as const;

export type WebServiceServiceSuspended = (typeof WebServiceServiceSuspended)[keyof typeof WebServiceServiceSuspended];

export const WebServiceServiceSuspendersItem = {
    Admin: "admin",
    Billing: "billing",
    User: "user",
    ParentService: "parent_service",
    Unknown: "unknown",
} as const;

export type WebServiceServiceSuspendersItem = (typeof WebServiceServiceSuspendersItem)[keyof typeof WebServiceServiceSuspendersItem];
