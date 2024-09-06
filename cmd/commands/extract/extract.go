package extract

import (
	"io"
	"math"
	"os"

	"github.com/ivanvanderbyl/extract/pkg/content"
	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
)

type MarshalBytes []byte

func (m MarshalBytes) MarshalJSON() ([]byte, error) {
	return m, nil
}

func ExtractCommand(c *cli.Context) error {
	filePath := c.Args().First()
	doc, err := content.GetReader(filePath)
	if err != nil {
		return err
	}
	all, err := io.ReadAll(doc)
	if err != nil {
		return err
	}

	schemaPath := c.String("schema")
	extractionSchema, err := os.ReadFile(schemaPath)
	if err != nil {
		return errors.Wrap(err, "failed to read schema file")
	}

	ctx := c.Context

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
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
				Content: string(all),
			},
		},

		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
			JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
				Name:   "result_schema",
				Schema: MarshalBytes(extractionSchema),
				Strict: true,
			},
		},
	})
	if err != nil {
		return err
	}

	os.Stdout.Write([]byte(resp.Choices[0].Message.Content))
	return nil
}
