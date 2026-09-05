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

var History = make([]mistral.ChatMessage, 0)

func New(apiKey string, modelString string) *Client {
	return &Client{
		Model:         modelString,
		MistralClient: mistral.NewMistralClientDefault(string(apiKey)),
	}
}

func (c *Client) message(prompt string) (string, error) {
	History = append(History, mistral.ChatMessage{
		Role:    "user",
		Content: prompt,
	})
	res, err := c.MistralClient.Chat(
		c.Model,
		History,
		nil,
	)
	if err != nil {
		return "", err
	}
	History = append(History, mistral.ChatMessage{
		Role:    res.Choices[0].Message.Role,
		Content: res.Choices[0].Message.Content,
	})

	return res.Choices[0].Message.Content, nil
}
