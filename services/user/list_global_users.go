package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
	"github.com/Adgytec/adgytec-flow/utils/payload"
)

func (s *userService) getGlobalUsers(
	ctx context.Context,
	userId string,
	params core.PaginationRequestParams,
) (*core.ResponsePagination[models.GlobalUser], error) {
	permissionErr := s.accessManagement.CheckPermission(
		ctx,
		core.CreatePermissionEntity(userId, core.PermissionEntityTypeUser),
		core.CreatePermssionRequiredFromManagementPermission(listAllUsersPermission, nil),
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
	userId, userIdOk := helpers.GetContextValue(reqCtx, helpers.ActorIDKey)
	if !userIdOk {
		payload.EncodeError(w, fmt.Errorf("Can't find current user."))
		return
	}

	paginationParams := helpers.GetPaginationParamsFromRequest(r)
	userList, userErr := m.service.getGlobalUsers(reqCtx, userId, paginationParams)
	if userErr != nil {
		payload.EncodeError(w, userErr)
		return
	}

	payload.EncodeJSON(w, http.StatusOK, userList)

}
