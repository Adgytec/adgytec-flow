package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func NewAWSConfig() (aws.Config, error) {
	log.Println("loading aws config")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return aws.Config{}, &InvalidAWSConfigError{
			cause: err,
		}
	}

	return cfg, nil
}
