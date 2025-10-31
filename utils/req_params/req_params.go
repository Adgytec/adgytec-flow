package reqparams

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetUserIDFromRequest(r *http.Request) (uuid.UUID, error) {
	userID := chi.URLParam(r, "userID")

	userUUID, userIDErr := uuid.Parse(userID)
	if userIDErr != nil {
		return uuid.Nil, &InvalidUserIDError{
			InvalidUserID: userID,
		}
	}

	return userUUID, nil
}

func GetUserGroupIDFromRequest(r *http.Request) (uuid.UUID, error) {
	groupID := chi.URLParam(r, "groupID")

	groupUUID, groupIDErr := uuid.Parse(groupID)
	if groupIDErr != nil {
		return uuid.Nil, &InvalidUserGroupIDError{
			InvalidUserGroupID: groupID,
		}
	}

	return groupUUID, nil
}
