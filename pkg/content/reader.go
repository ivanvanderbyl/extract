package content

import (
	"fmt"
	"io"
	"os"
)

func GetReader(filePath string) (io.Reader, error) {
	if filePath == "" {
		if isStdInPiped() {
			return os.Stdin, nil
		} else {
			return nil, fmt.Errorf("File path or stdin required")
		}
	}

	return os.Open(filePath)
}

// isStdInPiped checks if stdin is being piped to the process
func isStdInPiped() bool {
	fileInfo, _ := os.Stdin.Stat()
	return (fileInfo.Mode() & os.ModeCharDevice) == 0
}
