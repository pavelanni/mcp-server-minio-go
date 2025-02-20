package tools

import (
	"context"
	"fmt"
	"log"
	"os"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/pavelanni/mcp-server-minio-go/fsutils"
)

type ListLocalFilesArgs struct {
	Path string
}

func ListLocalFilesHandler(ctx context.Context, args ListLocalFilesArgs) (*mcp.ToolResponse, error) {
	if err := fsutils.ValidatePath(args.Path, allowedDirectories); err != nil {
		log.Printf("Invalid file path: %v", err)
		log.Printf("Allowed directories: %v", allowedDirectories)
		return nil, fmt.Errorf("invalid file path: %v", err)
	}

	localFiles, err := os.ReadDir(args.Path)
	if err != nil {
		return nil, err
	}
	var response string
	for _, file := range localFiles {
		response += file.Name() + "\n"
	}

	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Local files in directory %s: %s", args.Path, response))), nil
}
