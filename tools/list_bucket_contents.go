package tools

import (
	"context"
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/minio/minio-go/v7"
)

type ListBucketContentsArgs struct {
	BucketName string
}

func ListBucketContentsHandler(ctx context.Context, args ListBucketContentsArgs) (*mcp.ToolResponse, error) {
	client, err := NewMinioClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %v", err)
	}

	objectCh := client.ListObjects(ctx, args.BucketName, minio.ListObjectsOptions{
		Recursive: true,
	})

	var contents string
	for object := range objectCh {
		contents += fmt.Sprintf("- %s (Size: %d bytes)\n", object.Key, object.Size)
	}
	return mcp.NewToolResponse(mcp.NewTextContent(contents)), nil
}
