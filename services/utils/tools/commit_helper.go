package tools

import (
	ai "commit_helper/services/openai"
	"commit_helper/services/utils/auth"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/manifoldco/promptui"
)

type RealSelector struct{}

func (RealSelector) SelectCommitMessage(commitMessages []string) error {

	if len(commitMessages) == 0 || allEmpty(commitMessages) {
		RunCommit()
		return nil
	}

	prompt := promptui.Select{
		Label: "Select commit message",
		Items: append([]string{"Refresh"}, commitMessages...),
	}

	_, result, err := prompt.Run()
	if err != nil {
		return err
	}

	if result == "Refresh" {
		RunCommit()
		return nil
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("powershell", "-Command", fmt.Sprintf("git commit -m %q", result))
	case "darwin":
		cmd = exec.Command("sh", "-c", fmt.Sprintf("git commit -m %q", result))
	default:
		cmd = exec.Command("powershell", "-Command", fmt.Sprintf("git commit -m %q", result))
	}
	fmt.Printf("You executed: git commit -m %q\n", result)
	_, err = cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

func CheckStage() (string, error) {

	cmd := exec.Command("git", "--no-pager", "diff", "--staged")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to check staged changes: %v", err)
	}

	if len(output) == 0 {
		return "", fmt.Errorf("no staged changes found. Please stage your changes and try again")
	}

	return string(output), nil
}

func RunCommit() {
	output, err := CheckStage()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	token, _ := auth.GetToken()
	if token == "" {
		token = "default"
	}
	messageStatus := ai.GetCommitMessage(output, RealSelector{}, token)
	if messageStatus != "Ok" {
		fmt.Println("Something went wrong please try again")
		return
	}
}

func allEmpty(messages []string) bool {
	for _, msg := range messages {
		if msg != "" {
			return false
		}
	}
	return true
}
