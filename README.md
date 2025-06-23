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
# Cloud Dedicated
go get github.com/komminarlabs/influxdb3-management-go/cloud

# Core
go get github.com/komminarlabs/influxdb3-management-go/core

# Enterprise
go get github.com/komminarlabs/influxdb3-management-go/enterprise
```

## 🔧 Development

### Generate Go Code from OpenAPI Specs

```bash
go generate ./...
```
