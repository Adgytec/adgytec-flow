package user

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/db"
)

func (m *mux) disableGlobalUser(w http.ResponseWriter, r *http.Request) {
	m.updateUserStatusUtil(w, r, db.GlobalUserStatusDisabled)
}
