package user

import (
	"net/http"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
)

func (m *userServiceMux) enableGlobalUser(w http.ResponseWriter, r *http.Request) {
	m.updateUserStatusUtil(w, r, db_actions.GlobalUserStatusEnabled)
}
