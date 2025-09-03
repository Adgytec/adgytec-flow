package user

import (
	"github.com/Adgytec/adgytec-flow/config/cache"
	"github.com/Adgytec/adgytec-flow/config/serializer"
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/google/uuid"
)

type userServiceParams interface {
	Database() core.Database
	Auth() core.Auth
	AccessManagement() core.AccessManagementPC
	CDN() core.CDN
	CacheClient() core.CacheClient
}

type muxParams interface {
	userServiceParams
	Middleware() core.MiddlewarePC
}

type userService struct {
	db               core.Database
	auth             core.Auth
	accessManagement core.AccessManagementPC
	cdn              core.CDN
	getUserCache     core.Cache[models.GlobalUser]
	getUserListCache core.Cache[core.ResponsePagination[models.GlobalUser]]
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
		return uuid.Nil, &InvalidUserIDError{
			InvalidUserID: userID,
		}
	}

	return userUUID, nil
}

func newService(params userServiceParams) *userService {
	return &userService{
		db:               params.Database(),
		auth:             params.Auth(),
		accessManagement: params.AccessManagement(),
		cdn:              params.CDN(),
		getUserCache:     cache.NewCache[models.GlobalUser](params.CacheClient(), serializer.NewGobSerializer[models.GlobalUser](), "user"),
		getUserListCache: cache.NewCache[core.ResponsePagination[models.GlobalUser]](params.CacheClient(), serializer.NewGobSerializer[core.ResponsePagination[models.GlobalUser]](), "user-list"),
	}
}
