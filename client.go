package main

import (
	"log"

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

func (c *Client) chatMessage(prompt string) (string, error) {
	History = append(History, mistral.ChatMessage{
		Role:    "user",
		Content: prompt,
	})

	res, err := c.MistralClient.Chat(
		c.Model,
		History,
		&mistral.ChatRequestParams{
			Tools:          tools,
			TopP:           0.2,
			Temperature:    0.2,
			MaxTokens:      5000,
			SafePrompt:     false,
			ToolChoice:     "auto",
			ResponseFormat: mistral.ResponseFormatText,
		},
	)
	if err != nil {
		log.Fatal("Error, occur:", err)
		return "", err
	}
	History = append(History, mistral.ChatMessage{
		Role:    res.Choices[0].Message.Role,
		Content: res.Choices[0].Message.Content,
	})

	return res.Choices[0].Message.Content, nil
}

// =========================
// Mistral tools
// =========================

var tools = []mistral.Tool{
	{
		Type: "function",

		Function: mistral.Function{
			Name: "read_file",

			Description: "Read the contents of a file.",

			Parameters: map[string]interface{}{
				"type": "object",

				"properties": map[string]interface{}{
					"path": map[string]interface{}{
						"type":        "string",
						"description": "The path of the file to read.",
					},
				},

				"required": []string{
					"path",
				},
			},
		},
	},
}
