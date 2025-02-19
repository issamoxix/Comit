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

	cmd := exec.Command("git", "--no-pager", "diff", "--staged")
	output, err := cmd.Output()
	if len(string(output)) == 0 {
		fmt.Println("No staged changes found.\nPlease stage your changes and try again.")
		return
	}
	commitMessage := ai.GetCommitMessage(string(output))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(commitMessage)

}
