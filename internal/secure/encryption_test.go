package secure

import (
	"testing"
)

func TestEncryption(t *testing.T) {
	orig := []byte("test message")
	key := "abacus"

	encrypted, err := encrypt(orig, key)
	if err != nil {
		t.Fatal(err)
	}

	if string(encrypted) == string(orig) {
		t.Fatalf("encrypted string is identical to unencrypted")
	}

	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		t.Fatal(err)
	}

	if string(decrypted) != string(orig) {
		t.Fatalf("DecryptContents(EncryptContents(\"%s\")) = \"%s\"; expected \"%s\"", orig, decrypted, orig)
	}
}

func TestEncryptionFailedDecryption(t *testing.T) {
	orig := []byte("test message")
	key := "abacus"
	wrongKey := "wrong"

	encrypted, err := encrypt(orig, key)
	if err != nil {
		t.Fatal(err)
	}

	if string(encrypted) == string(orig) {
		t.Fatalf("encrypted string is identical to unencrypted")
	}

	if _, err := decrypt(encrypted, wrongKey); err == nil {
		t.Fatalf("DecryptContents(EncryptContents(\"%s\")) [wrong decrypt key] has no errors; expected error.\n", orig)
	}
}
