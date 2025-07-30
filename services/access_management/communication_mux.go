package access_management

import (
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/go-chi/chi/v5"
)

type accessManagementMux struct {
	db core.IDatabase
}

func (m *accessManagementMux) BasePath() string {
	return ""
}

func (m *accessManagementMux) Router() *chi.Mux {
	mux := chi.NewMux()
	return mux
}

type iAccessManagementMuxParams interface {
	Database() core.IDatabase
}

func CreateAccessManagementMux(params iAccessManagementMuxParams) core.IServiceMux {
	return &accessManagementMux{
		db: params.Database(),
	}
}
