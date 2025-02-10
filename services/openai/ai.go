package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type commitResponse struct {
	Message string `json:"message"`
}

type RequestData struct {
	Code string `json:"code"`
}

func GetCommitMessage(content string) string {
	url := "http://comit.issamcloud.com"

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
	return data.Message
}
