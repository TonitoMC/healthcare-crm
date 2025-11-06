package storage

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	Client *minio.Client
	Bucket string
}

// NewMinioClient initializes a MinIO client from env vars. Returns error if missing config.
func NewMinioClient() (*MinioClient, error) {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	if endpoint == "" {
		return nil, fmt.Errorf("MINIO_ENDPOINT not set")
	}
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := os.Getenv("MINIO_USE_SSL") == "true"
	bucket := os.Getenv("MINIO_BUCKET")
	if bucket == "" {
		bucket = "healthcare-crm"
	}

	creds := credentials.NewStaticV4(accessKey, secretKey, "")
	mc, err := minio.New(endpoint, &minio.Options{Creds: creds, Secure: useSSL})
	if err != nil {
		return nil, err
	}

	// Ensure bucket exists
	ctx := context.Background()
	exists, err := mc.BucketExists(ctx, bucket)
	if err != nil {
		return nil, err
	}
	if !exists {
		if err := mc.MakeBucket(ctx, bucket, minio.MakeBucketOptions{}); err != nil {
			return nil, err
		}
	}

	return &MinioClient{Client: mc, Bucket: bucket}, nil
}

// UploadObject uploads an object from reader to given key. size is required when reader is not seekable.
func (m *MinioClient) UploadObject(ctx context.Context, key string, reader io.Reader, size int64, contentType string) error {
	_, err := m.Client.PutObject(ctx, m.Bucket, key, reader, size, minio.PutObjectOptions{ContentType: contentType})
	return err
}

// PresignedGetURL returns a presigned GET URL valid for duration
func (m *MinioClient) PresignedGetURL(ctx context.Context, key string, expiry time.Duration) (string, error) {
	reqParams := make(url.Values)
	presigned, err := m.Client.PresignedGetObject(ctx, m.Bucket, key, expiry, reqParams)
	if err != nil {
		return "", err
	}
	return presigned.String(), nil
}
