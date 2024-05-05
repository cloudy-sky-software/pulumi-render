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
