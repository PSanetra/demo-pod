# Kubernetes Demo Pod

DON'T RUN THIS EVER ON PRODUCTION!!!

## Run

```bash
$ docker run --rm -p 8080:8080 psanetra/demo-pod:latest
```

## Features

* Save state in a volume
* Read state from a volume
* Show environment variables
* Show changing ConfigMap values (via `--watch` flag)
* Show changing Secret values (via `--watch` flag)
* Control liveness
* Control readiness
* Control cpu utilization
* Control memory utilization
* Show available memory and memory in use
* Show client ip
