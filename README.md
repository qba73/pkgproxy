[![Go Reference](https://pkg.go.dev/badge/github.com/qba73/pkgproxy.svg)](https://pkg.go.dev/github.com/qba73/pkgproxy)
[![CI](https://github.com/qba73/pkgproxy/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/qba73/pkgproxy/actions/workflows/ci.yml)
[![CVE Scan](https://github.com/qba73/pkgproxy/actions/workflows/cvescan.yml/badge.svg?branch=main)](https://github.com/qba73/pkgproxy/actions/workflows/cvescan.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/qba73/pkgproxy)](https://goreportcard.com/report/github.com/qba73/pkgproxy)
[![CodeQL](https://github.com/qba73/pkgproxy/actions/workflows/github-code-scanning/codeql/badge.svg?branch=main)](https://github.com/qba73/pkgproxy/actions/workflows/github-code-scanning/codeql)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/qba73/pkgproxy)
![GitHub License](https://img.shields.io/github/license/qba73/pkgproxy)




# pkgproxy

`pkgproxy` is a proxy and a cli for [pkg.go.dev](https://pkg.go.dev) service. It allows to query package and get info using REST API and a command line.

## CLI

Install the binary

```shell
go install github.com/qba73/pkgproxy/cmd/pkg@latest
```

Getting help

```shell
pkg
Usage: pkg package

Checks the Go package, and reports basic information.
```

### Examples

Getting information about Go packages

- [bitfield/weaver](https://pkg.go.dev/github.com/bitfield/weaver)

```shell
pkg github.com/bitfield/weaver | jq .
```
```json
{
  "name": "github.com/bitfield/weaver",
  "repository": "github.com/bitfield/weaver",
  "version": "v0.3.3",
  "publishedDate": "Apr 20, 2025",
  "license": "MIT",
  "imports": "15",
  "importedBy": "0",
  "validGomod": "Yes",
  "redistributableLicense": "Yes",
  "taggedVersion": "Yes",
  "stableVersion": "No"
}
```

- [Open Telemetry Go client](https://pkg.go.dev/go.opentelemetry.io/otel)

```shell
pkg go.opentelemetry.io/otel | jq .
```
```json
{
  "name": "go.opentelemetry.io/otel",
  "repository": "github.com/open-telemetry/opentelemetry-go",
  "version": "v1.37.0",
  "publishedDate": "Jun 25, 2025",
  "license": "Apache-2.0",
  "imports": "5",
  "importedBy": "12,010",
  "validGomod": "Yes",
  "redistributableLicense": "Yes",
  "taggedVersion": "Yes",
  "stableVersion": "Yes"
}
```

- [qba73/meteo](https://pkg.go.dev/github.com/qba73/meteo)

```shell
pkg github.com/qba73/meteo | jq .
```
```json
{
  "name": "github.com/qba73/meteo",
  "repository": "github.com/qba73/meteo",
  "version": "v0.0.0",
  "publishedDate": "Jun 10, 2025",
  "license": "MIT",
  "imports": "14",
  "importedBy": "0",
  "validGomod": "Yes",
  "redistributableLicense": "Yes",
  "taggedVersion": "No",
  "stableVersion": "No"
}
```
