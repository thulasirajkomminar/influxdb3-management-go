// Package enterprise provides a Go client for the InfluxDB 3 Enterprise API.
//
// This package contains generated types and client methods for interacting with
// InfluxDB 3 Enterprise databases, tables, queries, processing engine, and administrative functions.
//
// The InfluxDB HTTP API for InfluxDB 3 Enterprise provides a programmatic interface for
// interacting with InfluxDB 3 Enterprise databases and resources.
package enterprise

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate types -o types.gen.go -package enterprise ../specs/enterprise/ref.yml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate client -o client.gen.go -package enterprise ../specs/enterprise/ref.yml
