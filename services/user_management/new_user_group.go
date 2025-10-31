package usermanagement

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/core"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type newUserGroupData struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (groupDetails newUserGroupData) Validate() error {
	validationErr := validation.ValidateStruct(&groupDetails,
		validation.Field(&groupDetails.Name, validation.Required),
	)

	if validationErr != nil {
		return &core.FieldValidationError{
			ValidationErrors: validationErr,
		}
	}

	return nil
}

func (s *userManagementService) newUserGroup(ctx context.Context, groupDetails newUserGroupData) (*db.NewUserGroupRow, error) {
	return nil, nil
}

func (m *serviceMux) newUserGroup(w http.ResponseWriter, r *http.Request) {}
