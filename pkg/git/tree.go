package gitfame

import (
	"strings"
)

func (g *GitContext) LsTree() ([]string, error) {
	result, err := g.makeCommand("git", "ls-tree", "-r", "--name-only", g.Revision).Output()
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return []string{}, nil
	}

	filenames := strings.Split(string(result[:len(result)-1]), "\n")
	filteredFilenames := make([]string, 0, len(filenames))
	for _, file := range filenames {
		if g.predicate(file) {
			filteredFilenames = append(filteredFilenames, file)
		}
	}

	return filteredFilenames, nil
}
