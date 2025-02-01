package tools

import (
	"context"
	"fmt"
	"io"
	"time"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/minio/minio-go/v7"
)

type PromptObjectArgs struct {
	BucketName string
	ObjectName string
	Prompt     string
}

func PromptObjectHandler(ctx context.Context, args PromptObjectArgs) (*mcp.ToolResponse, error) {
	client, err := NewMinioClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %v", err)
	}

	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	result, err := client.PromptObject(ctxTimeout, args.BucketName, args.ObjectName, args.Prompt, minio.PromptObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object: %v", err)
	}
	defer result.Close()

	content, err := io.ReadAll(result)
	if err != nil {
		return nil, fmt.Errorf("failed to read object: %v", err)
	}

	return mcp.NewToolResponse(mcp.NewTextContent(string(content))), nil
}
