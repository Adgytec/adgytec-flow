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

type PaginationRequestSorting int

const (
	LatestFirst PaginationRequestSorting = iota
	OldestFirst
)

// priority is given to NextCursor if both the cursor are present
type PaginationRequestParams struct {
	NextCursor *string
	PrevCursor *string
	Sorting    PaginationRequestSorting
}
