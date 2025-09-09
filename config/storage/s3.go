package storage

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Client struct {
	client        *s3.Client
	presignClient *s3.PresignClient
}

func (s *s3Client) GetPresignUploadUrl(bucketPath string) (string, error) {
	return "", nil
}

func NewS3Client(awsConfig aws.Config) Storage {
	log.Println("creating s3 presign client")

	client := s3.NewFromConfig(awsConfig)
	return &s3Client{
		client:        client,
		presignClient: s3.NewPresignClient(client),
	}
}
