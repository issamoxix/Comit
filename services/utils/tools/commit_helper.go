package tools

import (
	"fmt"
	"os/exec"

	"github.com/manifoldco/promptui"
)

type RealSelector struct{}

func (RealSelector) SelectCommitMessage(commitMessages []string) error {

	prompt := promptui.Select{
		Label: "Select commit message",
		Items: commitMessages,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return err
	}

	cmd := exec.Command("sh", "-c", fmt.Sprintf("git commit -m %q", result))
	fmt.Printf("You executed: git commit -m %q\n", result)
	_, err = cmd.Output()
	if err != nil {
		return err
	}
	return nil
}
