package router

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/config/app"
	"github.com/Adgytec/adgytec-flow/services/access_management"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/go-chi/chi/v5"
)

type serviceFactory func(params app.IApp) core.IServiceMux

var services = []serviceFactory{
	func(appConfig app.IApp) core.IServiceMux {
		return access_management.CreateAccessManagementMux(appConfig)
	},
}

func handle400(mux *chi.Mux) {
	mux.NotFound(func(w http.ResponseWriter, _ *http.Request) {
		payload.EncodeJSON(w, http.StatusNotFound, core.ResponseHTTPError{
			Message: helpers.StringPtr(
				http.StatusText(http.StatusNotFound),
			),
		})
	})

	mux.MethodNotAllowed(func(w http.ResponseWriter, _ *http.Request) {
		payload.EncodeJSON(w, http.StatusMethodNotAllowed, core.ResponseHTTPError{
			Message: helpers.StringPtr(
				http.StatusText(http.StatusMethodNotAllowed),
			),
		})
	})
}

func CreateApplicationRouter(appConfig app.IApp) *chi.Mux {
	mux := chi.NewMux()
	return mux
}
