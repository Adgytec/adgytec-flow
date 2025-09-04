package user

import (
	"context"
	"net/http"

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

	switch {
	case params.SearchQuery != "":
		return s.getGlobalUsersByQuery(ctx, params)
	case params.NextCursor != "":
		return s.getGlobalUsersNextPage(ctx, params)
	case params.PrevCursor != "":
		return s.getGlobalUsersPrevPage(ctx, params)
	}

	return s.getGlobalUsersInitial(ctx, params)
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
