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
2024/12/03 21:57:04 Monitoring Restate data at: /Users/miguel_lemos/Desktop/restate-recipe-benchmark/restate-data
2024/12/03 21:57:04 Initial disk usage: 878547030016 bytes
2024/12/03 21:57:04 Testing with 20 users
2024/12/03 21:57:05 TPS: 0.00, In-flight Retries: 0, Users: 20
2024/12/03 21:57:06 TPS: 8.99, In-flight Retries: 0, Users: 20
2024/12/03 21:57:07 TPS: 10.99, In-flight Retries: 0, Users: 20
2024/12/03 21:57:08 TPS: 11.50, In-flight Retries: 0, Users: 20
2024/12/03 21:57:09 TPS: 11.40, In-flight Retries: 0, Users: 20
2024/12/03 21:57:10 TPS: 11.50, In-flight Retries: 0, Users: 20
2024/12/03 21:57:11 TPS: 11.57, In-flight Retries: 0, Users: 20
2024/12/03 21:57:12 TPS: 11.75, In-flight Retries: 0, Users: 20
2024/12/03 21:57:13 TPS: 12.00, In-flight Retries: 0, Users: 20
2024/12/03 21:57:14 TPS: 11.90, In-flight Retries: 0, Users: 20
2024/12/03 21:57:15 TPS: 12.00, In-flight Retries: 0, Users: 20
2024/12/03 21:57:16 TPS: 12.17, In-flight Retries: 0, Users: 20
2024/12/03 21:57:17 TPS: 12.15, In-flight Retries: 0, Users: 20
2024/12/03 21:57:18 TPS: 12.14, In-flight Retries: 0, Users: 20
2024/12/03 21:57:19 TPS: 12.20, In-flight Retries: 0, Users: 20
2024/12/03 21:57:20 TPS: 12.19, In-flight Retries: 0, Users: 20
2024/12/03 21:57:21 TPS: 12.12, In-flight Retries: 0, Users: 20
2024/12/03 21:57:22 TPS: 12.28, In-flight Retries: 0, Users: 20
2024/12/03 21:57:23 TPS: 12.32, In-flight Retries: 0, Users: 20
2024/12/03 21:57:24 TPS: 12.35, In-flight Retries: 0, Users: 20
2024/12/03 21:57:25 TPS: 12.09, In-flight Retries: 0, Users: 20
2024/12/03 21:57:26 TPS: 12.14, In-flight Retries: 0, Users: 20
2024/12/03 21:57:27 TPS: 12.22, In-flight Retries: 0, Users: 20
2024/12/03 21:57:28 TPS: 12.04, In-flight Retries: 0, Users: 20
2024/12/03 21:57:29 TPS: 12.00, In-flight Retries: 0, Users: 20
2024/12/03 21:57:30 TPS: 11.96, In-flight Retries: 0, Users: 20
2024/12/03 21:57:31 TPS: 11.85, In-flight Retries: 0, Users: 20
2024/12/03 21:57:32 TPS: 11.89, In-flight Retries: 0, Users: 20
2024/12/03 21:57:33 TPS: 11.86, In-flight Retries: 0, Users: 20
2024/12/03 21:57:34 TPS: 11.87, In-flight Retries: 0, Users: 20
2024/12/03 21:57:35 TPS: 11.74, In-flight Retries: 0, Users: 20
2024/12/03 21:57:36 TPS: 11.78, In-flight Retries: 0, Users: 20
2024/12/03 21:57:37 TPS: 11.88, In-flight Retries: 0, Users: 20
2024/12/03 21:57:38 TPS: 11.79, In-flight Retries: 0, Users: 20
2024/12/03 21:57:39 TPS: 11.86, In-flight Retries: 0, Users: 20
2024/12/03 21:57:40 TPS: 11.92, In-flight Retries: 0, Users: 20
2024/12/03 21:57:41 TPS: 12.08, In-flight Retries: 0, Users: 20
2024/12/03 21:57:42 TPS: 12.13, In-flight Retries: 0, Users: 20
2024/12/03 21:57:43 TPS: 12.20, In-flight Retries: 0, Users: 20
2024/12/03 21:57:44 TPS: 12.32, In-flight Retries: 0, Users: 20
2024/12/03 21:57:45 TPS: 12.34, In-flight Retries: 0, Users: 20
2024/12/03 21:57:46 TPS: 12.43, In-flight Retries: 0, Users: 20
2024/12/03 21:57:47 TPS: 12.37, In-flight Retries: 0, Users: 20
2024/12/03 21:57:48 TPS: 12.39, In-flight Retries: 1, Users: 20
2024/12/03 21:57:48 Detected retries, scaling down to 18 users
2024/12/03 21:58:22 Testing with 18 users
2024/12/03 21:58:22 TPS: 11.01, In-flight Retries: 0, Users: 18
2024/12/03 21:58:23 TPS: 10.77, In-flight Retries: 0, Users: 18
2024/12/03 21:58:24 TPS: 10.92, In-flight Retries: 0, Users: 18
2024/12/03 21:58:25 TPS: 10.97, In-flight Retries: 0, Users: 18
2024/12/03 21:58:26 TPS: 11.03, In-flight Retries: 0, Users: 18
2024/12/03 21:58:27 TPS: 11.05, In-flight Retries: 0, Users: 18
2024/12/03 21:58:28 TPS: 11.07, In-flight Retries: 0, Users: 18
2024/12/03 21:58:29 TPS: 11.20, In-flight Retries: 0, Users: 18
2024/12/03 21:58:30 TPS: 11.19, In-flight Retries: 0, Users: 18
2024/12/03 21:58:31 TPS: 11.19, In-flight Retries: 0, Users: 18
2024/12/03 21:58:32 TPS: 11.20, In-flight Retries: 0, Users: 18
2024/12/03 21:58:33 TPS: 11.22, In-flight Retries: 0, Users: 18
2024/12/03 21:58:34 TPS: 11.22, In-flight Retries: 0, Users: 18
2024/12/03 21:58:35 TPS: 11.19, In-flight Retries: 0, Users: 18
2024/12/03 21:58:36 TPS: 11.13, In-flight Retries: 0, Users: 18
2024/12/03 21:58:37 TPS: 11.06, In-flight Retries: 0, Users: 18
2024/12/03 21:58:38 TPS: 11.06, In-flight Retries: 0, Users: 18
2024/12/03 21:58:39 TPS: 11.08, In-flight Retries: 0, Users: 18
2024/12/03 21:58:40 TPS: 11.13, In-flight Retries: 0, Users: 18
2024/12/03 21:58:41 TPS: 11.13, In-flight Retries: 0, Users: 18
2024/12/03 21:58:42 TPS: 11.15, In-flight Retries: 0, Users: 18
2024/12/03 21:58:43 TPS: 11.15, In-flight Retries: 0, Users: 18
2024/12/03 21:58:44 TPS: 11.16, In-flight Retries: 0, Users: 18
2024/12/03 21:58:45 TPS: 11.14, In-flight Retries: 0, Users: 18
2024/12/03 21:58:46 TPS: 11.12, In-flight Retries: 0, Users: 18
2024/12/03 21:58:47 TPS: 11.15, In-flight Retries: 0, Users: 18
2024/12/03 21:58:48 TPS: 11.17, In-flight Retries: 0, Users: 18
2024/12/03 21:58:49 TPS: 11.08, In-flight Retries: 1, Users: 18
2024/12/03 21:58:49 Detected retries, scaling down to 16 users
2024/12/03 21:59:23 Testing with 16 users
2024/12/03 21:59:23 TPS: 10.74, In-flight Retries: 0, Users: 16
2024/12/03 21:59:24 TPS: 10.49, In-flight Retries: 0, Users: 16
2024/12/03 21:59:25 TPS: 10.56, In-flight Retries: 0, Users: 16
2024/12/03 21:59:26 TPS: 10.59, In-flight Retries: 0, Users: 16
2024/12/03 21:59:27 TPS: 10.58, In-flight Retries: 0, Users: 16
2024/12/03 21:59:28 TPS: 10.56, In-flight Retries: 0, Users: 16
2024/12/03 21:59:29 TPS: 10.55, In-flight Retries: 0, Users: 16
2024/12/03 21:59:30 TPS: 10.54, In-flight Retries: 0, Users: 16
2024/12/03 21:59:31 TPS: 10.60, In-flight Retries: 0, Users: 16
2024/12/03 21:59:32 TPS: 10.67, In-flight Retries: 0, Users: 16
2024/12/03 21:59:33 TPS: 10.68, In-flight Retries: 0, Users: 16
2024/12/03 21:59:34 TPS: 10.60, In-flight Retries: 0, Users: 16
2024/12/03 21:59:35 TPS: 10.59, In-flight Retries: 0, Users: 16
2024/12/03 21:59:36 TPS: 10.55, In-flight Retries: 0, Users: 16
2024/12/03 21:59:37 TPS: 10.56, In-flight Retries: 0, Users: 16
2024/12/03 21:59:38 TPS: 10.57, In-flight Retries: 0, Users: 16
2024/12/03 21:59:39 TPS: 10.52, In-flight Retries: 0, Users: 16
2024/12/03 21:59:40 TPS: 10.51, In-flight Retries: 0, Users: 16
2024/12/03 21:59:41 TPS: 10.48, In-flight Retries: 0, Users: 16
2024/12/03 21:59:42 TPS: 10.40, In-flight Retries: 1, Users: 16
2024/12/03 21:59:42 Detected retries, scaling down to 14 users
2024/12/03 22:00:16 Testing with 14 users
2024/12/03 22:00:16 TPS: 8.33, In-flight Retries: 0, Users: 14
2024/12/03 22:00:17 TPS: 8.20, In-flight Retries: 0, Users: 14
2024/12/03 22:00:18 TPS: 8.33, In-flight Retries: 0, Users: 14
2024/12/03 22:00:19 TPS: 8.35, In-flight Retries: 0, Users: 14
2024/12/03 22:00:20 TPS: 8.29, In-flight Retries: 0, Users: 14
2024/12/03 22:00:21 TPS: 8.26, In-flight Retries: 0, Users: 14
2024/12/03 22:00:22 TPS: 8.30, In-flight Retries: 0, Users: 14
2024/12/03 22:00:23 TPS: 8.32, In-flight Retries: 2, Users: 14
2024/12/03 22:00:23 Detected retries, scaling down to 12 users
2024/12/03 22:00:56 Testing with 12 users
2024/12/03 22:00:56 TPS: 7.99, In-flight Retries: 3, Users: 12
2024/12/03 22:00:56 Detected retries, scaling down to 10 users
2024/12/03 22:01:30 Testing with 10 users
2024/12/03 22:01:30 TPS: 6.24, In-flight Retries: 0, Users: 10
2024/12/03 22:01:31 TPS: 6.15, In-flight Retries: 0, Users: 10
2024/12/03 22:01:32 TPS: 6.20, In-flight Retries: 0, Users: 10
2024/12/03 22:01:33 TPS: 6.20, In-flight Retries: 0, Users: 10
2024/12/03 22:01:34 TPS: 6.14, In-flight Retries: 0, Users: 10
2024/12/03 22:01:35 TPS: 6.14, In-flight Retries: 0, Users: 10
2024/12/03 22:01:36 TPS: 6.16, In-flight Retries: 0, Users: 10
2024/12/03 22:01:37 TPS: 6.18, In-flight Retries: 0, Users: 10
2024/12/03 22:01:38 TPS: 6.20, In-flight Retries: 0, Users: 10
2024/12/03 22:01:39 TPS: 6.26, In-flight Retries: 0, Users: 10
2024/12/03 22:01:40 TPS: 6.23, In-flight Retries: 0, Users: 10
2024/12/03 22:01:41 TPS: 6.21, In-flight Retries: 0, Users: 10
2024/12/03 22:01:42 TPS: 6.25, In-flight Retries: 0, Users: 10
2024/12/03 22:01:43 TPS: 6.26, In-flight Retries: 0, Users: 10
2024/12/03 22:01:44 TPS: 6.22, In-flight Retries: 0, Users: 10
2024/12/03 22:01:45 TPS: 6.19, In-flight Retries: 0, Users: 10
2024/12/03 22:01:46 TPS: 6.17, In-flight Retries: 0, Users: 10
2024/12/03 22:01:47 TPS: 6.16, In-flight Retries: 1, Users: 10
2024/12/03 22:01:47 Detected retries, scaling down to 9 users
2024/12/03 22:02:19 Testing with 9 users
2024/12/03 22:02:19 TPS: 5.43, In-flight Retries: 2, Users: 9
2024/12/03 22:02:19 Detected retries, scaling down to 8 users
2024/12/03 22:02:54 Testing with 8 users
2024/12/03 22:02:54 TPS: 4.81, In-flight Retries: 0, Users: 8
2024/12/03 22:02:55 TPS: 4.69, In-flight Retries: 0, Users: 8
2024/12/03 22:02:56 TPS: 4.78, In-flight Retries: 0, Users: 8
2024/12/03 22:02:57 TPS: 4.87, In-flight Retries: 0, Users: 8
2024/12/03 22:02:58 TPS: 4.92, In-flight Retries: 0, Users: 8
2024/12/03 22:02:59 TPS: 4.88, In-flight Retries: 0, Users: 8
2024/12/03 22:03:00 TPS: 4.90, In-flight Retries: 0, Users: 8
2024/12/03 22:03:01 TPS: 4.95, In-flight Retries: 0, Users: 8
2024/12/03 22:03:02 TPS: 4.91, In-flight Retries: 0, Users: 8
2024/12/03 22:03:03 TPS: 4.91, In-flight Retries: 0, Users: 8
2024/12/03 22:03:04 TPS: 4.93, In-flight Retries: 0, Users: 8
2024/12/03 22:03:05 TPS: 4.89, In-flight Retries: 0, Users: 8
2024/12/03 22:03:06 TPS: 4.89, In-flight Retries: 0, Users: 8
2024/12/03 22:03:07 TPS: 4.92, In-flight Retries: 0, Users: 8
2024/12/03 22:03:08 TPS: 4.90, In-flight Retries: 0, Users: 8
2024/12/03 22:03:09 TPS: 4.92, In-flight Retries: 0, Users: 8
2024/12/03 22:03:10 TPS: 4.92, In-flight Retries: 0, Users: 8
2024/12/03 22:03:11 TPS: 4.94, In-flight Retries: 0, Users: 8
2024/12/03 22:03:12 TPS: 4.89, In-flight Retries: 0, Users: 8
2024/12/03 22:03:13 TPS: 4.93, In-flight Retries: 0, Users: 8
2024/12/03 22:03:14 TPS: 4.91, In-flight Retries: 0, Users: 8
2024/12/03 22:03:15 TPS: 4.87, In-flight Retries: 0, Users: 8
2024/12/03 22:03:16 TPS: 4.86, In-flight Retries: 0, Users: 8
2024/12/03 22:03:17 TPS: 4.81, In-flight Retries: 0, Users: 8
2024/12/03 22:03:18 TPS: 4.81, In-flight Retries: 0, Users: 8
2024/12/03 22:03:19 TPS: 4.78, In-flight Retries: 1, Users: 8
2024/12/03 22:03:19 Detected retries, scaling down to 7 users
2024/12/03 22:03:53 Testing with 7 users
2024/12/03 22:03:53 TPS: 4.08, In-flight Retries: 1, Users: 7
2024/12/03 22:03:53 Detected retries, scaling down to 6 users
2024/12/03 22:04:26 Testing with 6 users
2024/12/03 22:04:26 TPS: 4.24, In-flight Retries: 3, Users: 6
2024/12/03 22:04:26 Detected retries, scaling down to 5 users
2024/12/03 22:04:57 Testing with 5 users
2024/12/03 22:04:57 TPS: 3.76, In-flight Retries: 0, Users: 5
2024/12/03 22:04:58 TPS: 3.69, In-flight Retries: 0, Users: 5
2024/12/03 22:04:59 TPS: 3.73, In-flight Retries: 0, Users: 5
2024/12/03 22:05:00 TPS: 3.74, In-flight Retries: 0, Users: 5
2024/12/03 22:05:01 TPS: 3.78, In-flight Retries: 0, Users: 5
2024/12/03 22:05:02 TPS: 3.75, In-flight Retries: 0, Users: 5
2024/12/03 22:05:03 TPS: 3.76, In-flight Retries: 0, Users: 5
2024/12/03 22:05:04 TPS: 3.71, In-flight Retries: 0, Users: 5
2024/12/03 22:05:05 TPS: 3.67, In-flight Retries: 0, Users: 5
2024/12/03 22:05:06 TPS: 3.70, In-flight Retries: 0, Users: 5
2024/12/03 22:05:07 TPS: 3.74, In-flight Retries: 0, Users: 5
2024/12/03 22:05:08 TPS: 3.72, In-flight Retries: 0, Users: 5
2024/12/03 22:05:09 TPS: 3.68, In-flight Retries: 0, Users: 5
2024/12/03 22:05:10 TPS: 3.66, In-flight Retries: 0, Users: 5
2024/12/03 22:05:11 TPS: 3.65, In-flight Retries: 0, Users: 5
2024/12/03 22:05:12 TPS: 3.61, In-flight Retries: 0, Users: 5
2024/12/03 22:05:13 TPS: 3.62, In-flight Retries: 0, Users: 5
2024/12/03 22:05:14 TPS: 3.61, In-flight Retries: 0, Users: 5
2024/12/03 22:05:15 TPS: 3.57, In-flight Retries: 0, Users: 5
2024/12/03 22:05:16 TPS: 3.56, In-flight Retries: 0, Users: 5
2024/12/03 22:05:17 TPS: 3.57, In-flight Retries: 0, Users: 5
2024/12/03 22:05:18 TPS: 3.56, In-flight Retries: 0, Users: 5
2024/12/03 22:05:19 TPS: 3.59, In-flight Retries: 0, Users: 5
2024/12/03 22:05:20 TPS: 3.58, In-flight Retries: 0, Users: 5
2024/12/03 22:05:21 TPS: 3.60, In-flight Retries: 0, Users: 5
2024/12/03 22:05:22 TPS: 3.56, In-flight Retries: 0, Users: 5
2024/12/03 22:05:23 TPS: 3.56, In-flight Retries: 0, Users: 5
2024/12/03 22:05:24 TPS: 3.57, In-flight Retries: 0, Users: 5
2024/12/03 22:05:25 TPS: 3.58, In-flight Retries: 0, Users: 5
2024/12/03 22:05:26 TPS: 3.59, In-flight Retries: 0, Users: 5
2024/12/03 22:05:27 TPS: 3.56, In-flight Retries: 0, Users: 5
2024/12/03 22:05:28 TPS: 3.58, In-flight Retries: 0, Users: 5
2024/12/03 22:05:29 TPS: 3.59, In-flight Retries: 0, Users: 5
2024/12/03 22:05:30 TPS: 3.60, In-flight Retries: 0, Users: 5
2024/12/03 22:05:31 TPS: 3.57, In-flight Retries: 0, Users: 5
2024/12/03 22:05:32 TPS: 3.58, In-flight Retries: 0, Users: 5
2024/12/03 22:05:33 TPS: 3.57, In-flight Retries: 0, Users: 5
2024/12/03 22:05:34 TPS: 3.58, In-flight Retries: 0, Users: 5
2024/12/03 22:05:35 TPS: 3.58, In-flight Retries: 0, Users: 5
2024/12/03 22:05:36 TPS: 3.54, In-flight Retries: 0, Users: 5
2024/12/03 22:05:37 TPS: 3.55, In-flight Retries: 0, Users: 5
2024/12/03 22:05:38 TPS: 3.54, In-flight Retries: 0, Users: 5
2024/12/03 22:05:39 TPS: 3.54, In-flight Retries: 0, Users: 5
2024/12/03 22:05:40 TPS: 3.53, In-flight Retries: 0, Users: 5
2024/12/03 22:05:41 TPS: 3.52, In-flight Retries: 0, Users: 5
2024/12/03 22:05:42 TPS: 3.51, In-flight Retries: 0, Users: 5
2024/12/03 22:05:43 TPS: 3.48, In-flight Retries: 0, Users: 5
2024/12/03 22:05:44 TPS: 3.46, In-flight Retries: 0, Users: 5
2024/12/03 22:05:45 TPS: 3.44, In-flight Retries: 1, Users: 5
2024/12/03 22:05:45 Reached minimum users, stopping benchmark
2024/12/03 22:05:46 Report generated: reports/benchmark-report-2024-12-03-22-05-46.html
(base) miguel_lemos@Mac restate-recipe-benchmark % ls reports
benchmark-report-2024-12-03-21-09-05.html	benchmark-report-2024-12-03-21-46-09.html
benchmark-report-2024-12-03-21-37-03.html	benchmark-report-2024-12-03-22-05-46.html
```