package piconote

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func Exec(privateMode bool, command, noteName string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	baseDir := path.Join(homeDir, "Notes")
	if privateMode {
		baseDir = path.Join(baseDir, ".private")
	}

	noteFile := noteName
	if !strings.HasSuffix(noteName, ".md") {
		noteFile += ".md"
	}

	filePath := path.Join(baseDir, noteFile)
	switch command {
	case "view":
		return view(filePath)
	case "list":
		return list(privateMode, baseDir)
	case "write":
		return write(filePath)
	case "remove":
		return remove(filePath)
	}

	return fmt.Errorf("invalid command: %s", command)
}
