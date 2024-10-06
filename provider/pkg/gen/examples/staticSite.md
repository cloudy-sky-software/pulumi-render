{{% examples %}}

## Example Usage

{{% example %}}

### Create a StaticSite service

```typescript
import * as render from "@cloudyskysoftware/render";

const ownerId = render.owners
    .listOwnersOutput()
    .apply(
        (result) =>
            result.items.filter(
                (i) => i.owner?.email === "OWNER_EMAIL"
            )[0].owner?.id || ""
    );

const staticSite = new render.services.StaticSite("staticsite", {
    name: "My custom static site",
    ownerId,
    repo: "https://github.com/cloudy-sky-software/test-static-site",
    autoDeploy: "no",
    branch: "main",
    serviceDetails: {
    publishPath: "public",
    },
    type: "static_site",
});

export const url = staticSite.serviceDetails.apply((s) => s?.url);
```

{{% /example %}}
{{% /examples %}}
