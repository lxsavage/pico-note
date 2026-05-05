package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	"golang.org/x/term"
)

func passHash(password string) ([]byte, error) {
	res := sha256.Sum256([]byte(password))
	return res[:], nil
}

// Using the specified prompt, read a string from the specified input in secure mode and return the result
func ReadSecureString(f *os.File, prompt string) (string, error) {
	fmt.Print(prompt)

	res, err := term.ReadPassword(int(f.Fd()))
	if err != nil {
		return "", err
	}

	fmt.Println()
	return string(res), nil
}

// Encrypts the specified string into a []byte ciphertext using the provided key with AES-256 encryption
func EncryptString(key, s string) ([]byte, error) {
	keyHash, err := passHash(key)
	if err != nil {
		return nil, fmt.Errorf("key hashing error: %v", err)
	}

	block, err := aes.NewCipher(keyHash)
	if err != nil {
		return nil, fmt.Errorf("cipher creation error: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("GCM cipher creation error: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("GCM nonce creation error: %v", err)
	}

	encrypted := gcm.Seal(nil, nonce, []byte(s), nil)
	resultWithNonce := append(nonce, encrypted...)
	return resultWithNonce, nil
}

// Decrypts an AES-256 ciphertext with the provided key and returns the result
func DecryptCiphertext(key string, ciphertext []byte) (string, error) {
	keyHash, err := passHash(key)
	if err != nil {
		return "", fmt.Errorf("key hashing error: %v", err)
	}

	block, err := aes.NewCipher(keyHash)
	if err != nil {
		return "", fmt.Errorf("cipher creation error: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("GCM cipher creation error: %v", err)
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	res, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
