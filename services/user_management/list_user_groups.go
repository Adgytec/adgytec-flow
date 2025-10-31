package usermanagement

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/pagination"
)

func (s *userManagementService) listUserGroups(ctx context.Context,
	params pagination.PaginationRequestParams,
) (*pagination.ResponsePagination[any], error) {
	return nil, nil
}

func (m *serviceMux) listUserGroups(w http.ResponseWriter, r *http.Request) {}
