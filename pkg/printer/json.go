package printer

import (
	"encoding/json"
	"fmt"
	"io"

	gitfame "github.com/dendi239/gitfame/pkg/git"
)

type jsonFormatter struct{}

func (j jsonFormatter) Write(w io.Writer, authors []gitfame.AuthorInfo) error {
	data, err := json.Marshal(authors)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	if err != nil {
		return err
	}

	return nil
}

type jsonLinesFormatter struct{}

func (j jsonLinesFormatter) Write(w io.Writer, authors []gitfame.AuthorInfo) error {
	for _, a := range authors {
		data, err := json.Marshal(a)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(w, "%s\n", data)
		if err != nil {
			return err
		}
	}
	return nil
}
