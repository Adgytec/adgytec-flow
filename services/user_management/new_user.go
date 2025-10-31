package usermanagement

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type newUserData struct {
	Email string `json:"email"`
}

func (userData newUserData) Validate() error {
	validationErr := validation.ValidateStruct(&userData,
		validation.Field(&userData.Email, validation.Required, is.Email),
	)

	if validationErr != nil {
		return &core.FieldValidationError{
			ValidationErrors: validationErr,
		}
	}

	return nil
}

func (s *userManagementService) newUser(ctx context.Context, userData newUserData) (*uuid.UUID, error) {
	permissionErr := s.iam.CheckPermission(ctx,
		iam.NewPermissionRequiredFromManagementPermission(
			newManagementUserPermission,
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

	newUserID, userCreateErr := s.userService.NewUser(ctx, userData.Email)
	if userCreateErr != nil {
		return nil, userCreateErr
	}

	// add user to management
	dbErr := qtx.Queries().NewManagementUser(ctx, newUserID)
	if dbErr != nil {
		return nil, dbErr
	}

	commitErr := tx.Commit(ctx)
	if commitErr != nil {
		return nil, commitErr
	}

	return &newUserID, nil
}

func (m *serviceMux) newUser(w http.ResponseWriter, r *http.Request) {
	newUserDetails, payloadErr := payload.DecodeRequestBodyAndValidate[newUserData](w, r)
	if payloadErr != nil {
		payload.EncodeError(w, payloadErr)
		return
	}

	userID, newUserErr := m.service.newUser(r.Context(), newUserDetails)
	if newUserErr != nil {
		payload.EncodeError(w, newUserErr)
		return
	}

	payload.EncodeJSON(w, http.StatusCreated, map[string]any{
		"id":      userID.String(),
		"message": "user added successfully",
	})
}
