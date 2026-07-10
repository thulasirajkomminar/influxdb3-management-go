# Changelog

All notable changes to this project will automatically be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## v1.0.0 - 2026-07-10

### Breaking Changes

* Split the single root package into variant-specific packages. The root `influxdb3` package has been removed; import the package for your InfluxDB variant instead:
  * `github.com/thulasirajkomminar/influxdb3-management-go/cloud` — Cloud Dedicated
  * `github.com/thulasirajkomminar/influxdb3-management-go/core` — Core
  * `github.com/thulasirajkomminar/influxdb3-management-go/enterprise` — Enterprise

> [!Important]
>
> See [UPGRADING.md](UPGRADING.md) for migration instructions. Available API methods differ per variant depending on their upstream OpenAPI spec.

### What's Changed

* feat: added InfluxDB 3 Core and Enterprise API clients, generated from their upstream OpenAPI specs.
* feat: added a weekly GitHub Actions workflow that syncs the OpenAPI specs from [influxdata/docs-v2](https://github.com/influxdata/docs-v2), regenerates the clients, verifies the build, and opens a PR.
* feat: added runnable examples for all three variants under `examples/`.
* chore: upgraded to Go 1.26.5.
* chore: pinned all GitHub Actions to release commit SHAs, restricted workflow permissions, and added the `helpers:pinGitHubActionDigests` Renovate preset.
* chore: PR CI now checks formatting, runs `go vet`, and builds the examples module.
* chore: reworked the issue templates (variant dropdown, YAML feature request form) and disabled blank issues.

**Full Changelog**: https://github.com/thulasirajkomminar/influxdb3-management-go/compare/v0.6.0...v1.0.0

## v0.6.0 - 2026-06-09

### What's Changed

* fix: modified GH actions permissions.
* chore: migrated dependency management from Dependabot to Renovate.
* fix(deps): updated `oapi-codegen` to v2.7.1 and `runtime` to v1.4.2.

**Full Changelog**: https://github.com/thulasirajkomminar/influxdb3-management-go/compare/v0.5.0...v0.6.0

## v0.5.0 - 2026-03-04

### What's Changed

* chore: updated docs.
* chore(deps): updated dependencies.

**Full Changelog**: https://github.com/thulasirajkomminar/influxdb3-management-go/compare/v0.4.0...v0.5.0

## v0.4.0 - 2026-02-23

### What's Changed

* feat: updated OpenAPI spec and regenerated the client.

**Full Changelog**: https://github.com/thulasirajkomminar/influxdb3-management-go/compare/v0.3.0...v0.4.0

## v0.3.0 - 2025-11-04

### What's Changed

* feat: upgraded OpenAPI spec and regenerated the client.
* docs: updated README and example.
* chore(deps): updated `oapi-codegen` to v2.5.1.

**Full Changelog**: https://github.com/thulasirajkomminar/influxdb3-management-go/compare/v0.2.0...v0.3.0

## v0.2.0 - 2025-10-30

### What's Changed

* chore: moved the repo from org to personal account.
* chore: added example code snippet.
* chore(deps): updated dependencies.

**Full Changelog**: https://github.com/thulasirajkomminar/influxdb3-management-go/compare/v0.1.2...v0.2.0

## v0.1.2 - 2025-03-11

### What's Changed

* chore: updated OpenAPI spec and regenerated the client.

**Full Changelog**: https://github.com/thulasirajkomminar/influxdb3-management-go/compare/v0.1.1...v0.1.2

## v0.1.1 - 2024-09-30

### What's Changed

* chore: used tools package to generate code.
* fix: used `JSONMerge` since `JsonMerge` is deprecated.

**Full Changelog**: https://github.com/thulasirajkomminar/influxdb3-management-go/compare/v0.1.0...v0.1.1

## v0.1.0 - 2024-06-26

### What's Changed

* feat: initial version of the client library for the InfluxDB 3 Cloud Dedicated Management API.
