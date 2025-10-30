// Package cloud provides a Go client for the InfluxDB 3 Cloud Dedicated Management API.
//
// This package contains generated types and client methods for interacting with
// InfluxDB 3 Cloud Dedicated clusters, databases, and tokens.
//
// The InfluxDB v3 Management API Go client library lets you manage an InfluxDB Cloud Dedicated
// instance and integrate functions such as creating and managing databases, permissions,
// and tokens into your workflow or application.
package cloud

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate types -o types.gen.go -package cloud ../specs/cloud/openapi.yml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate client -o client.gen.go -package cloud ../specs/cloud/openapi.yml
