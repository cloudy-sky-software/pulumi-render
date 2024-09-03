# Pulumi Native Provider for Render

[Render](https://render.com/) is a unified cloud to build and run all your apps and websites with free TLS certificates, a global CDN, DDoS protection, private networks, and auto deploys from Git.

> This provider was generated using [`pulschema`](https://github.com/cloudy-sky-software/pulschema) and [`pulumi-provider-framework`](https://github.com/cloudy-sky-software/pulumi-provider-framework).

## Package SDKs

- Node.js: https://www.npmjs.com/package/@cloudyskysoftware/pulumi-render
- Python: https://pypi.org/project/pulumi-render/
- .NET: https://www.nuget.org/packages/Pulumi.Render
- Go: `import github.com/cloudyskysoftware/pulumi-render/sdk/go/render`

## Using The Provider

You'll need an API key. Follow Render's [docs](https://render.com/docs/api#getting-started) for creating one.
Then set the API key as a secret with `pulumi config set --secret render:apiKey`.

## Releasing A New Version

:info: Switch to the `main` branch first and get the latest `git pull origin main && git fetch`. Check what the last release tag was.

1. Regular releases should just increment the patch version unless a minor or a major (breaking changes) version bump is warranted.
1. Update the `CHANGELOG.md` with notes about what will be included in this release.
1. Commit the changelog with `git commit -am "vX.Y.Z"` or something similar and push `git push origin main`.
1. Tag the commit with the release version by running

   ```bash
   git tag vX.Y.Z
   git tag sdk/vX.Y.Z
   ```

1. Push the tags.

   ```bash
   git push --tags
   ```
