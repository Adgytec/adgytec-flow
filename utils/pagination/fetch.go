package pagination

import (
	"context"
	"slices"
	"time"
)

// GetPaginatedData T defines db response type and M defines model used in application
// this is the function called by the service which require paginated response
func GetPaginatedData[T any, M PaginationItem](
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

func getPageByQuery[T any, M PaginationItem](
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

func getInitialPage[T any, M PaginationItem](
	ctx context.Context,
	sort PaginationRequestSorting,
	actions *PaginationActions[T, M],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]

	var list []T
	var listErr error

	if sort == PaginationRequestSortingLatestFirst {
		list, listErr = actions.InitialLatestFirst(ctx, PaginationLimit)
	} else {
		list, listErr = actions.InitialOldestFirst(ctx, PaginationLimit)
	}
	if listErr != nil {
		return zero, listErr
	}

	models := actions.ToModel(list)
	var next *M

	// handle next page details
	if len(models) > PaginationLimit {
		models = models[:PaginationLimit]
		next = &models[len(models)-1]
	}
	return NewPaginationResponse(models, next, nil), nil
}

func getNextPage[T any, M PaginationItem](
	ctx context.Context,
	nextCursor string,
	sort PaginationRequestSorting,
	actions *PaginationActions[T, M],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]

	nextCursorVal := DecodeCursorValue(nextCursor)
	if nextCursorVal == nil {
		return zero, &InvalidCursorValueError{
			Cursor: nextCursor,
		}
	}

	if sort == PaginationRequestSortingLatestFirst {
		return getNextPageLatestFirst(ctx, *nextCursorVal, actions)
	}
	return getNextPageOldestFirst(ctx, *nextCursorVal, actions)
}

func getNextPageLatestFirst[T any, M PaginationItem](
	ctx context.Context,
	nextCursor time.Time,
	actions *PaginationActions[T, M],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]

	list, listErr := actions.LesserThanCursorLatestFirst(ctx, nextCursor, PaginationLimit)
	if listErr != nil {
		return zero, listErr
	}

	models := actions.ToModel(list)
	var next *M
	var prev *M

	// handle next page details
	if len(models) > PaginationLimit {
		models = models[:PaginationLimit]
		next = &models[len(models)-1]
	}

	if len(models) > 0 {
		prevCursor := models[0].GetCreatedAt()

		prevItem, prevItemErr := actions.GreaterThanCursorLatestFirst(ctx, prevCursor, 1)
		if prevItemErr == nil && len(prevItem) > 0 {
			prev = &models[0]
		}
	}

	return NewPaginationResponse(models, next, prev), nil
}

func getNextPageOldestFirst[T any, M PaginationItem](
	ctx context.Context,
	nextCursor time.Time,
	actions *PaginationActions[T, M],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]

	list, listErr := actions.GreaterThanCursorOldestFirst(ctx, nextCursor, PaginationLimit)
	if listErr != nil {
		return zero, listErr
	}

	models := actions.ToModel(list)
	var next *M
	var prev *M

	// handle next page
	if len(models) > PaginationLimit {
		models := models[:PaginationLimit]
		next = &models[len(models)-1]
	}

	// handle prev page
	if len(models) > 0 {
		prevCursor := models[0].GetCreatedAt()

		prevItem, prevItemErr := actions.LesserThanCursorOldestFirst(ctx, prevCursor, 1)
		if prevItemErr == nil && len(prevItem) > 0 {
			prev = &models[0]
		}
	}

	return NewPaginationResponse(models, next, prev), nil
}

func getPrevPage[T any, M PaginationItem](
	ctx context.Context,
	prevCursor string,
	sort PaginationRequestSorting,
	actions *PaginationActions[T, M],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]
	return zero, nil
}
