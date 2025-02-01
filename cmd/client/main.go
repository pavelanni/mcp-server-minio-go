package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

func main() {
	// Start the server
	cmdPath := path.Join(os.Getenv("HOME"), "Learning", "mcp-server-minio-go", "cmd", "server", "main.go")
	cmd := exec.Command("go", "run", cmdPath)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalf("Failed to get stdin pipe: %v", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Failed to get stdout pipe: %v", err)
	}

	log.Println("Running command: ", cmd.Path, cmd.Args)
	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer cmd.Process.Kill()

	// Create a new MCP client
	clientTransport := stdio.NewStdioServerTransportWithIO(stdout, stdin)
	client := mcp.NewClient(clientTransport)

	if _, err := client.Initialize(context.Background()); err != nil {
		log.Fatalf("Failed to initialize client: %v", err)
	}
	log.Println("Initialized client")

	// List available tools
	log.Println("Listing tools")
	tools, err := client.ListTools(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to list tools: %v", err)
	}

	log.Println("Available Tools:")
	for _, tool := range tools.Tools {
		desc := ""
		if tool.Description != nil {
			desc = *tool.Description
		}
		log.Printf("Tool: %s. Description: %s", tool.Name, desc)
	}
	// Get MinIO credentials from environment variables
	// or use default values for testing
	env := map[string]string{
		"MINIO_ENDPOINT":   getEnvOrDefault("MINIO_ENDPOINT", "localhost:9000"),
		"MINIO_ACCESS_KEY": getEnvOrDefault("MINIO_ACCESS_KEY", "minioadmin"),
		"MINIO_SECRET_KEY": getEnvOrDefault("MINIO_SECRET_KEY", "minioadmin"),
		"MINIO_USE_SSL":    getEnvOrDefault("MINIO_USE_SSL", "false"),
	}

	// Create the request payload
	args := map[string]interface{}{
		"environment": env,
	}

	// Make the request
	resp, err := client.CallTool(context.Background(), "list-buckets", args)
	if err != nil {
		log.Fatalf("Failed to execute tool: %v", err)
	}

	// Print the response
	for _, content := range resp.Content {
		fmt.Println(content.TextContent)
	}

	// List bucket contents
	log.Println("Listing bucket contents")
	args = map[string]interface{}{
		"environment": env,
		"bucketName":  "ai-data",
	}
	resp, err = client.CallTool(context.Background(), "list-bucket-contents", args)
	if err != nil {
		log.Fatalf("Failed to execute tool: %v", err)
	}

	// Print the response
	for _, content := range resp.Content {
		fmt.Println(content.TextContent)
	}

	// Prompt an object
	log.Println("Prompting an object")
	args = map[string]interface{}{
		"environment": env,
		"bucketName":  "ai-data",
		"objectName":  "aircraft.jpg",
		"prompt":      "What kind of aircraft is this?",
	}
	resp, err = client.CallTool(context.Background(), "prompt-object", args)
	if err != nil {
		log.Fatalf("Failed to execute tool: %v", err)
	}

	// Print the response
	for _, content := range resp.Content {
		fmt.Println(content.TextContent)
	}

	// Upload a file
	log.Println("Uploading a file")
	args = map[string]interface{}{
		"environment": env,
		"bucketName":  "ai-data",
		"filePath":    "go.sum",
		"objectName":  "go.sum",
	}
	resp, err = client.CallTool(context.Background(), "upload-file", args)
	if err != nil {
		log.Fatalf("Failed to execute tool: %v", err)
	}

	// Print the response
	for _, content := range resp.Content {
		fmt.Println(content.TextContent)
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
