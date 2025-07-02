[![Go Reference](https://pkg.go.dev/badge/github.com/qba73/pkgproxy.svg)](https://pkg.go.dev/github.com/qba73/pkgproxy)
[![CI](https://github.com/qba73/pkgproxy/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/qba73/pkgproxy/actions/workflows/ci.yml)
[![CVE Scan](https://github.com/qba73/pkgproxy/actions/workflows/cvescan.yml/badge.svg?branch=main)](https://github.com/qba73/pkgproxy/actions/workflows/cvescan.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/qba73/pkgproxy)](https://goreportcard.com/report/github.com/qba73/pkgproxy)
[![CodeQL](https://github.com/qba73/pkgproxy/actions/workflows/github-code-scanning/codeql/badge.svg?branch=main)](https://github.com/qba73/pkgproxy/actions/workflows/github-code-scanning/codeql)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/qba73/pkgproxy)
![GitHub License](https://img.shields.io/github/license/qba73/pkgproxy)
[![Scorecard](https://github.com/qba73/pkgproxy/actions/workflows/scorecard.yml/badge.svg)](https://github.com/qba73/pkgproxy/actions/workflows/scorecard.yml)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/10848/badge)](https://www.bestpractices.dev/projects/10848)




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

## Examples

### Getting information about Go packages

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

### Using `pkg` with `scorecard`

Checking OpenSSF score for a Go package.

1. Generate and export `GITHUB_AUTH_TOKEN` env var.

2. Verify `scorecard` is installed:

```shell
scorecard version
         __  ____     ____    ___    ____    _____    ____      _      ____    ____
        / / / ___|   / ___|  / _ \  |  _ \  | ____|  / ___|    / \    |  _ \  |  _ \
       / /  \___ \  | |     | | | | | |_) | |  _|   | |       / _ \   | |_) | | | | |
  _   / /    ___) | | |___  | |_| | |  _ <  | |___  | |___   / ___ \  |  _ <  | |_| |
 (_) /_/    |____/   \____|  \___/  |_| \_\ |_____|  \____| /_/   \_\ |_| \_\ |____/
./scorecard: OpenSSF Scorecard

GitVersion:    5.2.1
GitCommit:     ab2f6e92482462fe66246d9e32f642855a691dc1
GitTreeState:  clean
BuildDate:     2025-05-30T16:02:02Z
GoVersion:     go1.24.3
Compiler:      gc
Platform:      darwin/arm64
```

To check the score, we need to pass the Go package URL. But what if we have only the package name? This is where the `pkg` CLI comes in handy. `pkg` queries the `pkg.go.dev` service for information and returns package info in JSON format.

3. Send a query to `pkg.go.dev`:
```shell
pkg go.opentelemetry.io/otel | jq -r '.repository'
```
response:
```
github.com/open-telemetry/opentelemetry-go
```

4. Send a query to `scorecard`

```shell
scorecard --repo github.com/open-telemetry/opentelemetry-go --format json | jq .score
```
```shell
9.6
```

---

How to use `pkg` and `scorecard` together?

- [OpenTelemetry](https://github.com/open-telemetry/opentelemetry-go)

```shell
scorecard --repo $(pkg go.opentelemetry.io/otel | jq -r '.repository' ) --format json | jq .score
```
response:
```shell
9.6
```

- [Inspector](https://github.com/qba73/inspector)
```shell
scorecard --repo $(pkg github.com/qba73/inspector | jq -r '.repository' ) --format json | jq .score
```
response:
```shell
6.4
```
