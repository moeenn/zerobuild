package command

import (
	"fmt"
)

type CliCommand string

const (
	CliCommandInvalid CliCommand = "<invalid>"
	CliCommandHelp    CliCommand = "help"
	CliCommandInit    CliCommand = "init"
	CliCommandCompile CliCommand = "compile"
)

func ParseCliCommand(raw string) (CliCommand, error) {
	switch raw {
	case string(CliCommandHelp):
		return CliCommandHelp, nil

	case string(CliCommandInit):
		return CliCommandInit, nil

	case string(CliCommandCompile):
		return CliCommandCompile, nil
	}

	return CliCommandInvalid, fmt.Errorf("invalid command: %s.", raw)
}
