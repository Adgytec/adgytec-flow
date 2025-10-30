package cdn

import "time"

const defaultSignedURLExpiration = time.Hour
const cdnExpiryKey = "CDN_EXPIRY"

type CDN interface {
	GetSignedURL(*string) *string
}
