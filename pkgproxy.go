package pkgproxy

import (
	"strings"

	"github.com/gocolly/colly"
)

type Package struct {
	Name                   string
	Repository             string
	Version                string
	PublishedDate          string
	License                string
	Imports                string
	ImportedBy             string
	ValidGoMod             string
	RedistributableLicense string
	TaggedVersion          string
	StableVersion          string
}

type PkgCollector struct {
	BaseURL   string
	Collector *colly.Collector
}

func NewPkgCollector() *PkgCollector {
	c := colly.NewCollector(
		colly.AllowedDomains("pkg.go.dev", "www.pkg.go.dev"),
		colly.CacheDir("./.pkg_cache"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)
	return &PkgCollector{
		Collector: c,
		BaseURL:   "https://pkg.go.dev",
	}
}

func (p *PkgCollector) Get(pkgName string) Package {
	// Repository URL
	var repo string
	p.Collector.OnHTML(".UnitMeta-repo", func(e *colly.HTMLElement) {
		repo = e.ChildText("a")
	})

	// Version
	var version string
	p.Collector.OnHTML(".go-Main-headerDetails", func(e *colly.HTMLElement) {
		version = e.ChildAttr("a", "aria-label")
		version = sanitizeVersion(version)
	})

	// Published Date
	var published string
	p.Collector.OnHTML(".go-Main-headerDetailItem[data-test-id='UnitHeader-commitTime']", func(e *colly.HTMLElement) {
		published = sanitizeDate(e.Text)
	})

	// License
	var license string
	p.Collector.OnHTML(".go-Main-headerDetailItem[data-test-id='UnitHeader-licenses']", func(e *colly.HTMLElement) {
		license = strings.TrimSpace(e.ChildText("a"))
	})

	p.Collector.Visit(p.BaseURL + "/" + pkgName)

	return Package{
		Name:          pkgName,
		Repository:    repo,
		Version:       version,
		PublishedDate: published,
		License:       license,
	}
}

func sanitizeDate(date string) string {
	return strings.TrimSpace(strings.Split(strings.TrimSpace(date), ":")[1])
}

func sanitizeVersion(version string) string {
	chunks := strings.Split(strings.TrimSpace(version), "-")
	chunks = strings.Split(chunks[0], ":")
	return strings.TrimSpace(chunks[1])
}

func Get(name string) Package {
	p := NewPkgCollector()
	p.BaseURL = "https://pkg.go.dev"
	return p.Get(name)
}
