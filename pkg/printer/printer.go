package printer

import (
	"io"
	"os"
	"sort"

	gitfame "github.com/dendi239/gitfame/pkg/git"
)

type Printer struct {
	OrderBy string
	Format  string
}

type Formatter interface {
	Write(w io.Writer, authors []gitfame.AuthorInfo) error
}

var (
	formatters = map[string]Formatter{
		"tabular":    tabularFormatter{},
		"json":       jsonFormatter{},
		"json-lines": jsonLinesFormatter{},
		"csv":        csvFormatter{},
	}
)

func (p *Printer) Print(authors []gitfame.AuthorInfo) {
	sort.Sort(sortAuthors{
		key:     p.OrderBy,
		authors: authors,
	})

	f := formatters[p.Format]
	err := f.Write(os.Stdout, authors)
	if err != nil {
		panic(err)
	}
}
