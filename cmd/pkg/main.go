/*
pkg retrieves information about Go packages from pkg.go.dev.

By default, pkg prints the JSON formatted data to standard output.

Usage:

	pkg package_name

When pkg reads from standard input, it accepts a full Go package name.
A package name must be a syntactically valid string. When formatting
pkg outputs a JSON representation of the package data, allowing the output
to be further formatted or processed by piping it through, for example,
the jq utility and other standard command-line utilities.

# Examples

To check package info:

	pkg github.com/bitfield/script

To convert the package info with jq:

	pkg github.com/bitfield/script | jq .
*/
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/qba73/pkgproxy"
)

var usage = `Usage: pkg package

Checks the Go package, and reports basic information.
`

func main() {
	os.Exit(run(os.Stdout, os.Stderr))
}

func run(w, ew io.Writer) int {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println(usage)
		return 0
	}
	gopackage := flag.Args()[0]
	info, err := pkgproxy.GetJSON(gopackage)
	if err != nil {
		fmt.Fprint(ew, err)
		return 1
	}
	fmt.Fprintf(w, "%s\n", info)
	return 0
}
