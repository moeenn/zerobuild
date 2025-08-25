package zero

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type ObjectManager struct {
	currentDir string
	cacheDir   string
}

func newObjectManager() (*ObjectManager, error) {
	currentDir, err := getCurrentDir()
	if err != nil {
		return nil, fmt.Errorf("object manager: %w", err)
	}

	mngr := &ObjectManager{
		currentDir: currentDir,
		cacheDir:   ".cache",
	}

	return mngr, nil
}

func (m ObjectManager) validateOrCreateCacheDir() error {
	fullPath := path.Join(m.currentDir, m.cacheDir, "objects")
	if pathInfo, err := os.Stat(fullPath); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(fullPath, 0777); err != nil {
				return fmt.Errorf("failed to create object cache directory: %w", err)
			}
		}

		if !pathInfo.IsDir() {
			return fmt.Errorf("%s exists but is not a directory", fullPath)
		}
	}

	return nil
}

func (m ObjectManager) GetObjectFilePath(sourceFilePath string) string {
	objectDirPath := path.Join(m.currentDir, m.cacheDir, "objects")
	sourceFilePath = strings.ReplaceAll(sourceFilePath, m.currentDir, "")
	sourceFilePath = strings.TrimPrefix(sourceFilePath, "/")
	filename := strings.ReplaceAll(sourceFilePath, "/", "_")
	filepath := path.Join(objectDirPath, filename)
	return filepath + ".o"
}
