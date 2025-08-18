package user

import (
	"github.com/Adgytec/adgytec-flow/config/cache"
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type iUserServiceParams interface {
	Database() core.IDatabase
	Auth() core.IAuth
	AccessManagement() core.IAccessManagementPC
	CDN() core.ICDN
	CacheClient() core.ICacheClient
}

type userService struct {
	db               core.IDatabase
	auth             core.IAuth
	accessManagement core.IAccessManagementPC
	cdn              core.ICDN
	getUserCache     core.ICache[models.GlobalUser]
	getUserListCache core.ICache[[]models.GlobalUser]
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

func createUserService(params iUserServiceParams) *userService {
	return &userService{
		db:               params.Database(),
		auth:             params.Auth(),
		accessManagement: params.AccessManagement(),
		cdn:              params.CDN(),
		getUserCache:     cache.CreateNewCache[models.GlobalUser](params.CacheClient(), "user"),
		getUserListCache: cache.CreateNewCache[[]models.GlobalUser](params.CacheClient(), "user-list"),
	}
}
