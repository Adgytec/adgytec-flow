package duration

import (
	"os"
	"time"

	"github.com/rs/zerolog/log"
)

func GetFromEnv(key string, defaultDur time.Duration) time.Duration {
	durStr := os.Getenv(key)
	if durStr == "" {
		return defaultDur
	}

	dur, parseErr := time.ParseDuration(durStr)
	if parseErr != nil {
		log.Error().
			Err(parseErr).
			Str("action", "duration-parse").
			Str("key", key).
			Send()
		return defaultDur
	}

	return dur
}
