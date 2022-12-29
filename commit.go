package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/krish-r/go-commit/color"
)

var (
	quitChar = color.Colorize("q", color.BrightRed)
	quit     = color.ColorizeFirstChar("Quit", color.BrightRed)
	yes      = color.ColorizeFirstChar("yes", color.Green)
	no       = color.ColorizeFirstChar("no", color.Yellow)
)

func printOptions[T Options](message string, options []T) {
	fmt.Println(color.Colorize(message, color.Green))
	for _, option := range options {
		option.print()
	}
	fmt.Printf("%s: %s\n", quitChar, quit)
}

func getChangeType(reader *bufio.Reader) string {
	printOptions("Change Type:", changeTypes)
	fmt.Println(color.Colorize("Select a change type:", color.Green))

	changeType := readOptions(OptionRange{min: changeTypes[0].id, max: changeTypes[len(changeTypes)-1].id}, reader)
	return changeTypes[changeType-1].shortDesc + ":"
}

func getEmoji(reader *bufio.Reader) string {
	printOptions("Emoji:", emojis)
	fmt.Println(color.Colorize("Select an emoji:", color.Green))
	emoji := readOptions(OptionRange{min: emojis[0].id, max: emojis[len(emojis)-1].id}, reader)
	return emojis[emoji].code
}

func getCommitMessage(reader *bufio.Reader) string {
	fmt.Printf("%s (%s to %s)\n", color.Colorize("Enter a commit message:", color.Green), quitChar, quit)
	message := readMessage(fmt.Sprintf("%s (%s to %s) ...", color.Colorize("Please enter a valid commit message", color.Red), quitChar, quit), reader)
	return message
}

func performCommit(reader *bufio.Reader, commit *Commit) {
	commitMessage := strings.Join(toSlice(*commit), " ")
	fmt.Printf("%s %s\n", color.Colorize("Commit Message:", color.Green), color.Colorize(commitMessage, color.White))
	fmt.Printf("%s (%s)/(%s)/(%s): ", color.Colorize("Ready to commit?:", color.Green), yes, no, quit)

	confirmation := readMessage(fmt.Sprintf("%s (%s to %s) ...", color.Colorize("Please enter a valid option", color.Red), quitChar, quit), reader)
	switch strings.ToUpper(confirmation) {
	case "Y":
		cmd := exec.Command("git", "commit", "-m", commitMessage)
		fmt.Printf("Executing command: %s\n", color.Colorize(cmd.String(), color.Green))
		err := cmd.Run()
		if err != nil {
			fmt.Println(color.Colorize(fmt.Sprintf("Encountered error: %q while performing commit", err.Error()), color.Red))
		}
	case "N":
		updateMessage(reader, commit)
		performCommit(reader, commit)
	case "Q":
		os.Exit(0)
	default:
		fmt.Println(color.Colorize(fmt.Sprintf("Received %q. Retry ...", confirmation), color.Red))
		performCommit(reader, commit)
	}
}

func readOptions(optionRange OptionRange, reader *bufio.Reader) int {
	var option int
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(color.Colorize("Encountered an error while reading user input, Please Retry ...", color.Red))
			continue
		}
		input = strings.Replace(input, "\n", "", -1)
		if "Q" == strings.ToUpper(input) {
			os.Exit(0)
		}
		option, err = strconv.Atoi(input)
		if err != nil {
			fmt.Printf("%s (%s to %s) ...\n", color.Colorize("Please Enter a valid number", color.Red), quitChar, quit)
			continue
		}
		if option < optionRange.min || option > optionRange.max {
			fmt.Printf("%s (%s to %s) ...\n", color.Colorize(fmt.Sprintf("Please Enter a valid number between %d and %d", optionRange.min, optionRange.max), color.Red), quitChar, quit)
			continue
		}
		break
	}
	return option
}

func readMessage(prompt string, reader *bufio.Reader) string {
	var message string
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(color.Colorize("Encountered an error while reading user input, Please Retry ...", color.Red))
			continue
		}
		message = strings.TrimSpace(strings.Replace(input, "\n", "", -1))
		if message == "" {
			fmt.Println(prompt)
			continue
		}
		if "Q" == strings.ToUpper(message) {
			os.Exit(0)
		}
		break
	}
	return message
}

func updateMessage(reader *bufio.Reader, commit *Commit) {
	fmt.Printf("%s (%s)/(%s)/(%s)/(%s)/(%s)/(%s): ",
		color.Colorize("What would you like to change?:", color.Green),
		color.ColorizeFirstChar("type", color.Yellow),
		color.ColorizeFirstChar("emoji", color.Yellow),
		color.ColorizeFirstChar("message", color.Yellow),
		color.ColorizeFirstChar("all", color.Yellow),
		color.ColorizeFirstChar("cancel", color.Yellow),
		quit,
	)
	confirmation := readMessage(fmt.Sprintf("%s (%s to %s) ...", color.Colorize("Please enter a valid option", color.Red), quitChar, quit), reader)
	switch strings.ToUpper(confirmation) {
	case "T":
		commit.changeType = getChangeType(reader)
	case "E":
		commit.emoji = getEmoji(reader)
	case "M":
		commit.message = getCommitMessage(reader)
	case "A":
		commit.changeType = getChangeType(reader)
		commit.emoji = getEmoji(reader)
		commit.message = getCommitMessage(reader)
	case "C":
		return
	case "Q":
		return
	default:
		fmt.Println(color.Colorize(fmt.Sprintf("Received %q. Retry ...", confirmation), color.Red))
		updateMessage(reader, commit)
	}
}

func printWorkingDir() {
	presentWorkingDir, err := os.Getwd()
	if err != nil {
		fmt.Println(color.Colorize("Error encountered while retrieving current working directory", color.Red))
	}
	fmt.Printf("Current Directory: %s\n", presentWorkingDir)
}

func toSlice(cs Commit) []string {
	if cs.emoji == "" {
		return []string{cs.changeType, cs.message}
	}
	return []string{cs.changeType, cs.emoji, cs.message}
}
