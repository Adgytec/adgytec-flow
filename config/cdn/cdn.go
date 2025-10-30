package cdn

import "time"

const defaultSignedURLExpiration = time.Hour
const cdnExpiryKey = "CDN_EXPIRY"

type CDN interface {
	GetSignedURL(bucketPath *string) *string
	GetSignedURLWithDuration(bucketPath *string, expiryIn time.Duration) *string
}
