package cdn

import "time"

const validDuration = time.Hour

type CDN interface {
	GetSignedUrl(*string) *string
}
