package pagination

import (
	"net/http"
	"strings"
)

type PaginationKey string

const (
	NextCursor  PaginationKey = "next-cursor"
	PrevCursor  PaginationKey = "prev-cursor"
	Sort        PaginationKey = "sort"
	SearchQuery PaginationKey = "search"
)

func GetRequestQueryValue(r *http.Request, key PaginationKey) string {
	queryVal := r.URL.Query().Get(string(key))
	return strings.TrimSpace(queryVal)
}

func GetPaginationParamsFromRequest(r *http.Request) PaginationRequestParams {
	return PaginationRequestParams{
		NextCursor:  GetRequestQueryValue(r, NextCursor),
		PrevCursor:  GetRequestQueryValue(r, PrevCursor),
		Sorting:     PaginationRequestSorting(GetRequestQueryValue(r, Sort)).Value(),
		SearchQuery: GetRequestQueryValue(r, SearchQuery),
	}
}
