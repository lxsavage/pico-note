package secure

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

// Using the specified prompt, read a string from the specified input in secure
// mode and return the result
func ReadSecureString(f *os.File, prompt string) (string, error) {
	fmt.Print(prompt)
	res, err := term.ReadPassword(int(f.Fd()))
	if err != nil {
		return "", err
	}

	fmt.Println()
	return string(res), nil
}

// Decrypts fromPath using AES-256 and the provided key, and puts the result in
// toPath.
func DecryptFile(fromPath, toPath, key string) error {
	contents, err := os.ReadFile(fromPath)
	if err != nil {
		return fmt.Errorf("unable to open the from-file: %v", err)
	}

	decrypted, err := decrypt(contents, key)
	if err != nil {
		return fmt.Errorf("unable to decrypt the from-file: %v", err)
	}

	if err := os.WriteFile(toPath, []byte(decrypted), 0755); err != nil {
		return fmt.Errorf("unable to write to the to-file: %v", err)
	}

	return nil
}

// Takes the content of decryptedPath, encrypts it with the provided key under
// AES-256, then puts the result in encryptedPath.
func SyncEncryptedFile(decryptedPath, encryptedPath, key string) error {
	plaintext, err := os.ReadFile(decryptedPath)
	if err != nil {
		return err
	}

	encrypted, err := encrypt(plaintext, key)
	if err != nil {
		return err
	}

	if err := os.WriteFile(encryptedPath, encrypted, 0755); err != nil {
		return err
	}

	return nil
}
