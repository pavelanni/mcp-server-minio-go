package tools

type MinIOTool struct {
	Name        string
	Description string
	Handler     any
}

var MinIOTools = []MinIOTool{
	{
		Name:        "list-buckets",
		Description: "List all buckets",
		Handler:     ListBucketsHandler,
	},
	{
		Name:        "create-bucket",
		Description: "Create a new bucket",
		Handler:     CreateBucketHandler,
	},
	{
		Name:        "prompt-object",
		Description: "Prompt an object",
		Handler:     PromptObjectHandler,
	},
	{
		Name:        "list-bucket-contents",
		Description: "List the contents of a bucket",
		Handler:     ListBucketContentsHandler,
	},
	{
		Name:        "upload-object",
		Description: "Upload an object to a bucket",
		Handler:     UploadObjectHandler,
	},
	{
		Name:        "download-object",
		Description: "Download an object from a bucket",
		Handler:     DownloadObjectHandler,
	},
	{
		Name:        "delete-object",
		Description: "Delete an object from a bucket",
		Handler:     DeleteObjectHandler,
	},
	{
		Name:        "delete-bucket",
		Description: "Delete a bucket",
		Handler:     DeleteBucketHandler,
	},
	{
		Name:        "set-bucket-versioning",
		Description: "Set bucket versioning",
		Handler:     SetBucketVersioningHandler,
	},
	{
		Name:        "get-bucket-versioning",
		Description: "Get bucket versioning",
		Handler:     GetBucketVersioningHandler,
	},
	{
		Name:        "copy-object",
		Description: "Copy an object to a new bucket",
		Handler:     CopyObjectHandler,
	},
	{
		Name:        "move-object",
		Description: "Move an object to a new bucket",
		Handler:     MoveObjectHandler,
	},
	{
		Name:        "get-object-tags",
		Description: "Get the tags of an object",
		Handler:     GetObjectTagsHandler,
	},
	{
		Name:        "set-object-tags",
		Description: "Set the tags of an object",
		Handler:     SetObjectTagsHandler,
	},
	{
		Name:        "get-object-metadata",
		Description: "Get the metadata of an object",
		Handler:     GetObjectMetadataHandler,
	},
	{
		Name:        "get-admin-info",
		Description: "Get detailed technical info about the MinIO cluster",
		Handler:     GetAdminInfoHandler,
	},
}
