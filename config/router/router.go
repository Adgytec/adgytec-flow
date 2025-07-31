package router

import (
	"net/http"
	"os"

	"github.com/Adgytec/adgytec-flow/config/app"
	"github.com/Adgytec/adgytec-flow/services/access_management"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	mux.Use(middleware.Heartbeat("/health"))
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.StripSlashes)
	mux.Use(middleware.AllowContentType("application/json"))

	allowedOrigins := []string{
		"https://*",
	}

	if os.Getenv("ENV") == "development" {
		allowedOrigins = append(allowedOrigins, "http://*")
	}

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	handle400(mux)

	for _, factory := range services {
		serviceMux := factory(appConfig)
		mux.Mount(serviceMux.BasePath(), serviceMux.Router())
	}

	return mux
}
