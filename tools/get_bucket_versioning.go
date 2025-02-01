package tools

import (
	"context"
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
)

type GetBucketVersioningArgs struct {
	BucketName string
}

func GetBucketVersioningHandler(ctx context.Context, args GetBucketVersioningArgs) (*mcp.ToolResponse, error) {
	client, err := NewMinioClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %v", err)
	}

	versioningConfig, err := client.GetBucketVersioning(ctx, args.BucketName)
	if err != nil {
		return nil, fmt.Errorf("failed to get bucket versioning: %v", err)
	}

	enabled := versioningConfig.Status == "Enabled"

	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Bucket versioning is %v for bucket '%s'", enabled, args.BucketName))), nil
}
