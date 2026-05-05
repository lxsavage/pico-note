package secure

import (
	"testing"
)

func TestEncryption(t *testing.T) {
	orig := "test message"
	key := "abacus"

	encrypted, err := encryptString(key, orig)
	if err != nil {
		t.Fatal(err)
	}
	if string(encrypted) == orig {
		t.Fatalf("encrypted string is identical to unencrypted")
	}

	decrypted, err := decryptCiphertext(key, encrypted)
	if err != nil {
		t.Fatal(err)
	}
	if decrypted != orig {
		t.Fatalf("DecryptContents(EncryptContents(\"%s\")) = \"%s\"; expected \"%s\"", orig, decrypted, orig)
	}
}

func TestEncryptionFailedDecryption(t *testing.T) {
	orig := "test message"
	key := "abacus"
	wrongKey := "wrong"

	encrypted, err := encryptString(key, orig)
	if err != nil {
		t.Fatal(err)
	}
	if string(encrypted) == orig {
		t.Fatalf("encrypted string is identical to unencrypted")
	}

	if _, err := decryptCiphertext(wrongKey, encrypted); err == nil {
		t.Fatalf("DecryptContents(EncryptContents(\"%s\")) [wrong decrypt key] has no errors; expected error.\n", orig)
	}
}
