# Changelog

All notable changes to this project will automatically be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - Unreleased

### ⚠️ Breaking Changes

- Split the single root package into variant-specific packages. The root `influxdb3` package has been removed; import the package for your InfluxDB variant instead (see [UPGRADING.md](UPGRADING.md)):
  - `github.com/thulasirajkomminar/influxdb3-management-go/cloud` — Cloud Dedicated
  - `github.com/thulasirajkomminar/influxdb3-management-go/core` — Core
  - `github.com/thulasirajkomminar/influxdb3-management-go/enterprise` — Enterprise

### Added

- InfluxDB 3 Core and Enterprise API clients, generated from their upstream OpenAPI specs
- Per-variant OpenAPI specs under `specs/`
- Runnable examples for all three variants under `examples/`
- Weekly GitHub Actions workflow that syncs the OpenAPI specs from [influxdata/docs-v2](https://github.com/influxdata/docs-v2), regenerates the clients, verifies the build, and opens a PR

### Changed

- Upgraded to Go 1.26.5
- PR CI now checks formatting, runs `go vet`, and builds the examples module

## [0.6.0] - 2026-06-09

### Changed

- Updated GitHub Actions permissions
- Migrated dependency management from Dependabot to Renovate
- Updated `oapi-codegen` to v2.7.1 and `runtime` to v1.4.2

## [0.5.0] - 2026-03-04

### Changed

- Updated documentation
- Updated dependencies

## [0.4.0] - 2026-02-23

### Changed

- Updated OpenAPI spec and regenerated the client

## [0.3.0] - 2025-11-04

### Changed

- Upgraded OpenAPI spec and regenerated the client
- Updated README and example
- Updated `oapi-codegen` to v2.5.1

## [0.2.0] - 2025-10-30

### Changed

- Moved the repository from the `komminarlabs` organization to `thulasirajkomminar`
- Added example code snippet
- Updated dependencies

## [0.1.2] - 2025-03-11

### Changed

- Updated OpenAPI spec and regenerated the client

## [0.1.1] - 2024-09-30

### Changed

- Used tools package to generate code
- Used `JSONMerge` since `JsonMerge` is deprecated

## [0.1.0] - 2024-06-26

### Added

- Initial version of the client library for the InfluxDB 3 Cloud Dedicated Management API

[1.0.0]: https://github.com/thulasirajkomminar/influxdb3-management-go/compare/v0.6.0...HEAD
[0.6.0]: https://github.com/thulasirajkomminar/influxdb3-management-go/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/thulasirajkomminar/influxdb3-management-go/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/thulasirajkomminar/influxdb3-management-go/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/thulasirajkomminar/influxdb3-management-go/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/thulasirajkomminar/influxdb3-management-go/compare/v0.1.2...v0.2.0
[0.1.2]: https://github.com/thulasirajkomminar/influxdb3-management-go/compare/v0.1.1...v0.1.2
[0.1.1]: https://github.com/thulasirajkomminar/influxdb3-management-go/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/thulasirajkomminar/influxdb3-management-go/releases/tag/v0.1.0
