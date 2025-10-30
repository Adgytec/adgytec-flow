package cdn

import "time"

var validDuration = time.Hour * 1

type CDN interface {
	GetSignedUrl(*string) *string
}
