package chatgpt

import (
	"github.com/sashabaranov/go-openai"
)

type GPTClient struct {
	apiKey string
	client *openai.Client
}

func NewGPTClient(apiKey string) *GPTClient {
	return &GPTClient{
		apiKey: apiKey,
		client: openai.NewClient(apiKey),
	}
}
