package cdn

import "errors"

var (
	ErrInvalidCloudfrontConfig     = errors.New("invalid cloudfront config")
	ErrInvalidCloudfrontPrivateKey = errors.New("invalid cloudfront private key")
)
