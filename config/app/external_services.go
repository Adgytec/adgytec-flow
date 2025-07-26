package app

import (
	"github.com/Adgytec/adgytec-flow/config/auth"
	configAWS "github.com/Adgytec/adgytec-flow/config/aws"
	"github.com/Adgytec/adgytec-flow/utils/interfaces"
)

type externalServices struct {
	auth     interfaces.IAuth
	database interfaces.IDatabase
}

func (s *externalServices) Auth() interfaces.IAuth {
	return s.auth
}

func (s *externalServices) Database() interfaces.IDatabase {
	return s.database
}

func createExternalServices() IAppExternalServices {
	awsConfig := configAWS.CreateAWSConfig()

	return &externalServices{
		auth: auth.CreateCognitoAuthClient(awsConfig),
	}
}
