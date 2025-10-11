package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/rs/zerolog/log"
)

func NewAWSConfig() (aws.Config, error) {
	log.Info().Msg("loading aws config")

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return aws.Config{}, &InvalidAWSConfigError{
			cause: err,
		}
	}

	return cfg, nil
}
