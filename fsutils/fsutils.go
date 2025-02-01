package fsutils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func NormalizePath(path string) string {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Printf("Failed to get user home directory: %v", err)
			return path
		}
		path = filepath.Join(home, path[1:])
	}
	log.Printf("Normalized path: %s", filepath.Clean(path))
	return filepath.Clean(path)
}

func ValidatePath(path string, allowedPaths []string) error {
	normalizedPath := NormalizePath(path)
	for _, allowedPath := range allowedPaths {
		if strings.HasPrefix(normalizedPath, allowedPath) {
			return nil
		}
	}
	return fmt.Errorf("path %s is not allowed", path)
}
