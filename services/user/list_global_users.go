package user

import (
	"context"
	"net/http"

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
			Cache:                        s.getUserListCache,
			ToModel:                      s.getUserResponseModels,
			Query:                        s.getGlobalUsersQuery,
			InitialLatestFirst:           s.getGlobalUsersInitialLatestFirst,
			InitialOldestFirst:           s.getGlobalUserInitialOldestFirst,
			GreaterThanCursorLatestFirst: s.getGlobalUsersGreaterThanCursorLatestFirst,
			GreaterThanCursorOldestFirst: s.getGlobalUsersGreaterThanCursorOldestFirst,
			LesserThanCursorLatestFirst:  s.getGlobalUsersLesserThanCursorLatestFirst,
			LesserThanCursorOldestFirst:  s.getGlobalUsersLesserThanCursorOldestFirst,
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
