package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *userService) updateUserStatus(ctx context.Context, userID uuid.UUID, status db.GlobalUserStatus) error {
	requiredPermission := enableUserPermission
	if status == db.GlobalUserStatusDisabled {
		requiredPermission = disableUserPermission
	}

	permissionErr := s.iam.CheckPermission(
		ctx,
		iam.NewPermissionRequiredFromManagementPermission(
			requiredPermission,
			iam.PermissionRequiredResources{},
		),
	)
	if permissionErr != nil {
		return permissionErr
	}

	// start transaction
	tx, txErr := s.db.NewTransaction(ctx)
	if txErr != nil {
		return txErr
	}
	defer tx.Rollback(context.Background())
	qtx := s.db.Queries().WithTx(tx)

	userData, dbErr := qtx.UpdateGlobalUserStatus(
		ctx,
		db.UpdateGlobalUserStatusParams{
			ID:     userID,
			Status: status,
		},
	)
	if dbErr != nil {
		if errors.Is(dbErr, pgx.ErrNoRows) {
			return &UserNotFoundError{}
		}
		return dbErr
	}

	// update cognito
	var authErr error
	if status == db.GlobalUserStatusDisabled {
		authErr = s.auth.DisableUser(userData.Username)
	} else {
		authErr = s.auth.EnableUser(userData.Username)
	}
	if authErr != nil {
		return authErr
	}

	return tx.Commit(context.Background())
}

func (m *mux) updateUserStatusUtil(w http.ResponseWriter, r *http.Request, status db.GlobalUserStatus) {
	if !status.Valid() {
		payload.EncodeError(w, fmt.Errorf("invalid-status-value"))
		return
	}

	reqCtx := r.Context()
	userID := chi.URLParam(r, "userID")

	userUUID, userIdErr := m.service.getUserUUIDFromString(userID)
	if userIdErr != nil {
		payload.EncodeError(w, userIdErr)
		return
	}

	statusErr := m.service.updateUserStatus(reqCtx, userUUID, status)
	if statusErr != nil {
		payload.EncodeError(w, statusErr)
		return
	}

	payload.EncodeJSON(w, http.StatusOK, "user status updated successfully")
}
