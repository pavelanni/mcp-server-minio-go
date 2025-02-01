package tools

import (
	"context"
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/minio/minio-go/v7"
)

type CopyObjectArgs struct {
	SrcBucketName string
	SrcObjectName string
	DstBucketName string
	DstObjectName string
}

func CopyObjectHandler(ctx context.Context, args CopyObjectArgs) (*mcp.ToolResponse, error) {
	client, err := NewMinioClient()
	if err != nil {
		return nil, err
	}

	srcOptions := minio.CopySrcOptions{
		Bucket: args.SrcBucketName,
		Object: args.SrcObjectName,
	}

	dstOptions := minio.CopyDestOptions{
		Bucket: args.DstBucketName,
		Object: args.DstObjectName,
	}

	uploadInfo, err := client.CopyObject(ctx, dstOptions, srcOptions)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("File '%s' copied successfully to bucket '%s'. ETag: %s, VersionID: %s", args.SrcObjectName, args.DstBucketName, uploadInfo.ETag, uploadInfo.VersionID))), nil
}
