package filter

import (
	"path"
	"strings"
)

type Filter struct {
	Extensions []string
	Languages  []string
	Exclude    []string
	RestrictTo []string

	extensions map[string]struct{}
	langExt    map[string]struct{}
}

func (f *Filter) CanTake(filename string) bool {
	if len(f.Extensions) > 0 {
		if len(f.extensions) == 0 {
			f.extensions = make(map[string]struct{})
			for _, ext := range f.Extensions {
				f.extensions[ext] = struct{}{}
			}
		}

		ext := path.Ext(filename)
		if _, ok := f.extensions[ext]; !ok {
			return false
		}
	}

	if len(f.Languages) > 0 {
		if len(f.langExt) == 0 {
			f.langExt = make(map[string]struct{})
			for _, l := range f.Languages {
				for _, ext := range languages[strings.ToLower(l)].Extensions {
					f.langExt[ext] = struct{}{}
				}
			}
		}

		ext := path.Ext(filename)
		if _, ok := f.langExt[ext]; !ok {
			return false
		}
	}

	for _, exclude := range f.Exclude {
		match, err := path.Match(exclude, filename)
		if err != nil {
			panic(err)
		}
		if match {
			return false
		}
	}

	if len(f.RestrictTo) > 0 {
		restrictedFound := false
		for _, glob := range f.RestrictTo {
			match, err := path.Match(glob, filename)
			if err != nil {
				panic(err)
			}
			if match {
				restrictedFound = true
				break
			}
		}

		if !restrictedFound {
			return false
		}
	}

	return true
}
