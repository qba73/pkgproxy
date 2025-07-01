package main

import (
	"fmt"

	"github.com/qba73/pkgproxy"
)

func main() {
	p, err := pkgproxy.GetJSON("github.com/bitfield/weaver")
	if err != nil {
		// handle error
	}
	fmt.Printf("%+v\n", p)
	// {"name":"github.com/bitfield/weaver","repository":"github.com/bitfield/weaver","version":"v0.3.3","publishedDate":"Apr 20, 2025","license":"MIT","imports":"15","importedBy":"0","validGomod":"Yes","redistributableLicense":"Yes","taggedVersion":"Yes","stableVersion":"No"}
}
