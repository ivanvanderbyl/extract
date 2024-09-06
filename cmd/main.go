package main

import (
	"bytes"
	"io"
	"log/slog"
	"os"

	"github.com/ivanvanderbyl/extract/cmd/commands/extract"
	"github.com/ivanvanderbyl/extract/pkg/content"
	"github.com/ivanvanderbyl/extract/pkg/infer"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "extract",
		Usage:       "A tool for extracting structured data from unstructured documents",
		Description: "Extract structured data from a document according to the input JSON Schema",
		Commands: []*cli.Command{
			{
				Name:        "infer",
				Description: "Infer the JSON Schema of a document",
				Action: func(ctx *cli.Context) error {
					filePath := ctx.Args().First()
					doc, err := content.GetReader(filePath)
					if err != nil {
						return err
					}
					all, err := io.ReadAll(doc)
					if err != nil {
						return err
					}

					result, err := infer.InferSchemaFromDocument(ctx.Context, string(all))
					if err != nil {
						return err
					}

					_, err = io.Copy(os.Stdout, bytes.NewBufferString(result))
					if err != nil {
						return err
					}
					return nil
				},
			},
			// {
			// 	Name:        "extract",
			// 	Description: "Extract structured data from a document according to the input JSON Schema",
			// 	Hidden:      true,
			// 	Flags: []cli.Flag{
			// 		&cli.StringFlag{
			// 			Name:     "schema",
			// 			Usage:    "The JSON Schema to use for extraction",
			// 			Required: true,
			// 		},
			// 	},
			// 	Action: func(ctx *cli.Context) error {
			// 		return nil
			// 	},
			// },
		},
		Args:      true,
		ArgsUsage: "The document to extract data from",
		Action:    extract.ExtractCommand,
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error("Error running extraction", "error", err)
		os.Exit(1)
	}
}
