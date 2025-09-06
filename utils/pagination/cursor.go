package pagination

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

func decodeCursorValue(cursor string) *time.Time {
	byteCursor, decodeErr := base64.RawStdEncoding.DecodeString(cursor)
	if decodeErr != nil {
		return nil
	}

	var timeVal time.Time
	timeConvErr := timeVal.UnmarshalBinary(byteCursor)
	if timeConvErr != nil {
		return nil
	}

	return &timeVal
}
