package pkgproxy_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/pkgproxy"
)

func TestRetrievePackageRepositoryAddress(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("testdata/gopkg-in-tomb-v1.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
		}
		_, err = io.Copy(w, f)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
		}
	}))
	defer ts.Close()

	pc := pkgproxy.NewPkgCollector()
	pc.Collector.AllowedDomains = []string{}
	pc.BaseURL = ts.URL

	got := pc.Get("gopkg.in/tomb.v1")
	want := pkgproxy.Package{
		Name:                   "gopkg.in/tomb.v1",
		Repository:             "github.com/go-tomb/tomb",
		Version:                "v1.0.0",
		PublishedDate:          "Oct 24, 2014",
		License:                "BSD-3-Clause",
		Imports:                "3",
		ImportedBy:             "685",
		ValidGoMod:             "No",
		RedistributableLicense: "Yes",
		TaggedVersion:          "No",
		StableVersion:          "No",
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
