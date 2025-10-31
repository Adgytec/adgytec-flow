package usermanagement

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/pagination"
	"github.com/Adgytec/adgytec-flow/utils/payload"
)

func (s *userManagementService) listUserGroups(ctx context.Context,
	params pagination.PaginationRequestParams,
) (*pagination.ResponsePagination[any], error) {
	return nil, nil
}

func (m *serviceMux) listUserGroups(w http.ResponseWriter, r *http.Request) {
	paginationParams := pagination.GetPaginationParamsFromRequest(r)
	groupList, groupErr := m.service.listUserGroups(r.Context(), paginationParams)
	if groupErr != nil {
		payload.EncodeError(w, groupErr)
		return
	}

	payload.EncodeJSON(w, http.StatusOK, groupList)

}
