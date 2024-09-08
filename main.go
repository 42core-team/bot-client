package main

import (
	"fmt"
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
		fmt.Println(err)
	}

	exec.BuildBot()
	exec.RunBot()
}
