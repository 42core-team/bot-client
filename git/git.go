package git

import (
	"os"

	"github.com/go-git/go-git/v5"
)

func Clone(url, dir string) (*git.Repository, error) {
	return git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
		SingleBranch: true,
		Depth: 1,
		Progress: os.Stdout,
		// Auth: &,
	})
}
