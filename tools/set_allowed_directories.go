package tools

import (
	"log"

	"github.com/pavelanni/mcp-server-minio-go/fsutils"
)

var allowedDirectories []string

func SetAllowedDirectories(dirs []string) {
	log.Printf("Setting allowed directories: %v", dirs)
	for _, dir := range dirs {
		log.Printf("Allowed directory: %s", dir)
		allowedDirectories = append(allowedDirectories, fsutils.NormalizePath(dir))
	}
	log.Printf("Allowed directories: %v", allowedDirectories)
}
