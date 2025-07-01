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

func TestRetrievePackageSourceCodeURL(t *testing.T) {
	t.Parallel()

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	pc.BaseURL = ts.URL

	got := pc.Get("gopkg.in/tomb.v1")
	want := pkgproxy.Package{
		Address: "gopkg.in/tomb.v1",
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
