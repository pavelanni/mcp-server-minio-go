package tools

import (
	"context"
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
)

type ListAllowedDirectoriesArgs struct {
}

func ListAllowedDirectoriesHandler(ctx context.Context, args ListAllowedDirectoriesArgs) (*mcp.ToolResponse, error) {
	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Allowed directories: %v", allowedDirectories))), nil
}
