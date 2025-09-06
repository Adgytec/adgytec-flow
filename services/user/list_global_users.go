package user

import (
	"context"
	"net/http"
	"time"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/pagination"
	"github.com/Adgytec/adgytec-flow/utils/payload"
)

func (s *userService) getGlobalUsers(
	ctx context.Context,
	params pagination.PaginationRequestParams,
) (*pagination.ResponsePagination[models.GlobalUser], error) {
	permissionErr := s.iam.CheckPermission(
		ctx,
		iam.NewPermissionRequiredFromManagementPermission(
			listAllUsersPermission,
			iam.PermissionRequiredResources{},
		),
	)
	if permissionErr != nil {
		return nil, permissionErr
	}

	return pagination.GetPaginatedData(
		ctx,
		params,
		pagination.PaginationActions[
			db.GlobalUserDetail,
			models.GlobalUser,
		]{
			Cache:   s.getUserListCache,
			ToModel: s.getUserResponseModels,
			Query: func(
				ctx context.Context,
				searchQuery string,
				limit int32,
			) ([]db.GlobalUserDetail, error) {
				return s.db.Queries().GetGlobalUsersByQuery(
					ctx,
					db.GetGlobalUsersByQueryParams{
						Limit: limit,
						Query: searchQuery,
					},
				)
			},
			InitialLatestFirst: func(
				ctx context.Context,
				limit int32,
			) ([]db.GlobalUserDetail, error) {
				return s.db.Queries().GetGlobalUsersLatestFirst(
					ctx,
					limit,
				)
			},
			InitialOldestFirst: func(
				ctx context.Context,
				limit int32,
			) ([]db.GlobalUserDetail, error) {
				return s.db.Queries().GetGlobalUsersOldestFirst(
					ctx,
					limit,
				)
			},
			GreaterThanCursorLatestFirst: func(
				ctx context.Context,
				cursor time.Time,
				limit int32,
			) ([]db.GlobalUserDetail, error) {
				return s.db.Queries().GetGlobalUsersLatestFirstGreaterThanCursor(
					ctx,
					db.GetGlobalUsersLatestFirstGreaterThanCursorParams{
						Cursor: cursor,
						Limit:  limit,
					},
				)
			},
			GreaterThanCursorOldestFirst: func(
				ctx context.Context,
				cursor time.Time,
				limit int32,
			) ([]db.GlobalUserDetail, error) {
				return s.db.Queries().GetGlobalUsersOldestFirstGreaterThanCursor(
					ctx,
					db.GetGlobalUsersOldestFirstGreaterThanCursorParams{
						Cursor: cursor,
						Limit:  limit,
					},
				)
			},
			LesserThanCursorLatestFirst: func(
				ctx context.Context,
				cursor time.Time,
				limit int32,
			) ([]db.GlobalUserDetail, error) {
				return s.db.Queries().GetGlobalUsersLatestFirstLesserThanCursor(
					ctx,
					db.GetGlobalUsersLatestFirstLesserThanCursorParams{
						Cursor: cursor,
						Limit:  limit,
					},
				)
			},
			LesserThanCursorOldestFirst: func(
				ctx context.Context,
				cursor time.Time,
				limit int32,
			) ([]db.GlobalUserDetail, error) {
				return s.db.Queries().GetGlobalUsersOldestFirstLesserThanCursor(
					ctx,
					db.GetGlobalUsersOldestFirstLesserThanCursorParams{
						Cursor: cursor,
						Limit:  limit,
					},
				)
			},
		},
	)
}

func (m *userServiceMux) getGlobalUsers(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()

	paginationParams := pagination.GetPaginationParamsFromRequest(r)
	userList, userErr := m.service.getGlobalUsers(reqCtx, paginationParams)
	if userErr != nil {
		payload.EncodeError(w, userErr)
		return
	}

	payload.EncodeJSON(w, http.StatusOK, userList)
}
