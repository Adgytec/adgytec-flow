package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func (a *authCommon) CompareSignedHash(hash string, payload ...[]byte) error {
	mac := hmac.New(sha256.New, a.secret)

	// write data sequentially
	for _, data := range payload {
		_, err := mac.Write(data)
		if err != nil {
			return err
		}
	}

	expectedHash := mac.Sum(nil)

	decodedHash, decodeErr := hex.DecodeString(hash)
	if decodeErr != nil {
		return &InvalidHashError{}
	}

	if !hmac.Equal(decodedHash, expectedHash) {
		return &HashMismatchError{}
	}

	return nil
}
