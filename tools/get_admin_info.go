package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/minio/madmin-go/v3"
)

type GetAdminInfoArgs struct {
}

func GetAdminInfoHandler(ctx context.Context, args GetAdminInfoArgs) (*mcp.ToolResponse, error) {
	// Initialize admin client
	// Get MinIO endpoint and credentials from environment variables
	endpoint := getEnvOrDefault("MINIO_ENDPOINT", "ai-data.demo.minio.tech:31636")
	accessKey := getEnvOrDefault("MINIO_ACCESS_KEY", "pavel@minio.io")
	secretKey := getEnvOrDefault("MINIO_SECRET_KEY", "0bjectst0rage!")
	useSSL := true
	adminClient, err := madmin.New(
		endpoint,
		accessKey,
		secretKey,
		useSSL,
	)
	if err != nil {
		log.Fatalf("Error creating admin client: %v", err)
	}

	// Get server info
	serverInfo, err := adminClient.ServerInfo(ctx)
	if err != nil {
		return nil, err
	}

	jsonOutput := prettyPrint(serverInfo)
	printOutput := printClusterMetrics(serverInfo)

	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Admin info for the MinIO cluster in JSON: %s\n and in print format: %s\n", jsonOutput, printOutput))), nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func prettyPrint(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return ""
	}
	return string(b)
}

func printClusterMetrics(info madmin.InfoMessage) string {
	output := ""
	output += fmt.Sprintf("Mode: %s\n", info.Mode)
	output += fmt.Sprintf("Deployment ID: %s\n", info.DeploymentID)

	// Print bucket and object counts
	output += fmt.Sprintf("Total Buckets: %d\n", info.Buckets.Count)
	output += fmt.Sprintf("Total Objects: %d\n", info.Objects.Count)

	// Print backend information
	output += fmt.Sprintf("Backend Type: %s\n", info.Backend.Type)
	output += fmt.Sprintf("Online Disks: %d\n", info.Backend.OnlineDisks)
	output += fmt.Sprintf("Offline Disks: %d\n", info.Backend.OfflineDisks)

	// Print server status
	output += "\nServer Status:\n"
	for _, server := range info.Servers {
		output += fmt.Sprintf("  - Endpoint: %s\n", server.Endpoint)
		output += fmt.Sprintf("    State: %s\n", server.State)
		output += fmt.Sprintf("    Uptime: %s\n", time.Duration(server.Uptime)*time.Second)
		output += fmt.Sprintf("    Version: %s\n", server.Version)
		output += fmt.Sprintf("    Network: %v\n", server.Network)

		// Print drive information for each server
		output += "    Drives:\n"
		for _, drive := range server.Disks {
			output += fmt.Sprintf("      - Path: %s\n", drive.Endpoint)
			output += fmt.Sprintf("        State: %s\n", drive.State)
			output += fmt.Sprintf("        Total Space: %.2f GB\n", float64(drive.TotalSpace)/(1<<30))
			output += fmt.Sprintf("        Used Space: %.2f GB\n", float64(drive.UsedSpace)/(1<<30))
			output += fmt.Sprintf("        Available Space: %.2f GB\n", float64(drive.AvailableSpace)/(1<<30))
		}
		output += "\n"
	}
	return output
}
