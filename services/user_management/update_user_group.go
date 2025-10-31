package usermanagement

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	reqparams "github.com/Adgytec/adgytec-flow/utils/req_params"
	"github.com/Adgytec/adgytec-flow/utils/types"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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
					if name.Missing() {
						return nil
					}

					if name.Null() {
						return fmt.Errorf("cannot be null")
					}

					if strings.TrimSpace(name.Value) == "" {
						return fmt.Errorf("cannot be blank")
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
	permissionErr := s.iam.CheckPermission(ctx,
		iam.NewPermissionRequiredFromManagementPermission(updateUserGroupPermission,
			iam.PermissionRequiredResources{},
		),
	)
	if permissionErr != nil {
		return nil, permissionErr
	}

	// start tx
	qtx, tx, txErr := s.db.WithTransaction(ctx)
	if txErr != nil {
		return nil, txErr
	}
	defer tx.Rollback(context.Background())

	// get existing group details
	existingGroup, existingErr := qtx.Queries().GetUserGroupByIDForUpdate(ctx, groupID)
	if existingErr != nil {
		if errors.Is(existingErr, pgx.ErrNoRows) {
			return nil, &UserGroupNotFoundError{}
		}
		return nil, existingErr
	}

	// update group obj
	updatedGroupParams := db.UpdateUserGroupParams{
		ID: groupID,
	}

	// name check
	updatedGroupParams.Name = existingGroup.Name
	if groupDetails.Name.Present() {
		updatedGroupParams.Name = groupDetails.Name.Value
	}

	// description check
	updatedGroupParams.Description = existingGroup.Description
	if groupDetails.Description.Present() {
		if groupDetails.Description.Null() {
			updatedGroupParams.Description = nil
		} else {
			updatedGroupParams.Description = &groupDetails.Description.Value
		}
	}

	// update group
	updatedGroup, dbErr := qtx.Queries().UpdateUserGroup(ctx, updatedGroupParams)
	if dbErr != nil {
		var pgErr *pgconn.PgError
		if errors.As(dbErr, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return nil, &UserGroupWithSameNameExistsError{}
		}
		return nil, dbErr
	}

	commitErr := tx.Commit(ctx)
	if commitErr != nil {
		return nil, commitErr
	}

	return &updatedGroup, nil
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
