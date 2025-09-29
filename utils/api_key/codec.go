package apikey

import "encoding/base64"

func encodeApiKeyBytesToBase64(payload []byte) string {
	return base64.StdEncoding.EncodeToString(payload)
}

func decodeStringToApiKeyValue(apiKey string) ([]byte, error) {
	byteApiKey, decodeErr := base64.StdEncoding.DecodeString(apiKey)
	if decodeErr != nil {
		return nil, &InvalidApiKeyError{
			apiKey: apiKey,
		}
	}

	return byteApiKey, nil
}
