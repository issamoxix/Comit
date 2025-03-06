package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type VersionResponseType struct {
	Version string `json:"version"`
}

func TestVersionApi(t *testing.T) {

	resp, err := http.Get(VersionURL)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var versionResponse VersionResponseType
	err = json.Unmarshal(body, &versionResponse)
	if err != nil {
		t.Fatal(err)
	}

	if versionResponse.Version == "" {
		t.Fatal("Version is empty")
	}
	if Version != versionResponse.Version {
		t.Fatalf("Expected version %s, got %s", Version, versionResponse.Version)
	}
}
