package piconote

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func view(file string) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	fmt.Printf("%s", bytes)
	if bytes[len(bytes)-1] != '\n' {
		fmt.Println()
	}

	return nil
}

func list(privateMode bool, dir string) error {
	contents, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	if len(contents) == 0 {
		fmt.Println("No existing notes. Create one with the \"write\" command; see -h for more information.")
		return nil
	}

	fmt.Print("Existing notes")
	if privateMode {
		fmt.Print(" (private)")
	}
	fmt.Println()

	for _, entry := range contents {
		if !entry.IsDir() {
			noteName, _ := strings.CutSuffix(entry.Name(), ".md")
			fmt.Println("- " + noteName)
		}
	}
	return nil
}

func write(file string) error {
	editorPath := "/bin/vi"
	if visual, ok := os.LookupEnv("VISUAL"); ok {
		editorPath = visual
	} else if editor, ok := os.LookupEnv("EDITOR"); ok {
		editorPath = editor
	}
	cmd := exec.Command(editorPath, file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
