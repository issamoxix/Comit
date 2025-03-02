package main

import (
	ai "commit_helper/services/openai"
	"commit_helper/services/utils"
	"commit_helper/services/utils/tools"
	"fmt"
	"os"
	"strings"
)

func main() {
	messages := make(chan string)
	go func() {
		latestVersion := utils.GetLatestVersion()
		if strings.Compare(latestVersion, utils.Version) == -1 {
			messages <- "🚀 A new version (%s) is available! Please update."
		}
	}()
	msg := <-messages
	fmt.Println(msg)
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "update":
			err := utils.SelfUpdate()
			if err != nil {
				fmt.Println("Update failed:", err)
			} else {
				fmt.Println("Successfully updated to the latest version.")
			}
			return

		case "version":
			fmt.Printf("Current Version: %s\n", utils.Version)
			return

		case "-b":
			ai.GetBranchNames(os.Args[2])
			return

		case "help":
			fmt.Println("Usage: comit [command]\nCommands: \n\tupdate   : To update the app\n\tversion  : To see the version of the app\n\t-b       : To generate branch name\n\thelp     : To see this help message")
			return

		case "c":
			ai.GetPromptResponse(os.Args[2])
			return

		case "-c":
			ai.GetPromptResponse(os.Args[2])
			return

		default:
			fmt.Printf("Unknown command: %s\nUse \"comit help\" to see available commands", os.Args[1])
			return
		}
	}

	tools.RunCommit()
}
