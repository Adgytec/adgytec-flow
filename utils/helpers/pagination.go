package helpers

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

const (
	PaginationLimit  = 25
	SearchQueryLimit = PaginationLimit * 2
)

func encodeTimeToBase64(payload time.Time) string {
	bytePayload, convErr := payload.MarshalBinary()
	if convErr != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(bytePayload)
}

func DecodeCursorValue(cursor string) *time.Time {
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

func GetPaginationParamsFromRequest(r *http.Request) core.PaginationRequestParams {
	return core.PaginationRequestParams{
		NextCursor:  GetRequestQueryValue(r, NextCursor),
		PrevCursor:  GetRequestQueryValue(r, PrevCursor),
		Sorting:     core.PaginationRequestSorting(GetRequestQueryValue(r, Sort)).Value(),
		SearchQuery: GetRequestQueryValue(r, SearchQuery),
	}
}

func NewPaginationResponse[T core.IPaginationItem](items []T, next, prev *T) *core.ResponsePagination[T] {
	var pageInfo core.PageInfo

	if next != nil {
		// has next page
		pageInfo.HasNextPage = true
		pageInfo.NextCursor = encodeTimeToBase64((*next).GetCreatedAt())
	}

	if prev != nil {
		// has prev page
		pageInfo.HasPrevPage = true
		pageInfo.PrevCursor = encodeTimeToBase64((*prev).GetCreatedAt())
	}

	return &core.ResponsePagination[T]{
		PageInfo:  pageInfo,
		PageItems: items,
	}
}
