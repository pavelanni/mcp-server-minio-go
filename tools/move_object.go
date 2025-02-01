package tools

import (
	"context"
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/minio/minio-go/v7"
)

type MoveObjectArgs struct {
	SrcBucketName string
	SrcObjectName string
	DstBucketName string
	DstObjectName string
}

func MoveObjectHandler(ctx context.Context, args MoveObjectArgs) (*mcp.ToolResponse, error) {
	client, err := NewMinioClient()
	if err != nil {
		return nil, err
	}

	_, err = client.CopyObject(ctx, minio.CopyDestOptions{
		Bucket: args.DstBucketName,
		Object: args.DstObjectName,
	}, minio.CopySrcOptions{
		Bucket: args.SrcBucketName,
		Object: args.SrcObjectName,
	})
	if err != nil {
		return nil, err
	}

	err = client.RemoveObject(ctx, args.SrcBucketName, args.SrcObjectName, minio.RemoveObjectOptions{})
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("File '%s' moved successfully to bucket '%s'", args.SrcObjectName, args.DstBucketName))), nil
}
