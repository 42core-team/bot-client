package main

import (
	"fmt"

	"github.com/42core-team/bot-client/exec"
	"github.com/42core-team/bot-client/git"
)

func main() {
	_, err := git.Clone("https://github.com/42core-team/my-core-bot")
	if err != nil {
		fmt.Println(err)
	}

	exec.BuildBot()
	exec.RunBot()
}
