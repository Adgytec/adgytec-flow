package usermanagement

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
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

func (s *userManagementService) newUser(ctx context.Context, userData newUserData) error {
	permissionErr := s.iam.CheckPermission(ctx,
		iam.NewPermissionRequiredFromManagementPermission(
			newManagementUserPermission,
			iam.PermissionRequiredResources{},
		),
	)
	if permissionErr != nil {
		return permissionErr
	}

	qtx, tx, txErr := s.db.WithTransaction(ctx)
	if txErr != nil {
		return txErr
	}
	defer tx.Rollback(context.Background())

	rootUserID, rootUserCreateErr := s.userService.NewUser(ctx, userData.Email)
	if rootUserCreateErr != nil {
		return rootUserCreateErr
	}

	// add user to management
	dbErr := qtx.Queries().NewManagementUser(ctx, rootUserID)
	if dbErr != nil {
		return dbErr
	}

	return nil
}

func (m *serviceMux) newUser(w http.ResponseWriter, r *http.Request) {
	newUserDetails, payloadErr := payload.DecodeRequestBodyAndValidate[newUserData](w, r)
	if payloadErr != nil {
		payload.EncodeError(w, payloadErr)
		return
	}

	newUserErr := m.service.newUser(r.Context(), newUserDetails)
	if newUserErr != nil {
		payload.EncodeError(w, newUserErr)
		return
	}

	payload.EncodeJSON(w, http.StatusCreated, "user added successfully")
}
