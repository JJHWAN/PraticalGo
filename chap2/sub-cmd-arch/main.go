package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/JJHWAN/chap2/sub-cmd-arch/cmd"
)

var errInvalidSubCommand = errors.New("Invalid sub-command specified")

// func for cmd_a

// func for cmd_b

func printUsage(w io.Writer) {
	fmt.Fprint(w, "Usage: mync [http|grpc] -h\n")
	cmd.HandleHttp(w, []string{"-h"})
	cmd.HandleGrpc(w, []string{"-h"})
}

func handleCommand(w io.Writer, args []string) error {
	var err error
	if len(args) < 1 {
		err = errInvalidSubCommand
	} else {
		switch args[0] {
		case "http":
			// call handler for http command
			err = cmd.HandleHttp(w, args[1:])
		case "grpc":
			err = cmd.HandleGrpc(w, args[1:])
		case "-h", "-help":
			printUsage(w)
		default:
			err = errInvalidSubCommand
		}
	}

	if errors.Is(err, errInvalidSubCommand) || errors.Is(err, cmd.ErrorNoServerSpecified) {
		fmt.Fprintln(w, err)
		printUsage(w)
	}
	return err
}

func main() {
	if err := handleCommand(os.Stdout, os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
