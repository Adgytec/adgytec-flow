package pagination

import (
	"context"
	"crypto/sha256"
	"fmt"
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

type paginationRequestSorting string

func (val paginationRequestSorting) Value() paginationRequestSorting {
	switch val {
	case paginationRequestSortingLatestFirst, paginationRequestSortingOldestFirst:
		return val
	}
	return paginationRequestSortingLatestFirst
}

const (
	paginationRequestSortingLatestFirst paginationRequestSorting = "latest"
	paginationRequestSortingOldestFirst paginationRequestSorting = "oldest"
)

// if multiple conflicting values are presentin PaginationRequestParams values are chosen in following order
// SearchQuery
// NextCursor
// PrevCursor
// Note: Sorting defines the sorting for the actual data that is stored in persistent storage
// backend doesn't concerns itself for current page sorting (except for search query results as it doesn't have any next and prev page)
type PaginationRequestParams struct {
	NextCursor  string
	PrevCursor  string
	Sorting     paginationRequestSorting
	SearchQuery string
}

func (p PaginationRequestParams) cacheID() string {
	var id string

	switch {
	case p.SearchQuery != "":
		queryHash := sha256.Sum256([]byte(p.SearchQuery))
		id = fmt.Sprintf("query:%x", queryHash[:16])
	case p.NextCursor != "":
		id = fmt.Sprintf("next:%s", p.NextCursor)
	case p.PrevCursor != "":
		id = fmt.Sprintf("prev:%s", p.PrevCursor)
	default:
		id = "initial"
	}

	return fmt.Sprintf("%s:%s", id, p.Sorting)
}

// PaginationFuncQuery defines a func required for getting paginated data with search query
type PaginationFuncQuery[T any] func(ctx context.Context, searchQuery string, limit int32) ([]T, error)

// PaginationFuncInitial defines a func required for getting initial page data
type PaginationFuncInitial[T any] func(ctx context.Context, limit int32) ([]T, error)

// PaginationFuncCursor defines a func required for getting pages using cursor
// cursor actual evaluation is done by client providing the funcs
type PaginationFuncCursor[T any] func(ctx context.Context, cursor time.Time, limit int32) ([]T, error)

// PaginationFuncToModel converts db response model to actual item model which can be used by applications
type PaginationFuncToModel[T any, M PaginationItem] func(items []T) []M

// PaginationActions types T defines item type sent by database-query resolution, M defines model used in applications
type PaginationActions[T any, M PaginationItem] struct {
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

func (a *PaginationActions[T, M]) checkEssentials() error {
	// using ErrpaginationActionNotImplemented because we are only checking essential method that should always be present and should always lead to 500 response
	if a.ToModel == nil || a.Cache == nil {
		return ErrPaginationActionNotImplemented
	}

	return nil
}
