package communication

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

type communicationEmailSES struct {
	client *ses.Client
}

func (c *communicationEmailSES) SendMail(to []string, from string) error {
	return nil
}

func newSesClient(awsConfig aws.Config) iCommunicationEmail {
	log.Println("creating ses client")
	return &communicationEmailSES{
		client: ses.NewFromConfig(awsConfig),
	}
}
