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
	os.Exit(run(os.Stdout))
}

func run(w io.Writer) int {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println(usage)
		return 0
	}
	gopackage := flag.Args()[0]
	info := pkgproxy.Get(gopackage)
	fmt.Fprintf(w, "%+v\n", info)
	return 0
}
