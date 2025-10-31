package usermanagement

import (
	"context"
	"errors"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	reqparams "github.com/Adgytec/adgytec-flow/utils/req_params"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (s *userManagementService) newUserGroupUser(ctx context.Context, groupID uuid.UUID, userData newUserData) (*uuid.UUID, error) {
	permissionErr := s.iam.CheckPermission(ctx,
		iam.NewPermissionRequiredFromManagementPermission(
			addUserInUserGroupPermission,
			iam.PermissionRequiredResources{},
		),
	)
	if permissionErr != nil {
		return nil, permissionErr
	}

	qtx, tx, txErr := s.db.WithTransaction(ctx)
	if txErr != nil {
		return nil, txErr
	}
	defer tx.Rollback(context.Background())

	userID := core.GetUserIDFromUsername(userData.Email)

	dbErr := qtx.Queries().NewUserGroupUser(ctx,
		db.NewUserGroupUserParams{
			UserGroupID: groupID,
			UserID:      userID,
		},
	)
	if dbErr != nil {
		var pgErr *pgconn.PgError
		if errors.As(dbErr, &pgErr) {
			if pgErr.Code == pgerrcode.ForeignKeyViolation {
				switch pgErr.ColumnName {
				case "user_group_id":
					return nil, &UserGroupNotFoundError{}
				case "user_id":
					return nil, &UserNotExistsInManagementError{}
				}
			}
		}
		return nil, dbErr
	}

	commitErr := tx.Commit(ctx)
	if commitErr != nil {
		return nil, commitErr
	}

	return &userID, nil
}

func (m *serviceMux) newUserGroupUser(w http.ResponseWriter, r *http.Request) {
	groupID, groupIDErr := reqparams.GetUserGroupIDFromRequest(r)
	if groupIDErr != nil {
		payload.EncodeError(w, groupIDErr)
		return
	}

	newUserDetails, payloadErr := payload.DecodeRequestBodyAndValidate[newUserData](w, r)
	if payloadErr != nil {
		payload.EncodeError(w, payloadErr)
		return
	}

	userID, newUserErr := m.service.newUserGroupUser(r.Context(), groupID, newUserDetails)
	if newUserErr != nil {
		payload.EncodeError(w, newUserErr)
		return
	}

	payload.EncodeJSON(w, http.StatusCreated, map[string]any{
		"groupID": groupID,
		"userID":  userID,
		"message": "user added successfully",
	})
}
