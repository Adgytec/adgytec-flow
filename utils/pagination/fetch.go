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
	actions *PaginationActions[T, M],
) (*ResponsePagination[M], error) {
	res, resErr := actions.Cache.Get(reqParams.cacheID(), func() (ResponsePagination[M], error) {
		switch {
		case reqParams.SearchQuery != "":
			return getPageByQuery(ctx, reqParams.SearchQuery, reqParams.Sorting, actions)
		case reqParams.NextCursor != "":
			return getNextPage(ctx, reqParams.NextCursor, reqParams.Sorting, actions)
		case reqParams.PrevCursor != "":
			return getPrevPage(ctx, reqParams.PrevCursor, reqParams.Sorting, actions)
		default:
			return getInitialPage(ctx, reqParams.Sorting, actions)
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
	sort paginationRequestSorting,
	actions *PaginationActions[T, M],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]

	list, listErr := actions.Query(ctx, searchQuery, searchQueryLimit)
	if listErr != nil {
		return zero, listErr
	}

	models := actions.ToModel(list)

	// handle ordering
	if sort == paginationRequestSortingOldestFirst {
		slices.Reverse(models)
	}
	return newPaginationResponse(models, nil, nil), nil
}

func getInitialPage[T any, M PaginationItem](
	ctx context.Context,
	sort paginationRequestSorting,
	actions *PaginationActions[T, M],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]

	var list []T
	var listErr error

	if sort == paginationRequestSortingLatestFirst {
		list, listErr = actions.InitialLatestFirst(ctx, paginationLimit+1)
	} else {
		list, listErr = actions.InitialOldestFirst(ctx, paginationLimit+1)
	}
	if listErr != nil {
		return zero, listErr
	}

	models := actions.ToModel(list)
	var next *M

	// handle next page details
	if len(models) > paginationLimit {
		models = models[:paginationLimit]
		next = &models[len(models)-1]
	}
	return newPaginationResponse(models, next, nil), nil
}

func getNextPage[T any, M PaginationItem](
	ctx context.Context,
	nextCursor string,
	sort paginationRequestSorting,
	actions *PaginationActions[T, M],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]

	nextCursorVal := decodeCursorValue(nextCursor)
	if nextCursorVal == nil {
		return zero, &InvalidCursorValueError{
			Cursor: nextCursor,
		}
	}

	if sort == paginationRequestSortingLatestFirst {
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

	// require item created before next cursor
	list, listErr := actions.LesserThanCursorLatestFirst(ctx, nextCursor, paginationLimit+1)
	if listErr != nil {
		return zero, listErr
	}

	models := actions.ToModel(list)
	var next *M
	var prev *M

	// handle next page details
	if len(models) > paginationLimit {
		models = models[:paginationLimit]
		next = &models[len(models)-1]
	}

	if len(models) > 0 {
		prevCursor := models[0].GetCreatedAt()

		prevItem, prevItemErr := actions.GreaterThanCursorLatestFirst(ctx, prevCursor, 1)
		if prevItemErr == nil && len(prevItem) > 0 {
			prev = &models[0]
		}
	}

	return newPaginationResponse(models, next, prev), nil
}

func getNextPageOldestFirst[T any, M PaginationItem](
	ctx context.Context,
	nextCursor time.Time,
	actions *PaginationActions[T, M],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]

	// require items created after next cursor
	list, listErr := actions.GreaterThanCursorOldestFirst(ctx, nextCursor, paginationLimit+1)
	if listErr != nil {
		return zero, listErr
	}

	models := actions.ToModel(list)
	var next *M
	var prev *M

	// handle next page
	if len(models) > paginationLimit {
		models = models[:paginationLimit]
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

	return newPaginationResponse(models, next, prev), nil
}

func getPrevPage[T any, M PaginationItem](
	ctx context.Context,
	prevCursor string,
	sort paginationRequestSorting,
	actions *PaginationActions[T, M],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]

	prevCursorVal := decodeCursorValue(prevCursor)
	if prevCursorVal == nil {
		return zero, &InvalidCursorValueError{
			Cursor: prevCursor,
		}
	}

	if sort == paginationRequestSortingLatestFirst {
		return getPrevPageLatestFirst(ctx, *prevCursorVal, actions)
	}
	return getPrevPageOldestFirst(ctx, *prevCursorVal, actions)
}

func getPrevPageLatestFirst[T any, M PaginationItem](
	ctx context.Context,
	prevCursor time.Time,
	actions *PaginationActions[T, M],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]

	// cursor is time based so when fetching previous page with latest first,
	// we need items that are created after prev cursor with latest first
	list, listErr := actions.GreaterThanCursorLatestFirst(ctx, prevCursor, paginationLimit+1)
	if listErr != nil {
		return zero, listErr
	}

	models := actions.ToModel(list)
	var next *M
	var prev *M

	// handle prev page
	if len(models) > paginationLimit {
		models = models[1:]
		prev = &models[0]
	}

	// handle next page
	if len(models) > 0 {
		nextCursor := models[len(models)-1].GetCreatedAt()

		nextItem, nextItemErr := actions.LesserThanCursorLatestFirst(ctx, nextCursor, 1)
		if nextItemErr == nil && len(nextItem) > 0 {
			next = &models[len(models)-1]
		}
	}

	return newPaginationResponse(models, next, prev), nil
}

func getPrevPageOldestFirst[T any, M PaginationItem](
	ctx context.Context,
	prevCursor time.Time,
	actions *PaginationActions[T, M],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]

	// require items that are created prior to prev cursor
	list, listErr := actions.LesserThanCursorOldestFirst(ctx, prevCursor, paginationLimit+1)
	if listErr != nil {
		return zero, listErr
	}

	models := actions.ToModel(list)
	var next *M
	var prev *M

	// handle prev page
	if len(models) > paginationLimit {
		models = models[1:]
		prev = &models[0]
	}

	// handle next page
	if len(models) > 0 {
		nextCursor := models[len(models)-1].GetCreatedAt()

		nextItem, nextItemErr := actions.GreaterThanCursorOldestFirst(ctx, nextCursor, 1)
		if nextItemErr == nil && len(nextItem) > 0 {
			next = &models[len(models)-1]
		}
	}

	return newPaginationResponse(models, next, prev), nil
}
