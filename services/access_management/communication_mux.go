package access_management

import (
	"github.com/Adgytec/adgytec-flow/utils/interfaces"
	"github.com/go-chi/chi/v5"
)

type accessManagementMux struct {
	db interfaces.IDatabase
}

func (m *accessManagementMux) BasePath() string {
	return ""
}

func (m *accessManagementMux) Router() *chi.Mux {
	mux := chi.NewMux()
	return mux
}

type iAccessManagementMuxParams interface {
	Database() interfaces.IDatabase
}

func CreateAccessManagementMux(params iAccessManagementMuxParams) interfaces.IServiceMux {
	return &accessManagementMux{
		db: params.Database(),
	}
}
