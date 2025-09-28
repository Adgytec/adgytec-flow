package cdn

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidCloudfrontConfig = errors.New("invalid cloudfront config")
)

type InvalidCloudfrontPrivateKeyError struct {
	cause error
}

func (e *InvalidCloudfrontPrivateKeyError) Error() string {
	return fmt.Sprintf("invalid cloudfront private key: %v", e.cause)
}
