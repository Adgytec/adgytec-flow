package apikey

import (
	"encoding/hex"
	"os"
	"sync"
)

type apiKeyCircumfixValue struct {
	prefix byte
	suffix byte
}

var (
	circumfixValue     apiKeyCircumfixValue
	apiKeyOnce         sync.Once
	apiKeyCircumfixErr error
)

// only single byte value is required for prefix and suffix
func apiKeyCircumfix() (apiKeyCircumfixValue, error) {
	apiKeyOnce.Do(func() {
		prefixString := os.Getenv("API_KEY_PREFIX")
		suffixString := os.Getenv("API_KEY_SUFFIX")

		prefixByte, prefixErr := hex.DecodeString(prefixString)
		if prefixErr != nil || len(prefixByte) != 1 {
			apiKeyCircumfixErr = ErrInvalidApiKeyPrefixValue
			return
		}
		circumfixValue.prefix = prefixByte[0]

		suffixByte, suffixErr := hex.DecodeString(suffixString)
		if suffixErr != nil || len(suffixByte) != 1 {
			apiKeyCircumfixErr = ErrInvalidApiKeySuffixValue
			return
		}
		circumfixValue.suffix = suffixByte[0]
	})

	return circumfixValue, apiKeyCircumfixErr
}
