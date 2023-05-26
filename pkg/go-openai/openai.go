package openai

import (
	"context"

	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
)

type OpenAIClient struct {
	apiKey string
	model  string
	logger *logrus.Entry
	Client *openai.Client
}

func NewOpenAIClient(logger *logrus.Entry, apiKey string, modelID string) *OpenAIClient {
	return &OpenAIClient{
		logger: logger,
		apiKey: apiKey,
		Client: openai.NewClient(apiKey),
		model:  modelID,
	}
}

func (g *OpenAIClient) Completion(ctx context.Context, prompt string) (string, error) {
	completionRequest := openai.CompletionRequest{
		Model:       g.model,
		Prompt:      prompt,
		MaxTokens:   10,
		Temperature: 0,
		TopP:        0.1,
		N:           1,
		LogProbs:    16,
	}

	completion, err := g.Client.CreateCompletion(ctx, completionRequest)
	if err != nil {
		g.logger.Errorf("error creating completion: %v", err)
	}

	return completion.Choices[0].Text, nil
}

func (g *OpenAIClient) CreateFile(ctx context.Context, fileName, filePath string) (string, error) {
	fileRequest := openai.FileRequest{
		FileName: fileName,
		FilePath: fileName,
		Purpose:  "fine-tune",
	}

	file, err := g.Client.CreateFile(ctx, fileRequest)
	if err != nil {
		g.logger.Errorf("error creating file: %v", err)
	}
	g.logger.Infof("training file created: %+v", file)

	return file.ID, nil
}

func (g *OpenAIClient) CreateFineTune(ctx context.Context, fileID, prompt string) (string, error) {
	fineTuneRequest := openai.FineTuneRequest{
		TrainingFile: fileID,
		Model:        "davinci",
	}

	fineTune, err := g.Client.CreateFineTune(ctx, fineTuneRequest)
	if err != nil {
		g.logger.Errorf("error creating fine tune: %v", err)
	}
	g.logger.Infof("fine tune created: %+v", fineTune)

	return fineTune.ID, nil
}
