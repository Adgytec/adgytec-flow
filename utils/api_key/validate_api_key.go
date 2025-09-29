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

	circumfixVal, circumfixErr := apiKeyCircumfix()
	if circumfixErr != nil {
		return nil, circumfixErr
	}

	// check prefix
	if payload[0] != circumfixVal.prefix {
		return nil, invalidApiKeyErr
	}

	// check suffix
	if payload[len(payload)-1] != circumfixVal.suffix {
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
