package usermanagement

import (
	"context"
	"errors"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	reqparams "github.com/Adgytec/adgytec-flow/utils/req_params"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *userManagementService) getUserGroupByID(ctx context.Context, groupID uuid.UUID) (*models.UserGroup, error) {
	permissionErr := s.iam.CheckPermission(ctx,
		iam.NewPermissionRequiredFromManagementPermission(getUserGroupPermission,
			iam.PermissionRequiredResources{},
		),
	)
	if permissionErr != nil {
		return nil, permissionErr
	}

	userGroup, groupErr := s.userGroupCache.Get(groupID.String(), func() (models.UserGroup, error) {
		var zero models.UserGroup

		group, dbErr := s.db.Queries().GetUserGroupByID(ctx, groupID)
		if dbErr != nil {
			if errors.Is(dbErr, pgx.ErrNoRows) {
				return zero, &UserGroupNotFoundError{}
			}
			return zero, dbErr
		}

		return getUserGroupResponseModel(group), nil
	})

	if groupErr != nil {
		return nil, groupErr
	}

	return &userGroup, nil
}

func (m *serviceMux) getUserGroupByID(w http.ResponseWriter, r *http.Request) {
	groupID, groupIDErr := reqparams.GetUserGroupIDFromRequest(r)
	if groupIDErr != nil {
		payload.EncodeError(w, groupIDErr)
		return
	}

	groupDetails, groupErr := m.service.getUserGroupByID(r.Context(), groupID)
	if groupErr != nil {
		payload.EncodeError(w, groupErr)
		return
	}

	payload.EncodeJSON(w, http.StatusOK, groupDetails)
}
