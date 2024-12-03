import * as restate from "npm:@restatedev/restate-sdk/fetch";

const handler = restate
  .endpoint()
  .bind(
    restate.service({
      name: "Greeter",
      handlers: {
        greet: async (ctx: restate.Context, name: string) => {
          const greetingId = ctx.rand.uuidv4();
          
          // Inline notification function with 0.00001% failure rate
          await ctx.run(() => {
            if (Math.random() < 0.00001) {
              console.error(`👻 Failed to send notification: ${greetingId} - ${name}`);
              throw new Error(`Failed to send notification ${greetingId} - ${name}`);
            }
            console.log(`Notification sent: ${greetingId} - ${name}`);
          });
          
          await ctx.sleep(500);
          
          // Inline reminder function with 0.00001% failure rate
          await ctx.run(() => {
            if (Math.random() < 0.00001) {
              console.error(`👻 Failed to send reminder: ${greetingId}`);
              throw new Error(`Failed to send reminder: ${greetingId}`);
            }
            console.log(`Reminder sent: ${greetingId}`);
          });
          
          return `You said hi to ${name}!`;
        },
      },
    }),
  )
  .bidirectional()
  .handler();

Deno.serve({ port: 9080 }, handler.fetch);