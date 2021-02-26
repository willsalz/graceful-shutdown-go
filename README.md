# graceful-shutdown-go
graceful shutdown patterns for golang

```console
$ go run main.go
2021/02/26 11:36:51 [main] starting services
2021/02/26 11:36:51 [service] starting
2021/02/26 11:36:51 [service] doing some work…
2021/02/26 11:36:56 [service] doing some work…
^C2021/02/26 11:37:00 [main] system waiting up to 10s before exit.
2021/02/26 11:37:01 [service] shutting down.
2021/02/26 11:37:01 [main] cleanup done; exiting now
```
