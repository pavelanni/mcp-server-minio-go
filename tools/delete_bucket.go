package tools

import (
	"context"
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
)

type DeleteBucketArgs struct {
	BucketName string
}

func DeleteBucketHandler(ctx context.Context, args DeleteBucketArgs) (*mcp.ToolResponse, error) {
	client, err := NewMinioClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %v", err)
	}

	err = client.RemoveBucket(ctx, args.BucketName)
	if err != nil {
		return nil, fmt.Errorf("failed to delete bucket: %v", err)
	}

	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Bucket '%s' deleted successfully", args.BucketName))), nil
}
