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
