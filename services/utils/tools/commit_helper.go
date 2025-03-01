package tools

import (
	ai "commit_helper/services/openai"
	"fmt"
	"os/exec"

	"github.com/manifoldco/promptui"
)

type RealSelector struct{}

func (RealSelector) SelectCommitMessage(commitMessages []string) error {

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
	cmd := exec.Command("sh", "-c", fmt.Sprintf("git commit -m %q", result))
	fmt.Printf("You executed: git commit -m %q\n", result)
	_, err = cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

func CheckStage() string {

	cmd := exec.Command("git", "--no-pager", "diff", "--staged")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if len(output) == 0 {
		fmt.Println("No staged changes found.\nPlease stage your changes and try again.")
		return ""
	}

	return string(output)
}

func RunCommit() {
	stageStatus := CheckStage()
	if stageStatus == "" {
		return
	}

	messageStatus := ai.GetCommitMessage(CheckStage(), RealSelector{})
	if messageStatus != "Ok" {
		fmt.Println("Something went wrong please try again")
		return
	}
}
