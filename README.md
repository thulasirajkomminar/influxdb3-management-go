# InfluxDB v3 Management Go Client Library

A Go client library for InfluxDB v3 products, supporting Cloud Dedicated, Core, and Enterprise deployments.

## Supported products

### InfluxDB 3 Cloud Dedicated

Management API for InfluxDB Cloud Dedicated clusters, databases, and tokens.

**Import**: `import "github.com/komminarlabs/influxdb3-management-go/cloud"`

### InfluxDB 3 Core

Core API for writing data, querying with SQL/InfluxQL, and managing local instances.

**Import**: `import "github.com/komminarlabs/influxdb3-management-go/core"`

### InfluxDB 3 Enterprise

Enterprise API with advanced features including processing engine and fine-grained permissions.

**Import**: `import "github.com/komminarlabs/influxdb3-management-go/enterprise"`

## 🚀 Installation

Install the product(s) you need:

```bash
export INFLUXDB3_ACCOUNT_ID="4ade9b2e-0a52-4a46-b3b8-1b43ea493a98"
export INFLUXDB3_CLUSTER_ID="a379c48a-791e-47fe-ba64-628ba19507e8"
export INFLUXDB3_TOKEN="1e0f14063eb14a9e94fe765bf999a90cb7962f8e0f394110b91053ea26cdce5071d6bca29e4d4684bed463cf2ea9f381"
export INFLUXDB3_URL="https://console.influxdata.com/api/v0"
# Cloud Dedicated
go get github.com/komminarlabs/influxdb3-management-go/cloud

# Core
go get github.com/komminarlabs/influxdb3-management-go/core

# Enterprise
go get github.com/komminarlabs/influxdb3-management-go/enterprise
```

## 🔧 Development

### Generate Go Code from OpenAPI Specs

import (
    "context"
    "io"
    "net/http"
    
    "github.com/caarlos0/env/v11"
    "github.com/thulasirajkomminar/influxdb3-management-go"
)

type InfluxdbConfig struct {
    AccountId influxdb3.UuidV4 `env:"INFLUXDB3_ACCOUNT_ID"`
    ClusterId influxdb3.UuidV4 `env:"INFLUXDB3_CLUSTER_ID"`
    Token     string           `env:"INFLUXDB3_TOKEN"`
    Url       string           `env:"INFLUXDB3_URL"`
}

func main() {
    cfg := InfluxdbConfig{}
    opts := env.Options{RequiredIfNoDef: true}

    err := env.ParseWithOptions(&cfg, opts)
    if err != nil {
        panic(err)
    }

    ctx := context.Background()
    client, err := influxdb3.NewClient(cfg.Url, influxdb3.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
        req.Header.Set("Accept", "application/json")
        req.Header.Set("Authorization", "Bearer "+cfg.Token)
        return nil

    }))
    if err != nil {
        panic(err)
    }

    resp, err := client.GetDatabaseTokens(ctx, cfg.AccountId, cfg.ClusterId)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        bodyBytes, err := io.ReadAll(resp.Body)
        if err != nil {
            panic(err)
        }
        bodyString := string(bodyBytes)
        println(bodyString)
    }
}
```bash
go generate ./...
```
