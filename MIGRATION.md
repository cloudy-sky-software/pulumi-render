# Migration Guide

This doc serves as a migration guide for breaking changes.

## General

If a version requires the use of [`aliases`](https://www.pulumi.com/docs/concepts/options/aliases/),
they can be removed after you run at least one update with the `aliases`
resource option.

## v0.2.x

- Changes to types used by resources.

- Use aliases for the following resources:

### From v0.0.x:

- `render:services:CustomDomains` has changed to `render:services:CustomDomain`
- `render:services:EnvVars` has changed to `render:services:EnvVarForService`

### From v0.1.x:

- `render:services:EnvVar` has changed to `render:services:EnvVarForService`
- `render:services:AutoScaling` has changed to `render:services:AutoscaleService`
