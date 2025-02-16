package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

type VersionResponse struct {
	Version string `json:"version"`
}

func GetVersion() string {
	url := "https://comit.issamcloud.com"
	resp, err := http.Get(url)
	if err != nil {
		return "0.2"
	}

	if resp.StatusCode != 200 {
		return "0.2"
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "0.2"
	}

	var versionResponse VersionResponse
	err = json.Unmarshal(body, &versionResponse)

	if err != nil {
		return "0.2"
	}
	return versionResponse.Version
}
