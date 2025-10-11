package communication

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/rs/zerolog/log"
)

type communicationEmailSES struct {
	client *ses.Client
}

func (c *communicationEmailSES) SendMail(to []string, from string) error {
	return nil
}

func newSesClient(awsConfig aws.Config) communicationEmail {
	log.Info().Msg("new ses client")
	return &communicationEmailSES{
		client: ses.NewFromConfig(awsConfig),
	}
}
