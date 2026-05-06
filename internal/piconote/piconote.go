package piconote

import (
	"fmt"
	"os"
	"path"
	"strings"

	"logansavage.dev/piconote/internal/secure"
)

type ExecParams struct {
	Private  bool
	Command  string
	NoteName string
}

func Exec(p ExecParams) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	baseDir := path.Join(homeDir, "Notes")
	if p.Private {
		baseDir = path.Join(baseDir, ".private")
	}

	noteFile := p.NoteName
	if !strings.HasSuffix(p.NoteName, ".md") {
		noteFile += ".md"
	}

	decryptedPath := path.Join(baseDir, noteFile)
	encryptedPath := ""
	key := ""
	if p.Private && !BypassFileCommands[p.Command] {
		var err error
		key, err = secure.ReadSecureString(os.Stdin, "Password:")
		if err != nil {
			return err
		}

		encryptedPath = decryptedPath
		if decryptedPath, err = decryptIntoTempFile(encryptedPath, key); err != nil {
			return err
		}
	}

	defer func() {
		if !p.Private || BypassFileCommands[p.Command] {
			return
		}
		_ = syncAndCleanupTempFile(decryptedPath, encryptedPath, key)
	}()

	var resultErr error
	switch p.Command {
	case "view":
		resultErr = view(decryptedPath)
	case "list":
		resultErr = list(p.Private, baseDir)
	case "write":
		resultErr = write(decryptedPath)
	case "remove":
		_ = os.Remove(decryptedPath)
	default:
		return fmt.Errorf("invalid command: %s", p.Command)
	}

	return resultErr
}

func decryptIntoTempFile(encryptedPath, key string) (string, error) {
	tempFile, err := os.CreateTemp(path.Dir(encryptedPath), ".~*.md")
	if err != nil {
		return "", err
	}

	decryptedPath := tempFile.Name()
	tempFile.Close()

	if _, err := os.Stat(encryptedPath); err != nil {
		f, err := os.Create(encryptedPath)
		if err != nil {
			return "", err
		}
		f.Close()
	} else if err := secure.DecryptFile(encryptedPath, decryptedPath, key); err != nil {
		os.Remove(decryptedPath)
		return "", err
	}

	return decryptedPath, nil
}

func syncAndCleanupTempFile(decryptedPath, encryptedPath, pass string) error {
	if err := secure.SyncEncryptedFile(decryptedPath, encryptedPath, pass); err != nil {
		return fmt.Errorf("unable to sync encrypted file: %v", err)
	}

	return os.Remove(decryptedPath)
}
