package pagination

import (
	"context"
	"time"

	"github.com/Adgytec/adgytec-flow/config/cache"
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

// PaginationFuncQuery defines a func required for getting paginated data with search query
type PaginationFuncQuery[T any] func(ctx context.Context, searchQuery string, limit int) ([]T, error)

// PaginationFuncInitial defines a func requried for getting initial page data
type PaginationFuncInitial[T any] func(ctx context.Context, limit int) ([]T, error)

// PaginationFuncCursor defines a func required for ggetting pages using cursor
// cursor actual evaluation is done by client providing the funcs
type PaginationFuncCursor[T any] func(ctx context.Context, cursor string, limit int) ([]T, error)

// PaginationFuncToModel converts db response model to acutal item model which can be used by applications
type PaginationFuncToModel[T any, M any] func(items []T) []M

// PaginationActions types T defines item type send by database-query resolution, M defines model used in applications
type PaginationActions[T any, M any] struct {
	Query                        PaginationFuncQuery[T]
	InitialLatestFirst           PaginationFuncInitial[T]
	InitialOldestFirst           PaginationFuncInitial[T]
	GreaterThanCursorLatestFirst PaginationFuncCursor[T]
	GreaterThanCursorOldestFirst PaginationFuncCursor[T]
	LesserThanCursorLatestFirst  PaginationFuncCursor[T]
	LesserThanCursorOldestFirst  PaginationFuncCursor[T]
	ToModel                      PaginationFuncToModel[T, M]
	Cache                        cache.Cache[ResponsePagination[M]]
}
