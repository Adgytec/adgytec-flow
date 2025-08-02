package communication

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type iCommunicationEmail interface {
	SendMail([]string, string) error
}

type communicationImpl struct {
	email iCommunicationEmail
}

func (c *communicationImpl) SendMail(to []string, from string) error {
	return c.email.SendMail(to, from)
}

func CreateSESAndSNSCommunicationClient(awsConfig aws.Config) core.ICommunicaiton {
	log.Println("creating ses and sns communication client")
	return &communicationImpl{
		email: createSesClient(awsConfig),
	}
}
