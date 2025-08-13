package core

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
	case LatestFirst, OldestFirst:
		return val
	}
	return LatestFirst
}

const (
	LatestFirst PaginationRequestSorting = "latest"
	OldestFirst PaginationRequestSorting = "oldest"
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
