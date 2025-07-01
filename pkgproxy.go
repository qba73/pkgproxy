package pkgproxy

import (
	"github.com/gocolly/colly"
)

type Package struct {
	Address       string
	Repository    string
	Published     string
	ValidGoMod    string
	TaggedVersion string
	StableVersion string
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
	}
}

func (p *PkgCollector) Get(pkgName string) Package {
	var repo string
	p.Collector.OnHTML(".UnitMeta-repo", func(e *colly.HTMLElement) {
		repo = e.ChildText("a")
	})

	p.Collector.Visit(p.BaseURL + "/" + pkgName)

	return Package{
		Address:    pkgName,
		Repository: repo,
	}
}

func Get(name string) Package {
	p := NewPkgCollector()
	p.BaseURL = "https://pkg.go.dev"
	return p.Get(name)
}
