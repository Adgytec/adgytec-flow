package auth

import (
	apikey "github.com/Adgytec/adgytec-flow/utils/api_key"
	"github.com/google/uuid"
)

func (a *authCommon) ValidateAPIKey(apiKey string) (uuid.UUID, error) {
	return apikey.ValidateApiKey(apiKey)
}
