package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	reqparams "github.com/Adgytec/adgytec-flow/utils/req_params"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
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
	qtx, tx, txErr := s.db.WithTransaction(ctx)
	if txErr != nil {
		return txErr
	}
	defer tx.Rollback(context.Background())

	username, dbErr := qtx.Queries().UpdateGlobalUserStatus(
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

	// handle enable from auth provider here
	// db act as source of truth
	// it should be run in transaction as enable is required for user login in client application
	if status == db.GlobalUserStatusEnabled {
		authErr := s.auth.EnableUser(ctx, username)
		if authErr != nil {
			return authErr
		}
	}

	commitErr := tx.Commit(ctx)
	if commitErr != nil {
		return commitErr
	}

	// handle disable from auth provider here
	// db act as source of truth
	// disabling user from auth provider act as welcome addition to also prevent user login in client application
	if status == db.GlobalUserStatusDisabled {
		authErr := s.auth.DisableUser(ctx, username)
		if authErr != nil {
			log.Error().Err(authErr).Str("action", "auth provider disable user").Send()
		}
	}

	// update cache
	s.userStatusCache.Set(userID.String(), status)
	return nil
}

func (m *userServiceMux) updateUserStatusUtil(w http.ResponseWriter, r *http.Request, status db.GlobalUserStatus) {
	if !status.Valid() {
		payload.EncodeError(w, fmt.Errorf("invalid-status-value"))
		return
	}

	reqCtx := r.Context()

	userID, userIdErr := reqparams.GetUserIDFromRequest(r)
	if userIdErr != nil {
		payload.EncodeError(w, userIdErr)
		return
	}

	statusErr := m.service.updateUserStatus(reqCtx, userID, status)
	if statusErr != nil {
		payload.EncodeError(w, statusErr)
		return
	}

	payload.EncodeJSON(w, http.StatusOK, "user status updated successfully")
}
