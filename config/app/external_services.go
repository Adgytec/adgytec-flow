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
	auth          core.Auth
	database      core.DatabaseWithShutdown
	communicaiton core.Communicaiton
	storage       core.Storage
	cdn           core.CDN
	cacheClient   core.CacheClient
}

func (s *externalServices) Auth() core.Auth {
	return s.auth
}

func (s *externalServices) Database() core.Database {
	return s.database
}

func (s *externalServices) Communication() core.Communicaiton {
	return s.communicaiton
}

func (s *externalServices) Storage() core.Storage {
	return s.storage
}

func (s *externalServices) CDN() core.CDN {
	return s.cdn
}

func (s *externalServices) Shutdown() {
	s.database.Shutdown()
}

func (s *externalServices) CacheClient() core.CacheClient {
	return s.cacheClient
}

func newExternalServices() appExternalServices {
	awsConfig := configAWS.NewAWSConfig()

	return &externalServices{
		auth:          auth.NewCognitoAuthClient(awsConfig),
		database:      database.NewPgxDbConnectionPool(),
		communicaiton: communication.NewAWSCommunicationClient(awsConfig),
		storage:       storage.NewS3Client(awsConfig),
		cdn:           cdn.NewCloudfrontCDNSigner(),
		cacheClient:   cache.NewInMemoryCacheClient(),
	}
}
