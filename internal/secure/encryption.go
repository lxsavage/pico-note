package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

func hashKey(key string) ([]byte, error) {
	res := sha256.Sum256([]byte(key))
	return res[:], nil
}

func encrypt(plaintext []byte, key string) ([]byte, error) {
	keyHash, err := hashKey(key)
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

	encrypted := gcm.Seal(nil, nonce, plaintext, nil)
	resultWithNonce := append(nonce, encrypted...)
	return resultWithNonce, nil
}

func decrypt(ciphertext []byte, key string) ([]byte, error) {
	keyHash, err := hashKey(key)
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

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	res, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}
