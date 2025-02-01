package tools

import (
	"context"
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/minio/minio-go/v7"
)

type GetObjectTagsArgs struct {
	BucketName string
	ObjectName string
}

func GetObjectTagsHandler(ctx context.Context, args GetObjectTagsArgs) (*mcp.ToolResponse, error) {
	client, err := NewMinioClient()
	if err != nil {
		return nil, err
	}

	tags, err := client.GetObjectTagging(ctx, args.BucketName, args.ObjectName, minio.GetObjectTaggingOptions{})
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Received tags for object '%s' in bucket '%s': %v", args.ObjectName, args.BucketName, tags))), nil
}
