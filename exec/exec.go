package exec

import (
	"log"
	"os"
	"os/exec"

	"github.com/42core-team/bot-client/fail"
	"github.com/42core-team/bot-client/rabbitmq"
)

func BuildBot() error {
	log.Println("Building bot...")
	cmd := exec.Command("make", "-C", "repo", "bot")
	stdout, err := cmd.StdoutPipe()
	fail.OnError(err, "Failed to get stdout pipe")
	stderr, err := cmd.StderrPipe()
	fail.OnError(err, "Failed to get stderr pipe")
	go rabbitmq.StreamLogs(stdout)
	go rabbitmq.StreamLogs(stderr)
	err = cmd.Start()
	fail.OnError(err, "Failed to start build")
	rabbitmq.SendStatus(rabbitmq.STATUS_BUILDING)
	return cmd.Wait()
}

func RunBot() error {
	log.Println("Running bot...")
	cmd := exec.Command("./repo/bot", os.Getenv("PLAYER_ID"))
	stdout, err := cmd.StdoutPipe()
	fail.OnError(err, "Failed to get stdout pipe")
	stderr, err := cmd.StderrPipe()
	fail.OnError(err, "Failed to get stderr pipe")
	go rabbitmq.StreamLogs(stdout)
	go rabbitmq.StreamLogs(stderr)
	err = cmd.Start()
	fail.OnError(err, "Failed to start bot")
	rabbitmq.SendStatus(rabbitmq.STATUS_RUNNING)
	log.Println("Bot is running")
	return cmd.Wait()
}
