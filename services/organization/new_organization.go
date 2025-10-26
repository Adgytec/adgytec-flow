package org

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/services/media"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type newOrganizationResponse struct {
	NextStep           string                     `json:"nextStep"`
	MediaUploadDetails []media.MediaUploadDetails `json:"mediaUploadDetails,omitempty"`
}

func (res *newOrganizationResponse) AddNextStep() {
	if res.NextStep == "" {
		res.NextStep = "Done"
	}
}

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

func (restrictionItem newOrganizationRestrictionItem) getDbRestrictionParams(orgID uuid.UUID) db.AddOrganizationRestrictionsParams {
	return db.AddOrganizationRestrictionsParams{
		OrgID:         orgID,
		RestrictionID: restrictionItem.ID,
		Value:         restrictionItem.Value,
		Info:          restrictionItem.Info,
	}
}

type newOrganizationData struct {
	Name            string                           `json:"name"`
	RootUser        string                           `json:"rootUser"`
	Description     *string                          `json:"description"`
	Logo            *media.NewMediaItemInfo          `json:"logo"`
	CoverMedia      *media.NewMediaItemInfo          `json:"coverMedia"`
	RestrictionInfo []newOrganizationRestrictionItem `json:"restrictionInfo"`
}

func (orgDetails newOrganizationData) Validate() error {
	validationErr := validation.ValidateStruct(&orgDetails,
		validation.Field(&orgDetails.Name, validation.Required),
		validation.Field(&orgDetails.RootUser, is.Email),
		validation.Field(orgDetails.Logo),
		validation.Field(orgDetails.CoverMedia),
		validation.Field(orgDetails.RestrictionInfo),
	)

	if validationErr != nil {
		return &core.FieldValidationError{
			ValidationErrors: validationErr,
		}
	}

	return nil
}

func (orgDetails newOrganizationData) providedRestrictionsMap() map[uuid.UUID]struct{} {
	provided := make(map[uuid.UUID]struct{}, len(orgDetails.RestrictionInfo))
	for _, r := range orgDetails.RestrictionInfo {
		provided[r.ID] = struct{}{}
	}

	return provided
}

func (orgDetails newOrganizationData) compareRestrictions(coreServiceRestrictions []db.AddServiceRestrictionIntoStagingParams) error {
	provided := orgDetails.providedRestrictionsMap()

	var missing []db.AddServiceRestrictionIntoStagingParams
	for _, required := range coreServiceRestrictions {
		if _, ok := provided[required.ID]; !ok {
			missing = append(missing, required)
		}
	}

	if len(missing) > 0 {
		return &MissingRequiredCoreServicesRestrictionsError{
			missingServicesRestrictions: missing,
		}
	}
	return nil
}

func (orgDetails newOrganizationData) getOrgRestrictions(orgID uuid.UUID) []db.AddOrganizationRestrictionsParams {
	if len(orgDetails.RestrictionInfo) == 0 {
		return nil
	}

	restrictions := make([]db.AddOrganizationRestrictionsParams, 0, len(orgDetails.RestrictionInfo))
	for _, restrictionItem := range orgDetails.RestrictionInfo {
		restrictions = append(restrictions, restrictionItem.getDbRestrictionParams(orgID))
	}

	return restrictions
}

func (s *orgService) newOrganization(ctx context.Context, orgDetails newOrganizationData) (*newOrganizationResponse, error) {
	permissionErr := s.iam.CheckPermission(ctx,
		iam.NewPermissionRequiredFromManagementPermission(
			createOrganizationPermission,
			iam.PermissionRequiredResources{},
		),
	)
	if permissionErr != nil {
		return nil, permissionErr
	}

	// check if all core service restrictions are provided
	restrictionMissingErr := orgDetails.compareRestrictions(s.serviceDetails.GetCoreServiceRestrictions())
	if restrictionMissingErr != nil {
		return nil, restrictionMissingErr
	}

	// start tx
	qtx, tx, txErr := s.db.WithTransaction(ctx)
	if txErr != nil {
		return nil, txErr
	}
	defer tx.Rollback(context.Background())

	// new org response
	newOrgRes := newOrganizationResponse{}

	// handle provided media info
	var mediaItemDetails []media.NewMediaItemInfoWithStorageDetails

	var logo *uuid.UUID
	var coverMedia *uuid.UUID

	if orgDetails.Logo != nil {
		logo = pointer.New(orgDetails.Logo.ID)
		mediaItemDetails = append(mediaItemDetails,
			media.NewMediaItemInfoWithStorageDetails{
				NewMediaItemInfo: *orgDetails.Logo,
				RequiredMime:     media.ImageMime,
				BucketPrefix:     path.Join("new-organization", "logo"), // for new organizations only
			},
		)
	}

	if orgDetails.CoverMedia != nil {
		coverMedia = pointer.New(orgDetails.CoverMedia.ID)
		mediaItemDetails = append(mediaItemDetails,
			media.NewMediaItemInfoWithStorageDetails{
				NewMediaItemInfo: *orgDetails.CoverMedia,
				RequiredMime:     media.VisualMime,
				BucketPrefix:     path.Join("new-organization", "cover-media"), // for new organizations only
			},
		)
	}

	if len(mediaItemDetails) > 0 {
		mediaService := s.media.WithTransaction(qtx)
		uploadDetails, mediaUploadErr := mediaService.NewMediaItems(ctx, mediaItemDetails)
		if mediaUploadErr != nil {
			return nil, mediaUploadErr
		}

		newOrgRes.MediaUploadDetails = uploadDetails
		newOrgRes.NextStep = "Upload media items"
	}

	// create root user
	// user creation is always independent of parent transaction
	rootUserID, rootUserCreateErr := s.userService.NewUser(ctx, orgDetails.RootUser)
	if rootUserCreateErr != nil {
		return nil, rootUserCreateErr
	}

	// create new organization
	orgID, dbErr := qtx.Queries().NewOrganization(ctx,
		db.NewOrganizationParams{
			RootUser:    rootUserID,
			Name:        orgDetails.Name,
			Description: orgDetails.Description,
			Logo:        logo,
			CoverMedia:  coverMedia,
		},
	)
	if dbErr != nil {
		return nil, dbErr
	}

	// add new org restrictions
	orgRestrictionsParams := orgDetails.getOrgRestrictions(orgID)
	if len(orgRestrictionsParams) > 0 {
		_, dbErr := qtx.Queries().AddOrganizationRestrictions(ctx, orgRestrictionsParams)
		if dbErr != nil {
			return nil, dbErr
		}
	}

	// tx commit
	commitErr := tx.Commit(ctx)
	if commitErr != nil {
		return nil, commitErr
	}

	newOrgRes.AddNextStep()
	return &newOrgRes, nil
}

func (s *orgServiceMux) newOrganization(w http.ResponseWriter, r *http.Request) {
	newOrgDetails, payloadErr := payload.DecodeRequestBodyAndValidate[newOrganizationData](w, r)
	if payloadErr != nil {
		payload.EncodeError(w, payloadErr)
		return
	}

	newOrgRes, newOrgErr := s.service.newOrganization(r.Context(), newOrgDetails)
	if newOrgErr != nil {
		payload.EncodeError(w, newOrgErr)
		return
	}

	payload.EncodeJSON(w, http.StatusCreated, newOrgRes)
}
