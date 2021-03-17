package gitfame

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

type AuthorInfo struct {
	Name    string `json:"name"`
	Commits int    `json:"commits"`
	Lines   int    `json:"lines"`
	Files   int    `json:"files"`
}

type GitContext struct {
	Revision       string
	Directory      string
	UseCommiter    bool
	predicate      func(string) bool
	maxOpenedFiles int
}

type GitContextInterface interface {
	Gitfame() ([]AuthorInfo, error)
}

func NewContext(predicate func(string) bool) GitContext {
	return GitContext{
		Revision:       "",
		Directory:      "",
		predicate:      predicate,
		maxOpenedFiles: runtime.NumCPU(),
	}
}

func (g *GitContext) makeCommand(name string, args ...string) *exec.Cmd {
	command := exec.Command(name, args...)
	command.Dir = g.Directory
	return command
}

func (g *GitContext) Gitfame() ([]AuthorInfo, error) {
	start := time.Now()

	filenames, err := g.LsTree()
	if err != nil {
		return nil, err
	}

	fmt.Fprintf(
		os.Stderr, "[%5.2fs] %d filenames generated!\n",
		time.Since(start).Seconds(), len(filenames))

	commits := make(map[string]*CommitInfo)
	authoredFiles := make(map[string]int)
	authoredCommits := make(map[string]map[string]struct{})

	var mu sync.Mutex

	runRunRun := make(chan struct{}, g.maxOpenedFiles)
	for j := 0; j < g.maxOpenedFiles; j++ {
		runRunRun <- struct{}{}
	}

	var wg sync.WaitGroup
	wg.Add(len(filenames))

	for _, file := range filenames {
		go func(file string) {
			<-runRunRun

			fileSha, fileAuthor, err := g.Log(file)
			if err != nil {
				panic(err)
			}
			desc, err := g.Blame(file)
			if err != nil {
				panic(err)
			}

			mu.Lock()
			lookedAuthors := make(map[string]struct{})

			if len(desc) == 0 {
				authoredFiles[fileAuthor.Author]++
				if _, ok := authoredCommits[fileAuthor.Author]; !ok {
					authoredCommits[fileAuthor.Author] = make(map[string]struct{})
				}
				authoredCommits[fileAuthor.Author][fileSha] = struct{}{}
			}

			for sha, fileVal := range desc {
				if _, ok := lookedAuthors[fileVal.Author]; !ok {
					lookedAuthors[fileVal.Author] = struct{}{}
					authoredFiles[fileVal.Author]++
				}

				if _, ok := authoredCommits[fileVal.Author]; !ok {
					authoredCommits[fileVal.Author] = make(map[string]struct{})
				}
				authoredCommits[fileVal.Author][sha] = struct{}{}

				if val, ok := commits[sha]; ok {
					val.Lines += fileVal.Lines
				} else {
					commits[sha] = fileVal
				}
			}
			mu.Unlock()

			wg.Done()
			runRunRun <- struct{}{}
		}(file)
	}

	wg.Wait()
	for j := 0; j < g.maxOpenedFiles; j++ {
		<-runRunRun
	}
	fmt.Fprintf(os.Stderr, "[%5.2fs] files analyzed!\n", time.Since(start).Seconds())

	authors := make(map[string]*AuthorInfo)
	for author, authorCommits := range authoredCommits {
		authors[author] = &AuthorInfo{
			Name:    author,
			Commits: len(authorCommits),
			Lines:   0,
			Files:   authoredFiles[author],
		}
	}
	for author, files := range authoredFiles {
		if _, ok := authors[author]; !ok {
			authors[author] = &AuthorInfo{
				Name:    author,
				Commits: 0,
				Lines:   0,
				Files:   files,
			}
		}
	}

	for _, info := range commits {
		if author, ok := authors[info.Author]; ok {
			author.Lines += info.Lines
		} else {
			authors[info.Author] = &AuthorInfo{
				Name:    info.Author,
				Commits: 0,
				Lines:   info.Lines,
			}
		}
	}

	authorsList := make([]AuthorInfo, 0)
	for _, author := range authors {
		authorsList = append(authorsList, *author)
	}

	fmt.Fprintf(os.Stderr, "[%5.2fs] build list!\n", time.Since(start).Seconds())

	return authorsList, nil
}
