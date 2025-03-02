package main

import (
	ai "commit_helper/services/openai"
	"commit_helper/services/utils"
	"commit_helper/services/utils/tools"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/inconshreveable/go-update"
)

var UpdateLink = "https://github.com/issamoxix/Comit/releases/download/%s/%s"

func selfUpdate() error {
	fmt.Println("Checking for updates...")
	updateVersion := utils.GetVersion()
	var fileName string
	switch runtime.GOOS {
	case "windows":
		fileName = "comit.exe"
	case "darwin":
		fileName = "comit"
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	latestBinaryURL := fmt.Sprintf(UpdateLink, updateVersion, fileName)
	resp, err := http.Get(latestBinaryURL)
	if err != nil {
		return fmt.Errorf("failed to download update: %v", err)
	}
	defer resp.Body.Close()

	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		return fmt.Errorf("failed to apply update: %v", err)
	}

	fmt.Println("Update successful!")
	return nil
}

func main() {

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "update":
			err := selfUpdate()
			if err != nil {
				fmt.Println("Update failed:", err)
			} else {
				fmt.Println("Successfully updated to the latest version.")
			}
			return

		case "version":
			fmt.Printf("Version: %s", utils.Version)
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
