## 0.4.7

Upgrade `pulumi-provider-framework` to `v0.0.0-20250428021316-0a5ad29dff6f`.

## 0.4.3

Upgrade `pulumi-provider-framework` to `v0.0.0-20250403053300-d893597c4dc5`.

## 0.4.2

Fix a regression introduced in 0.4.0/0.4.1. (#453)

Other changes:

- Exclude some more `/logs` endpoints.
- New individual resources for adding a single SecretFile, EnvVar and EnvGroup for a service.

## 0.4.0, 0.4.1

Update to the latest Render OpenAPI spec.

New resource `KeyValue` supports Render's
managed [key-value store backed by Valkey](https://render.com/docs/key-value).

**Note**: See [[MIGRATION.md]] for migrating from 0.3.x to 0.4.x.

## 0.3.12

Update to the latest Render OpenAPI spec. (Fixes #367.)

## 0.3.11

Upgrade pulumi-provider-framework to the latest to fix import issue caused by a leading '/'.

## 0.3.10

Fix issue with resource imports by upgrading `pulumi-provider-framework` to the latest.

## 0.3.9

Fix input type for `SecretFilesForService`.

## 0.3.8

Fix input type for `EnvVarsForService`. This was broken in the 0.3.0 release.

## 0.3.6, 0.3.7

Fix execution error of list-style invokes.

## 0.3.5

Remove unnecessary envelope properties from `get*` and `list*` functions.

## 0.3.3, 0.3.4

Upgrade pulumi-provider-framework to fix a bug with validating response codes for DELETE calls.

## 0.3.2

Fix resource name for Redis.

## 0.3.0, 0.3.1

Updated the OpenAPI spec to the latest which adds many new resources including
Postgres, Redis, environment groups.

## 0.2.1

- Fix request validation error when updating `EnvVarsForService` resource.
- Check for all service types in OnPostUpdate to trigger a service deployment.

## 0.2.0

Switch to Render's official OpenAPI spec. See the [migration guide](./MIGRATION.md) for more details.
For the most part some resources require an alias to avoid a replacement.

Upgrade `pulumi-provider-framework` to the latest which now allows
importing resources that are children of other resources.

## 0.1.0, 0.1.1

This release contains some breaking changes related to the names of functions but does not change any resource names.

- Upgrade pulschema and pulumi-provider-framework which changes the names of certain get/list functions.
- New resources have been added:
  - `RegistryCredential`
  - `AutoScaling`

## 0.0.47

- Upgrade pulschema and pulumi-provider-framework modules to add auto-naming feature.

## 0.0.46

- Fix `CustomDomain` and `Job` resource inputs

## 0.0.45

- Add support for EnvVar resource ([#55](cloudy-sky-software/pulumi-render#55))

## 0.0.42 - 0.0.44

- Fix resource import

## 0.0.41

- Fix resource reads

## 0.0.40

- Fix the plugin download URL

## 0.0.39

- Switch to the [Pulumi provider framework](https://github.com/cloudy-sky-software/pulumi-provider-framework)

## 0.0.37

- Fix ambiguous type error in service details.
- Add envVars and secretFiles to the base service type.

## 0.0.36

- Fix discriminator mapping bug.
- Add missing output props for all service types.

## 0.0.35

Add OpenAPI schemas for all Render service types.

## 0.0.34

Implement support for get/list\* funcs.

## 0.0.33

Fix for handling updates vs. replacements.

## 0.0.32

Fix for path param handling when dealing with output properties.

## 0.0.31

Change package names under which the language SDKs are published.
From here on the following are the respective package names:

Go: `github.com/cloudy-sky-software/pulumi-render/sdk/go/render`
.NET: `Pulumi.Render`
npm (Node.js): `@cloudyskysoftware/pulumi-render`
PyPi (Python): `pulumi_render`

## 0.0.30

- `Service` resource is now split to each individual service type if the POST operation
  uses a discriminator.
- Gracefully handle `allOf` in the OpenAPI spec for Pulumi schema generation.

## 0.0.29

Fix plugin download URL

## 0.0.28

First release
