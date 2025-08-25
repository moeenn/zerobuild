package main

import (
	"errors"
	"fmt"
	"os"
	"zerobuild/internal/command"
)

func run(args []string) error {
	if len(args) < 2 {
		return errors.New("please provide a command")
	}

	var err error
	parsedCommand, err := command.ParseCliCommand(args[1])
	if err != nil {
		return err
	}

	switch parsedCommand {
	case command.CliCommandHelp:
		printHelp()

	case command.CliCommandInit:
		remainingArgs := args[1:]
		err = initProject(remainingArgs)

	case command.CliCommandCompile:
		remainingArgs := args[1:]
		err = compileProject(remainingArgs)
	}

	return err
}

func printHelp() {
	// TODO: implement.
}

func initProject(args []string) error {
	// TODO: implement.
	return nil
}

func compileProject(args []string) error {
	// TODO: implement.
	return nil
}

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}
