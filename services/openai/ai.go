package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/alecthomas/chroma/quick"
)

type commitResponse struct {
	Message []string `json:"message"`
}

type RequestData struct {
	Code string `json:"code"`
}

type RequestAgentResponse struct {
	Prompt string `json:"prompt"`
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

	fmt.Println("\nBranches:")

	for _, branch := range data.Branch {
		fmt.Printf("  - " + branch + "\n")
	}
	return ""
}

func GetPromptResponse(prompt string) {
	url = url + "/agent"
	payload := RequestAgentResponse{
		Prompt: prompt,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("Error: " + resp.Status)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data RequestAgentResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(data.Prompt, "\n")
	var codeStartIndex int
	var language string
	for index, line := range lines {
		if strings.Contains(line, "```") && len(line) > 3 {
			codeStartIndex = index + 1
			language = line[3:]
			continue
		}
		if strings.Contains(line, "```") && len(line) == 3 {
			err := quick.Highlight(os.Stdout, strings.Join(lines[codeStartIndex:index], "\n"), language, "terminal16m", "monokai")
			if err != nil {
				fmt.Println("Error:", err)
			}
			codeStartIndex = 0
			continue
		}
		if codeStartIndex > 0 && index >= codeStartIndex {
			continue
		}
		fmt.Println(line)
	}
	return
}
