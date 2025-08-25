package zero

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type SourceManager struct {
	currentDir       string
	sourceExtensions []string
	headerExtensions []string
	logger           *log.Logger
}

func NewSourceManager(logger *log.Logger) (*SourceManager, error) {
	currentDir, err := getCurrentDir()
	if err != nil {
		return nil, fmt.Errorf("source manager: %w", err)
	}

	mngr := &SourceManager{
		currentDir:       currentDir,
		sourceExtensions: []string{"cpp"},
		headerExtensions: []string{"hpp", "h"},
		logger:           logger,
	}

	return mngr, nil
}

func (m SourceManager) isSourceFile(path string) bool {
	for _, ext := range m.sourceExtensions {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	return false
}

func (m SourceManager) DiscoverSourceFiles() ([]string, error) {
	files := []string{}
	err := filepath.Walk(m.currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			m.logger.Printf("failed to read directory. Error: %s", err.Error())
			return nil
		}

		if !info.IsDir() {
			if m.isSourceFile(path) {
				files = append(files, path)
			}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to find source files: %w", err)
	}

	return files, nil
}

// TODO: in the future this will include path of dependencies as well.
func (m SourceManager) GetIncludeDirectory() []string {
	includeDirs := []string{}

	fullPath := path.Join(m.currentDir, "include")
	if _, err := os.Stat(fullPath); err == nil {
		includeDirs = append(includeDirs, fullPath)
	}

	return includeDirs
}
