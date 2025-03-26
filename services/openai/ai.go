package ai

import (
	"bufio"
	"bytes"
	"commit_helper/services/utils"
	"commit_helper/services/utils/auth"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

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

type CommitMessageSelector interface {
	SelectCommitMessage(messages []string) error
}

type BranchSelector interface {
	SelectBranchMessage(messages []string, context string) error
}

func GetCommitMessage(content string, selector CommitMessageSelector, token string) string {
	var url = utils.ComitURL + "/commit" + "?token=" + token
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

func GetBranchNames(context string, selector BranchSelector) string {
	var url = utils.ComitURL + "/branch"
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

	if err := selector.SelectBranchMessage(data.Branch, context); err != nil {
		return "Error: " + err.Error()
	}
	return "Ok"
}

func GetPromptResponse(prompt string) {
	var context = "/agent"
	data := ApiResponse(prompt, context)
	if data == "" {
		fmt.Println("something went wrong please try again")
	}
	lines := strings.Split(data, "\n")
	PretterPromptResponse(lines)
}

func GetLivePromptResponse(token string) {
	context := "/live"
	context_id := []string{GenerateComitId()}
	fmt.Print("Hi! What would you like to do? (Type 'q' to quit): ")
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "q" {
			fmt.Println("You chose to quit. Goodbye!")
			break
		}
		data := ApiResponse(input, context, context_id...)
		if data == "" {
			fmt.Println("something went wrong please try again")
		}
		lines := strings.Split(data, "\n")
		fmt.Println("")
		PretterPromptResponse(lines)
		fmt.Print("\n(Type 'q' to quit): ")
	}
}

func ApiResponse(prompt string, context string, context_id ...string) string {
	token, _ := auth.GetToken()
	if token == "" {
		token = "default"
	}
	var url = utils.ComitURL + context + "?token=" + token

	if context_id != nil {
		url += "&context_id=" + context_id[0]
	}

	payload := RequestAgentResponse{
		Prompt: prompt,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return ""
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("Error: " + resp.Status)
		return ""
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var data RequestAgentResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return ""
	}
	return data.Prompt
}

func PretterPromptResponse(lines []string) {
	var codeStartIndex int
	var language string

	for index, line := range lines {
		line = strings.TrimSpace(line)
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
}

func GenerateComitId() string {
	now := time.Now()
	timestamp := fmt.Sprintf("%d", now.Unix())
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)
	suffix := make([]byte, 3)
	for i := range suffix {
		suffix[i] = letters[r.Intn(len(letters))]
	}
	return timestamp + string(suffix)
}
