package user

import (
	"github.com/Adgytec/adgytec-flow/config/auth"
	"github.com/Adgytec/adgytec-flow/config/cache"
	"github.com/Adgytec/adgytec-flow/config/cdn"
	"github.com/Adgytec/adgytec-flow/config/database"
	"github.com/Adgytec/adgytec-flow/config/serializer"
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/pagination"
	"github.com/google/uuid"
)

type userServiceParams interface {
	Database() database.Database
	Auth() auth.Auth
	IAMService() iam.IAMServicePC
	CDN() cdn.CDN
	CacheClient() cache.CacheClient
}

type userServiceMuxParams interface {
	userServiceParams
	Middleware() core.MiddlewarePC
}

type userService struct {
	db               database.Database
	auth             auth.Auth
	iam              iam.IAMServicePC
	cdn              cdn.CDN
	getUserCache     cache.Cache[models.GlobalUser]
	getUserListCache cache.Cache[pagination.ResponsePagination[models.GlobalUser]]
}

func (s *userService) getUserResponseModel(user db.GlobalUserDetail) models.GlobalUser {
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

func (s *userService) getUserResponseModels(users []db.GlobalUserDetail) []models.GlobalUser {
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

func newUserService(params userServiceParams) *userService {
	return &userService{
		db:               params.Database(),
		auth:             params.Auth(),
		iam:              params.IAMService(),
		cdn:              params.CDN(),
		getUserCache:     cache.NewCache[models.GlobalUser](params.CacheClient(), serializer.NewGobSerializer[models.GlobalUser](), "user"),
		getUserListCache: cache.NewCache[pagination.ResponsePagination[models.GlobalUser]](params.CacheClient(), serializer.NewGobSerializer[pagination.ResponsePagination[models.GlobalUser]](), "user-list"),
	}
}
