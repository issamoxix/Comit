package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type commitResponse struct {
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

var url = "https://comit.issamcloud.com"

type CommitMessageSelector interface {
	SelectCommitMessage(messages []string) error
}

func GetCommitMessage(content string, selector CommitMessageSelector) string {

	payload := RequestData{
		Code: content,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "Error: " + err.Error()
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		return "Error: " + err.Error()
	}

	if resp.StatusCode != 200 {
		return "Error: " + resp.Status
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error: " + err.Error()
	}

	var data commitResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "Error: " + err.Error()
	}

	if err := selector.SelectCommitMessage(data.Message); err != nil {
		return "Error: " + err.Error()
	}

	return "Ok"
}

func GetBranchNames(context string) string {
	url = url + "/branch"
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
