package user

import (
	"net/http"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
)

func (m *userServiceMux) disableGlobalUser(w http.ResponseWriter, r *http.Request) {
	m.service.updateUserStatusHandler(w, r, db_actions.GlobalUserStatusDisabled)
}
