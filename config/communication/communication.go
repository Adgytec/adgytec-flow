package communication

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type Communication interface {
	SendMail([]string, string) error
}

type communicationEmail interface {
	SendMail([]string, string) error
}

type communicationImpl struct {
	email communicationEmail
}

func (c *communicationImpl) SendMail(to []string, from string) error {
	return c.email.SendMail(to, from)
}

func NewAWSCommunicationClient(awsConfig aws.Config) Communication {
	log.Println("creating aws communication client")
	return &communicationImpl{
		email: newSesClient(awsConfig),
	}
}
