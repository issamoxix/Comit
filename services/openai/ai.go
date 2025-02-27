package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"

	"github.com/manifoldco/promptui"
)

type commitResponse struct {
	// change the message type to a table of string
	Message []string `json:"message"`
}

type RequestData struct {
	Code string `json:"code"`
}

type RequestBranchName struct {
	Context string `json:"context"`
}

type BranchResponse struct {
	Branch []string `json:"branch"`
}

func GetCommitMessage(content string) string {
	url := "https://comit.issamcloud.com"

	payload := RequestData{
		Code: content,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return "Error: " + err.Error()
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println(err)
		return "Error: " + err.Error()
	}

	if resp.StatusCode != 200 {
		fmt.Println("Error: " + resp.Status)
		return "Error: " + resp.Status
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "Error: " + err.Error()
	}

	var data commitResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return "Error: " + err.Error()
	}

	// declaring a table to store the commit messages
	var commitMessages []string
	// geting the messages form the response and adding it to the commitMessages
	commitMessages = append(commitMessages, data.Message...)
	// calling the function to select the commit type and passing the commitMessages
	SelectCommitType(commitMessages)

	return ""
}

func GetBranchNames(context string) string {
	url := "https://comit.issamcloud.com/branch"
	payload := RequestBranchName{
		Context: context,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return "Error: " + err.Error()
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return "Error: " + err.Error()
	}

	if resp.StatusCode != 200 {
		fmt.Println("Error: " + resp.Status)
		return "Error: " + resp.Status
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "Error: " + err.Error()
	}

	var data BranchResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return "Error: " + err.Error()
	}

	for _, branch := range data.Branch {
		fmt.Println(branch)
	}
	return ""
}

func SelectCommitType(commitMessages []string) string {

	prompt := promptui.Select{
		Label: "Select commit message",
		Items: commitMessages,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return ""
	}
	// runing the commit that the user selected {result}
	cmd := exec.Command("sh", "-c", fmt.Sprintf("git commit -m %q", result))
	fmt.Printf("You executed: git commit -m %q\n", result)
	_, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
		return "Error: " + err.Error()
	}
	return ""
}
