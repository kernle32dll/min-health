[![Build Status](https://travis-ci.com/kernle32dll/min-health.svg?branch=master)](https://travis-ci.com/kernle32dll/min-health)
[![GoDoc](https://godoc.org/github.com/kernle32dll/min-health?status.svg)](http://godoc.org/github.com/kernle32dll/min-health)
[![Go Report Card](https://goreportcard.com/badge/github.com/kernle32dll/min-health)](https://goreportcard.com/report/github.com/kernle32dll/min-health)
[![codecov](https://codecov.io/gh/kernle32dll/min-health/branch/master/graph/badge.svg)](https://codecov.io/gh/kernle32dll/min-health)

# min-health

min-health is a minimal health checker.

Its designed to simply call a given url with a given HTTP method, and check the status code (if any).

min-health can be used as either a command, or as a library.

Detailed documentation can be found on [GoDoc](https://godoc.org/github.com/kernle32dll/min-health).

## Usage as command

Grab the [latest binary](https://github.com/kernle32dll/min-health/releases/latest), and use that as your health check.

The binary takes the following parameters:

|  Parameter | Default             |  Description                                |
|------------|---------------------|---------------------------------------------|
| `--url`    | `http://localhost/` | The full URL of the health endpoint to call |
| `--method` | `GET`               | The HTTP method to use for the health check |


An example, on how to use min-health in a Dockerfile:

```Dockerfile
...

RUN wget https://github.com/kernle32dll/min-health/releases/download/v1.0.0/min-health-linux-amd64 && chmod +x min-health-linux-amd64
HEALTHCHECK --start-period=5s CMD ["/min-health", "--url=http://localhost:9024/health"]

...
```

## Usage as library

Download:

```
go get github.com/kernle32dll/min-health
```

The most simple usage is as follows:

```go
import "github.com/kernle32dll/min-health/health"

...

c := &health.Config{
    Method: http.MethodGet,
    URL:    "http://localhost/",
}

if !health.DoRequest(c) {
    // oh no, health check failed!
}
```

The `health.Config` also allows usage of a custom http client, and a log function. If not set, they default to the default http client,
and `log.Println` respectively.