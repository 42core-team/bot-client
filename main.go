package main

import (
	"log"
	"os"

	"github.com/42core-team/bot-client/exec"
	"github.com/42core-team/bot-client/git"
	"github.com/42core-team/bot-client/rabbitmq"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	rabbitmq.Init()
	defer rabbitmq.Close()

	pullBuild()
}

func pullBuild() {
	_, err := git.Clone(os.Getenv("REPO_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	exec.BuildBot()
	exec.RunBot()
}
