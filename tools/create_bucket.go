package tools

import (
	"context"
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/minio/minio-go/v7"
)

type CreateBucketArgs struct {
	BucketName string
}

func CreateBucketHandler(ctx context.Context, args CreateBucketArgs) (*mcp.ToolResponse, error) {
	client, err := NewMinioClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %v", err)
	}

	err = client.MakeBucket(ctx, args.BucketName, minio.MakeBucketOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to create bucket: %v", err)
	}

	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Bucket '%s' created successfully", args.BucketName))), nil
}
