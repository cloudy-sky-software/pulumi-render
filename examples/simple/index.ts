import * as render from "@css/render";

import * as o from "@css/render/types/output";

const service = new render.services.Service("my-service", {
  name: "My test service",
  ownerId: "",
  repo: "https://github.com/praneetloke/test",
  type: "static_site",
  serviceDetails: {},
});

export const staticSiteUrl = service.serviceDetails.apply(
  (d) => (d as o.services.ServiceStaticSite).url
);
