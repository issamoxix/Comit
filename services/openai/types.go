package ai

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
