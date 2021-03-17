package printer

import (
	"encoding/csv"
	"io"
	"strconv"

	gitfame "github.com/dendi239/gitfame/pkg/git"
)

type csvFormatter struct{}

func (c csvFormatter) Write(w io.Writer, authors []gitfame.AuthorInfo) error {
	wc := csv.NewWriter(w)
	defer wc.Flush()
	err := wc.Write([]string{
		"Name",
		"Lines",
		"Commits",
		"Files",
	})
	if err != nil {
		return err
	}

	for _, a := range authors {
		err := wc.Write([]string{
			a.Name,
			strconv.Itoa(a.Lines),
			strconv.Itoa(a.Commits),
			strconv.Itoa(a.Files),
		})
		if err != nil {
			return err
		}
	}
	return nil
}
