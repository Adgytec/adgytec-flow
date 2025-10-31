package usermanagement

import (
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/database/models"
)

func getUserGroupResponseModel(group db.ManagementUserGroupDetails) models.UserGroup {
	return models.UserGroup{
		ID:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		CreatedAt:   group.CreatedAt,
		UserCount:   group.UserCount,
	}
}

func getUserGroupResponseModels(groups []db.ManagementUserGroupDetails) []models.UserGroup {
	userGroupModels := make([]models.UserGroup, 0, len(groups))
	for _, group := range groups {
		userGroupModels = append(userGroupModels, getUserGroupResponseModel(group))
	}
	return userGroupModels
}
