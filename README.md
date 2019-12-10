# Kubernetes Demo Pod

DON'T RUN THIS EVER ON PRODUCTION!!!

## Run

```bash
$ docker run --rm -p 8080:8080 psanetra/demo-pod:latest
```

## Features

* Delay startup (via `--startup-delay` option)
* Save state in a volume
* Read state from a volume
* Show environment variables
* Show changing ConfigMap values (via `--watch` option)
* Show changing Secret values (via `--watch` option)
* Control liveness
* Control readiness
* Control cpu utilization
* Control memory utilization
* Show available and used memory
* Show client IP
* Show Pod IPs
