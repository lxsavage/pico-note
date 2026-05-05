package piconote

import (
	"fmt"
	"log"
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

	pass := ""
	encryptedPath := ""
	if p.Private && !BypassFileCommands[p.Command] {
		secured, err := secure.ReadSecureString(os.Stdin, "Password:")
		if err != nil {
			return err
		}

		pass = secured
		encryptedPath = decryptedPath

		tempFile, err := os.CreateTemp(path.Dir(encryptedPath), ".~*.md")
		if err != nil {
			return err
		}

		decryptedPath = tempFile.Name()
		tempFile.Close()

		if _, err := os.Stat(encryptedPath); err != nil {
			f, err := os.Create(encryptedPath)
			if err != nil {
				return err
			}
			f.Close()
		} else {
			if err := secure.DecryptFile(encryptedPath, decryptedPath, pass); err != nil {
				os.Remove(decryptedPath)
				return err
			}
		}
	}

	// Ensure that the decrypted temp file gets cleaned up regardless of how the command goes
	defer func() {
		if !p.Private || BypassFileCommands[p.Command] {
			return
		}

		if err := secure.SyncEncryptedFile(decryptedPath, encryptedPath, pass); err != nil {
			log.Printf("unable to sync encrypted file: %v", err)
		}

		_ = os.Remove(decryptedPath)
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
