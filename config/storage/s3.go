package storage

import (
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Client struct {
	client            *s3.Client
	presignClient     *s3.PresignClient
	bucket            string
	bucketRegion      string
	presignExpiration time.Duration
}

func NewS3Client(awsConfig aws.Config) (Storage, error) {
	log.Println("creating s3 client")

	bucket := os.Getenv("AWS_S3_STUDIO_BUCKET")
	bucketRegion := os.Getenv("AWS_S3_STUDIO_BUCKET_REGION")

	if bucket == "" || bucketRegion == "" {
		return nil, ErrInvalidS3ConfigValue
	}

	client := s3.NewFromConfig(awsConfig)
	return &s3Client{
		client:            client,
		presignClient:     s3.NewPresignClient(client),
		bucket:            bucket,
		bucketRegion:      bucketRegion,
		presignExpiration: time.Hour,
	}, nil
}
