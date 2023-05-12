package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetGoModPath() (string, error) {
	// Get the absolute path of the current test file.
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %v", err)
	}

	// Loop through the parent directories until we find the go.mod file or reach the root directory.
	for {
		goModPath := filepath.Join(wd, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			// The go.mod file exists, return its path.
			return goModPath, nil
		} else if os.IsNotExist(err) {
			// The go.mod file does not exist in this directory, go up one level.
			wd = filepath.Dir(wd)
		} else {
			// An error occurred, return it.
			return "", fmt.Errorf("failed to check for go.mod file: %v", err)
		}

		// We have reached the root directory, and the go.mod file does not exist.
		if wd == filepath.Dir(wd) {
			return "", fmt.Errorf("go.mod file not found in directory tree")
		}
	}
}
