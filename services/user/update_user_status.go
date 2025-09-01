package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/core"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func (s *userService) updateUserStatus(ctx context.Context, userID string, status db_actions.GlobalUserStatus) error {
	requiredPermission := enableUserPermission
	if status == db_actions.GlobalUserStatusDisabled {
		requiredPermission = disableUserPermission
	}

	requiredPermissions := []core.IPermissionRequired{
		helpers.CreatePermissionRequiredFromManagementPermission(requiredPermission, core.PermissionRequiredResources{}),
	}
	permissionErr := s.accessManagement.CheckPermission(
		ctx,
		requiredPermissions,
	)
	if permissionErr != nil {
		return permissionErr
	}

	userUUID, userIdErr := s.getUserUUIDFromString(userID)
	if userIdErr != nil {
		return userIdErr
	}

	// start transaction
	tx, txErr := s.db.NewTransaction(ctx)
	if txErr != nil {
		return txErr
	}
	defer tx.Rollback(context.Background())
	qtx := s.db.Queries().WithTx(tx)

	_, dbErr := qtx.UpdateGlobalUserStatus(
		ctx,
		db_actions.UpdateGlobalUserStatusParams{
			ID:     userUUID,
			Status: status,
		},
	)
	if dbErr != nil {
		if errors.Is(dbErr, pgx.ErrNoRows) {
			return &app_errors.UserNotFoundError{}
		}
		return dbErr
	}

	// update cognito
	var authErr error
	if status == db_actions.GlobalUserStatusDisabled {
		authErr = s.auth.DisableUser(userID)
	} else {
		authErr = s.auth.EnableUser(userID)
	}
	if authErr != nil {
		return authErr
	}

	return tx.Commit(context.Background())
}

func (s *userService) updateUserStatusHandler(w http.ResponseWriter, r *http.Request, status db_actions.GlobalUserStatus) {
	if !status.Valid() {
		payload.EncodeError(w, fmt.Errorf("invalid-status-value"))
		return
	}

	reqCtx := r.Context()
	userID := chi.URLParam(r, "userID")
	enableErr := s.updateUserStatus(reqCtx, userID, status)
	if enableErr != nil {
		payload.EncodeError(w, enableErr)
	}

	payload.EncodeJSON(w, http.StatusOK, "user status updated successfully")
}
