package ai

import (
	"context"
	"encoding/json"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

const model = "gpt-4o-mini"

type Client struct {
	oai *openai.Client
}

func NewClient() *Client {
	return &Client{
		oai: openai.NewClient(os.Getenv("OPENAI_API_KEY")),
	}
}

func (c *Client) chat(ctx context.Context, system string, user string, out any) error {
	resp, err := c.oai.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: model,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: system},
			{Role: openai.ChatMessageRoleUser, Content: user},
		},
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONObject,
		},
	})
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(resp.Choices[0].Message.Content), out)
}
