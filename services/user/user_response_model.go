package user

import (
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

func (s *userService) getUserResponseModel(user db.GlobalUserDetails) models.GlobalUser {
	userModel := models.GlobalUser{
		ID:          user.ID,
		Email:       user.Email,
		Name:        user.Name,
		About:       user.About,
		DateOfBirth: user.DateOfBirth,
		CreatedAt:   user.CreatedAt,
		Status:      user.Status,
	}

	if user.ProfilePictureID != nil {
		// all the media fields will always be present
		profilePictureModel := &models.ImageDetails{
			MediaID:       *user.ProfilePictureID,
			OriginalImage: s.cdn.GetSignedUrl(user.UncompressedProfilePicture),
			Size:          user.ProfilePictureSize,
			Status:        pointer.New(string(user.ProfilePictureStatus.GlobalMediaStatus)),
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

func (s *userService) getUserResponseModels(users []db.GlobalUserDetails) []models.GlobalUser {
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

func (pc *userServicePC) GetUserReponseModel(user db.GlobalUserDetails) models.GlobalUser {
	return pc.service.getUserResponseModel(user)
}

func (pc *userServicePC) GetUserResponseModels(users []db.GlobalUserDetails) []models.GlobalUser {
	return pc.service.getUserResponseModels(users)
}
