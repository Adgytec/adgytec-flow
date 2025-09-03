package communication

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type communicationEmail interface {
	SendMail([]string, string) error
}

type communicationImpl struct {
	email communicationEmail
}

func (c *communicationImpl) SendMail(to []string, from string) error {
	return c.email.SendMail(to, from)
}

func NewAWSCommunicationClient(awsConfig aws.Config) core.Communication {
	log.Println("creating aws communication client")
	return &communicationImpl{
		email: newSesClient(awsConfig),
	}
}
