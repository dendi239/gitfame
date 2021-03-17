package printer

import (
	"fmt"
	"io"
	"text/tabwriter"

	gitfame "github.com/dendi239/gitfame/pkg/git"
)

type tabularFormatter struct{}

func (t tabularFormatter) Write(w io.Writer, authors []gitfame.AuthorInfo) error {
	wt := tabwriter.NewWriter(w, 0, 0, 1, ' ', 0)

	_, err := fmt.Fprintf(wt, "Name\tLines\tCommits\tFiles\n")
	if err != nil {
		return err
	}

	for _, author := range authors {
		_, err := fmt.Fprintf(
			wt, "%s\t%d\t%d\t%d\n",
			author.Name, author.Lines, author.Commits, author.Files)

		if err != nil {
			return err
		}
	}

	wt.Flush()
	return nil
}
