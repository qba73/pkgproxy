package main

import (
	"fmt"

	"github.com/qba73/pkgproxy"
)

func main() {
	p := pkgproxy.Get("github.com/bitfield/weaver")
	fmt.Printf("%+v\n", p)
	// {Name:github.com/bitfield/weaver Repository:github.com/bitfield/weaver Version:v0.3.3 PublishedDate:Apr 20, 2025 License:MIT Imports:15 ImportedBy:0 ValidGoMod:Yes RedistributableLicense:Yes TaggedVersion:Yes StableVersion:No}
}
