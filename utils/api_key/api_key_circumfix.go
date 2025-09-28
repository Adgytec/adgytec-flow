package apikey

import (
	"encoding/hex"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	apiKeyPrefix byte
	apiKeySuffix byte
	apiKeyOnce   sync.Once
)

// apiKeyCircumfix() panics when invalid hex values are found for prefix and suffix
// only single byte value is required
func apiKeyCircumfix() (byte, byte) {
	apiKeyOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println("error loading .env file")
		}

		prefixString := os.Getenv("API_KEY_PREFIX")
		suffixString := os.Getenv("API_KEY_SUFFIX")

		prefixByte, prefixErr := hex.DecodeString(prefixString)
		if prefixErr != nil || len(prefixByte) != 1 {
			log.Fatal("invalid hex value for api key prefix. Only single byte value is required")
		}
		apiKeyPrefix = prefixByte[0]

		suffixByte, suffixErr := hex.DecodeString(suffixString)
		if suffixErr != nil || len(suffixByte) != 1 {
			log.Fatal("invalid hex value for api key suffix. Only single byte value is required")
		}
		apiKeySuffix = suffixByte[0]
	})

	return apiKeyPrefix, apiKeySuffix
}
