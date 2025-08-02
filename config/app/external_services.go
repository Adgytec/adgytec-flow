package app

import (
	"github.com/Adgytec/adgytec-flow/config/auth"
	configAWS "github.com/Adgytec/adgytec-flow/config/aws"
	"github.com/Adgytec/adgytec-flow/config/communication"
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type externalServices struct {
	auth          core.IAuth
	database      core.IDatabase
	communicaiton core.ICommunicaiton
}

func (s *externalServices) Auth() core.IAuth {
	return s.auth
}

func (s *externalServices) Database() core.IDatabase {
	return s.database
}

func (s *externalServices) Communication() core.ICommunicaiton {
	return s.communicaiton
}

func createExternalServices() iAppExternalServices {
	awsConfig := configAWS.CreateAWSConfig()

	return &externalServices{
		auth:          auth.CreateCognitoAuthClient(awsConfig),
		database:      database.CreatePgxDbConnectionPool(),
		communicaiton: communication.CreateCommunicationClient(awsConfig),
	}
}
