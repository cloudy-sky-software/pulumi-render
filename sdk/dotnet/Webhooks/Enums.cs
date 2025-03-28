// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.Render.Webhooks
{
    [EnumType]
    public readonly struct EventFilterItem : IEquatable<EventFilterItem>
    {
        private readonly string _value;

        private EventFilterItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EventFilterItem AutoscalingConfigChanged { get; } = new EventFilterItem("autoscaling_config_changed");
        public static EventFilterItem AutoscalingEnded { get; } = new EventFilterItem("autoscaling_ended");
        public static EventFilterItem AutoscalingStarted { get; } = new EventFilterItem("autoscaling_started");
        public static EventFilterItem BranchDeleted { get; } = new EventFilterItem("branch_deleted");
        public static EventFilterItem BuildEnded { get; } = new EventFilterItem("build_ended");
        public static EventFilterItem BuildStarted { get; } = new EventFilterItem("build_started");
        public static EventFilterItem CommitIgnored { get; } = new EventFilterItem("commit_ignored");
        public static EventFilterItem CronJobRunEnded { get; } = new EventFilterItem("cron_job_run_ended");
        public static EventFilterItem CronJobRunStarted { get; } = new EventFilterItem("cron_job_run_started");
        public static EventFilterItem DeployEnded { get; } = new EventFilterItem("deploy_ended");
        public static EventFilterItem DeployStarted { get; } = new EventFilterItem("deploy_started");
        public static EventFilterItem DiskCreated { get; } = new EventFilterItem("disk_created");
        public static EventFilterItem DiskUpdated { get; } = new EventFilterItem("disk_updated");
        public static EventFilterItem DiskDeleted { get; } = new EventFilterItem("disk_deleted");
        public static EventFilterItem ImagePullFailed { get; } = new EventFilterItem("image_pull_failed");
        public static EventFilterItem InstanceCountChanged { get; } = new EventFilterItem("instance_count_changed");
        public static EventFilterItem JobRunEnded { get; } = new EventFilterItem("job_run_ended");
        public static EventFilterItem MaintenanceModeEnabled { get; } = new EventFilterItem("maintenance_mode_enabled");
        public static EventFilterItem MaintenanceModeUriUpdated { get; } = new EventFilterItem("maintenance_mode_uri_updated");
        public static EventFilterItem MaintenanceEnded { get; } = new EventFilterItem("maintenance_ended");
        public static EventFilterItem MaintenanceStarted { get; } = new EventFilterItem("maintenance_started");
        public static EventFilterItem PlanChanged { get; } = new EventFilterItem("plan_changed");
        public static EventFilterItem PreDeployEnded { get; } = new EventFilterItem("pre_deploy_ended");
        public static EventFilterItem PreDeployStarted { get; } = new EventFilterItem("pre_deploy_started");
        public static EventFilterItem ServerAvailable { get; } = new EventFilterItem("server_available");
        public static EventFilterItem ServerFailed { get; } = new EventFilterItem("server_failed");
        public static EventFilterItem ServerHardwareFailure { get; } = new EventFilterItem("server_hardware_failure");
        public static EventFilterItem ServerRestarted { get; } = new EventFilterItem("server_restarted");
        public static EventFilterItem ServerUnhealthy { get; } = new EventFilterItem("server_unhealthy");
        public static EventFilterItem ServiceResumed { get; } = new EventFilterItem("service_resumed");
        public static EventFilterItem ServiceSuspended { get; } = new EventFilterItem("service_suspended");
        public static EventFilterItem ZeroDowntimeRedeployEnded { get; } = new EventFilterItem("zero_downtime_redeploy_ended");
        public static EventFilterItem ZeroDowntimeRedeployStarted { get; } = new EventFilterItem("zero_downtime_redeploy_started");
        public static EventFilterItem PostgresAvailable { get; } = new EventFilterItem("postgres_available");
        public static EventFilterItem PostgresBackupCompleted { get; } = new EventFilterItem("postgres_backup_completed");
        public static EventFilterItem PostgresBackupStarted { get; } = new EventFilterItem("postgres_backup_started");
        public static EventFilterItem PostgresClusterLeaderChanged { get; } = new EventFilterItem("postgres_cluster_leader_changed");
        public static EventFilterItem PostgresCreated { get; } = new EventFilterItem("postgres_created");
        public static EventFilterItem PostgresDiskSizeChanged { get; } = new EventFilterItem("postgres_disk_size_changed");
        public static EventFilterItem PostgresHaStatusChanged { get; } = new EventFilterItem("postgres_ha_status_changed");
        public static EventFilterItem PostgresRestarted { get; } = new EventFilterItem("postgres_restarted");
        public static EventFilterItem PostgresUnavailable { get; } = new EventFilterItem("postgres_unavailable");
        public static EventFilterItem PostgresUpgradeFailed { get; } = new EventFilterItem("postgres_upgrade_failed");
        public static EventFilterItem PostgresUpgradeStarted { get; } = new EventFilterItem("postgres_upgrade_started");
        public static EventFilterItem PostgresUpgradeSucceeded { get; } = new EventFilterItem("postgres_upgrade_succeeded");
        public static EventFilterItem PostgresReadReplicasChanged { get; } = new EventFilterItem("postgres_read_replicas_changed");
        public static EventFilterItem KeyValueAvailable { get; } = new EventFilterItem("key_value_available");
        public static EventFilterItem KeyValueConfigRestart { get; } = new EventFilterItem("key_value_config_restart");
        public static EventFilterItem KeyValueUnhealthy { get; } = new EventFilterItem("key_value_unhealthy");

        public static bool operator ==(EventFilterItem left, EventFilterItem right) => left.Equals(right);
        public static bool operator !=(EventFilterItem left, EventFilterItem right) => !left.Equals(right);

        public static explicit operator string(EventFilterItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EventFilterItem other && Equals(other);
        public bool Equals(EventFilterItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct WebhookWithCursorWebhookPropertiesEventFilterItem : IEquatable<WebhookWithCursorWebhookPropertiesEventFilterItem>
    {
        private readonly string _value;

        private WebhookWithCursorWebhookPropertiesEventFilterItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WebhookWithCursorWebhookPropertiesEventFilterItem AutoscalingConfigChanged { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("autoscaling_config_changed");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem AutoscalingEnded { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("autoscaling_ended");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem AutoscalingStarted { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("autoscaling_started");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem BranchDeleted { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("branch_deleted");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem BuildEnded { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("build_ended");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem BuildStarted { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("build_started");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem CommitIgnored { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("commit_ignored");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem CronJobRunEnded { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("cron_job_run_ended");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem CronJobRunStarted { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("cron_job_run_started");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem DeployEnded { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("deploy_ended");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem DeployStarted { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("deploy_started");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem DiskCreated { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("disk_created");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem DiskUpdated { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("disk_updated");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem DiskDeleted { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("disk_deleted");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem ImagePullFailed { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("image_pull_failed");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem InstanceCountChanged { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("instance_count_changed");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem JobRunEnded { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("job_run_ended");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem MaintenanceModeEnabled { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("maintenance_mode_enabled");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem MaintenanceModeUriUpdated { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("maintenance_mode_uri_updated");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem MaintenanceEnded { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("maintenance_ended");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem MaintenanceStarted { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("maintenance_started");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem PlanChanged { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("plan_changed");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem PreDeployEnded { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("pre_deploy_ended");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem PreDeployStarted { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("pre_deploy_started");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem ServerAvailable { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("server_available");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem ServerFailed { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("server_failed");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem ServerHardwareFailure { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("server_hardware_failure");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem ServerRestarted { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("server_restarted");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem ServerUnhealthy { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("server_unhealthy");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem ServiceResumed { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("service_resumed");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem ServiceSuspended { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("service_suspended");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem ZeroDowntimeRedeployEnded { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("zero_downtime_redeploy_ended");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem ZeroDowntimeRedeployStarted { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("zero_downtime_redeploy_started");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem PostgresAvailable { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("postgres_available");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem PostgresBackupCompleted { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("postgres_backup_completed");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem PostgresBackupStarted { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("postgres_backup_started");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem PostgresClusterLeaderChanged { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("postgres_cluster_leader_changed");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem PostgresCreated { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("postgres_created");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem PostgresDiskSizeChanged { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("postgres_disk_size_changed");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem PostgresHaStatusChanged { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("postgres_ha_status_changed");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem PostgresRestarted { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("postgres_restarted");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem PostgresUnavailable { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("postgres_unavailable");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem PostgresUpgradeFailed { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("postgres_upgrade_failed");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem PostgresUpgradeStarted { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("postgres_upgrade_started");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem PostgresUpgradeSucceeded { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("postgres_upgrade_succeeded");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem PostgresReadReplicasChanged { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("postgres_read_replicas_changed");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem KeyValueAvailable { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("key_value_available");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem KeyValueConfigRestart { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("key_value_config_restart");
        public static WebhookWithCursorWebhookPropertiesEventFilterItem KeyValueUnhealthy { get; } = new WebhookWithCursorWebhookPropertiesEventFilterItem("key_value_unhealthy");

        public static bool operator ==(WebhookWithCursorWebhookPropertiesEventFilterItem left, WebhookWithCursorWebhookPropertiesEventFilterItem right) => left.Equals(right);
        public static bool operator !=(WebhookWithCursorWebhookPropertiesEventFilterItem left, WebhookWithCursorWebhookPropertiesEventFilterItem right) => !left.Equals(right);

        public static explicit operator string(WebhookWithCursorWebhookPropertiesEventFilterItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WebhookWithCursorWebhookPropertiesEventFilterItem other && Equals(other);
        public bool Equals(WebhookWithCursorWebhookPropertiesEventFilterItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct WebhookWithCursorpropertieswebhookEventFilterItem : IEquatable<WebhookWithCursorpropertieswebhookEventFilterItem>
    {
        private readonly string _value;

        private WebhookWithCursorpropertieswebhookEventFilterItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WebhookWithCursorpropertieswebhookEventFilterItem AutoscalingConfigChanged { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("autoscaling_config_changed");
        public static WebhookWithCursorpropertieswebhookEventFilterItem AutoscalingEnded { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("autoscaling_ended");
        public static WebhookWithCursorpropertieswebhookEventFilterItem AutoscalingStarted { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("autoscaling_started");
        public static WebhookWithCursorpropertieswebhookEventFilterItem BranchDeleted { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("branch_deleted");
        public static WebhookWithCursorpropertieswebhookEventFilterItem BuildEnded { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("build_ended");
        public static WebhookWithCursorpropertieswebhookEventFilterItem BuildStarted { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("build_started");
        public static WebhookWithCursorpropertieswebhookEventFilterItem CommitIgnored { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("commit_ignored");
        public static WebhookWithCursorpropertieswebhookEventFilterItem CronJobRunEnded { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("cron_job_run_ended");
        public static WebhookWithCursorpropertieswebhookEventFilterItem CronJobRunStarted { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("cron_job_run_started");
        public static WebhookWithCursorpropertieswebhookEventFilterItem DeployEnded { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("deploy_ended");
        public static WebhookWithCursorpropertieswebhookEventFilterItem DeployStarted { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("deploy_started");
        public static WebhookWithCursorpropertieswebhookEventFilterItem DiskCreated { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("disk_created");
        public static WebhookWithCursorpropertieswebhookEventFilterItem DiskUpdated { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("disk_updated");
        public static WebhookWithCursorpropertieswebhookEventFilterItem DiskDeleted { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("disk_deleted");
        public static WebhookWithCursorpropertieswebhookEventFilterItem ImagePullFailed { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("image_pull_failed");
        public static WebhookWithCursorpropertieswebhookEventFilterItem InstanceCountChanged { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("instance_count_changed");
        public static WebhookWithCursorpropertieswebhookEventFilterItem JobRunEnded { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("job_run_ended");
        public static WebhookWithCursorpropertieswebhookEventFilterItem MaintenanceModeEnabled { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("maintenance_mode_enabled");
        public static WebhookWithCursorpropertieswebhookEventFilterItem MaintenanceModeUriUpdated { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("maintenance_mode_uri_updated");
        public static WebhookWithCursorpropertieswebhookEventFilterItem MaintenanceEnded { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("maintenance_ended");
        public static WebhookWithCursorpropertieswebhookEventFilterItem MaintenanceStarted { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("maintenance_started");
        public static WebhookWithCursorpropertieswebhookEventFilterItem PlanChanged { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("plan_changed");
        public static WebhookWithCursorpropertieswebhookEventFilterItem PreDeployEnded { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("pre_deploy_ended");
        public static WebhookWithCursorpropertieswebhookEventFilterItem PreDeployStarted { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("pre_deploy_started");
        public static WebhookWithCursorpropertieswebhookEventFilterItem ServerAvailable { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("server_available");
        public static WebhookWithCursorpropertieswebhookEventFilterItem ServerFailed { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("server_failed");
        public static WebhookWithCursorpropertieswebhookEventFilterItem ServerHardwareFailure { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("server_hardware_failure");
        public static WebhookWithCursorpropertieswebhookEventFilterItem ServerRestarted { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("server_restarted");
        public static WebhookWithCursorpropertieswebhookEventFilterItem ServerUnhealthy { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("server_unhealthy");
        public static WebhookWithCursorpropertieswebhookEventFilterItem ServiceResumed { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("service_resumed");
        public static WebhookWithCursorpropertieswebhookEventFilterItem ServiceSuspended { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("service_suspended");
        public static WebhookWithCursorpropertieswebhookEventFilterItem ZeroDowntimeRedeployEnded { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("zero_downtime_redeploy_ended");
        public static WebhookWithCursorpropertieswebhookEventFilterItem ZeroDowntimeRedeployStarted { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("zero_downtime_redeploy_started");
        public static WebhookWithCursorpropertieswebhookEventFilterItem PostgresAvailable { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("postgres_available");
        public static WebhookWithCursorpropertieswebhookEventFilterItem PostgresBackupCompleted { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("postgres_backup_completed");
        public static WebhookWithCursorpropertieswebhookEventFilterItem PostgresBackupStarted { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("postgres_backup_started");
        public static WebhookWithCursorpropertieswebhookEventFilterItem PostgresClusterLeaderChanged { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("postgres_cluster_leader_changed");
        public static WebhookWithCursorpropertieswebhookEventFilterItem PostgresCreated { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("postgres_created");
        public static WebhookWithCursorpropertieswebhookEventFilterItem PostgresDiskSizeChanged { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("postgres_disk_size_changed");
        public static WebhookWithCursorpropertieswebhookEventFilterItem PostgresHaStatusChanged { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("postgres_ha_status_changed");
        public static WebhookWithCursorpropertieswebhookEventFilterItem PostgresRestarted { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("postgres_restarted");
        public static WebhookWithCursorpropertieswebhookEventFilterItem PostgresUnavailable { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("postgres_unavailable");
        public static WebhookWithCursorpropertieswebhookEventFilterItem PostgresUpgradeFailed { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("postgres_upgrade_failed");
        public static WebhookWithCursorpropertieswebhookEventFilterItem PostgresUpgradeStarted { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("postgres_upgrade_started");
        public static WebhookWithCursorpropertieswebhookEventFilterItem PostgresUpgradeSucceeded { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("postgres_upgrade_succeeded");
        public static WebhookWithCursorpropertieswebhookEventFilterItem PostgresReadReplicasChanged { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("postgres_read_replicas_changed");
        public static WebhookWithCursorpropertieswebhookEventFilterItem KeyValueAvailable { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("key_value_available");
        public static WebhookWithCursorpropertieswebhookEventFilterItem KeyValueConfigRestart { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("key_value_config_restart");
        public static WebhookWithCursorpropertieswebhookEventFilterItem KeyValueUnhealthy { get; } = new WebhookWithCursorpropertieswebhookEventFilterItem("key_value_unhealthy");

        public static bool operator ==(WebhookWithCursorpropertieswebhookEventFilterItem left, WebhookWithCursorpropertieswebhookEventFilterItem right) => left.Equals(right);
        public static bool operator !=(WebhookWithCursorpropertieswebhookEventFilterItem left, WebhookWithCursorpropertieswebhookEventFilterItem right) => !left.Equals(right);

        public static explicit operator string(WebhookWithCursorpropertieswebhookEventFilterItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WebhookWithCursorpropertieswebhookEventFilterItem other && Equals(other);
        public bool Equals(WebhookWithCursorpropertieswebhookEventFilterItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
