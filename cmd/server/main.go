package main

import (
	"log"
	"os"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
	"github.com/pavelanni/mcp-server-minio-go/tools"
	"github.com/spf13/pflag"
)

func main() {
	// Define command line flags
	var allowedDirs []string
	pflag.StringSliceVar(&allowedDirs, "allowed-directories", []string{}, "List of allowed directories for MinIO operations")
	pflag.Parse()

	server := mcp.NewServer(stdio.NewStdioServerTransport(),
		mcp.WithName("minio-mcp-server"),
		mcp.WithVersion("1.0.0"))

	if server == nil {
		log.Fatalf("Server is nil")
	}

	// Print server environment
	for _, env := range os.Environ() {
		log.Println(env)
	}

	// Make allowed directories available to tools
	tools.SetAllowedDirectories(allowedDirs)

	// Register tools
	for _, tool := range tools.MinIOTools {
		if err := server.RegisterTool(tool.Name, tool.Description, tool.Handler); err != nil {
			log.Fatalf("Failed to register %s tool: %v", tool.Name, err)
		}
	}

	// Start the server
	if err := server.Serve(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
	// Keep the server running
	select {}

}
