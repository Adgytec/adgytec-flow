package pagination

import (
	"context"
)

// GetPaginatedData T defines db response type and M defines model used in application
// this is the function called by the service which require paginated response
func GetPaginatedData[T any, M any](
	ctx context.Context,
	reqParams PaginationRequestParams,
	actions PaginationActions[T, M],
) (*ResponsePagination[M], error) {
	return nil, nil
}

func getPageByQuery[T any, M any](
	ctx context.Context,
	searchQuery string,
	actions *PaginationActions[T, M],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]
	return zero, nil
}

func getInitialPage[T any, M any](
	ctx context.Context,
	sort PaginationRequestSorting,
	actions *PaginationActions[T, M],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]
	return zero, nil
}

func getNextPage[T any, M any](
	ctx context.Context,
	nextCursor string,
	sort PaginationRequestSorting,
	actions *PaginationActions[T, M],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]
	return zero, nil
}

func getPrevPage[T any, M any](
	ctx context.Context,
	prevCursor string,
	sort PaginationRequestSorting,
	actions *PaginationActions[T, M],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]
	return zero, nil
}
