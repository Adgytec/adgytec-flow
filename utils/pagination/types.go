package pagination

import "time"

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
