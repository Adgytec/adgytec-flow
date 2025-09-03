package access_management

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/go-chi/chi/v5"
)

type accessManagementMux struct {
	service    *accessManagement
	middleware core.MiddlewarePC
}

func (m *accessManagementMux) BasePath() string {
	return "/access-management"
}

func (m *accessManagementMux) Router() *chi.Mux {
	mux := chi.NewMux()
	return mux
}

func NewAccessManagementMux(params accessManagementMuxParams) core.ServiceMux {
	log.Println("adding access-managment mux")
	return &accessManagementMux{
		service:    newAccessManagementService(params),
		middleware: params.Middleware(),
	}
}
