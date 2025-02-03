package tools

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/minio/minio-go/v7"
	"github.com/pavelanni/mcp-server-minio-go/fsutils"
)

type DownloadObjectArgs struct {
	BucketName string
	ObjectName string
	FilePath   string
}

func DownloadObjectHandler(ctx context.Context, args DownloadObjectArgs) (*mcp.ToolResponse, error) {
	client, err := NewMinioClient()
	if err != nil {
		log.Printf("Failed to create MinIO client: %v", err)
		return nil, fmt.Errorf("failed to create MinIO client: %v", err)
	}

	if err := fsutils.ValidatePath(args.FilePath, allowedDirectories); err != nil {
		log.Printf("Invalid file path: %v", err)
		log.Printf("Allowed directories: %v", allowedDirectories)
		return nil, fmt.Errorf("invalid file path: %v", err)
	}

	object, err := client.GetObject(ctx, args.BucketName, args.ObjectName, minio.GetObjectOptions{})
	if err != nil {
		log.Printf("Failed to get object: %v", err)
		return nil, fmt.Errorf("failed to get object: %v", err)
	}
	defer object.Close()

	if err := os.MkdirAll(filepath.Dir(args.FilePath), 0755); err != nil {
		log.Printf("Failed to create directories: %v", err)
		return nil, fmt.Errorf("failed to create directories: %v", err)
	}

	outFile, err := os.Create(args.FilePath)
	if err != nil {
		log.Printf("Failed to create file: %v", err)
		return nil, fmt.Errorf("failed to create file: %v", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, object)
	if err != nil {
		log.Printf("Failed to copy object to file: %v", err)
		return nil, fmt.Errorf("failed to copy object to file: %v", err)
	}

	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Object '%s' downloaded successfully from bucket '%s' and saved to '%s'", args.ObjectName, args.BucketName, args.FilePath))), nil
}
