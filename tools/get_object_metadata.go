package tools

import (
	"context"
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/minio/minio-go/v7"
)

type GetObjectMetadataArgs struct {
	BucketName string
	ObjectName string
}

func GetObjectMetadataHandler(ctx context.Context, args GetObjectMetadataArgs) (*mcp.ToolResponse, error) {
	client, err := NewMinioClient()
	if err != nil {
		return nil, err
	}

	objectInfo, err := client.StatObject(ctx, args.BucketName, args.ObjectName, minio.StatObjectOptions{})
	if err != nil {
		return nil, err
	}
	metadata := objectInfo.Metadata
	userMetadata := objectInfo.UserMetadata

	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("User metadata for object '%s' in bucket '%s': %v\n\nMetadata: %v", args.ObjectName, args.BucketName, userMetadata, metadata))), nil
}
