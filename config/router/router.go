package router

import (
	"log"
	"net/http"
	"os"

	"github.com/Adgytec/adgytec-flow/config/app"
	"github.com/Adgytec/adgytec-flow/services/access_management"
	"github.com/Adgytec/adgytec-flow/services/user"
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
		return access_management.NewAccessManagementMux(appConfig)
	},
	func(appConfig app.IApp) core.IServiceMux {
		return user.NewUserServiceMux(appConfig)
	},
}

func handle400(mux *chi.Mux) {
	mux.NotFound(func(w http.ResponseWriter, _ *http.Request) {
		payload.EncodeJSON(w, http.StatusNotFound, core.ResponseHTTPError{
			Message: helpers.ValuePtr(
				http.StatusText(http.StatusNotFound),
			),
		})
	})

	mux.MethodNotAllowed(func(w http.ResponseWriter, _ *http.Request) {
		payload.EncodeJSON(w, http.StatusMethodNotAllowed, core.ResponseHTTPError{
			Message: helpers.ValuePtr(
				http.StatusText(http.StatusMethodNotAllowed),
			),
		})
	})
}

func CreateApplicationRouter(appConfig app.IApp) *chi.Mux {
	log.Println("adding application mux")
	mux := chi.NewMux()

	mux.Use(middleware.Heartbeat("/health"))
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.StripSlashes)
	mux.Use(middleware.AllowContentType("application/json"))
	mux.Use(middleware.Compress(5, "application/json"))

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

	mux.Use(appConfig.Middleware().ValidateAndGetActorDetailsFromHttpRequest)
	mux.Use(appConfig.Middleware().ValidateActorTypeUserGlobalStatus)
	for _, factory := range services {
		serviceMux := factory(appConfig)
		mux.Mount(serviceMux.BasePath(), serviceMux.Router())
	}

	return mux
}
