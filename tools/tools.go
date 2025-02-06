package tools

type MinIOTool struct {
	Name        string
	Description string
	Handler     any
}

var ReadOnlyTools = []MinIOTool{
	{
		Name:        "list-buckets",
		Description: "List all buckets",
		Handler:     ListBucketsHandler,
	},
	{
		Name:        "list-bucket-contents",
		Description: "List the contents of a bucket",
		Handler:     ListBucketContentsHandler,
	},
	{
		Name:        "prompt-object",
		Description: "Prompt an object",
		Handler:     PromptObjectHandler,
	},
	{
		Name:        "download-object",
		Description: "Download an object from a bucket",
		Handler:     DownloadObjectHandler,
	},
	{
		Name:        "get-bucket-versioning",
		Description: "Get bucket versioning",
		Handler:     GetBucketVersioningHandler,
	},
	{
		Name:        "get-object-tags",
		Description: "Get the tags of an object",
		Handler:     GetObjectTagsHandler,
	},
	{
		Name:        "get-object-metadata",
		Description: "Get the metadata of an object",
		Handler:     GetObjectMetadataHandler,
	},
}

var WriteTools = []MinIOTool{
	{
		Name:        "create-bucket",
		Description: "Create a new bucket",
		Handler:     CreateBucketHandler,
	},
	{
		Name:        "upload-object",
		Description: "Upload an object to a bucket",
		Handler:     UploadObjectHandler,
	},
	{
		Name:        "copy-object",
		Description: "Copy an object to a new bucket",
		Handler:     CopyObjectHandler,
	},
	{
		Name:        "set-object-tags",
		Description: "Set the tags of an object",
		Handler:     SetObjectTagsHandler,
	},
	{
		Name:        "set-bucket-versioning",
		Description: "Set bucket versioning",
		Handler:     SetBucketVersioningHandler,
	},
}

var DeleteTools = []MinIOTool{
	{
		Name:        "delete-bucket",
		Description: "Delete a bucket",
		Handler:     DeleteBucketHandler,
	},
	{
		Name:        "delete-object",
		Description: "Delete an object from a bucket",
		Handler:     DeleteObjectHandler,
	},
	{
		Name:        "move-object",
		Description: "Move an object to a new bucket",
		Handler:     MoveObjectHandler,
	},
}

var AdminTools = []MinIOTool{
	{
		Name:        "get-admin-info",
		Description: "Get detailed technical info about the MinIO cluster",
		Handler:     GetAdminInfoHandler,
	},
	{
		Name:        "set-bucket-versioning",
		Description: "Set bucket versioning",
		Handler:     SetBucketVersioningHandler,
	},
}
