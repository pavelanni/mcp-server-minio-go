package tools

import (
	"context"
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
)

type SetBucketVersioningArgs struct {
	BucketName string
	Enabled    bool
}

func SetBucketVersioningHandler(ctx context.Context, args SetBucketVersioningArgs) (*mcp.ToolResponse, error) {
	client, err := NewMinioClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %v", err)
	}

	if args.Enabled {
		err = client.EnableVersioning(ctx, args.BucketName)
	} else {
		err = client.SuspendVersioning(ctx, args.BucketName)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to set bucket versioning: %v", err)
	}

	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Bucket versioning set to %v for bucket '%s'", args.Enabled, args.BucketName))), nil
}
