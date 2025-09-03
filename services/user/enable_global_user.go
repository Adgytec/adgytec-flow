package user

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/db"
)

func (m *mux) enableGlobalUser(w http.ResponseWriter, r *http.Request) {
	m.updateUserStatusUtil(w, r, db.GlobalUserStatusEnabled)
}
