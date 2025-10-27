package usermanagement

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/core"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type newUserData struct {
	Email string `json:"email"`
}

func (userData newUserData) Validate() error {
	validationErr := validation.ValidateStruct(&userData,
		validation.Field(&userData.Email, is.Email),
	)

	if validationErr != nil {
		return &core.FieldValidationError{
			ValidationErrors: validationErr,
		}
	}

	return nil
}

func (s *userManagementService) newUser(ctx context.Context) error {
	return nil
}

func (m *serviceMux) newUser(w http.ResponseWriter, r *http.Request) {}
