package pagination

import (
	"context"
	"slices"
)

// GetPaginatedData T defines db response type and M defines model used in application
// this is the function called by the service which require paginated response
func GetPaginatedData[T any, M any](
	ctx context.Context,
	reqParams PaginationRequestParams,
	actions PaginationActions[T, M],
) (*ResponsePagination[M], error) {

	res, resErr := actions.Cache.Get(reqParams.cacheID(), func() (ResponsePagination[M], error) {
		switch {
		case reqParams.SearchQuery != "":
			return getPageByQuery(ctx, reqParams.SearchQuery, reqParams.Sorting, &actions)
		case reqParams.NextCursor != "":
			return getNextPage(ctx, reqParams.NextCursor, reqParams.Sorting, &actions)
		case reqParams.PrevCursor != "":
			return getPrevPage(ctx, reqParams.PrevCursor, reqParams.Sorting, &actions)
		default:
			return getInitialPage(ctx, reqParams.Sorting, &actions)
		}
	})
	if resErr != nil {
		return nil, resErr
	}

	return &res, nil
}

// utilities funcs to fetch actual data

func getPageByQuery[T any, M any](
	ctx context.Context,
	searchQuery string,
	sort PaginationRequestSorting,
	actions *PaginationActions[T, M],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]

	list, listErr := actions.Query(ctx, searchQuery, SearchQueryLimit)
	if listErr != nil {
		return zero, listErr
	}

	models := actions.ToModel(list)

	// handle ordering
	if sort == PaginationRequestSortingOldestFirst {
		slices.Reverse(models)
	}
	return NewPaginationResponse(models, nil, nil), nil
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
