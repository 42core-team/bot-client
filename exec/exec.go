package exec

import (
	"log"
	"os/exec"
)

func BuildBot() error {
	log.Println("Building bot...")
	cmd := exec.Command("make", "build")
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()
	return cmd.Run()
}

func RunBot() error {
	log.Println("Running bot...")
	cmd := exec.Command("./bot")
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()
	return cmd.Run()
}
