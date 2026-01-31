package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

const appName = "seman"

func DataDir() (string, error) {
	dataRoot := os.Getenv("XDG_DATA_HOME")
	if dataRoot == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("resolve home directory: %w", err)
		}
		dataRoot = filepath.Join(home, ".local", "share")
	}
	return filepath.Join(dataRoot, appName), nil
}
