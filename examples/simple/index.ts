import * as render from "@pulumi/render";

const random = new render.Random("my-random", { length: 24 });

export const output = random.result;
