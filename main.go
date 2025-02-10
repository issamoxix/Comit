package main

import (
	ai "commit_helper/services/openai"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "--no-pager", "diff", "--staged")
	output, err := cmd.Output()
	commitMessage := ai.GetCommitMessage(string(output))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(commitMessage)

}
