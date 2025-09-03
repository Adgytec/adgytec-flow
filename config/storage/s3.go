package storage

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3PresignClient struct {
	client *s3.PresignClient
}

func (s *s3PresignClient) GetPresignUploadUrl(bucketPath string) (string, error) {
	return "", nil
}

func NewS3Client(awsConfig aws.Config) core.IStorage {
	log.Println("creating s3 presign client")
	return &s3PresignClient{
		client: s3.NewPresignClient(
			s3.NewFromConfig(awsConfig),
		),
	}
}
