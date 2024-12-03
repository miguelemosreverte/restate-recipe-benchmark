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

-----

##### Benchmark

```
go mod tidy
go run benchmark.go
```

```
ls -t reports/benchmark-report-*.html | head -n 1 | xargs -I {} mv {} index.html
```

Logs:

```
go run benchmark.go
# command-line-arguments
ld: warning: ignoring duplicate libraries: '-lrocksdb'
Starting benchmark...
2024/12/04 00:11:17 Monitoring Restate data at: /Users/miguel_lemos/Desktop/restate-recipe-benchmark/restate-data
2024/12/04 00:11:17 Initial disk usage: 878464024576 bytes
2024/12/04 00:11:17 Starting new test iteration
2024/12/04 00:11:17 Testing with 20 users (range: 5-100)
2024/12/04 00:11:17 Starting workers...
2024/12/04 00:11:17 Workers started
2024/12/04 00:11:18 TPS: 0.00, In-flight Retries: 0, Users: 20
2024/12/04 00:11:19 TPS: 9.00, In-flight Retries: 0, Users: 20
2024/12/04 00:11:20 TPS: 9.00, In-flight Retries: 0, Users: 20
2024/12/04 00:11:21 TPS: 11.00, In-flight Retries: 0, Users: 20
2024/12/04 00:11:22 TPS: 11.40, In-flight Retries: 0, Users: 20
2024/12/04 00:11:23 TPS: 11.33, In-flight Retries: 0, Users: 20
2024/12/04 00:11:23 Current configuration stable with avg TPS: 11.24
2024/12/04 00:11:24 TPS: 11.00, In-flight Retries: 0, Users: 20
2024/12/04 00:11:24 Current configuration stable with avg TPS: 11.24
2024/12/04 00:11:25 TPS: 11.25, In-flight Retries: 0, Users: 20
2024/12/04 00:11:25 Current configuration stable with avg TPS: 11.19
2024/12/04 00:11:26 TPS: 11.55, In-flight Retries: 0, Users: 20
2024/12/04 00:11:26 Current configuration stable with avg TPS: 11.27
2024/12/04 00:11:27 TPS: 11.70, In-flight Retries: 0, Users: 20
2024/12/04 00:11:27 Current configuration stable with avg TPS: 11.50
2024/12/04 00:11:27 Test duration complete
2024/12/04 00:11:27 Cleanup: Closing stop channel
2024/12/04 00:11:27 Cleanup: Waiting for workers with timeout
2024/12/04 00:11:32 Test iteration cleanup timeout, forcing continuation
2024/12/04 00:11:32 Starting cooldown
2024/12/04 00:11:37 Cooldown complete
2024/12/04 00:11:37 Test was stable with TPS 11.50
2024/12/04 00:11:37 New best configuration found!
2024/12/04 00:11:37 Aggressively doubling users from 20 to 40
2024/12/04 00:11:37 Starting next iteration
2024/12/04 00:11:37 Starting new test iteration
2024/12/04 00:11:37 Testing with 40 users (range: 21-80)
2024/12/04 00:11:37 Starting workers...
2024/12/04 00:11:37 Cleanup: Worker wait timeout, continuing anyway
2024/12/04 00:11:37 Workers started
2024/12/04 00:11:38 TPS: 0.00, In-flight Retries: 0, Users: 40
2024/12/04 00:11:39 TPS: 18.49, In-flight Retries: 0, Users: 40
2024/12/04 00:11:40 TPS: 21.99, In-flight Retries: 0, Users: 40
2024/12/04 00:11:41 TPS: 24.49, In-flight Retries: 0, Users: 40
2024/12/04 00:11:42 TPS: 25.19, In-flight Retries: 0, Users: 40
2024/12/04 00:11:42 Current configuration stable with avg TPS: 23.89
2024/12/04 00:11:43 TPS: 25.33, In-flight Retries: 0, Users: 40
2024/12/04 00:11:43 Current configuration stable with avg TPS: 25.00
2024/12/04 00:11:44 TPS: 25.14, In-flight Retries: 0, Users: 40
2024/12/04 00:11:44 Current configuration stable with avg TPS: 25.22
2024/12/04 00:11:45 TPS: 24.62, In-flight Retries: 0, Users: 40
2024/12/04 00:11:45 Current configuration stable with avg TPS: 25.03
2024/12/04 00:11:46 TPS: 24.55, In-flight Retries: 0, Users: 40
2024/12/04 00:11:46 Current configuration stable with avg TPS: 24.77
2024/12/04 00:11:47 TPS: 24.40, In-flight Retries: 0, Users: 40
2024/12/04 00:11:47 Current configuration stable with avg TPS: 24.52
2024/12/04 00:11:47 Test duration complete
2024/12/04 00:11:47 Cleanup: Closing stop channel
2024/12/04 00:11:47 Cleanup: Waiting for workers with timeout
2024/12/04 00:11:52 Test iteration cleanup timeout, forcing continuation
2024/12/04 00:11:52 Starting cooldown
2024/12/04 00:11:57 Cooldown complete
2024/12/04 00:11:57 Test was stable with TPS 24.52
2024/12/04 00:11:57 New best configuration found!
2024/12/04 00:11:57 Aggressively doubling users from 40 to 80
2024/12/04 00:11:57 Starting next iteration
2024/12/04 00:11:57 Starting new test iteration
2024/12/04 00:11:57 Testing with 80 users (range: 41-160)
2024/12/04 00:11:57 Starting workers...
2024/12/04 00:11:57 Workers started
2024/12/04 00:11:57 Cleanup: Worker wait timeout, continuing anyway
2024/12/04 00:11:58 TPS: 0.00, In-flight Retries: 0, Users: 80
2024/12/04 00:11:59 TPS: 35.99, In-flight Retries: 0, Users: 80
2024/12/04 00:12:00 TPS: 45.32, In-flight Retries: 0, Users: 80
2024/12/04 00:12:01 TPS: 44.99, In-flight Retries: 0, Users: 80
2024/12/04 00:12:02 TPS: 45.19, In-flight Retries: 0, Users: 80
2024/12/04 00:12:02 Current configuration stable with avg TPS: 45.17
2024/12/04 00:12:03 TPS: 46.00, In-flight Retries: 0, Users: 80
2024/12/04 00:12:03 Current configuration stable with avg TPS: 45.39
2024/12/04 00:12:04 TPS: 46.28, In-flight Retries: 0, Users: 80
2024/12/04 00:12:04 Current configuration stable with avg TPS: 45.82
2024/12/04 00:12:05 TPS: 45.75, In-flight Retries: 0, Users: 80
2024/12/04 00:12:05 Current configuration stable with avg TPS: 46.01
2024/12/04 00:12:06 TPS: 45.55, In-flight Retries: 0, Users: 80
2024/12/04 00:12:06 Current configuration stable with avg TPS: 45.86
2024/12/04 00:12:07 TPS: 45.80, In-flight Retries: 0, Users: 80
2024/12/04 00:12:07 Current configuration stable with avg TPS: 45.70
2024/12/04 00:12:07 Test duration complete
2024/12/04 00:12:07 Cleanup: Closing stop channel
2024/12/04 00:12:07 Cleanup: Waiting for workers with timeout
2024/12/04 00:12:12 Test iteration cleanup timeout, forcing continuation
2024/12/04 00:12:12 Starting cooldown
2024/12/04 00:12:17 Cooldown complete
2024/12/04 00:12:17 Test was stable with TPS 45.70
2024/12/04 00:12:17 New best configuration found!
2024/12/04 00:12:17 Aggressively doubling users from 80 to 160
2024/12/04 00:12:17 Starting next iteration
2024/12/04 00:12:17 Starting new test iteration
2024/12/04 00:12:17 Testing with 160 users (range: 81-320)
2024/12/04 00:12:17 Starting workers...
2024/12/04 00:12:17 Workers started
2024/12/04 00:12:17 Cleanup: Worker wait timeout, continuing anyway
2024/12/04 00:12:18 TPS: 2.00, In-flight Retries: 0, Users: 160
2024/12/04 00:12:19 TPS: 70.97, In-flight Retries: 0, Users: 160
2024/12/04 00:12:20 TPS: 85.64, In-flight Retries: 0, Users: 160
2024/12/04 00:12:21 TPS: 89.23, In-flight Retries: 0, Users: 160
2024/12/04 00:12:22 TPS: 91.19, In-flight Retries: 0, Users: 160
2024/12/04 00:12:22 Current configuration stable with avg TPS: 88.69
2024/12/04 00:12:23 TPS: 91.16, In-flight Retries: 0, Users: 160
2024/12/04 00:12:23 Current configuration stable with avg TPS: 90.52
2024/12/04 00:12:24 TPS: 91.28, In-flight Retries: 0, Users: 160
2024/12/04 00:12:24 Current configuration stable with avg TPS: 91.21
2024/12/04 00:12:25 TPS: 89.62, In-flight Retries: 0, Users: 160
2024/12/04 00:12:25 Current configuration stable with avg TPS: 90.68
2024/12/04 00:12:26 TPS: 89.66, In-flight Retries: 0, Users: 160
2024/12/04 00:12:26 Current configuration stable with avg TPS: 90.18
2024/12/04 00:12:27 TPS: 89.79, In-flight Retries: 0, Users: 160
2024/12/04 00:12:27 Current configuration stable with avg TPS: 89.69
2024/12/04 00:12:27 Test duration complete
2024/12/04 00:12:27 Cleanup: Closing stop channel
2024/12/04 00:12:27 Cleanup: Waiting for workers with timeout
2024/12/04 00:12:32 Test iteration cleanup timeout, forcing continuation
2024/12/04 00:12:32 Starting cooldown
2024/12/04 00:12:37 Cooldown complete
2024/12/04 00:12:37 Test was stable with TPS 89.69
2024/12/04 00:12:37 New best configuration found!
2024/12/04 00:12:37 Aggressively doubling users from 160 to 320
2024/12/04 00:12:37 Starting next iteration
2024/12/04 00:12:37 Starting new test iteration
2024/12/04 00:12:37 Testing with 320 users (range: 161-640)
2024/12/04 00:12:37 Starting workers...
2024/12/04 00:12:37 Workers started
2024/12/04 00:12:37 Cleanup: Worker wait timeout, continuing anyway
2024/12/04 00:12:38 TPS: 0.00, In-flight Retries: 50, Users: 320
2024/12/04 00:12:38 Detected retries, marking as unstable
2024/12/04 00:12:38 Cleanup: Closing stop channel
2024/12/04 00:12:38 Cleanup: Waiting for workers with timeout
2024/12/04 00:12:48 Cleanup: Worker wait timeout, continuing anyway
2024/12/04 00:12:48 Test iteration cleanup completed
2024/12/04 00:12:48 Starting cooldown
2024/12/04 00:12:53 Cooldown complete
2024/12/04 00:12:53 Test was unstable, adjusting load
2024/12/04 00:12:53 Reducing users by 25% to 240
2024/12/04 00:12:53 Search complete - tested up to 320 users (2x best stable) or hit max unstable attempts
2024/12/04 00:12:53 Final optimal configuration: 160 users, 89.69 TPS
2024/12/04 00:12:53 Benchmark complete, generating report
2024/12/04 00:12:53 Report generated: reports/benchmark-report-2024-12-04-00-12-53.html
(base) miguel_lemos@Mac restate-recipe-benchmark % go run benchmark.go
# command-line-arguments
ld: warning: ignoring duplicate libraries: '-lrocksdb'
Starting benchmark...
2024/12/04 00:17:13 Monitoring Restate data at: /Users/miguel_lemos/Desktop/restate-recipe-benchmark/restate-data
2024/12/04 00:17:13 Initial disk usage: 878493245440 bytes
2024/12/04 00:17:13 Starting test iteration with 20 users
2024/12/04 00:17:14 TPS: 0.00, In-flight Retries: 0, Users: 20
2024/12/04 00:17:15 TPS: 9.00, In-flight Retries: 0, Users: 20
2024/12/04 00:17:16 TPS: 11.00, In-flight Retries: 0, Users: 20
2024/12/04 00:17:17 TPS: 10.75, In-flight Retries: 0, Users: 20
2024/12/04 00:17:18 TPS: 10.20, In-flight Retries: 0, Users: 20
2024/12/04 00:17:19 TPS: 10.83, In-flight Retries: 0, Users: 20
2024/12/04 00:17:20 TPS: 11.28, In-flight Retries: 0, Users: 20
2024/12/04 00:17:21 TPS: 11.87, In-flight Retries: 0, Users: 20
2024/12/04 00:17:22 TPS: 11.55, In-flight Retries: 0, Users: 20
2024/12/04 00:17:23 TPS: 11.80, In-flight Retries: 0, Users: 20
2024/12/04 00:17:33 Test was stable with TPS 11.74
2024/12/04 00:17:33 New best configuration found!
2024/12/04 00:17:33 Doubling users from 20 to 40
2024/12/04 00:17:33 Starting test iteration with 40 users
2024/12/04 00:17:34 TPS: 0.00, In-flight Retries: 0, Users: 40
2024/12/04 00:17:35 TPS: 17.98, In-flight Retries: 0, Users: 40
2024/12/04 00:17:36 TPS: 22.32, In-flight Retries: 0, Users: 40
2024/12/04 00:17:37 TPS: 23.74, In-flight Retries: 0, Users: 40
2024/12/04 00:17:38 TPS: 24.19, In-flight Retries: 0, Users: 40
2024/12/04 00:17:39 TPS: 24.33, In-flight Retries: 0, Users: 40
2024/12/04 00:17:40 TPS: 24.28, In-flight Retries: 0, Users: 40
2024/12/04 00:17:41 TPS: 24.37, In-flight Retries: 0, Users: 40
2024/12/04 00:17:42 TPS: 24.88, In-flight Retries: 0, Users: 40
2024/12/04 00:17:43 TPS: 24.60, In-flight Retries: 0, Users: 40
2024/12/04 00:17:53 Worker cleanup timed out
2024/12/04 00:17:58 Test was stable with TPS 24.62
2024/12/04 00:17:58 New best configuration found!
2024/12/04 00:17:58 Doubling users from 40 to 80
2024/12/04 00:17:58 Starting test iteration with 80 users
2024/12/04 00:17:59 TPS: 0.00, In-flight Retries: 0, Users: 80
2024/12/04 00:18:00 TPS: 36.97, In-flight Retries: 0, Users: 80
2024/12/04 00:18:01 TPS: 44.32, In-flight Retries: 0, Users: 80
2024/12/04 00:18:02 TPS: 45.98, In-flight Retries: 0, Users: 80
2024/12/04 00:18:03 TPS: 46.98, In-flight Retries: 0, Users: 80
2024/12/04 00:18:04 TPS: 47.65, In-flight Retries: 0, Users: 80
2024/12/04 00:18:05 TPS: 47.85, In-flight Retries: 0, Users: 80
2024/12/04 00:18:06 TPS: 47.50, In-flight Retries: 0, Users: 80
2024/12/04 00:18:07 TPS: 46.55, In-flight Retries: 0, Users: 80
2024/12/04 00:18:08 TPS: 46.49, In-flight Retries: 0, Users: 80
2024/12/04 00:18:18 Worker cleanup timed out
2024/12/04 00:18:23 Test was stable with TPS 46.85
2024/12/04 00:18:23 New best configuration found!
2024/12/04 00:18:23 Doubling users from 80 to 160
2024/12/04 00:18:23 Starting test iteration with 160 users
2024/12/04 00:18:24 TPS: 1.00, In-flight Retries: 0, Users: 160
2024/12/04 00:18:25 TPS: 69.98, In-flight Retries: 0, Users: 160
2024/12/04 00:18:26 TPS: 86.30, In-flight Retries: 0, Users: 160
2024/12/04 00:18:27 TPS: 92.47, In-flight Retries: 0, Users: 160
2024/12/04 00:18:28 TPS: 92.79, In-flight Retries: 0, Users: 160
2024/12/04 00:18:29 TPS: 91.47, In-flight Retries: 0, Users: 160
2024/12/04 00:18:30 TPS: 91.84, In-flight Retries: 0, Users: 160
2024/12/04 00:18:31 TPS: 90.73, In-flight Retries: 0, Users: 160
2024/12/04 00:18:32 TPS: 91.32, In-flight Retries: 0, Users: 160
2024/12/04 00:18:33 TPS: 90.69, In-flight Retries: 0, Users: 160
2024/12/04 00:18:43 Worker cleanup timed out
2024/12/04 00:18:48 Test was stable with TPS 90.91
2024/12/04 00:18:48 New best configuration found!
2024/12/04 00:18:48 Doubling users from 160 to 320
2024/12/04 00:18:48 Starting test iteration with 320 users
2024/12/04 00:18:49 TPS: 1.00, In-flight Retries: 94, Users: 320
2024/12/04 00:18:59 Worker cleanup timed out
2024/12/04 00:19:05 Test was unstable
2024/12/04 00:19:05 Found bounds: 160 (stable) to 320 (unstable)
2024/12/04 00:19:05 Starting fine-tuning phase between 160 and 320 users
2024/12/04 00:19:05 Fine-tuning attempt 1/2 with 240 users
2024/12/04 00:19:06 Fine-tuning TPS: 0.99, Retries: 8, Users: 240
2024/12/04 00:19:16 Fine-tuning worker cleanup timed out
2024/12/04 00:19:21 Fine-tuning attempt 2/2 with 200 users
2024/12/04 00:19:22 Fine-tuning TPS: 0.00, Retries: 0, Users: 200
2024/12/04 00:19:23 Fine-tuning TPS: 85.87, Retries: 0, Users: 200
2024/12/04 00:19:24 Fine-tuning TPS: 106.59, Retries: 0, Users: 200
2024/12/04 00:19:25 Fine-tuning TPS: 111.93, Retries: 0, Users: 200
2024/12/04 00:19:26 Fine-tuning TPS: 113.34, Retries: 0, Users: 200
2024/12/04 00:19:27 Fine-tuning TPS: 111.96, Retries: 0, Users: 200
2024/12/04 00:19:28 Fine-tuning TPS: 112.53, Retries: 0, Users: 200
2024/12/04 00:19:29 Fine-tuning TPS: 113.46, Retries: 0, Users: 200
2024/12/04 00:19:30 Fine-tuning TPS: 114.07, Retries: 0, Users: 200
2024/12/04 00:19:31 Fine-tuning TPS: 115.17, Retries: 0, Users: 200
2024/12/04 00:19:32 Fine-tuning TPS: 115.97, Retries: 0, Users: 200
2024/12/04 00:19:33 Fine-tuning TPS: 115.89, Retries: 0, Users: 200
2024/12/04 00:19:34 Fine-tuning TPS: 115.67, Retries: 0, Users: 200
2024/12/04 00:19:35 Fine-tuning TPS: 115.48, Retries: 0, Users: 200
2024/12/04 00:19:36 Fine-tuning TPS: 114.71, Retries: 0, Users: 200
2024/12/04 00:19:37 Fine-tuning TPS: 114.17, Retries: 0, Users: 200
2024/12/04 00:19:38 Fine-tuning TPS: 113.98, Retries: 0, Users: 200
2024/12/04 00:19:39 Fine-tuning TPS: 113.21, Retries: 0, Users: 200
2024/12/04 00:19:40 Fine-tuning TPS: 112.93, Retries: 0, Users: 200
2024/12/04 00:19:41 Fine-tuning TPS: 112.14, Retries: 0, Users: 200
2024/12/04 00:19:42 Fine-tuning TPS: 111.65, Retries: 0, Users: 200
2024/12/04 00:19:43 Fine-tuning TPS: 111.08, Retries: 0, Users: 200
2024/12/04 00:19:44 Fine-tuning TPS: 110.99, Retries: 0, Users: 200
2024/12/04 00:19:45 Fine-tuning TPS: 110.03, Retries: 0, Users: 200
2024/12/04 00:19:46 Fine-tuning TPS: 109.83, Retries: 0, Users: 200
2024/12/04 00:19:47 Fine-tuning TPS: 109.53, Retries: 0, Users: 200
2024/12/04 00:19:48 Fine-tuning TPS: 109.07, Retries: 0, Users: 200
2024/12/04 00:19:49 Fine-tuning TPS: 108.56, Retries: 0, Users: 200
2024/12/04 00:19:50 Fine-tuning TPS: 108.34, Retries: 0, Users: 200
2024/12/04 00:19:51 Fine-tuning TPS: 108.09, Retries: 0, Users: 200
2024/12/04 00:20:01 Fine-tuning worker cleanup timed out
2024/12/04 00:20:06 New best configuration in fine-tuning: 200 users at 108.33 TPS
2024/12/04 00:20:06 Running final 5-minute stability test with 200 users
2024/12/04 00:20:07 Final stability test - TPS: 5.99, Retries: 0
2024/12/04 00:20:08 Final stability test - TPS: 87.92, Retries: 0
2024/12/04 00:20:09 Final stability test - TPS: 102.97, Retries: 0
2024/12/04 00:20:10 Final stability test - TPS: 107.95, Retries: 0
2024/12/04 00:20:11 Final stability test - TPS: 108.77, Retries: 0
2024/12/04 00:20:12 Final stability test - TPS: 111.31, Retries: 0
2024/12/04 00:20:13 Final stability test - TPS: 110.56, Retries: 0
2024/12/04 00:20:14 Final stability test - TPS: 109.73, Retries: 0
2024/12/04 00:20:15 Final stability test - TPS: 109.32, Retries: 0
2024/12/04 00:20:16 Final stability test - TPS: 109.48, Retries: 0
2024/12/04 00:20:17 Final stability test - TPS: 109.35, Retries: 0
2024/12/04 00:20:18 Final stability test - TPS: 109.41, Retries: 0
2024/12/04 00:20:19 Final stability test - TPS: 109.99, Retries: 0
2024/12/04 00:20:20 Final stability test - TPS: 109.99, Retries: 0
2024/12/04 00:20:21 Final stability test - TPS: 109.45, Retries: 0
2024/12/04 00:20:22 Final stability test - TPS: 109.37, Retries: 0
2024/12/04 00:20:23 Final stability test - TPS: 109.17, Retries: 0
2024/12/04 00:20:24 Final stability test - TPS: 108.49, Retries: 0
2024/12/04 00:20:25 Final stability test - TPS: 108.26, Retries: 0
2024/12/04 00:20:26 Final stability test - TPS: 107.49, Retries: 0
2024/12/04 00:20:27 Final stability test - TPS: 107.28, Retries: 0
2024/12/04 00:20:28 Final stability test - TPS: 107.08, Retries: 0
2024/12/04 00:20:29 Final stability test - TPS: 106.77, Retries: 0
2024/12/04 00:20:30 Final stability test - TPS: 106.37, Retries: 0
2024/12/04 00:20:31 Final stability test - TPS: 105.67, Retries: 0
2024/12/04 00:20:32 Final stability test - TPS: 105.11, Retries: 0
2024/12/04 00:20:33 Final stability test - TPS: 105.26, Retries: 0
2024/12/04 00:20:34 Final stability test - TPS: 104.99, Retries: 0
2024/12/04 00:20:35 Final stability test - TPS: 104.51, Retries: 0
2024/12/04 00:20:36 Final stability test - TPS: 104.29, Retries: 0
2024/12/04 00:20:37 Final stability test - TPS: 104.29, Retries: 0
2024/12/04 00:20:38 Final stability test - TPS: 104.09, Retries: 0
2024/12/04 00:20:39 Final stability test - TPS: 103.90, Retries: 0
2024/12/04 00:20:40 Final stability test - TPS: 103.82, Retries: 0
2024/12/04 00:20:41 Final stability test - TPS: 103.80, Retries: 0
2024/12/04 00:20:42 Final stability test - TPS: 103.58, Retries: 0
2024/12/04 00:20:43 Final stability test - TPS: 103.43, Retries: 0
2024/12/04 00:20:44 Final stability test - TPS: 103.26, Retries: 0
2024/12/04 00:20:45 Final stability test - TPS: 103.18, Retries: 0
2024/12/04 00:20:46 Final stability test - TPS: 103.20, Retries: 0
2024/12/04 00:20:47 Final stability test - TPS: 103.02, Retries: 0
2024/12/04 00:20:48 Final stability test - TPS: 102.97, Retries: 0
2024/12/04 00:20:49 Final stability test - TPS: 102.76, Retries: 0
2024/12/04 00:20:50 Final stability test - TPS: 102.56, Retries: 0
2024/12/04 00:20:51 Final stability test - TPS: 102.44, Retries: 0
2024/12/04 00:20:52 Final stability test - TPS: 102.37, Retries: 0
2024/12/04 00:20:53 Final stability test - TPS: 102.25, Retries: 0
2024/12/04 00:20:54 Final stability test - TPS: 102.18, Retries: 0
2024/12/04 00:20:55 Final stability test - TPS: 101.88, Retries: 0
2024/12/04 00:20:56 Final stability test - TPS: 101.62, Retries: 0
2024/12/04 00:20:57 Final stability test - TPS: 101.59, Retries: 0
2024/12/04 00:20:58 Final stability test - TPS: 101.46, Retries: 0
2024/12/04 00:20:59 Final stability test - TPS: 101.60, Retries: 0
2024/12/04 00:21:00 Final stability test - TPS: 101.66, Retries: 0
2024/12/04 00:21:01 Final stability test - TPS: 101.63, Retries: 0
2024/12/04 00:21:02 Final stability test - TPS: 101.68, Retries: 0
2024/12/04 00:21:03 Final stability test - TPS: 101.44, Retries: 0
2024/12/04 00:21:04 Final stability test - TPS: 101.50, Retries: 0
2024/12/04 00:21:05 Final stability test - TPS: 101.35, Retries: 0
2024/12/04 00:21:06 Final stability test - TPS: 101.25, Retries: 0
2024/12/04 00:21:07 Final stability test - TPS: 101.15, Retries: 0
2024/12/04 00:21:08 Final stability test - TPS: 100.90, Retries: 0
2024/12/04 00:21:09 Final stability test - TPS: 100.66, Retries: 0
2024/12/04 00:21:10 Final stability test - TPS: 100.39, Retries: 0
2024/12/04 00:21:11 Final stability test - TPS: 100.23, Retries: 0
2024/12/04 00:21:12 Final stability test - TPS: 99.98, Retries: 0
2024/12/04 00:21:13 Final stability test - TPS: 99.79, Retries: 0
2024/12/04 00:21:14 Final stability test - TPS: 99.48, Retries: 0
2024/12/04 00:21:15 Final stability test - TPS: 99.42, Retries: 0
2024/12/04 00:21:16 Final stability test - TPS: 99.26, Retries: 0
2024/12/04 00:21:17 Final stability test - TPS: 99.24, Retries: 0
2024/12/04 00:21:18 Final stability test - TPS: 99.19, Retries: 0
2024/12/04 00:21:19 Final stability test - TPS: 99.00, Retries: 0
2024/12/04 00:21:20 Final stability test - TPS: 98.73, Retries: 0
2024/12/04 00:21:21 Final stability test - TPS: 98.57, Retries: 0
2024/12/04 00:21:22 Final stability test - TPS: 98.49, Retries: 0
2024/12/04 00:21:23 Final stability test - TPS: 98.38, Retries: 0
2024/12/04 00:21:24 Final stability test - TPS: 98.36, Retries: 0
2024/12/04 00:21:25 Final stability test - TPS: 98.26, Retries: 0
2024/12/04 00:21:26 Final stability test - TPS: 98.11, Retries: 0
2024/12/04 00:21:27 Final stability test - TPS: 97.90, Retries: 0
2024/12/04 00:21:28 Final stability test - TPS: 97.73, Retries: 0
2024/12/04 00:21:29 Final stability test - TPS: 97.56, Retries: 0
2024/12/04 00:21:30 Final stability test - TPS: 97.59, Retries: 0
2024/12/04 00:21:31 Final stability test - TPS: 97.47, Retries: 0
2024/12/04 00:21:32 Final stability test - TPS: 97.45, Retries: 0
2024/12/04 00:21:33 Final stability test - TPS: 97.34, Retries: 0
2024/12/04 00:21:34 Final stability test - TPS: 97.29, Retries: 0
2024/12/04 00:21:35 Final stability test - TPS: 97.29, Retries: 0
2024/12/04 00:21:36 Final stability test - TPS: 97.32, Retries: 0
2024/12/04 00:21:37 Final stability test - TPS: 97.40, Retries: 0
2024/12/04 00:21:38 Final stability test - TPS: 97.47, Retries: 0
2024/12/04 00:21:39 Final stability test - TPS: 97.42, Retries: 0
2024/12/04 00:21:40 Final stability test - TPS: 97.41, Retries: 0
2024/12/04 00:21:41 Final stability test - TPS: 97.46, Retries: 0
2024/12/04 00:21:42 Final stability test - TPS: 97.54, Retries: 0
2024/12/04 00:21:43 Final stability test - TPS: 97.62, Retries: 0
2024/12/04 00:21:44 Final stability test - TPS: 97.71, Retries: 0
2024/12/04 00:21:45 Final stability test - TPS: 97.70, Retries: 0
2024/12/04 00:21:46 Final stability test - TPS: 97.69, Retries: 0
2024/12/04 00:21:47 Final stability test - TPS: 97.71, Retries: 0
2024/12/04 00:21:48 Final stability test - TPS: 97.75, Retries: 0
2024/12/04 00:21:49 Final stability test - TPS: 97.87, Retries: 0
2024/12/04 00:21:50 Final stability test - TPS: 97.89, Retries: 0
2024/12/04 00:21:51 Final stability test - TPS: 97.88, Retries: 0
2024/12/04 00:21:52 Final stability test - TPS: 97.90, Retries: 0
2024/12/04 00:21:53 Final stability test - TPS: 97.93, Retries: 0
2024/12/04 00:21:54 Final stability test - TPS: 97.84, Retries: 0
2024/12/04 00:21:55 Final stability test - TPS: 97.89, Retries: 0
2024/12/04 00:21:56 Final stability test - TPS: 97.95, Retries: 0
2024/12/04 00:21:57 Final stability test - TPS: 98.01, Retries: 0
2024/12/04 00:21:58 Final stability test - TPS: 97.96, Retries: 0
2024/12/04 00:21:59 Final stability test - TPS: 97.93, Retries: 0
2024/12/04 00:22:00 Final stability test - TPS: 97.96, Retries: 0
2024/12/04 00:22:01 Final stability test - TPS: 97.88, Retries: 0
2024/12/04 00:22:02 Final stability test - TPS: 97.95, Retries: 0
2024/12/04 00:22:03 Final stability test - TPS: 97.99, Retries: 0
2024/12/04 00:22:04 Final stability test - TPS: 97.96, Retries: 0
2024/12/04 00:22:05 Final stability test - TPS: 97.92, Retries: 0
2024/12/04 00:22:06 Final stability test - TPS: 97.98, Retries: 0
2024/12/04 00:22:07 Final stability test - TPS: 97.92, Retries: 0
2024/12/04 00:22:08 Final stability test - TPS: 97.95, Retries: 0
2024/12/04 00:22:09 Final stability test - TPS: 98.00, Retries: 0
2024/12/04 00:22:10 Final stability test - TPS: 98.04, Retries: 0
2024/12/04 00:22:11 Final stability test - TPS: 98.04, Retries: 0
2024/12/04 00:22:12 Final stability test - TPS: 98.05, Retries: 0
2024/12/04 00:22:13 Final stability test - TPS: 98.18, Retries: 0
2024/12/04 00:22:14 Final stability test - TPS: 98.22, Retries: 0
2024/12/04 00:22:15 Final stability test - TPS: 98.18, Retries: 0
2024/12/04 00:22:16 Final stability test - TPS: 98.13, Retries: 0
2024/12/04 00:22:17 Final stability test - TPS: 98.14, Retries: 0
2024/12/04 00:22:18 Final stability test - TPS: 98.17, Retries: 0
2024/12/04 00:22:19 Final stability test - TPS: 98.07, Retries: 0
2024/12/04 00:22:20 Final stability test - TPS: 98.07, Retries: 0
2024/12/04 00:22:21 Final stability test - TPS: 98.12, Retries: 0
2024/12/04 00:22:22 Final stability test - TPS: 98.05, Retries: 0
2024/12/04 00:22:23 Final stability test - TPS: 98.06, Retries: 0
2024/12/04 00:22:24 Final stability test - TPS: 97.98, Retries: 0
2024/12/04 00:22:25 Final stability test - TPS: 97.88, Retries: 0
2024/12/04 00:22:26 Final stability test - TPS: 97.89, Retries: 0
2024/12/04 00:22:27 Final stability test - TPS: 97.80, Retries: 0
2024/12/04 00:22:28 Final stability test - TPS: 97.76, Retries: 0
2024/12/04 00:22:29 Final stability test - TPS: 97.68, Retries: 0
2024/12/04 00:22:30 Final stability test - TPS: 97.59, Retries: 0
2024/12/04 00:22:31 Final stability test - TPS: 97.60, Retries: 0
2024/12/04 00:22:32 Final stability test - TPS: 97.47, Retries: 0
2024/12/04 00:22:33 Final stability test - TPS: 97.41, Retries: 0
2024/12/04 00:22:34 Final stability test - TPS: 97.34, Retries: 0
2024/12/04 00:22:35 Final stability test - TPS: 97.35, Retries: 0
2024/12/04 00:22:36 Final stability test - TPS: 97.31, Retries: 0
2024/12/04 00:22:37 Final stability test - TPS: 97.34, Retries: 0
2024/12/04 00:22:38 Final stability test - TPS: 97.28, Retries: 0
2024/12/04 00:22:39 Final stability test - TPS: 97.26, Retries: 0
2024/12/04 00:22:40 Final stability test - TPS: 97.25, Retries: 0
2024/12/04 00:22:41 Final stability test - TPS: 97.30, Retries: 0
2024/12/04 00:22:42 Final stability test - TPS: 97.29, Retries: 0
2024/12/04 00:22:43 Final stability test - TPS: 97.27, Retries: 0
2024/12/04 00:22:44 Final stability test - TPS: 97.32, Retries: 0
2024/12/04 00:22:45 Final stability test - TPS: 97.34, Retries: 0
2024/12/04 00:22:46 Final stability test - TPS: 97.36, Retries: 0
2024/12/04 00:22:47 Final stability test - TPS: 97.37, Retries: 0
2024/12/04 00:22:48 Final stability test - TPS: 97.39, Retries: 0
2024/12/04 00:22:49 Final stability test - TPS: 97.38, Retries: 0
2024/12/04 00:22:50 Final stability test - TPS: 97.38, Retries: 0
2024/12/04 00:22:51 Final stability test - TPS: 97.38, Retries: 0
2024/12/04 00:22:52 Final stability test - TPS: 97.33, Retries: 0
2024/12/04 00:22:53 Final stability test - TPS: 97.30, Retries: 0
2024/12/04 00:22:54 Final stability test - TPS: 97.36, Retries: 0
2024/12/04 00:22:55 Final stability test - TPS: 97.35, Retries: 0
2024/12/04 00:22:56 Final stability test - TPS: 97.36, Retries: 0
2024/12/04 00:22:57 Final stability test - TPS: 97.37, Retries: 0
2024/12/04 00:22:58 Final stability test - TPS: 97.39, Retries: 0
2024/12/04 00:22:59 Final stability test - TPS: 97.42, Retries: 0
2024/12/04 00:23:00 Final stability test - TPS: 97.36, Retries: 0
2024/12/04 00:23:01 Final stability test - TPS: 97.30, Retries: 0
2024/12/04 00:23:02 Final stability test - TPS: 97.24, Retries: 0
2024/12/04 00:23:03 Final stability test - TPS: 97.21, Retries: 0
2024/12/04 00:23:04 Final stability test - TPS: 97.15, Retries: 0
2024/12/04 00:23:05 Final stability test - TPS: 97.14, Retries: 0
2024/12/04 00:23:06 Final stability test - TPS: 97.19, Retries: 0
2024/12/04 00:23:07 Final stability test - TPS: 97.22, Retries: 0
2024/12/04 00:23:08 Final stability test - TPS: 97.28, Retries: 0
2024/12/04 00:23:09 Final stability test - TPS: 97.29, Retries: 0
2024/12/04 00:23:10 Final stability test - TPS: 97.31, Retries: 0
2024/12/04 00:23:11 Final stability test - TPS: 97.29, Retries: 0
2024/12/04 00:23:12 Final stability test - TPS: 97.33, Retries: 0
2024/12/04 00:23:13 Final stability test - TPS: 97.35, Retries: 0
2024/12/04 00:23:14 Final stability test - TPS: 97.41, Retries: 0
2024/12/04 00:23:15 Final stability test - TPS: 97.44, Retries: 0
2024/12/04 00:23:16 Final stability test - TPS: 97.47, Retries: 0
2024/12/04 00:23:17 Final stability test - TPS: 97.53, Retries: 0
2024/12/04 00:23:18 Final stability test - TPS: 97.52, Retries: 0
2024/12/04 00:23:19 Final stability test - TPS: 97.54, Retries: 0
2024/12/04 00:23:20 Final stability test - TPS: 97.52, Retries: 0
2024/12/04 00:23:21 Final stability test - TPS: 97.50, Retries: 0
2024/12/04 00:23:22 Final stability test - TPS: 97.58, Retries: 0
2024/12/04 00:23:23 Final stability test - TPS: 97.57, Retries: 0
2024/12/04 00:23:24 Final stability test - TPS: 97.58, Retries: 0
2024/12/04 00:23:25 Final stability test - TPS: 97.55, Retries: 0
2024/12/04 00:23:26 Final stability test - TPS: 97.53, Retries: 0
2024/12/04 00:23:27 Final stability test - TPS: 97.61, Retries: 0
2024/12/04 00:23:28 Final stability test - TPS: 97.65, Retries: 0
2024/12/04 00:23:29 Final stability test - TPS: 97.68, Retries: 0
2024/12/04 00:23:30 Final stability test - TPS: 97.70, Retries: 0
2024/12/04 00:23:31 Final stability test - TPS: 97.74, Retries: 0
2024/12/04 00:23:32 Final stability test - TPS: 97.77, Retries: 0
2024/12/04 00:23:33 Final stability test - TPS: 97.83, Retries: 0
2024/12/04 00:23:34 Final stability test - TPS: 97.84, Retries: 0
2024/12/04 00:23:35 Final stability test - TPS: 97.90, Retries: 0
2024/12/04 00:23:36 Final stability test - TPS: 97.95, Retries: 0
2024/12/04 00:23:37 Final stability test - TPS: 97.95, Retries: 0
2024/12/04 00:23:38 Final stability test - TPS: 97.94, Retries: 0
2024/12/04 00:23:39 Final stability test - TPS: 97.92, Retries: 0
2024/12/04 00:23:40 Final stability test - TPS: 97.93, Retries: 0
2024/12/04 00:23:41 Final stability test - TPS: 97.90, Retries: 0
2024/12/04 00:23:42 Final stability test - TPS: 97.87, Retries: 0
2024/12/04 00:23:43 Final stability test - TPS: 97.83, Retries: 0
2024/12/04 00:23:44 Final stability test - TPS: 97.81, Retries: 0
2024/12/04 00:23:45 Final stability test - TPS: 97.79, Retries: 0
2024/12/04 00:23:46 Final stability test - TPS: 97.79, Retries: 0
2024/12/04 00:23:47 Final stability test - TPS: 97.77, Retries: 0
2024/12/04 00:23:48 Final stability test - TPS: 97.71, Retries: 0
2024/12/04 00:23:49 Final stability test - TPS: 97.71, Retries: 0
2024/12/04 00:23:50 Final stability test - TPS: 97.66, Retries: 0
2024/12/04 00:23:51 Final stability test - TPS: 97.66, Retries: 0
2024/12/04 00:23:52 Final stability test - TPS: 97.64, Retries: 0
2024/12/04 00:23:53 Final stability test - TPS: 97.65, Retries: 0
2024/12/04 00:23:54 Final stability test - TPS: 97.62, Retries: 0
2024/12/04 00:23:55 Final stability test - TPS: 97.61, Retries: 0
2024/12/04 00:23:56 Final stability test - TPS: 97.59, Retries: 0
2024/12/04 00:23:57 Final stability test - TPS: 97.57, Retries: 0
2024/12/04 00:23:58 Final stability test - TPS: 97.53, Retries: 0
2024/12/04 00:23:59 Final stability test - TPS: 97.51, Retries: 0
2024/12/04 00:24:00 Final stability test - TPS: 97.47, Retries: 0
2024/12/04 00:24:01 Final stability test - TPS: 97.51, Retries: 0
2024/12/04 00:24:02 Final stability test - TPS: 97.54, Retries: 0
2024/12/04 00:24:03 Final stability test - TPS: 97.53, Retries: 0
2024/12/04 00:24:04 Final stability test - TPS: 97.52, Retries: 0
2024/12/04 00:24:05 Final stability test - TPS: 97.52, Retries: 0
2024/12/04 00:24:06 Final stability test - TPS: 97.51, Retries: 0
2024/12/04 00:24:07 Final stability test - TPS: 97.50, Retries: 0
2024/12/04 00:24:08 Final stability test - TPS: 97.44, Retries: 0
2024/12/04 00:24:09 Final stability test - TPS: 97.42, Retries: 0
2024/12/04 00:24:10 Final stability test - TPS: 97.43, Retries: 0
2024/12/04 00:24:11 Final stability test - TPS: 97.41, Retries: 0
2024/12/04 00:24:12 Final stability test - TPS: 97.44, Retries: 0
2024/12/04 00:24:13 Final stability test - TPS: 97.47, Retries: 0
2024/12/04 00:24:14 Final stability test - TPS: 97.41, Retries: 0
2024/12/04 00:24:15 Final stability test - TPS: 97.35, Retries: 0
2024/12/04 00:24:16 Final stability test - TPS: 97.32, Retries: 0
2024/12/04 00:24:17 Final stability test - TPS: 97.29, Retries: 0
2024/12/04 00:24:18 Final stability test - TPS: 97.25, Retries: 0
2024/12/04 00:24:19 Final stability test - TPS: 97.20, Retries: 0
2024/12/04 00:24:20 Final stability test - TPS: 97.16, Retries: 0
2024/12/04 00:24:21 Final stability test - TPS: 97.15, Retries: 0
2024/12/04 00:24:22 Final stability test - TPS: 97.16, Retries: 0
2024/12/04 00:24:23 Final stability test - TPS: 97.14, Retries: 0
2024/12/04 00:24:24 Final stability test - TPS: 97.08, Retries: 0
2024/12/04 00:24:25 Final stability test - TPS: 97.03, Retries: 0
2024/12/04 00:24:26 Final stability test - TPS: 97.01, Retries: 0
2024/12/04 00:24:27 Final stability test - TPS: 96.99, Retries: 0
2024/12/04 00:24:28 Final stability test - TPS: 97.01, Retries: 0
2024/12/04 00:24:29 Final stability test - TPS: 97.03, Retries: 0
2024/12/04 00:24:30 Final stability test - TPS: 97.00, Retries: 0
2024/12/04 00:24:31 Final stability test - TPS: 97.03, Retries: 0
2024/12/04 00:24:32 Final stability test - TPS: 97.01, Retries: 0
2024/12/04 00:24:33 Final stability test - TPS: 97.00, Retries: 0
2024/12/04 00:24:34 Final stability test - TPS: 96.98, Retries: 0
2024/12/04 00:24:35 Final stability test - TPS: 96.98, Retries: 0
2024/12/04 00:24:36 Final stability test - TPS: 97.00, Retries: 0
2024/12/04 00:24:37 Final stability test - TPS: 96.99, Retries: 0
2024/12/04 00:24:38 Final stability test - TPS: 96.96, Retries: 0
2024/12/04 00:24:39 Final stability test - TPS: 96.95, Retries: 0
2024/12/04 00:24:40 Final stability test - TPS: 96.92, Retries: 0
2024/12/04 00:24:41 Final stability test - TPS: 96.92, Retries: 0
2024/12/04 00:24:42 Final stability test - TPS: 96.89, Retries: 0
2024/12/04 00:24:43 Final stability test - TPS: 96.87, Retries: 0
2024/12/04 00:24:44 Final stability test - TPS: 96.87, Retries: 0
2024/12/04 00:24:45 Final stability test - TPS: 96.90, Retries: 0
2024/12/04 00:24:46 Final stability test - TPS: 96.94, Retries: 0
2024/12/04 00:24:47 Final stability test - TPS: 96.91, Retries: 0
2024/12/04 00:24:48 Final stability test - TPS: 96.93, Retries: 0
2024/12/04 00:24:49 Final stability test - TPS: 96.94, Retries: 0
2024/12/04 00:24:50 Final stability test - TPS: 96.97, Retries: 0
2024/12/04 00:24:51 Final stability test - TPS: 96.97, Retries: 0
2024/12/04 00:24:52 Final stability test - TPS: 96.96, Retries: 0
2024/12/04 00:24:53 Final stability test - TPS: 96.99, Retries: 0
2024/12/04 00:24:54 Final stability test - TPS: 96.98, Retries: 0
2024/12/04 00:24:55 Final stability test - TPS: 96.99, Retries: 0
2024/12/04 00:24:56 Final stability test - TPS: 97.00, Retries: 0
2024/12/04 00:24:57 Final stability test - TPS: 97.00, Retries: 0
2024/12/04 00:24:58 Final stability test - TPS: 97.00, Retries: 0
2024/12/04 00:24:59 Final stability test - TPS: 96.99, Retries: 0
2024/12/04 00:25:00 Final stability test - TPS: 97.01, Retries: 0
2024/12/04 00:25:01 Final stability test - TPS: 96.99, Retries: 0
2024/12/04 00:25:02 Final stability test - TPS: 96.98, Retries: 0
2024/12/04 00:25:03 Final stability test - TPS: 97.01, Retries: 0
2024/12/04 00:25:04 Final stability test - TPS: 97.00, Retries: 0
2024/12/04 00:25:05 Final stability test - TPS: 97.02, Retries: 0
2024/12/04 00:25:06 Final stability test - TPS: 97.03, Retries: 0
2024/12/04 00:26:05 Benchmark complete. Final configuration: 200 users, 108.33 TPS
2024/12/04 00:26:05 Report generated: reports/benchmark-report-2024-12-04-00-26-05.html
(base) miguel_lemos@Mac restate-recipe-benchmark %
```