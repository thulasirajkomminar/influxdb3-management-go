# Upgrading Notes

This document captures required refactoring on your part when upgrading to a module version that contains breaking changes.

## Upgrading to v1.0.0

### Key Changes

#### Before

```go
import "github.com/thulasirajkomminar/influxdb3"

// Single client for Cloud Dedicated only
client, err := influxdb3.NewClient(baseURL)
```

### After

```go
// Choose your InfluxDB variant
import "github.com/thulasirajkomminar/influxdb3-management-go/cloud"      // Cloud Dedicated
import "github.com/thulasirajkomminar/influxdb3-management-go/core"       // Core
import "github.com/thulasirajkomminar/influxdb3-management-go/enterprise" // Enterprise

// Create variant-specific clients
cloudClient, err := influxdb3cloud.NewClient(baseURL)
coreClient, err := influxdb3core.NewClient(baseURL)
enterpriseClient, err := influxdb3enterprise.NewClient(baseURL)
```
