import * as restate from "npm:@restatedev/restate-sdk/fetch";

const handler = restate
  .endpoint()
  .bind(
    restate.service({
      name: "Greeter",
      handlers: {
        greet: async (ctx: restate.Context, name: string) => {
          const greetingId = ctx.rand.uuidv4();
          
          
          return `You said hi to ${name}!`;
        },
      },
    }),
  )
  .bidirectional()
  .handler();

Deno.serve({ port: 9080 }, handler.fetch);