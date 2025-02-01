package tools

import (
	"context"
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/minio/minio-go/v7"
)

type DeleteObjectArgs struct {
	BucketName string
	ObjectName string
}

func DeleteObjectHandler(ctx context.Context, args DeleteObjectArgs) (*mcp.ToolResponse, error) {
	client, err := NewMinioClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %v", err)
	}

	err = client.RemoveObject(ctx, args.BucketName, args.ObjectName, minio.RemoveObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to delete object: %v", err)
	}

	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Object '%s' deleted successfully from bucket '%s'", args.ObjectName, args.BucketName))), nil
}
