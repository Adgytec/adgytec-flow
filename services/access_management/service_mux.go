package access_management

import (
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/go-chi/chi/v5"
)

type accessManagementMux struct {
	service *accessManagement
}

func (m *accessManagementMux) BasePath() string {
	return ""
}

func (m *accessManagementMux) Router() *chi.Mux {
	mux := chi.NewMux()
	return mux
}

func CreateAccessManagementMux(params iAccessManagementParams) core.IServiceMux {
	return &accessManagementMux{
		service: &accessManagement{
			db: params.Database(),
		},
	}
}
