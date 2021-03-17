package gitfame

func (g *GitContext) Log(filename string) (string, CommitInfo, error) {
	res, err := g.makeCommand("git", "log", g.Revision, "--format=%H%an", "-n", "1", "--", filename).Output()
	if err != nil {
		return "", CommitInfo{}, err
	}

	sha := string(res[:40])
	author := res[40 : len(res)-1]

	return sha, CommitInfo{
		Author: string(author),
		Lines:  0,
	}, nil
}
