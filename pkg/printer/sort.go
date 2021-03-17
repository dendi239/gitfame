package printer

import (
	"fmt"

	gitfame "github.com/dendi239/gitfame/pkg/git"
)

type sortAuthors struct {
	key     string
	authors []gitfame.AuthorInfo
}

func (a sortAuthors) Len() int { return len(a.authors) }
func (a sortAuthors) Swap(i, j int) {
	a.authors[i], a.authors[j] = a.authors[j], a.authors[i]
}

func (a sortAuthors) Less(i, j int) bool {
	for _, key := range []string{
		a.key, "lines", "commits", "files", "name",
	} {
		switch key {
		case "lines":
			if a.authors[i].Lines != a.authors[j].Lines {
				return a.authors[i].Lines > a.authors[j].Lines
			}

		case "commits":
			if a.authors[i].Commits != a.authors[j].Commits {
				return a.authors[i].Commits > a.authors[j].Commits
			}

		case "files":
			if a.authors[i].Files != a.authors[j].Files {
				return a.authors[i].Files > a.authors[j].Files
			}

		case "name":
			if a.authors[i].Name != a.authors[j].Name {
				return a.authors[i].Name < a.authors[j].Name
			}

		default:
			panic(fmt.Errorf("unknown key to sort by: %s", key))
		}
	}

	return false
}
