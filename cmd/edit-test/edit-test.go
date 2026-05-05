//
// edit-test
//
// Written by Logan Savage. Licensed under the MIT license.
//
// This program is just used to test editing files using the same encryption
// system as the main program.
//
// Note that just like the main program, this is just a basic encryption setup
// which should protect from general navigation on the filesystem, but will not
// be secure enough for protecting against anyone with intermediate digital
// forensic skills. For that case, use a more secure system, such as what was
// provided with your operating system (FileVault, BitLocker, etc.), or a
// reputable third party tool such as VeraCrypt.
//

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"logansavage.dev/piconote/internal/secure"
)

const ENCRYPTED_FILE_PATH = "abacus.md.enc"
const DECRYPTED_FILE_PATH = "abacus.md"
const ORIGINAL_MESSAGE = "test message"
const ENCRYPTION_KEY = "abacus"
const EDITOR = "/usr/bin/vi"

func createFile(outpath, orig, key string) error {
	encrypted, err := secure.EncryptString(key, orig)
	if err != nil {
		return err
	}

	if err := os.WriteFile(outpath, encrypted, 0755); err != nil {
		return err
	}

	return nil
}

func decryptFile(fromPath, toPath, key string) error {
	contents, err := os.ReadFile(fromPath)
	if err != nil {
		return fmt.Errorf("[FAIL] unable to open .enc file: %v", err)
	}

	decrypted, err := secure.DecryptCiphertext(key, contents)
	if err != nil {
		return fmt.Errorf("[FAIL] unable to decrypt the .enc file: %v", err)
	}

	if err := os.WriteFile(toPath, []byte(decrypted), 0755); err != nil {
		return fmt.Errorf("[FAIL] unable to write decrypted temp file: %v", err)
	}
	return nil
}

func syncEncryptedFile(decryptedPath, encryptedPath, key string) error {
	plaintext, err := os.ReadFile(decryptedPath)
	if err != nil {
		return err
	}

	encrypted, err := secure.EncryptString(key, string(plaintext))
	if err != nil {
		return err
	}

	if err := os.WriteFile(encryptedPath, encrypted, 0755); err != nil {
		return err
	}

	return nil
}

func openFile(filepath, editorpath string) error {
	cmd := exec.Command(editorpath, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func main() {
	skipCreation := flag.Bool("skip-creation", false, "skip creating/overwriting the test file")
	flag.Parse()

	if !*skipCreation {
		if err := createFile(ENCRYPTED_FILE_PATH, ORIGINAL_MESSAGE, ENCRYPTION_KEY); err != nil {
			log.Fatalf("[FAIL] unable to write encrypted test file: %v", err)
		}
		log.Printf("[SUCCESS] Wrote a test file to '%s'.\nPassword: '%s'\n", ENCRYPTED_FILE_PATH, ENCRYPTION_KEY)
	}

	if err := decryptFile(ENCRYPTED_FILE_PATH, DECRYPTED_FILE_PATH, ENCRYPTION_KEY); err != nil {
		log.Fatal(err)
	}

	// Ensure decrypted temp file does not stay present if program has fatal error
	defer func() {
		_ = os.Remove(DECRYPTED_FILE_PATH)
	}()

	if err := openFile(DECRYPTED_FILE_PATH, EDITOR); err != nil {
		log.Fatal(err)
	}

	if err := syncEncryptedFile(DECRYPTED_FILE_PATH, ENCRYPTED_FILE_PATH, ENCRYPTION_KEY); err != nil {
		log.Fatal(err)
	}
}
