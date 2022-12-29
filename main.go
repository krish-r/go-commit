package main

import (
	"bufio"
	"os"
)

func main() {
	printWorkingDir()

	reader := bufio.NewReader(os.Stdin)
	commit := Commit{
		changeType: getChangeType(reader),
		emoji:      getEmoji(reader),
		message:    getCommitMessage(reader),
	}

	performCommit(reader, &commit)
}
