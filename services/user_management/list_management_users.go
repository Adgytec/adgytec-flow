package usermanagement

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/pagination"
	"github.com/Adgytec/adgytec-flow/utils/payload"
)

func (s *userManagementService) getManagementUsers(
	ctx context.Context,
	params pagination.PaginationRequestParams,
) (*pagination.ResponsePagination[models.GlobalUser], error) {
	permissionErr := s.iam.CheckPermission(
		ctx,
		iam.NewPermissionRequiredFromManagementPermission(
			listManagementUsersPermission,
			iam.PermissionRequiredResources{},
		),
	)
	if permissionErr != nil {
		return nil, permissionErr
	}

	return pagination.GetPaginatedData(
		ctx,
		params,
		&pagination.PaginationActions[
			db.GlobalUserDetails,
			models.GlobalUser,
		]{
			Cache:                        s.getUserListCache,
			ToModel:                      s.userService.GetUserResponseModels,
			Query:                        s.getManagementUsersQuery,
			InitialLatestFirst:           s.getManagementUsersInitialLatestFirst,
			InitialOldestFirst:           s.getManagementUsersInitialOldestFirst,
			GreaterThanCursorLatestFirst: s.getManagementUsersGreaterThanCursorLatestFirst,
			GreaterThanCursorOldestFirst: s.getManagementUsersGreaterThanCursorOldestFirst,
			LesserThanCursorLatestFirst:  s.getManagementUsersLesserThanCursorLatestFirst,
			LesserThanCursorOldestFirst:  s.getManagementUsersLesserThanCursorOldestFirst,
		},
	)
}

func (m *serviceMux) getManagementUsers(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()

	paginationParams := pagination.GetPaginationParamsFromRequestNormalizeQuery(r)
	userList, userErr := m.service.getManagementUsers(reqCtx, paginationParams)
	if userErr != nil {
		payload.EncodeError(w, userErr)
		return
	}

	payload.EncodeJSON(w, http.StatusOK, userList)
}
