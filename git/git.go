package git

import (
	"os"

	"github.com/go-git/go-git/v5"
)

func Clone(url string) (*git.Repository, error) {
	dir := "./my-core-bot/repo"
	os.RemoveAll(dir)
	return git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
		SingleBranch: true,
		Depth: 1,
		Progress: os.Stdout,
		// Auth: &,
	})
}
