package tools

import (
	"context"
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
)

type ListBucketsArgs struct {
}

func ListBucketsHandler(ctx context.Context, args ListBucketsArgs) (*mcp.ToolResponse, error) {
	client, err := NewMinioClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %v", err)
	}

	buckets, err := client.ListBuckets(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list buckets: %v", err)
	}
	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Buckets: %v", buckets))), nil
}
