package gitfame

import (
	"strings"
)

type CommitInfo struct {
	Author string
	Lines  int
}

func (g *GitContext) Blame(filename string) (map[string]*CommitInfo, error) {
	command := g.makeCommand("git", "blame", filename, "--porcelain", "-c", g.Revision)

	result, err := command.Output()
	if err != nil {
		return nil, err
	}

	commits := make(map[string]*CommitInfo)

	lines := strings.Split(string(result), "\n")
	for i := 0; i+1 < len(lines); {
		line := lines[i]
		entries := strings.Split(line, " ")

		sha := entries[0]
		if _, ok := commits[sha]; !ok {
			if g.UseCommiter {
				commits[sha] = &CommitInfo{
					Author: lines[i+5][10:],
				}
			} else {
				commits[sha] = &CommitInfo{
					Author: lines[i+1][7:],
				}
			}
			if lines[i+10][:8] != "filename" {
				i++
			}
			i += 12
		} else {
			i += 2
		}

		commits[sha].Lines++
	}

	return commits, err
}
