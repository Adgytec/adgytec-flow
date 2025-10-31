package usermanagement

import (
	"context"
	"errors"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
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
	permissionErr := s.iam.CheckPermission(ctx,
		iam.NewPermissionRequiredFromManagementPermission(newUserGroupPermission,
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

	newGroup, dbErr := qtx.Queries().NewUserGroup(ctx,
		db.NewUserGroupParams{
			Name:        groupDetails.Name,
			Description: groupDetails.Description,
		},
	)
	if dbErr != nil {
		var pgErr *pgconn.PgError
		if errors.As(dbErr, &pgErr) &&
			pgErr.Code == pgerrcode.UniqueViolation {
			return nil, &UserGroupWithSameNameExistsError{}
		}

		return nil, dbErr
	}

	commitErr := tx.Commit(ctx)
	if commitErr != nil {
		return nil, commitErr
	}

	return &newGroup, nil
}

func (m *serviceMux) newUserGroup(w http.ResponseWriter, r *http.Request) {
	newGroupDetails, payloadErr := payload.DecodeRequestBodyAndValidate[newUserGroupData](w, r)
	if payloadErr != nil {
		payload.EncodeError(w, payloadErr)
		return
	}

	newGroup, newGroupErr := m.service.newUserGroup(r.Context(), newGroupDetails)
	if newGroupErr != nil {
		payload.EncodeError(w, newGroupErr)
		return
	}

	payload.EncodeJSON(w, http.StatusCreated, newGroup)
}
