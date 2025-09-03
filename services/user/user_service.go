package user

import (
	"github.com/Adgytec/adgytec-flow/config/cache"
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/utils/core"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
	"github.com/google/uuid"
)

type userServiceParams interface {
	Database() core.Database
	Auth() core.Auth
	AccessManagement() core.AccessManagementPC
	CDN() core.CDN
	CacheClient() core.CacheClient
}

type userServiceMuxParams interface {
	userServiceParams
	Middleware() core.MiddlewarePC
}

type userService struct {
	db               core.Database
	auth             core.Auth
	accessManagement core.AccessManagementPC
	cdn              core.CDN
	getUserCache     core.Cache[models.GlobalUser]
	getUserListCache core.Cache[[]models.GlobalUser]
}

func (s *userService) getUserResponseModel(user db_actions.GlobalUserDetail) models.GlobalUser {
	userModel := models.GlobalUser{
		ID:          user.ID,
		Email:       user.Email,
		Name:        user.Name,
		About:       user.About,
		DateOfBirth: user.DateOfBirth,
		CreatedAt:   user.CreatedAt,
	}

	if user.ProfilePictureID != nil {
		profilePictureModel := &models.ImageQueryType{
			OriginalImage: s.cdn.GetSignedUrl(user.UncompressedProfilePicture),
			Size:          user.ProfilePictureSize,
			Status:        string(user.Status.GlobalMediaStatus),
			Thumbnail:     s.cdn.GetSignedUrl(user.Thumbnail),
			Small:         s.cdn.GetSignedUrl(user.Small),
			Medium:        s.cdn.GetSignedUrl(user.Medium),
			Large:         s.cdn.GetSignedUrl(user.Large),
			ExtraLarge:    s.cdn.GetSignedUrl(user.ExtraLarge),
		}
		userModel.ProfilePicture = profilePictureModel
	}

	return userModel
}

func (s *userService) getUserResponseModels(users []db_actions.GlobalUserDetail) []models.GlobalUser {
	usersLen := len(users)
	if usersLen == 0 {
		return nil
	}

	userModels := make([]models.GlobalUser, 0, usersLen)
	for _, user := range users {
		userModels = append(userModels, s.getUserResponseModel(user))
	}
	return userModels
}

func (s *userService) getUserUUIDFromString(userID string) (uuid.UUID, error) {
	userUUID, userIdErr := uuid.Parse(userID)
	if userIdErr != nil {
		return uuid.UUID{}, &app_errors.InvalidUserIDError{
			InvalidUserID: userID,
		}
	}

	return userUUID, nil
}

func newUserService(params userServiceParams) *userService {
	return &userService{
		db:               params.Database(),
		auth:             params.Auth(),
		accessManagement: params.AccessManagement(),
		cdn:              params.CDN(),
		getUserCache:     cache.NewCache[models.GlobalUser](params.CacheClient(), "user"),
		getUserListCache: cache.NewCache[[]models.GlobalUser](params.CacheClient(), "user-list"),
	}
}
