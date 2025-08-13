package helpers

import (
	"encoding/base64"
	"time"
)

func encodeTimeToBase64(payload time.Time) string {
	bytePayload, convErr := payload.MarshalBinary()
	if convErr != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(bytePayload)
}

func decodeTimeFromBase64(payload string) *time.Time {
	bytePayload, decodeErr := base64.RawStdEncoding.DecodeString(payload)
	if decodeErr != nil {
		return nil
	}

	var timeVal time.Time
	timeConvErr := timeVal.UnmarshalBinary(bytePayload)
	if timeConvErr != nil {
		return nil
	}

	return &timeVal
}
