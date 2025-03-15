package main

import (
	ai "commit_helper/services/openai"
	"commit_helper/services/utils"
	"commit_helper/services/utils/auth"
	"commit_helper/services/utils/tools"
	"fmt"
	"os"
)

func main() {

	messages := make(chan string)
	go func() {
		latestVersion := utils.GetLatestVersion()
		if latestVersion > utils.Version {
			messages <- fmt.Sprintf("🚀 A new version (%s) is available! Please update.", latestVersion)
		}
		close(messages)
	}()
	msg := <-messages
	if msg != "" {
		fmt.Println(msg)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "update", "u", "-u", "--update":
			_, err := utils.SelfUpdate()
			if err != nil {
				fmt.Println("Update failed:", err)
				return
			}
			return
		case "version", "v", "-v", "--version":
			fmt.Printf("Current Version: %s", utils.Version)
			return

		case "-b":
			ai.GetBranchNames(os.Args[2])
			return

		case "help", "h", "-h", "--help":
			fmt.Println("Usage: comit [command]\n" +
				"Commands:\n" +
				"    update, u, -u, --update  	    : Update the application to the latest version.\n" +
				"    version, v, -v, --version	    : Display the current version of the application.\n" +
				"    -b <arg>                 	    : Generate branch name using the given argument.\n" +
				"    -c, c <arg>              	    : Get a prompt response based on the given argument.\n" +
				"    login, -login, --login <token> : Login to the application.\n" +
				"    -l, l, live, --live       	    : Get a live prompt response.\n" +
				"    help, h, -h, --help      	    : Show this help message.")
			return

		case "-c", "c":
			ai.GetPromptResponse(os.Args[2])
			return

		case "-l", "l", "--live", "live":
			token, err := auth.GetToken()
			if err != nil {
				fmt.Println("Something went wrong please try again")
				return
			}
			if token == "" {
				fmt.Println("Please login first using \"comit login <token>\"")
				return
			}
			ai.GetLivePromptResponse(token)
			return

		case "login", "-login", "--login":
			auth.StoreToken(os.Args[2])
			token, err := auth.GetToken()
			if err != nil {
				fmt.Println("Something went wrong please try again")
				return
			}
			fmt.Println("Login successful! Token stored. ", token)
			return

		default:
			fmt.Printf("Unknown command: %s\nUse \"comit help\" to see available commands", os.Args[1])
			return
		}
	}

	tools.RunCommit()
}
