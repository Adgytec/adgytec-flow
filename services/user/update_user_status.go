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
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *userService) updateUserStatus(ctx context.Context, currentUserId, userId string, status db_actions.GlobalUserStatus) error {
	requiredPermission := enableUserPermission
	if status == db_actions.GlobalUserStatusDisabled {
		requiredPermission = disableUserPermission
	}

	permissionErr := s.accessManagement.CheckPermission(
		ctx,
		core.CreatePermissionEntity(currentUserId, core.PermissionEntityTypeUser),
		core.CreatePermssionRequiredFromManagementPermission(requiredPermission, nil),
	)
	if permissionErr != nil {
		return permissionErr
	}

	userUUID, userIdErr := uuid.Parse(userId)
	if userIdErr != nil {
		return &app_errors.InvalidUserIdError{
			InvalidUserId: userId,
		}
	}

	// start transaction
	tx, txErr := s.db.NewTransaction(ctx, currentUserId)
	if txErr != nil {
		return txErr
	}
	defer tx.Rollback(ctx)
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
		authErr = s.auth.DisableUser(userId)
	} else {
		authErr = s.auth.EnableUser(userId)
	}
	if authErr != nil {
		return authErr
	}

	return tx.Commit(ctx)
}

func (s *userService) updateUserStatusHandler(w http.ResponseWriter, r *http.Request, status db_actions.GlobalUserStatus) {
	if !status.Valid() {
		payload.EncodeError(w, fmt.Errorf("invalid-status-value"))
		return
	}

	reqCtx := r.Context()
	currentUserID, userIDOk := helpers.GetContextValue(reqCtx, helpers.UserID)
	if !userIDOk {
		payload.EncodeError(w, fmt.Errorf("Can't find current user."))
		return
	}

	userID := chi.URLParam(r, "userID")
	enableErr := s.updateUserStatus(reqCtx, currentUserID, userID, status)
	if enableErr != nil {
		payload.EncodeError(w, enableErr)
	}

	payload.EncodeJSON(w, http.StatusOK, "user status updated successfully")
}
