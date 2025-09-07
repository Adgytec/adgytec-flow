package pagination

import (
	"net/http"
	"strings"
)

type paginationKey string

const (
	NextCursor  paginationKey = "next-cursor"
	PrevCursor  paginationKey = "prev-cursor"
	Sort        paginationKey = "sort"
	SearchQuery paginationKey = "search"
)

func getRequestQueryValue(r *http.Request, key paginationKey) string {
	queryVal := r.URL.Query().Get(string(key))
	return strings.TrimSpace(queryVal)
}

func GetPaginationParamsFromRequest(r *http.Request) PaginationRequestParams {
	return PaginationRequestParams{
		NextCursor:  getRequestQueryValue(r, NextCursor),
		PrevCursor:  getRequestQueryValue(r, PrevCursor),
		Sorting:     paginationRequestSorting(getRequestQueryValue(r, Sort)).Value(),
		SearchQuery: getRequestQueryValue(r, SearchQuery),
	}
}
