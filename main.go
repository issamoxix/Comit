package main

import (
	ai "commit_helper/services/openai"
	"commit_helper/services/utils"
	"fmt"
	"net/http"
	"os"
	"os/exec"
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

	if len(os.Args) > 1 && os.Args[1] == "update" {
		err := selfUpdate()
		if err != nil {
			fmt.Println("Update failed:", err)
		} else {
			fmt.Println("Successfully updated to the latest version.")
		}
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("Version: %s", utils.Version)
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "-b" {
		ai.GetBranchNames(os.Args[2])
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "help" {
		fmt.Println("Usage: comit [command]\nCommands: \n\tupdate   : To update the app\n\tversion  : To see the version of the app\n\t-b       : To generate branch name\n\thelp     : To see this help message")
		return
	}

	cmd := exec.Command("git", "--no-pager", "diff", "--staged")
	output, err := cmd.Output()
	if len(string(output)) == 0 {
		fmt.Println("No staged changes found.\nPlease stage your changes and try again.")
		return
	}

	// GetCommitMessage function do not return nothing know we call inside of it SelectCommitType function
	ai.GetCommitMessage(string(output))
	if err != nil {
		fmt.Println(err)
		return
	}
}
