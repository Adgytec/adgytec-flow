package usermanagement

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	reqparams "github.com/Adgytec/adgytec-flow/utils/req_params"
	"github.com/Adgytec/adgytec-flow/utils/types"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type updateUserGroupData struct {
	Name        types.NullableString `json:"name"`
	Description types.NullableString `json:"description"`
}

func (groupDetails updateUserGroupData) Validate() error {
	validationErr := validation.ValidateStruct(&groupDetails,
		validation.Field(
			&groupDetails.Name,
			validation.By(
				func(val any) error {
					name := val.(types.NullableString)
					if name.Null() {
						return nil
					}

					nameLen := utf8.RuneCountInString(strings.TrimSpace(name.Value))
					if nameLen < 1 {
						return fmt.Errorf("Missing required name value")
					}

					return nil
				},
			),
		),
	)

	if validationErr != nil {
		return &core.FieldValidationError{
			ValidationErrors: validationErr,
		}
	}

	return nil
}

func (s *userManagementService) updateUserGroup(ctx context.Context, groupID uuid.UUID, groupDetails updateUserGroupData) (*db.UpdateUserGroupRow, error) {
	return nil, nil
}

func (m *serviceMux) updateUserGroup(w http.ResponseWriter, r *http.Request) {
	groupID, groupIDErr := reqparams.GetUserGroupIDFromRequest(r)
	if groupIDErr != nil {
		payload.EncodeError(w, groupIDErr)
		return
	}

	groupDetails, payloadErr := payload.DecodeRequestBodyAndValidate[updateUserGroupData](w, r)
	if payloadErr != nil {
		payload.EncodeError(w, payloadErr)
		return
	}

	updatedGroup, updateErr := m.service.updateUserGroup(r.Context(), groupID, groupDetails)
	if updateErr != nil {
		payload.EncodeError(w, updateErr)
		return
	}

	payload.EncodeJSON(w, http.StatusOK, updatedGroup)
}
