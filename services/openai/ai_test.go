package ai

import (
	"commit_helper/services/utils"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock implementation of CommitMessageSelector
type mockSelector struct {
	MockSelectCommitMessage func([]string) error
}

func (m *mockSelector) SelectCommitMessage(messages []string) error {
	if m.MockSelectCommitMessage != nil {
		return m.MockSelectCommitMessage(messages)
	}
	return nil
}

func TestGetCommitMessage(t *testing.T) {
	tests := []struct {
		name           string
		content        string
		mockResponse   string
		mockStatusCode int
		mockError      error
		mockSelectErr  error
		want           string
	}{
		{
			name:           "Success",
			content:        "test content",
			mockResponse:   `{"message": ["test message"]}`,
			mockStatusCode: 200,
			want:           "Ok",
		},
		{
			name:      "JSON Marshal Error",
			content:   string([]byte{0x80, 0x81}),
			mockError: errors.New("json marshal error"),
			want:      "Error: 500 Internal Server Error",
		},
		{
			name:      "HTTP Request Error",
			content:   "test content",
			mockError: errors.New("http post error"),
			want:      "Error: 500 Internal Server Error",
		},
		{
			name:           "Non-200 Status Code",
			content:        "test content",
			mockStatusCode: 404,
			want:           "Error: 404 Not Found",
		},
		{
			name:           "Read Response Body Error",
			content:        "test content",
			mockStatusCode: 200,
			mockResponse:   `{"message": ["test message"]}`,
			mockError:      errors.New("read response body error"),
			want:           "Error: 500 Internal Server Error",
		},
		{
			name:           "JSON Unmarshal Error",
			content:        "test content",
			mockStatusCode: 200,
			mockResponse:   `invalid json`,
			want:           "Error: invalid character 'i' looking for beginning of value",
		},
		{
			name:           "Select Commit Message Error",
			content:        "test content",
			mockStatusCode: 200,
			mockResponse:   `{"message": ["test message"]}`,
			mockSelectErr:  errors.New("select commit message error"),
			want:           "Error: select commit message error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock HTTP server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tt.mockError != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.WriteHeader(tt.mockStatusCode)
				w.Write([]byte(tt.mockResponse))
			}))
			defer server.Close()

			// Mock tools dependency
			mockTool := &mockSelector{
				MockSelectCommitMessage: func(messages []string) error {
					return tt.mockSelectErr
				},
			}
			utils.ComitURL = server.URL
			// Call GetCommitMessage with the test server URL and mock tool
			got := GetCommitMessage(tt.content, mockTool)

			// Validate result
			fmt.Println("got: ", got)
			assert.Equal(t, tt.want, got)
		})
	}
}
