package reqparams

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type idType string

const (
	idTypeUserID      idType = "User ID"
	idTypeUserGroupID idType = "User Group ID"
)

func getIDFromRequest(r *http.Request, key string, typeID idType) (uuid.UUID, error) {
	id := chi.URLParam(r, key)

	uuidVal, idErr := uuid.Parse(id)
	if idErr != nil {
		return uuid.Nil, &InvalidIDError{
			IDType:    typeID,
			InvalidID: id,
		}
	}

	return uuidVal, nil
}

func GetUserIDFromRequest(r *http.Request) (uuid.UUID, error) {
	return getIDFromRequest(r, "userID", idTypeUserID)
}

func GetUserGroupIDFromRequest(r *http.Request) (uuid.UUID, error) {
	return getIDFromRequest(r, "groupID", idTypeUserGroupID)
}
