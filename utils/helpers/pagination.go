package helpers

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/Adgytec/adgytec-flow/utils/core"
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

func GetPaginationParamsFromRequest(r *http.Request) core.PaginationRequestParams {
	return core.PaginationRequestParams{
		NextCursor:  GetRequestQueryValue(r, NextCursor),
		PrevCursor:  GetRequestQueryValue(r, PrevCursor),
		Sorting:     core.PaginationRequestSorting(GetRequestQueryValue(r, Sort)).Value(),
		SearchQuery: GetRequestQueryValue(r, SearchQuery),
	}
}
