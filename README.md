[![Go Reference](https://pkg.go.dev/badge/github.com/qba73/pkgproxy.svg)](https://pkg.go.dev/github.com/qba73/pkgproxy)
[![CI](https://github.com/qba73/pkgproxy/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/qba73/pkgproxy/actions/workflows/ci.yml)
[![CVE Scan](https://github.com/qba73/pkgproxy/actions/workflows/cvescan.yml/badge.svg?branch=main)](https://github.com/qba73/pkgproxy/actions/workflows/cvescan.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/qba73/pkgproxy)](https://goreportcard.com/report/github.com/qba73/pkgproxy)



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
  "Name": "github.com/bitfield/weaver",
  "Repository": "github.com/bitfield/weaver",
  "Version": "v0.3.3",
  "PublishedDate": "Apr 20, 2025",
  "License": "MIT",
  "Imports": "15",
  "ImportedBy": "0",
  "ValidGoMod": "Yes",
  "RedistributableLicense": "Yes",
  "TaggedVersion": "Yes",
  "StableVersion": "No"
}
```

- [qba73/meteo](https://pkg.go.dev/github.com/qba73/meteo)

```shell
pkg github.com/qba73/meteo | jq .
```
```json
{
  "Name": "github.com/qba73/meteo",
  "Repository": "github.com/qba73/meteo",
  "Version": "v0.0.0",
  "PublishedDate": "Jun 10, 2025",
  "License": "MIT",
  "Imports": "14",
  "ImportedBy": "0",
  "ValidGoMod": "Yes",
  "RedistributableLicense": "Yes",
  "TaggedVersion": "No",
  "StableVersion": "No"
}
```
