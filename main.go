package main

import (
	"fmt"
	"time"

	"github.com/42core-team/bot-client/git"
)

func main() {
	start := time.Now()

	_, err := git.Clone("git@github.com:GREAULouen/miniRT.git", "asd")
	if err != nil {
		fmt.Println(err)
	}

	elapsed := time.Since(start)
	fmt.Printf("git.Clone took %s\n", elapsed)
}
