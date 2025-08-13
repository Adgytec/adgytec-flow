package user

import (
	"context"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

func (s *userService) getGlobalUsersByQuery(ctx context.Context, params core.PaginationRequestParams) (*core.ResponsePagination[models.GlobalUser], error) {
	userList, userErr := s.db.Queries().GetGlobalUsersByQuery(
		ctx,
		db_actions.GetGlobalUsersByQueryParams{
			Limit: helpers.SearchQueryLimit,
			Query: params.SearchQuery,
		},
	)

	if userErr != nil {
		return nil, userErr
	}

	userModels := s.getUserResponseModels(userList)
	return helpers.CreatePaginationResponse(userModels, nil, nil), nil
}

func (s *userService) getGlobalUsersInitial(ctx context.Context, params core.PaginationRequestParams) (*core.ResponsePagination[models.GlobalUser], error) {
	var userList []db_actions.GlobalUserDetail
	var userErr error

	if params.Sorting == core.PaginationRequestSortingLatestFirst {
		userList, userErr = s.db.Queries().GetGlobalUsersLatestFirst(ctx, helpers.PaginationLimit+1)
	} else {
		userList, userErr = s.db.Queries().GetGlobalUsersOldestFirst(ctx, helpers.PaginationLimit+1)
	}

	if userErr != nil {
		return nil, userErr
	}

	userModels := s.getUserResponseModels(userList)
	var next *models.GlobalUser

	// handle next page details
	if len(userList) == helpers.PaginationLimit+1 {
		userLen := len(userModels)
		next = &userModels[userLen-2]
		userModels = userModels[:userLen-1]
	}

	return helpers.CreatePaginationResponse(userModels, next, nil), nil
}

func (s *userService) getGlobalUsersNextPage(ctx context.Context, params core.PaginationRequestParams) (*core.ResponsePagination[models.GlobalUser], error) {
	return nil, nil
}

func (s *userService) getGlobalUsersPrevPage(ctx context.Context, params core.PaginationRequestParams) (*core.ResponsePagination[models.GlobalUser], error) {
	return nil, nil
}
