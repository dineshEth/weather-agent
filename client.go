package main

import (
	mistral "github.com/gage-technologies/mistral-go"
)

type TypeAPIKEY string
type TypeModel string

type Client struct {
	Model         string
	MistralClient *mistral.MistralClient
}

func New(apiKey string, modelString string) *Client {
	return &Client{
		Model:         modelString,
		MistralClient: mistral.NewMistralClientDefault(string(apiKey)),
	}
}

func (c *Client) message(prompt string) (*mistral.ChatCompletionResponse, error) {
	return c.MistralClient.Chat(
		c.Model,
		[]mistral.ChatMessage{
			{
				Content: prompt,
				Role:    "user",
			},
		},
		nil,
	)
}
