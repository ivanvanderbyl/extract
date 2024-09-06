package infer

import (
	"context"
	"encoding/json"
	"math"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type Result struct {
	Schema string `json:"schema"`
}

func InferSchemaFromDocument(ctx context.Context, document string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	resultFormat := &jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"schema": {
				Type:        jsonschema.String,
				Description: "Raw JSON Schema according to your json schema formatting rules, as a string, omitting indentation and newlines",
			},
		},
		AdditionalProperties: false,
		Required:             []string{"schema"},
	}

	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:       openai.GPT4o20240806,
		Temperature: math.SmallestNonzeroFloat32, // 0.0, workaround for golang json marshalling of 0.0
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPrompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: document,
			},
		},

		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
			JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
				Name:   "new_schema",
				Schema: resultFormat,
				Strict: true,
			},
		},
	})
	if err != nil {
		return "", err
	}

	result := new(Result)
	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), result)
	if err != nil {
		return "", err
	}

	return result.Schema, nil
}
