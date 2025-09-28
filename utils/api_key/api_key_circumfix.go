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
// only single byte value is required, if multiple bytes are provided only first byte is used
func apiKeyCircumfix() (byte, byte) {
	apiKeyOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println("error loading .env file")
		}

		prefixString := os.Getenv("API_KEY_PREFIX")
		suffixString := os.Getenv("API_KEY_SUFFIX")

		prefixByte, prefixErr := hex.DecodeString(prefixString)
		if prefixErr != nil {
			log.Fatal("invalid hex value for api key prefx")
		}
		apiKeyPrefix = prefixByte[0]

		suffixByte, suffixErr := hex.DecodeString(suffixString)
		if suffixErr != nil {
			log.Fatal("invalid hex value for api key suffix")
		}
		apiKeySuffix = suffixByte[0]
	})

	return apiKeyPrefix, apiKeySuffix
}

func GetApiKeyCircumfix() (byte, byte) {
	return apiKeyCircumfix()
}
