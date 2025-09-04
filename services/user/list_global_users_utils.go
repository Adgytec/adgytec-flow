package user

import (
	"context"
	"slices"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/database/models"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
	"github.com/Adgytec/adgytec-flow/utils/pagination"
)

func (s *userService) getGlobalUsersByQuery(ctx context.Context, params pagination.PaginationRequestParams) (*pagination.ResponsePagination[models.GlobalUser], error) {
	userList, userErr := s.db.Queries().GetGlobalUsersByQuery(
		ctx,
		db.GetGlobalUsersByQueryParams{
			Limit: pagination.SearchQueryLimit,
			Query: params.SearchQuery,
		},
	)

	if userErr != nil {
		return nil, userErr
	}

	userModels := s.getUserResponseModels(userList)

	// handle ordering default is latest first
	if params.Sorting == pagination.PaginationRequestSortingOldestFirst {
		slices.Reverse(userModels)
	}

	return pagination.NewPaginationResponse(userModels, nil, nil), nil
}

func (s *userService) getGlobalUsersInitial(ctx context.Context, params pagination.PaginationRequestParams) (*pagination.ResponsePagination[models.GlobalUser], error) {
	var userList []db.GlobalUserDetail
	var userErr error

	if params.Sorting == pagination.PaginationRequestSortingLatestFirst {
		userList, userErr = s.db.Queries().GetGlobalUsersLatestFirst(ctx, pagination.PaginationLimit+1)
	} else {
		userList, userErr = s.db.Queries().GetGlobalUsersOldestFirst(ctx, pagination.PaginationLimit+1)
	}

	if userErr != nil {
		return nil, userErr
	}

	userModels := s.getUserResponseModels(userList)
	var next *models.GlobalUser

	// handle next page details
	if len(userList) == pagination.PaginationLimit+1 {
		userLen := len(userModels)
		next = &userModels[userLen-2]
		userModels = userModels[:userLen-1]
	}

	return pagination.NewPaginationResponse(userModels, next, nil), nil
}

func (s *userService) getGlobalUsersNextPage(ctx context.Context, params pagination.PaginationRequestParams) (*pagination.ResponsePagination[models.GlobalUser], error) {
	if params.Sorting == pagination.PaginationRequestSortingLatestFirst {
		return s.getGlobalUsersNextPageLatestFirst(ctx, params)
	}

	return s.getGlobalUsersNextPageOldestFirst(ctx, params)
}

func (s *userService) getGlobalUsersNextPageLatestFirst(ctx context.Context, params pagination.PaginationRequestParams) (*pagination.ResponsePagination[models.GlobalUser], error) {
	nextCursorVal := pagination.DecodeCursorValue(params.NextCursor)
	if nextCursorVal == nil {
		return nil, &app_errors.InvalidCursorValueError{
			Cursor: params.NextCursor,
		}
	}

	userList, userErr := s.db.Queries().GetGlobalUsersLatestFirstLesserThanCursor(
		ctx,
		db.GetGlobalUsersLatestFirstLesserThanCursorParams{
			Limit:  pagination.PaginationLimit + 1,
			Cursor: *nextCursorVal,
		},
	)

	if userErr != nil {
		return nil, userErr
	}

	userModels := s.getUserResponseModels(userList)
	var next *models.GlobalUser
	var prev *models.GlobalUser

	// handle next page details
	if len(userModels) > pagination.PaginationLimit {
		userModels = userModels[:pagination.PaginationLimit]
		next = &userModels[len(userModels)-1]
	}

	// handle prev page details
	if len(userModels) > 0 {
		prevCursor := userModels[0].GetCreatedAt()
		prevUser, prevUserErr := s.db.Queries().GetGlobalUsersLatestFirstGreaterThanCursor(
			ctx,
			db.GetGlobalUsersLatestFirstGreaterThanCursorParams{
				Limit:  1,
				Cursor: prevCursor,
			},
		)

		if prevUserErr == nil && len(prevUser) > 0 {
			prev = &userModels[0]
		}

	}

	return pagination.NewPaginationResponse(userModels, next, prev), nil
}

func (s *userService) getGlobalUsersNextPageOldestFirst(ctx context.Context, params pagination.PaginationRequestParams) (*pagination.ResponsePagination[models.GlobalUser], error) {
	nextCursorVal := pagination.DecodeCursorValue(params.NextCursor)
	if nextCursorVal == nil {
		return nil, &app_errors.InvalidCursorValueError{
			Cursor: params.NextCursor,
		}
	}

	userList, userErr := s.db.Queries().GetGlobalUsersOldestFirstGreaterThanCursor(
		ctx,
		db.GetGlobalUsersOldestFirstGreaterThanCursorParams{
			Limit:  pagination.PaginationLimit + 1,
			Cursor: *nextCursorVal,
		},
	)

	if userErr != nil {
		return nil, userErr
	}

	userModels := s.getUserResponseModels(userList)
	var next *models.GlobalUser
	var prev *models.GlobalUser

	// handle next page details
	if len(userModels) > pagination.PaginationLimit {
		userModels = userModels[:pagination.PaginationLimit]
		next = &userModels[len(userModels)-1]
	}

	// handle prev page details
	if len(userModels) > 0 {
		prevCursor := userModels[0].GetCreatedAt()
		prevUser, prevUserErr := s.db.Queries().GetGlobalUsersOldestFirstLesserThanCursor(
			ctx,
			db.GetGlobalUsersOldestFirstLesserThanCursorParams{
				Limit:  1,
				Cursor: prevCursor,
			},
		)

		if prevUserErr == nil && len(prevUser) > 0 {
			prev = &userModels[0]
		}

	}

	return pagination.NewPaginationResponse(userModels, next, prev), nil
}

func (s *userService) getGlobalUsersPrevPage(ctx context.Context, params pagination.PaginationRequestParams) (*pagination.ResponsePagination[models.GlobalUser], error) {
	if params.Sorting == pagination.PaginationRequestSortingLatestFirst {
		return s.getGlobalUsersPrevPageLatestFirst(ctx, params)
	}

	return s.getGlobalUsersPrevPageOldestFirst(ctx, params)

}

func (s *userService) getGlobalUsersPrevPageLatestFirst(ctx context.Context, params pagination.PaginationRequestParams) (*pagination.ResponsePagination[models.GlobalUser], error) {
	prevCursorVal := pagination.DecodeCursorValue(params.PrevCursor)
	if prevCursorVal == nil {
		return nil, &app_errors.InvalidCursorValueError{
			Cursor: params.PrevCursor,
		}
	}

	userList, userErr := s.db.Queries().GetGlobalUsersLatestFirstGreaterThanCursor(
		ctx,
		db.GetGlobalUsersLatestFirstGreaterThanCursorParams{
			Limit:  pagination.PaginationLimit + 1,
			Cursor: *prevCursorVal,
		},
	)
	if userErr != nil {
		return nil, userErr
	}

	userModels := s.getUserResponseModels(userList)
	var next *models.GlobalUser
	var prev *models.GlobalUser

	// handle prev page
	if len(userModels) > pagination.PaginationLimit {
		userModels = userModels[1:]
		prev = &userModels[0]
	}

	// handle next page
	if len(userModels) > 0 {
		nextCursor := userModels[len(userModels)-1].GetCreatedAt()
		nextUser, nextUserErr := s.db.Queries().GetGlobalUsersLatestFirstLesserThanCursor(
			ctx,
			db.GetGlobalUsersLatestFirstLesserThanCursorParams{
				Limit:  1,
				Cursor: nextCursor,
			},
		)

		if nextUserErr == nil && len(nextUser) > 0 {
			next = &userModels[len(userModels)-1]
		}
	}

	return pagination.NewPaginationResponse(userModels, next, prev), nil
}

func (s *userService) getGlobalUsersPrevPageOldestFirst(ctx context.Context, params pagination.PaginationRequestParams) (*pagination.ResponsePagination[models.GlobalUser], error) {
	prevCursorVal := pagination.DecodeCursorValue(params.PrevCursor)
	if prevCursorVal == nil {
		return nil, &app_errors.InvalidCursorValueError{
			Cursor: params.PrevCursor,
		}
	}

	userList, userErr := s.db.Queries().GetGlobalUsersOldestFirstLesserThanCursor(
		ctx,
		db.GetGlobalUsersOldestFirstLesserThanCursorParams{
			Limit:  pagination.PaginationLimit + 1,
			Cursor: *prevCursorVal,
		},
	)
	if userErr != nil {
		return nil, userErr
	}

	userModels := s.getUserResponseModels(userList)
	var next *models.GlobalUser
	var prev *models.GlobalUser

	// handle prev page
	if len(userModels) > pagination.PaginationLimit {
		userModels = userModels[1:]
		prev = &userModels[0]
	}

	// handle next page
	if len(userModels) > 0 {
		nextCursor := userModels[len(userModels)-1].GetCreatedAt()
		nextUser, nextUserErr := s.db.Queries().GetGlobalUsersOldestFirstGreaterThanCursor(
			ctx,
			db.GetGlobalUsersOldestFirstGreaterThanCursorParams{
				Limit:  1,
				Cursor: nextCursor,
			},
		)

		if nextUserErr == nil && len(nextUser) > 0 {
			next = &userModels[len(userModels)-1]
		}
	}

	return pagination.NewPaginationResponse(userModels, next, prev), nil
}
