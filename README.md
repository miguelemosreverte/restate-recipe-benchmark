# restate-recipe-benchmark
Will provide a recipe for use of restate and benchmark it


##### STEP 0: Quickstart
https://docs.restate.dev/get_started/quickstart/
```
brew install restatedev/tap/restate-server &&
brew install restatedev/tap/restate
restate-server 
```

```
restate example typescript-deno-hello-world
cd typescript-deno-hello-world
deno run --allow-net --allow-env --watch main.ts
```

```
restate dep add http://localhost:9080
```

```
curl localhost:8080/Greeter/greet -H 'content-type: application/json' -d '"Sarah"'
```

> "You said hi to Sarah!