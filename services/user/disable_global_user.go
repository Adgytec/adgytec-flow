package user

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/db"
)

func (m *userServiceMux) disableGlobalUser(w http.ResponseWriter, r *http.Request) {
	m.updateUserStatusUtil(w, r, db.GlobalUserStatusDisabled)
}
