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

	var fetchPageAction, prevPageAction PaginationFuncCursor[T]
	if sort == paginationRequestSortingLatestFirst {
		fetchPageAction = actions.LesserThanCursorLatestFirst
		prevPageAction = actions.GreaterThanCursorLatestFirst
	} else {
		fetchPageAction = actions.GreaterThanCursorOldestFirst
		prevPageAction = actions.LesserThanCursorOldestFirst
	}

	// requires items created after next cursor
	return getNextPageUtil(
		ctx,
		*nextCursorVal,
		actions.ToModel,
		fetchPageAction,
		prevPageAction,
	)
}

func getNextPageUtil[T any, M PaginationItem](
	ctx context.Context,
	nextCursor time.Time,
	toModel PaginationFuncToModel[T, M],
	fetchPageAction PaginationFuncCursor[T],
	prevPageAction PaginationFuncCursor[T],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]

	list, listErr := fetchPageAction(ctx, nextCursor, paginationLimit+1)
	if listErr != nil {
		return zero, listErr
	}

	models := toModel(list)
	var next *M
	var prev *M

	// handle next page details
	if len(models) > paginationLimit {
		models = models[:paginationLimit]
		next = &models[len(models)-1]
	}

	if len(models) > 0 {
		prevCursor := models[0].GetCreatedAt()

		prevItem, prevItemErr := prevPageAction(ctx, prevCursor, 1)
		if prevItemErr != nil {
			return zero, prevItemErr
		}

		if len(prevItem) > 0 {
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

	var fetchPageAction, nextPageAction PaginationFuncCursor[T]
	if sort == paginationRequestSortingLatestFirst {
		fetchPageAction = actions.GreaterThanCursorLatestFirst
		nextPageAction = actions.LesserThanCursorLatestFirst
	} else {
		fetchPageAction = actions.LesserThanCursorOldestFirst
		nextPageAction = actions.GreaterThanCursorOldestFirst
	}

	// requires items created before prev cursor
	return getPrevPageUtil(
		ctx,
		*prevCursorVal,
		actions.ToModel,
		fetchPageAction,
		nextPageAction,
	)
}

func getPrevPageUtil[T any, M PaginationItem](
	ctx context.Context,
	prevCursor time.Time,
	toModel PaginationFuncToModel[T, M],
	fetchPageAction PaginationFuncCursor[T],
	nextPageAction PaginationFuncCursor[T],
) (ResponsePagination[M], error) {
	var zero ResponsePagination[M]

	list, listErr := fetchPageAction(ctx, prevCursor, paginationLimit+1)
	if listErr != nil {
		return zero, listErr
	}

	models := toModel(list)
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

		nextItem, nextItemErr := nextPageAction(ctx, nextCursor, 1)
		if nextItemErr != nil {
			return zero, nextItemErr
		}

		if len(nextItem) > 0 {
			next = &models[len(models)-1]
		}
	}

	return newPaginationResponse(models, next, prev), nil
}
