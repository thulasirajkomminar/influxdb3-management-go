# Upgrading Notes

This document captures required refactoring on your part when upgrading to a module version that contains breaking changes.

## Upgrading to v1.0.0

### Key Changes

#### Before (Old Structure)

```go
import "github.com/komminarlabs/influxdb3"

// Single client for Cloud Dedicated only
client, err := influxdb3.NewClient(baseURL)
```

### After (New Structure)

```go
// Choose your InfluxDB variant
import "github.com/komminarlabs/influxdb3-management-go/cloud"      // Cloud Dedicated
import "github.com/komminarlabs/influxdb3-management-go/core"       // Core
import "github.com/komminarlabs/influxdb3-management-go/enterprise" // Enterprise

// Create variant-specific clients
cloudClient, err := cloud.NewClient(baseURL)
coreClient, err := core.NewClient(baseURL)
enterpriseClient, err := enterprise.NewClient(baseURL)
```
