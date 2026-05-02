package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"logansavage.dev/piconote/internal/piconote"
)

const Version = "localbuild"

type args struct {
	command  string
	noteName string
}

func getArgs(argv []string) (args, error) {
	if len(argv) == 0 {
		return args{}, errors.New("specify a subcommand.")
	}

	res := args{command: argv[0]}

	if len(argv) >= 2 {
		res.noteName = argv[1]
	} else if _, ok := piconote.BypassFileCommands[res.command]; !ok {
		return args{}, errors.New("command \"" + res.command + "\" requires a file to be specified.")
	}

	return res, nil
}

func main() {
	privateMode := flag.Bool("private", false, "read from the private note store")
	version := flag.Bool("version", false, "show version information for this version of PicoNote then exit successfully")

	flag.Usage = func() {
		nameRequiredCommands := []string{}
		for _, command := range piconote.Commands {
			if _, ok := piconote.BypassFileCommands[command]; !ok {
				nameRequiredCommands = append(nameRequiredCommands, command)
			}
		}

		fmt.Fprintf(os.Stderr, "Usage of %s:\n\n%s [-private] <command> [<name>]\n\n", os.Args[0], os.Args[0])
		fmt.Fprintf(os.Stderr, "  command\n        %s\n", strings.Join(piconote.Commands, " | "))
		fmt.Fprintf(os.Stderr, "  name\n"+
			"        name of the note to reference (if applicable).\n"+
			"        required in the commands: %s\n\n", strings.Join(nameRequiredCommands, ", "))

		flag.PrintDefaults()
	}
	flag.Parse()

	if *version {
		fmt.Fprintf(os.Stderr, "PicoNote %s\n", Version)
		os.Exit(0)
	}

	arguments, err := getArgs(flag.Args())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	if err := piconote.Exec(*privateMode, arguments.command, arguments.noteName); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)

		// If the editor exited with a non-zero exit code, pass that up the chain
		if exiterr, ok := errors.AsType[*exec.ExitError](err); ok {
			os.Exit(exiterr.ExitCode())
		}

		os.Exit(1)
	}
}
