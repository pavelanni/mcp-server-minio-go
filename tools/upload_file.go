package tools

import (
	"context"
	"fmt"
	"os"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/minio/minio-go/v7"
)

type UploadFileArgs struct {
	BucketName string
	FilePath   string
	ObjectName string
}

func UploadFileHandler(ctx context.Context, args UploadFileArgs) (*mcp.ToolResponse, error) {
	client, err := NewMinioClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %v", err)
	}

	file, err := os.Open(args.FilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %v", err)
	}

	_, err = client.PutObject(ctx, args.BucketName, args.ObjectName, file, fileInfo.Size(), minio.PutObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %v", err)
	}

	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("File '%s' uploaded successfully to bucket '%s'", args.ObjectName, args.BucketName))), nil
}
