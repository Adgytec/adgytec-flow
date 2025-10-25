package org

import (
	"context"
	"fmt"

	"github.com/Adgytec/adgytec-flow/services/media"
	"github.com/Adgytec/adgytec-flow/utils/core"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type newOrganizationRestrictionItem struct {
	ID    uuid.UUID `json:"id"`
	Info  *string   `json:"info"`
	Value int32     `json:"value"`
}

func (restrictionItem newOrganizationRestrictionItem) Validate() error {
	validationErr := validation.ValidateStruct(&restrictionItem,
		validation.Field(&restrictionItem.ID,
			validation.By(func(val any) error {
				id := val.(uuid.UUID)

				// nil uuid check is explictly required
				// validation.required doesn't do anything for uuid
				if id == uuid.Nil {
					return fmt.Errorf("id cannot be nil UUID")
				}
				return nil
			}),
		),
		validation.Field(&restrictionItem.Value, validation.Min(-1)),
	)

	if validationErr != nil {
		return validationErr
	}

	return nil
}

type newOrganizationData struct {
	Name            string                           `json:"name"`
	RootUser        string                           `json:"rootUser"`
	Description     *string                          `json:"description"`
	Logo            *media.NewMediaItemInfo          `json:"logo"`
	CoverPhoto      *media.NewMediaItemInfo          `json:"coverPhoto"`
	RestrictionInfo []newOrganizationRestrictionItem `json:"restrictionInfo"`
}

func (orgDetails newOrganizationData) Validate() error {
	validationErr := validation.ValidateStruct(&orgDetails,
		validation.Field(&orgDetails.Name, validation.Required),
		validation.Field(&orgDetails.RootUser, is.Email),
		validation.Field(orgDetails.Logo),
		validation.Field(orgDetails.CoverPhoto),
		validation.Field(orgDetails.RestrictionInfo),
	)

	if validationErr != nil {
		return &core.FieldValidationError{
			ValidationErrors: validationErr,
		}
	}

	return nil
}

func (s *orgService) newOrganization(ctx context.Context) error {
	return nil
}
