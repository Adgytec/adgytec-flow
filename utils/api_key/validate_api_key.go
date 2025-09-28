package apikey

import (
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/google/uuid"
)

func checkCircumfixBytes(payload []byte, apiKey string) ([]byte, error) {
	invalidApiKeyErr := &InvalidApiKeyError{
		apiKey: apiKey,
	}

	// payload len check at least 3 (prefix + base + suffix)
	if len(payload) < 3 {
		return nil, invalidApiKeyErr
	}

	prefixByte, suffixByte := apiKeyCircumfix()

	// check prefix
	if payload[0] != prefixByte {
		return nil, invalidApiKeyErr
	}

	// check suffix
	if payload[len(payload)-1] != suffixByte {
		return nil, invalidApiKeyErr
	}

	// get base payload
	base := payload[1 : len(payload)-1]
	return base, nil
}

func ValidateApiKey(apiKey string) (uuid.UUID, error) {
	apiKeyBytes, decodeErr := decodeStringToApiKeyValue(apiKey)
	if decodeErr != nil {
		return uuid.Nil, decodeErr
	}

	baseBytes, formatErr := checkCircumfixBytes(apiKeyBytes, apiKey)
	if formatErr != nil {
		return uuid.Nil, formatErr
	}

	return core.GetIDFromPayload(baseBytes), nil
}
