package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"

	"github.com/inconshreveable/go-update"
)

type VersionResponse struct {
	Version string `json:"version"`
}

func GetLatestVersion() string {
	resp, err := http.Get(VersionURL)
	if err != nil {
		return Version
	}

	if resp.StatusCode != 200 {
		return Version
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Version
	}

	var versionResponse VersionResponse
	err = json.Unmarshal(body, &versionResponse)

	if err != nil {
		return Version
	}
	return versionResponse.Version
}

func SelfUpdate() (string, error) {
	fmt.Println("Checking for updates...")
	latestVersion := GetLatestVersion()

	if latestVersion == Version {
		fmt.Println("You already have the latest version.")
		return "", nil
	}

	var fileName string
	switch runtime.GOOS {
	case "windows":
		fileName = "comit.exe"
	case "darwin":
		fileName = "comit"
	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	latestBinaryURL := fmt.Sprintf(UpdateLink, latestVersion, fileName)
	resp, err := http.Get(latestBinaryURL)
	if err != nil {
		return "", fmt.Errorf("failed to download update: %v", err)
	}
	defer resp.Body.Close()

	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		return "", fmt.Errorf("failed to apply update: %v", err)
	}
	return "Ok", nil
}
