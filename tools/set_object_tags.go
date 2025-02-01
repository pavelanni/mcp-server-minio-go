package tools

import (
	"context"
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/tags"
)

type SetObjectTagsArgs struct {
	BucketName string
	ObjectName string
	Tags       map[string]string
}

func SetObjectTagsHandler(ctx context.Context, args SetObjectTagsArgs) (*mcp.ToolResponse, error) {
	client, err := NewMinioClient()
	if err != nil {
		return nil, err
	}

	tags, err := tags.NewTags(args.Tags, true)
	if err != nil {
		return nil, err
	}
	err = client.PutObjectTagging(ctx, args.BucketName, args.ObjectName, tags, minio.PutObjectTaggingOptions{})
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Tags set for object '%s' in bucket '%s': %v", args.ObjectName, args.BucketName, args.Tags))), nil
}
