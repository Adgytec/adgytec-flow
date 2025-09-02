package app

import (
	"github.com/Adgytec/adgytec-flow/config/auth"
	configAWS "github.com/Adgytec/adgytec-flow/config/aws"
	"github.com/Adgytec/adgytec-flow/config/cache"
	"github.com/Adgytec/adgytec-flow/config/cdn"
	"github.com/Adgytec/adgytec-flow/config/communication"
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/config/storage"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type externalServices struct {
	auth          core.IAuth
	database      core.IDatabaseWithShutdown
	communicaiton core.ICommunicaiton
	storage       core.IStorage
	cdn           core.ICDN
	cacheClient   core.ICacheClient
	serializer    core.ISerializer[any]
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

func (s *externalServices) Storage() core.IStorage {
	return s.storage
}

func (s *externalServices) CDN() core.ICDN {
	return s.cdn
}

func (s *externalServices) Shutdown() {
	s.database.Shutdown()
}

func (s *externalServices) CacheClient() core.ICacheClient {
	return s.cacheClient
}

func (s *externalServices) Serializer() core.ISerializer[any] {
	return s.serializer
}

func createExternalServices() iAppExternalServices {
	awsConfig := configAWS.CreateAWSConfig()

	return &externalServices{
		auth:          auth.CreateCognitoAuthClient(awsConfig),
		database:      database.CreatePgxDbConnectionPool(),
		communicaiton: communication.CreateAWSCommunicationClient(awsConfig),
		storage:       storage.CreateS3Client(awsConfig),
		cdn:           cdn.CreateCloudfrontCDNSigner(),
		cacheClient:   cache.CreateInMemoryCacheClient(),
	}
}
