package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func (a *authCommon) NewSignedHash(payload ...[]byte) (string, error) {
	mac := hmac.New(sha256.New, a.secret)

	// write data sequentially
	for _, data := range payload {
		_, err := mac.Write(data)
		if err != nil {
			return "", err
		}
	}

	hashBytes := mac.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}
