// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.Render.NotificationSettings
{
    [EnumType]
    public readonly struct GetOwnerNotificationSettingPropertiesNotificationsToSend : IEquatable<GetOwnerNotificationSettingPropertiesNotificationsToSend>
    {
        private readonly string _value;

        private GetOwnerNotificationSettingPropertiesNotificationsToSend(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static GetOwnerNotificationSettingPropertiesNotificationsToSend None { get; } = new GetOwnerNotificationSettingPropertiesNotificationsToSend("none");
        public static GetOwnerNotificationSettingPropertiesNotificationsToSend Failure { get; } = new GetOwnerNotificationSettingPropertiesNotificationsToSend("failure");
        public static GetOwnerNotificationSettingPropertiesNotificationsToSend All { get; } = new GetOwnerNotificationSettingPropertiesNotificationsToSend("all");

        public static bool operator ==(GetOwnerNotificationSettingPropertiesNotificationsToSend left, GetOwnerNotificationSettingPropertiesNotificationsToSend right) => left.Equals(right);
        public static bool operator !=(GetOwnerNotificationSettingPropertiesNotificationsToSend left, GetOwnerNotificationSettingPropertiesNotificationsToSend right) => !left.Equals(right);

        public static explicit operator string(GetOwnerNotificationSettingPropertiesNotificationsToSend value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GetOwnerNotificationSettingPropertiesNotificationsToSend other && Equals(other);
        public bool Equals(GetOwnerNotificationSettingPropertiesNotificationsToSend other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct GetServiceNotificationOverridePropertiesNotificationsToSend : IEquatable<GetServiceNotificationOverridePropertiesNotificationsToSend>
    {
        private readonly string _value;

        private GetServiceNotificationOverridePropertiesNotificationsToSend(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static GetServiceNotificationOverridePropertiesNotificationsToSend Default { get; } = new GetServiceNotificationOverridePropertiesNotificationsToSend("default");
        public static GetServiceNotificationOverridePropertiesNotificationsToSend None { get; } = new GetServiceNotificationOverridePropertiesNotificationsToSend("none");
        public static GetServiceNotificationOverridePropertiesNotificationsToSend Failure { get; } = new GetServiceNotificationOverridePropertiesNotificationsToSend("failure");
        public static GetServiceNotificationOverridePropertiesNotificationsToSend All { get; } = new GetServiceNotificationOverridePropertiesNotificationsToSend("all");

        public static bool operator ==(GetServiceNotificationOverridePropertiesNotificationsToSend left, GetServiceNotificationOverridePropertiesNotificationsToSend right) => left.Equals(right);
        public static bool operator !=(GetServiceNotificationOverridePropertiesNotificationsToSend left, GetServiceNotificationOverridePropertiesNotificationsToSend right) => !left.Equals(right);

        public static explicit operator string(GetServiceNotificationOverridePropertiesNotificationsToSend value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GetServiceNotificationOverridePropertiesNotificationsToSend other && Equals(other);
        public bool Equals(GetServiceNotificationOverridePropertiesNotificationsToSend other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled : IEquatable<GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled>
    {
        private readonly string _value;

        private GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled Default { get; } = new GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled("default");
        public static GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled False { get; } = new GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled("false");
        public static GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled True { get; } = new GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled("true");

        public static bool operator ==(GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled left, GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled right) => left.Equals(right);
        public static bool operator !=(GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled left, GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled right) => !left.Equals(right);

        public static explicit operator string(GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled other && Equals(other);
        public bool Equals(GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct NotificationOverrideWithCursorOverridePropertiesNotificationsToSend : IEquatable<NotificationOverrideWithCursorOverridePropertiesNotificationsToSend>
    {
        private readonly string _value;

        private NotificationOverrideWithCursorOverridePropertiesNotificationsToSend(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static NotificationOverrideWithCursorOverridePropertiesNotificationsToSend Default { get; } = new NotificationOverrideWithCursorOverridePropertiesNotificationsToSend("default");
        public static NotificationOverrideWithCursorOverridePropertiesNotificationsToSend None { get; } = new NotificationOverrideWithCursorOverridePropertiesNotificationsToSend("none");
        public static NotificationOverrideWithCursorOverridePropertiesNotificationsToSend Failure { get; } = new NotificationOverrideWithCursorOverridePropertiesNotificationsToSend("failure");
        public static NotificationOverrideWithCursorOverridePropertiesNotificationsToSend All { get; } = new NotificationOverrideWithCursorOverridePropertiesNotificationsToSend("all");

        public static bool operator ==(NotificationOverrideWithCursorOverridePropertiesNotificationsToSend left, NotificationOverrideWithCursorOverridePropertiesNotificationsToSend right) => left.Equals(right);
        public static bool operator !=(NotificationOverrideWithCursorOverridePropertiesNotificationsToSend left, NotificationOverrideWithCursorOverridePropertiesNotificationsToSend right) => !left.Equals(right);

        public static explicit operator string(NotificationOverrideWithCursorOverridePropertiesNotificationsToSend value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NotificationOverrideWithCursorOverridePropertiesNotificationsToSend other && Equals(other);
        public bool Equals(NotificationOverrideWithCursorOverridePropertiesNotificationsToSend other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled : IEquatable<NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled>
    {
        private readonly string _value;

        private NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled Default { get; } = new NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled("default");
        public static NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled False { get; } = new NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled("false");
        public static NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled True { get; } = new NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled("true");

        public static bool operator ==(NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled left, NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled right) => left.Equals(right);
        public static bool operator !=(NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled left, NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled right) => !left.Equals(right);

        public static explicit operator string(NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled other && Equals(other);
        public bool Equals(NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
