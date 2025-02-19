package tools

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient() (*minio.Client, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		// It's often okay to continue if .env doesn't exist in production
		// as variables might be set through other means
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("error loading .env file: %v", err)
		}
	}

	endpoint := os.Getenv("MINIO_ENDPOINT")
	if endpoint == "" {
		return nil, fmt.Errorf("MINIO_ENDPOINT is required")
	}
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	if accessKey == "" {
		return nil, fmt.Errorf("MINIO_ACCESS_KEY is required")
	}
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	if secretKey == "" {
		return nil, fmt.Errorf("MINIO_SECRET_KEY is required")
	}
	useSSL := os.Getenv("MINIO_USE_SSL")
	if useSSL == "" {
		return nil, fmt.Errorf("MINIO_USE_SSL is required")
	}

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL == "true",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %v", err)
	}
	return client, nil
}
