package zero

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type CppStandard string

const (
	CppStandardGnu23 CppStandard = "gnu++23"
)

func CppStandardFromString(raw string) (CppStandard, error) {
	switch raw {
	case string(CppStandardGnu23):
		return CppStandardGnu23, nil
	}

	return "", fmt.Errorf("unrecognized C++ standard: %s", raw)
}

type CompileCommandOptions struct {
	Compiler     string
	Std          CppStandard
	Flags        string
	Includes     []string
	InputFile    string
	OutputObject string
	Optimization uint
	IsRelease    bool
}

func (opts CompileCommandOptions) String() string {
	var sb strings.Builder
	sb.WriteString(opts.Compiler + " ")
	for _, inc := range opts.Includes {
		sb.WriteString("-I" + inc + " ")
	}

	sb.WriteString(fmt.Sprintf("-O%d ", opts.Optimization))
	if opts.IsRelease {
		sb.WriteString("-DNDEBUG ")
	}

	sb.WriteString("-std=" + string(opts.Std))
	sb.WriteString(opts.Flags + " ")
	sb.WriteString("-o " + opts.OutputObject)
	sb.WriteString("-c " + opts.InputFile)
	return sb.String()
}

type CompileCommand struct {
	Directory string                `json:"directory"`
	Command   CompileCommandOptions `json:"command"`
	File      string                `json:"file"`
	Output    string                `json:"output"`
}

type ProjectFile struct {
	Name     string      `json:"file"`
	Compiler string      `json:"compiler"`
	Std      CppStandard `json:"std"`
	Flags    string      `json:"flags"`
}

func getCurrentDir() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to determine current dir: %w", err)
	}
	return filepath.Dir(ex), nil
}
