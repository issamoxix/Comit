package ai

import (
	config "commit_helper/services/utils"
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func Run(Content string) string {
	client := openai.NewClient(config.Envs.OpenAiKey)
	var gptContext string

	gptContext = "You are a Senior Software Engineer. Give me multiple  commit messages for these changes following the best practices (The Conventional Commits)"
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: gptContext,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: Content,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "Please set the OPENAI_KEY environment variable"
	}

	return resp.Choices[0].Message.Content
}
