package main

import (
	"fmt"

	"github.com/42core-team/bot-client/exec"
	"github.com/42core-team/bot-client/git"
	"github.com/42core-team/bot-client/rabbitmq"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	rabbitmq.Init()
	pullBuild()
	rabbitmq.Close()
}

func pullBuild() {
	_, err := git.Clone("https://github.com/42core-team/my-core-bot")
	if err != nil {
		fmt.Println(err)
	}

	exec.BuildBot()
	exec.RunBot()
}
