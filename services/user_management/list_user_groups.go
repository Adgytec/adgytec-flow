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

func (s *userManagementService) listUserGroups(ctx context.Context,
	params pagination.PaginationRequestParams,
) (*pagination.ResponsePagination[models.UserGroup], error) {
	permissionErr := s.iam.CheckPermission(
		ctx,
		iam.NewPermissionRequiredFromManagementPermission(
			listUserGroupsPermission,
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
			db.ManagementUserGroupDetails,
			models.UserGroup,
		]{
			Cache:                        s.userGroupListCache,
			ToModel:                      getUserGroupResponseModels,
			Query:                        s.getUserGroupsQuery,
			InitialLatestFirst:           s.getUserGroupsInitialLatestFirst,
			InitialOldestFirst:           s.getUserGroupsInitialOldestFirst,
			GreaterThanCursorLatestFirst: s.getUserGroupsGreaterThanCursorLatestFirst,
			GreaterThanCursorOldestFirst: s.getUserGroupsGreaterThanCursorOldestFirst,
			LesserThanCursorLatestFirst:  s.getUserGroupsLesserThanCursorLatestFirst,
			LesserThanCursorOldestFirst:  s.getUserGroupsLesserThanCursorOldestFirst,
		},
	)
}

func (m *serviceMux) listUserGroups(w http.ResponseWriter, r *http.Request) {
	paginationParams := pagination.GetPaginationParamsFromRequest(r)
	groupList, groupErr := m.service.listUserGroups(r.Context(), paginationParams)
	if groupErr != nil {
		payload.EncodeError(w, groupErr)
		return
	}

	payload.EncodeJSON(w, http.StatusOK, groupList)

}
