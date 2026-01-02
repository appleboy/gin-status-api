# gin-status-api

A system status monitoring API handler for the [Gin](https://github.com/gin-gonic/gin) framework that provides CPU, memory, garbage collection, and other runtime information for your Go applications.

[![GoDoc](https://godoc.org/github.com/appleboy/gin-status-api?status.svg)](https://godoc.org/github.com/appleboy/gin-status-api)
[![Run Tests](https://github.com/appleboy/gin-status-api/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/appleboy/gin-status-api/actions/workflows/go.yml)
[![Trivy Security Scan](https://github.com/appleboy/gin-status-api/actions/workflows/trivy-scan.yml/badge.svg?branch=master)](https://github.com/appleboy/gin-status-api/actions/workflows/trivy-scan.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/gin-status-api)](https://goreportcard.com/report/github.com/appleboy/gin-status-api)
[![codecov](https://codecov.io/gh/appleboy/gin-status-api/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/gin-status-api)

## Table of Contents

- [gin-status-api](#gin-status-api)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Installation](#installation)
  - [Quick Start](#quick-start)
    - [Basic Usage](#basic-usage)
    - [Running the Example](#running-the-example)
  - [API Response](#api-response)
  - [Response Example](#response-example)
  - [Use Cases](#use-cases)
  - [Related Projects](#related-projects)
  - [License](#license)
  - [Contributing](#contributing)

## Features

- **Simple Integration**: Add system monitoring with just one line of code
- **Comprehensive Metrics**: Provides complete runtime information including CPU, memory, and GC stats
- **JSON Format**: Easy to integrate with monitoring systems and dashboards
- **Lightweight**: Minimal performance impact on your application
- **Multiple Use Cases**: Perfect for health checks, performance monitoring, debugging, and alerting

## Installation

Download and install the package using `go get`:

```bash
go get github.com/appleboy/gin-status-api
```

## Quick Start

### Basic Usage

Import the package and register a route in your Gin application:

```go
package main

import (
    status "github.com/appleboy/gin-status-api"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.New()

    // Register the status API endpoint
    r.GET("/api/status", status.GinHandler)

    // Start the server
    r.Run() // Listen and serve on 0.0.0.0:8080
}
```

### Running the Example

The project includes a complete example that you can run:

```bash
cd _example
go run server.go
```

Then access the status endpoint using curl or your browser:

```bash
curl http://localhost:8080/api/status
```

## API Response

The API returns a JSON object containing comprehensive system status information:

| Field                 | Description                                            |
| --------------------- | ------------------------------------------------------ |
| `time`                | Current Unix timestamp                                 |
| `go_version`          | Go runtime version                                     |
| `go_os`               | Operating system                                       |
| `go_arch`             | System architecture                                    |
| `cpu_num`             | Number of CPU cores                                    |
| `goroutine_num`       | Current number of goroutines                           |
| `gomaxprocs`          | GOMAXPROCS setting                                     |
| `cgo_call_num`        | Number of CGO calls                                    |
| `memory_alloc`        | Bytes of allocated heap objects                        |
| `memory_total_alloc`  | Cumulative bytes allocated for heap objects            |
| `memory_sys`          | Total bytes of memory obtained from the OS             |
| `memory_lookups`      | Number of pointer lookups performed                    |
| `memory_mallocs`      | Cumulative count of heap objects allocated             |
| `memory_frees`        | Cumulative count of heap objects freed                 |
| `memory_stack`        | Bytes in stack spans                                   |
| `heap_alloc`          | Bytes of allocated heap objects (same as memory_alloc) |
| `heap_sys`            | Bytes of heap memory obtained from the OS              |
| `heap_idle`           | Bytes in idle (unused) spans                           |
| `heap_inuse`          | Bytes in in-use spans                                  |
| `heap_released`       | Bytes of physical memory returned to the OS            |
| `heap_objects`        | Number of allocated heap objects                       |
| `gc_next`             | Target heap size for next GC cycle                     |
| `gc_last`             | Time of last garbage collection                        |
| `gc_num`              | Number of completed GC cycles                          |
| `gc_per_second`       | GC cycles per second                                   |
| `gc_pause_per_second` | Average GC pause time per second                       |
| `gc_pause`            | Circular buffer of recent GC pause times               |

For more detailed information about these metrics, please refer to the [golang-stats-api-handler documentation](https://github.com/fukata/golang-stats-api-handler#toc3).

## Response Example

![response screenshot](screenshot/response.png)

## Use Cases

- **Health Checks**: Integrate with Kubernetes liveness/readiness probes
- **Performance Monitoring**: Track memory usage and GC behavior over time
- **Debugging**: Quickly inspect goroutine counts and resource usage
- **Alerting Systems**: Set up alerts based on memory thresholds or goroutine counts
- **DevOps Dashboards**: Feed metrics into Grafana, Prometheus, or other monitoring tools
- **Capacity Planning**: Analyze application resource consumption patterns

## Related Projects

- [Gin Web Framework](https://github.com/gin-gonic/gin) - HTTP web framework written in Go
- [golang-stats-api-handler](https://github.com/fukata/golang-stats-api-handler) - The underlying stats collection library

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request or open an Issue for bugs, feature requests, or questions.
