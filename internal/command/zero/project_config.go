package zero

import (
	"errors"
	"fmt"
	"os"
	"path"
)

type ProjectConfig struct {
	projectFileName string
	currentDir      string
	defaultFlags    string
}

func newProjectConfig() (*ProjectConfig, error) {
	currentDir, err := getCurrentDir()
	if err != nil {
		return nil, fmt.Errorf("project config: %w", err)
	}

	config := &ProjectConfig{
		currentDir:      currentDir,
		projectFileName: "project.json",
		defaultFlags:    "-Wextra -Werror -Wall -Wpedantic",
	}

	return config, nil
}

func (c ProjectConfig) isProjectDir() (bool, error) {
	files, err := os.ReadDir(c.currentDir)
	if err != nil {
		return false, fmt.Errorf("failed to list files in current directory: %w", err)
	}

	for _, file := range files {
		if file.Name() == c.projectFileName {
			return true, nil
		}
	}

	return false, nil
}

func (c ProjectConfig) ReadProjectFile() (*ProjectFile, error) {
	isProjectDir, err := c.isProjectDir()
	if err != nil {
		return nil, err
	}

	if !isProjectDir {
		return nil, errors.New("current directory is not a C++ project.")
	}

	projectFilePath := path.Join(c.currentDir, c.projectFileName)
	if _, err := os.Stat(projectFilePath); err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("current directory does not contain a %s file", c.projectFileName)
		}
		return nil, fmt.Errorf("failed to read %s file: %w", c.projectFileName, err)
	}

	// TODO: complete implementation.
}
