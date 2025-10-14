package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type newUserSocialLinkData struct {
	PlatformName string `json:"platformName"`
	ProfileLink  string `json:"profileLink"`
}

func (socialLink newUserSocialLinkData) Validate() error {
	validationErr := validation.ValidateStruct(
		&socialLink,
		validation.Field(&socialLink.PlatformName, validation.Required),
		validation.Field(&socialLink.ProfileLink, validation.Required, is.URL),
	)
	if validationErr != nil {
		return &core.FieldValidationError{
			ValidationErrors: validationErr,
		}
	}

	return nil
}

func (s *userService) newUserSocialLink(ctx context.Context, userID uuid.UUID, socialLinkDetails newUserSocialLinkData) (*db.GlobalUserSocialLinks, error) {
	permissionErr := s.iam.CheckPermission(ctx, iam.NewPermissionRequiredFromSelfPermission(
		updateSelfProfilePermission,
		iam.PermissionRequiredResources{
			UserID: &userID,
		},
	))
	if permissionErr != nil {
		return nil, permissionErr
	}

	// create tx
	qtx, tx, txErr := s.db.WithTransaction(ctx)
	if txErr != nil {
		return nil, txErr
	}
	defer tx.Rollback(context.Background())

	newSocialLink, dbErr := qtx.Queries().NewUserSocialLink(ctx, db.NewUserSocialLinkParams{
		UserID:       userID,
		PlatformName: socialLinkDetails.PlatformName,
		ProfileLink:  socialLinkDetails.ProfileLink,
	})
	if dbErr != nil {
		var pgErr *pgconn.PgError
		if errors.As(dbErr, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.ForeignKeyViolation:
				return nil, &UserNotFoundError{}
			case pgerrcode.UniqueViolation:
				return nil, &UserSocialPlatformDetailsAlreadyExistsError{
					PlatformName: socialLinkDetails.PlatformName,
				}
			}
		}
		return nil, dbErr
	}

	commitErr := tx.Commit(ctx)
	if commitErr != nil {
		return nil, commitErr
	}

	return &newSocialLink, nil
}

func (m *userServiceMux) newUserSelfSocialLink(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()

	userID, userIDErr := actor.GetActorIdFromContext(reqCtx)
	if userIDErr != nil {
		payload.EncodeError(w, userIDErr)
		return
	}

	newSocialLinkDetails, payloadErr := payload.DecodeRequestBodyAndValidate[newUserSocialLinkData](w, r)
	if payloadErr != nil {
		payload.EncodeError(w, payloadErr)
		return
	}

	newSocialLink, socialLinkErr := m.service.newUserSocialLink(reqCtx, userID, newSocialLinkDetails)
	if socialLinkErr != nil {
		payload.EncodeError(w, socialLinkErr)
		return
	}

	payload.EncodeJSON(w, http.StatusCreated, newSocialLink)
}
