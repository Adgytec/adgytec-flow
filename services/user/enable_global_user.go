package user

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/actions"
)

func (m *mux) enableGlobalUser(w http.ResponseWriter, r *http.Request) {
	m.updateUserStatusUtil(w, r, actions.GlobalUserStatusEnabled)
}
