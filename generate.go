package influxdb3

// InfluxDB Cloud Dedicated Management API
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate types -o cloud/types.gen.go -package cloud specs/cloud/openapi.yml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate client -o cloud/client.gen.go -package cloud specs/cloud/openapi.yml

// InfluxDB Core API
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate types -o core/types.gen.go -package core specs/core/ref.yml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate client -o core/client.gen.go -package core specs/core/ref.yml

// InfluxDB Enterprise API
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate types -o enterprise/types.gen.go -package enterprise specs/enterprise/ref.yml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate client -o enterprise/client.gen.go -package enterprise specs/enterprise/ref.yml
