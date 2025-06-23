// Package core provides a Go client for the InfluxDB 3 Core API.
//
// This package contains generated types and client methods for interacting with
// InfluxDB 3 Core databases, tables, queries, and administrative functions.
//
// The InfluxDB HTTP API for InfluxDB 3 Core provides a programmatic interface for
// interacting with InfluxDB 3 Core databases and resources.
package core

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate types -o types.gen.go -package core ../specs/core/ref.yml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate client -o client.gen.go -package core ../specs/core/ref.yml
