package pkgproxy

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type Package struct {
	Name                   string `json:"name"`
	Repository             string `json:"repository"`
	Version                string `json:"version"`
	PublishedDate          string `json:"publishedDate"`
	License                string `json:"license"`
	Imports                string `json:"imports"`
	ImportedBy             string `json:"importedBy"`
	ValidGoMod             string `json:"validGomod"`
	RedistributableLicense string `json:"redistributableLicense"`
	TaggedVersion          string `json:"taggedVersion"`
	StableVersion          string `json:"stableVersion"`
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

	// Imports
	var imports string
	p.Collector.OnHTML(".go-Main-headerDetailItem[data-test-id='UnitHeader-imports']", func(e *colly.HTMLElement) {
		imports = sanitizeImports(strings.TrimSpace(e.ChildText("a")))
	})

	// ImportedBy
	var importedBy string
	p.Collector.OnHTML(".go-Main-headerDetailItem[data-test-id='UnitHeader-importedby']", func(e *colly.HTMLElement) {
		importedBy = sanitizeImports(strings.TrimSpace(e.ChildText("a")))
	})

	// Package details
	var (
		details            []string
		validGoMod         string
		redistributableLic string
		taggedVersion      string
		stableVersion      string
	)

	p.Collector.OnHTML(".UnitMeta-details", func(e *colly.HTMLElement) {
		var v string
		e.ForEach("li", func(i int, h *colly.HTMLElement) {
			v = fmt.Sprintf("%s", h.ChildAttr("details>summary>img", "alt"))
			details = append(details, strings.TrimSpace(v))
		})
		if len(details) < 4 {
			return
		}
		validGoMod = sanitizeTrueFalse(details[0])
		redistributableLic = sanitizeTrueFalse(details[1])
		taggedVersion = sanitizeTrueFalse(details[2])
		stableVersion = sanitizeTrueFalse(details[3])
	})

	p.Collector.Visit(p.BaseURL + "/" + pkgName)

	return Package{
		Name:                   pkgName,
		Repository:             repo,
		Version:                version,
		PublishedDate:          published,
		License:                license,
		Imports:                imports,
		ImportedBy:             importedBy,
		ValidGoMod:             validGoMod,
		RedistributableLicense: redistributableLic,
		TaggedVersion:          taggedVersion,
		StableVersion:          stableVersion,
	}
}

func sanitizeTrueFalse(v string) string {
	if v == "unchecked" {
		return "No"
	}
	if v == "checked" {
		return "Yes"
	}
	return "Undetected"
}

func sanitizeDate(date string) string {
	chunks := strings.Split(strings.TrimSpace(date), ":")
	if len(chunks) < 2 {
		return "Undetected"
	}
	return strings.TrimSpace(chunks[1])
}

func sanitizeVersion(version string) string {
	chunks := strings.Split(strings.TrimSpace(version), ":")
	if len(chunks) < 2 {
		return "Undetected"
	}
	v := strings.TrimSpace(chunks[1])
	if strings.Contains(v, "-...-") {
		return strings.TrimSpace(strings.Split(v, "-")[0])
	}
	return v
}

func sanitizeImports(imports string) string {
	chunks := strings.Split(imports, ":")
	if len(chunks) < 2 {
		return "Undetected"
	}
	return strings.TrimSpace(chunks[1])
}

func Get(name string) Package {
	return NewPkgCollector().Get(name)
}

func GetJSON(name string) (string, error) {
	p := Get(name)
	data, err := json.Marshal(p)
	if err != nil {
		return "", fmt.Errorf("marshaling data: %w", err)
	}
	return string(data), nil
}
