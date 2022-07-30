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
