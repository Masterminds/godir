package pathutil

import (
	"go/build"
	"os"
	"path/filepath"
	"strings"
)

// Gopath gets GOPATH from environment and return the most relevant path.
//
// A GOPATH can contain a colon-separated list of paths. This retrieves the
// GOPATH and returns only the FIRST ("most relevant") path.
//
// This should be used carefully. If, for example, you are looking for a package,
// you may be better off using Gopaths.
func Gopath() string {
	gopaths := Gopaths()
	if len(gopaths) == 0 {
		return ""
	}
	return gopaths[0]
}

// Gopaths retrieves the Gopath as a list when there is more than one path
// listed in the Gopath.
func Gopaths() []string {
	p := os.Getenv("GOPATH")
	p = strings.Trim(p, string(filepath.ListSeparator))
	return filepath.SplitList(p)
}

// Name returns the name of the current package.
func Name(wd string) string {
	for _, gp := range Gopaths() {
		gp = filepath.Join(gp, "src")
		if strings.HasPrefix(wd, gp) {
			rel, err := filepath.Rel(gp, wd)
			if err == nil {
				return rel
			}
		}
	}
	return wd
}

func WhichGopath(wd string) string {
	for _, gp := range Gopaths() {
		sp := filepath.Join(gp, "src")
		if strings.HasPrefix(wd, sp) {
			return gp
		}
	}
	return ""
}

var IgnorePrefix = []string{
	".",
	"_",
	"vendor/",
	"testdata/",
}

func Ignore(wd string) bool {
	if wd == "vendor" || wd == "testdata" {
		return true
	}
	for _, s := range IgnorePrefix {
		if strings.HasPrefix(wd, s) {
			return true
		}
	}
	return false
}

func HasSrc(wd string, any bool) bool {
	p, err := build.Default.ImportDir(wd, 0)
	if err != nil {
		return false
	}
	if len(p.GoFiles) > 0 {
		return true
	}

	if !any {
		return false
	}

	if len(p.TestGoFiles) > 0 {
		return true
	}
	if len(p.XTestGoFiles) > 0 {
		return true
	}
	return false
}

func Subpaths(base string, onlySrc bool) []string {
	paths := []string{}
	filepath.Walk(base, func(p string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			if base == p {
				return nil
			}
			tp, err := filepath.Rel(base, p)
			if err != nil {
				return err
			}
			if Ignore(tp) {
				return filepath.SkipDir
			}
			if onlySrc && !HasSrc(p, false) {
				return nil
			}
			paths = append(paths, p)
		}
		return nil
	})
	return paths
}
