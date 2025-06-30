package main

import (
	"fmt"

	"github.com/qba73/pkgproxy"
)

func main() {
	repo := pkgproxy.Get("gopkg.in/tomb.v1")
	fmt.Printf("%+v\n", repo)
}
