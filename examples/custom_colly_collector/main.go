package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
	"github.com/qba73/pkgproxy"
)

func main() {
	pc := pkgproxy.NewPkgCollector()

	// Add custom Colly collector.
	// Note, no cache dir is created.
	pc.Collector = colly.NewCollector(
		colly.AllowedDomains("pkg.go.dev", "www.pkg.go.dev"),
		colly.UserAgent("User-Agent: my_collector/123"),
	)

	// Get info about Go package 'github.com/bitfield/weaver'
	weaver := pc.Get("github.com/bitfield/weaver")
	fmt.Printf("%+v\n", weaver)
	// {Name:github.com/bitfield/weaver Repository:github.com/bitfield/weaver Version:v0.3.3 PublishedDate:Apr 20, 2025 License:MIT Imports:15 ImportedBy:0 ValidGoMod:Yes RedistributableLicense:Yes TaggedVersion:Yes StableVersion:No}

	d, err := json.MarshalIndent(weaver, "", "  ")
	if err != nil {
		// handle error
	}
	fmt.Printf("%+v\n", string(d))
	//	{
	//	  "name": "github.com/bitfield/weaver",
	//	  "repository": "github.com/bitfield/weaver",
	//	  "version": "v0.3.3",
	//	  "publishedDate": "Apr 20, 2025",
	//	  "license": "MIT",
	//	  "imports": "15",
	//	  "importedBy": "0",
	//	  "validGomod": "Yes",
	//	  "redistributableLicense": "Yes",
	//	  "taggedVersion": "Yes",
	//	  "stableVersion": "No"
	//	}
}
