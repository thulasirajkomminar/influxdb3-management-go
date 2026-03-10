# InfluxDB v3 Management Go Client Library

A Go client library for InfluxDB v3 products, supporting Cloud Dedicated, Core, and Enterprise.

## Supported products

### InfluxDB 3 Cloud Dedicated

`import "github.com/thulasirajkomminar/influxdb3-management-go/cloud"`

### InfluxDB 3 Core

`import "github.com/thulasirajkomminar/influxdb3-management-go/core"`

### InfluxDB 3 Enterprise

`import "github.com/thulasirajkomminar/influxdb3-management-go/enterprise"`

## Generated types and API client

This library is generated using [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from these [OpenAPI specs](https://github.com/influxdata/docs-v2/tree/master/api-docs)

### Generate

```go
go generate ./...
```

## Usage

```bash
export INFLUXDB3_ACCOUNT_ID="4ade9b2e-0a52-4a46-b3b8-1b43ea493a98"
export INFLUXDB3_CLUSTER_ID="a379c48a-791e-47fe-ba64-628ba19507e8"
export INFLUXDB3_TOKEN="1e0f14063eb14a9e94fe765bf999a90cb7962f8e0f394110b91053ea26cdce5071d6bca29e4d4684bed463cf2ea9f381"
export INFLUXDB3_URL="https://console.influxdata.com/api/v0"
```

### Sample code to list databases in a cloud dedicated cluster

```go
package main

import (
    "context"
    "net/http"

    "github.com/caarlos0/env/v11"
    influxdb3cloud "github.com/thulasirajkomminar/influxdb3-management-go/cloud"
)

type InfluxdbConfig struct {
    AccountId influxdb3cloud.UuidV4 `env:"INFLUXDB3_ACCOUNT_ID"`
    ClusterId influxdb3cloud.UuidV4 `env:"INFLUXDB3_CLUSTER_ID"`
    Token     string                `env:"INFLUXDB3_TOKEN"`
    Url       string                `env:"INFLUXDB3_URL"`
}

func main() {
    cfg := InfluxdbConfig{}
    opts := env.Options{RequiredIfNoDef: true}
    
    err := env.ParseWithOptions(&cfg, opts)
    if err != nil {
        panic(err)
    }
    
    ctx := context.Background()
    client, err := influxdb3cloud.NewClientWithResponses(cfg.Url, influxdb3cloud.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
        req.Header.Set("Accept", "application/json")
        req.Header.Set("Authorization", "Bearer "+cfg.Token)
        return nil
    
    }))
    if err != nil {
        panic(err)
    }
    
    resp, err := client.GetClusterDatabasesWithResponse(ctx, cfg.AccountId, cfg.ClusterId)
    if err != nil {
        panic(err)
    }
    
    if resp.StatusCode() == 200 {
        for _, database := range *resp.JSON200 {
            println("Database name: " + database.Name)
        }
    }
}

```
