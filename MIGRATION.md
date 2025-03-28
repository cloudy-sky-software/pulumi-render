# Migration Guide

This doc serves as a migration guide for breaking changes.

## General

If a version requires the use of [`aliases`](https://www.pulumi.com/docs/concepts/options/aliases/),
they can be removed after you run at least one update with the `aliases`
resource option.

## v0.4.x

- `PostgresBackup` has been renamed to `PostgresExport` by Render.
  If you were using PostgresBackup, it's already likely broken
  because Render also renamed the endpoint from `/backup` to
  `/export`. So if you are using this resource today, you should
  switch to `PostgresExport` and add an alias in the resource
  options for the resource so that the state is migrated to the
  new resource.
- Render has switched to providing Valkey as a replacement
  for Redis as its key-value store. While they are still
  supporting Redis instances, no new Redis instances will
  allowed to be created. A new resource called `KeyValue`
  should be used instead. Read https://render.com/docs/key-value
  for more information about Render's plans for supporting
  existing Redis instances.

## v0.2.x

Changes to types used by resources.

Use aliases for the following resources:

### From v0.0.x:

- `render:services:CustomDomains` has changed to `render:services:CustomDomain`
- `render:services:EnvVars` has changed to `render:services:EnvVarForService`

### From v0.1.x:

- `render:services:EnvVar` has changed to `render:services:EnvVarForService`
- `render:services:AutoScaling` has changed to `render:services:AutoscaleService`
