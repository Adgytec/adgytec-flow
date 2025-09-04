package pagination

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

const (
	PaginationLimit  = 25
	SearchQueryLimit = PaginationLimit * 2
)

type PaginationItem interface {
	GetCreatedAt() time.Time
}

type PageInfo struct {
	HasNextPage bool   `json:"hasNextPage"`
	NextCursor  string `json:"nextCursor"`
	HasPrevPage bool   `json:"hasPrevPage"`
	PrevCursor  string `json:"prevCursor"`
}

type ResponsePagination[T any] struct {
	PageInfo  PageInfo `json:"pageInfo"`
	PageItems []T      `json:"pageItems"`
}

type PaginationRequestSorting string

func (val PaginationRequestSorting) Value() PaginationRequestSorting {
	switch val {
	case PaginationRequestSortingLatestFirst, PaginationRequestSortingOldestFirst:
		return val
	}
	return PaginationRequestSortingLatestFirst
}

const (
	PaginationRequestSortingLatestFirst PaginationRequestSorting = "latest"
	PaginationRequestSortingOldestFirst PaginationRequestSorting = "oldest"
)

// if multiple conflicting values are presentin PaginationRequestParams values are chosen in following order
// SearchQuery
// NextCursor
// PrevCursor
type PaginationRequestParams struct {
	NextCursor  string
	PrevCursor  string
	Sorting     PaginationRequestSorting
	SearchQuery string
}

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

func GetPaginationParamsFromRequest(r *http.Request) PaginationRequestParams {
	return PaginationRequestParams{
		NextCursor:  helpers.GetRequestQueryValue(r, helpers.NextCursor),
		PrevCursor:  helpers.GetRequestQueryValue(r, helpers.PrevCursor),
		Sorting:     PaginationRequestSorting(helpers.GetRequestQueryValue(r, helpers.Sort)).Value(),
		SearchQuery: helpers.GetRequestQueryValue(r, helpers.SearchQuery),
	}
}

func NewPaginationResponse[T PaginationItem](items []T, next, prev *T) *ResponsePagination[T] {
	var pageInfo PageInfo

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

	return &ResponsePagination[T]{
		PageInfo:  pageInfo,
		PageItems: items,
	}
}
