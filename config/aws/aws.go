package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func CreateAWSConfig() aws.Config {
	log.Println("loading aws config")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("error loading AWS config: %s", err)
	}

	return cfg
}
