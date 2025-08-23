package access_management

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/go-chi/chi/v5"
)

type accessManagementMux struct {
	service    *accessManagement
	middleware core.IMiddlewarePC
}

func (m *accessManagementMux) BasePath() string {
	return "/access-management"
}

func (m *accessManagementMux) Router() *chi.Mux {
	mux := chi.NewMux()
	return mux
}

func CreateAccessManagementMux(params iAccessManagementMuxParams) core.IServiceMux {
	log.Println("adding access-managment mux")
	return &accessManagementMux{
		service:    createAccessManagementService(params),
		middleware: params.Middleware(),
	}
}
