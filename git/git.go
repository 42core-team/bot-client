package git

import (
	"log"
	"os"

	"github.com/42core-team/bot-client/rabbitmq"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func Clone(url string) (*git.Repository, error) {
	log.Println("Cloning repository...")
	rabbitmq.SendStatus(rabbitmq.STATUS_PULLING)
	dir := "./repo"
	os.RemoveAll(dir)
	return git.PlainClone(dir, false, &git.CloneOptions{
		URL:          url,
		SingleBranch: true,
		Depth:        1,
		Progress:     os.Stdout,
		Auth: &http.BasicAuth{
			Username: "nil",
			Password: os.Getenv("GIT_ACCESS_TOKEN"),
		},
	})
}
