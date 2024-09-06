package extract

import (
	"io"
	"log/slog"

	"github.com/ivanvanderbyl/extract/pkg/content"
	"github.com/urfave/cli/v2"
)

func ExtractCommand(ctx *cli.Context) error {
	filePath := ctx.Args().First()
	doc, err := content.GetReader(filePath)
	if err != nil {
		return err
	}
	all, err := io.ReadAll(doc)
	if err != nil {
		return err
	}

	slog.Info("Extracting data from document", "document", all)
	return nil
}
