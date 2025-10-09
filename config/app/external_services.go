package app

import (
	"context"
	"errors"
	"net/url"
	"os"
	"strings"

	"github.com/Adgytec/adgytec-flow/config/auth"
	configAWS "github.com/Adgytec/adgytec-flow/config/aws"
	"github.com/Adgytec/adgytec-flow/config/cache"
	"github.com/Adgytec/adgytec-flow/config/cdn"
	"github.com/Adgytec/adgytec-flow/config/communication"
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/config/storage"
)

type externalServices struct {
	auth          auth.Auth
	database      database.DatabaseWithShutdown
	communication communication.Communication
	storage       storage.Storage
	cdn           cdn.CDN
	cacheClient   cache.CacheClient
}

func (s *externalServices) Auth() auth.Auth {
	return s.auth
}

func (s *externalServices) Database() database.Database {
	return s.database
}

func (s *externalServices) Communication() communication.Communication {
	return s.communication
}

func (s *externalServices) Storage() storage.Storage {
	return s.storage
}

func (s *externalServices) CDN() cdn.CDN {
	return s.cdn
}

func (s *externalServices) Shutdown(ctx context.Context) {
	s.database.Shutdown()
}

func (s *externalServices) CacheClient() cache.CacheClient {
	return s.cacheClient
}

func newExternalServices() (appExternalServices, error) {
	awsConfig, awsConfigErr := configAWS.NewAWSConfig()
	if awsConfigErr != nil {
		return nil, awsConfigErr
	}

	// parse api endpoint
	urlString := os.Getenv("API_ENDPOINT")
	if strings.TrimSpace(urlString) == "" {
		return nil, errors.New("missing API_ENDPOINT env variable")
	}

	apiURL, parseErr := url.Parse(urlString)
	if parseErr != nil {
		return nil, parseErr
	}

	authClient, authErr := auth.NewCognitoAuthClient(awsConfig, apiURL)
	if authErr != nil {
		return nil, authErr
	}

	dbPool, dbErr := database.NewPgxDbConnectionPool()
	if dbErr != nil {
		return nil, dbErr
	}

	storageClient, storageErr := storage.NewS3Client(awsConfig)
	if storageErr != nil {
		return nil, storageErr
	}

	cdnClient, cdnErr := cdn.NewCloudfrontCDNSigner()
	if cdnErr != nil {
		return nil, cdnErr
	}

	return &externalServices{
		auth:          authClient,
		database:      dbPool,
		storage:       storageClient,
		communication: communication.NewAWSCommunicationClient(awsConfig),
		cdn:           cdnClient,
		cacheClient:   cache.NewInMemoryCacheClient(),
	}, nil
}
