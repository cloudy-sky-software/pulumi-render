// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.Render.Environments
{
    /// <summary>
    /// Indicates whether an environment is `unprotected` or `protected`. Only admin users can perform destructive actions in `protected` environments.
    /// </summary>
    [EnumType]
    public readonly struct EnvironmentProtectedStatus : IEquatable<EnvironmentProtectedStatus>
    {
        private readonly string _value;

        private EnvironmentProtectedStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EnvironmentProtectedStatus Unprotected { get; } = new EnvironmentProtectedStatus("unprotected");
        public static EnvironmentProtectedStatus Protected { get; } = new EnvironmentProtectedStatus("protected");

        public static bool operator ==(EnvironmentProtectedStatus left, EnvironmentProtectedStatus right) => left.Equals(right);
        public static bool operator !=(EnvironmentProtectedStatus left, EnvironmentProtectedStatus right) => !left.Equals(right);

        public static explicit operator string(EnvironmentProtectedStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnvironmentProtectedStatus other && Equals(other);
        public bool Equals(EnvironmentProtectedStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicates whether an environment is `unprotected` or `protected`. Only admin users can perform destructive actions in `protected` environments.
    /// </summary>
    [EnumType]
    public readonly struct ProtectedStatus : IEquatable<ProtectedStatus>
    {
        private readonly string _value;

        private ProtectedStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ProtectedStatus Unprotected { get; } = new ProtectedStatus("unprotected");
        public static ProtectedStatus Protected { get; } = new ProtectedStatus("protected");

        public static bool operator ==(ProtectedStatus left, ProtectedStatus right) => left.Equals(right);
        public static bool operator !=(ProtectedStatus left, ProtectedStatus right) => !left.Equals(right);

        public static explicit operator string(ProtectedStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProtectedStatus other && Equals(other);
        public bool Equals(ProtectedStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}