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
	log.Println("Bot run successfully")
}

func pullBuild() {
	_, err := git.Clone(os.Getenv("REPO_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	err = exec.BuildBot()
	if err != nil {
		log.Fatalln("Build failed!\n" + err.Error())
	}

	err = exec.RunBot()
	if err != nil {
		log.Fatalln("Run failed!\n" + err.Error())
	}
}
