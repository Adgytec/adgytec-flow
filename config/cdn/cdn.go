package cdn

import "time"

const signedURLExpiration = time.Hour

type CDN interface {
	GetSignedUrl(*string) *string
}
