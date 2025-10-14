package user

import (
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/services"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type userServiceMux struct {
	service    *userService
	middleware core.MiddlewarePC
}

func (m *userServiceMux) BasePath() string {
	return "/user"
}

func (m *userServiceMux) Router() *chi.Mux {
	mux := chi.NewMux()

	mux.Use(m.middleware.ValidateAndGetActorDetailsFromHttpRequest)
	mux.Use(m.middleware.ValidateActorTypeUserGlobalStatus)

	mux.Group(func(router chi.Router) {
		router.Use(m.middleware.EnsureActorTypeUserOnly)

		router.Get("/profile", m.getUserSelfProfileHandler)

		router.Patch("/profile/update", m.updateSelfProfile)

		router.Post("/profile/add-social-link", m.newUserSelfSocialLink)

		router.Delete("/profile/social-link/{socialLinkID}", m.removeUserSelfSocialLink)

		router.Patch("/profile/social-link/{socialLinkID}", m.updateUserSelfSocialLink)
	})

	mux.Group(func(router chi.Router) {
		router.Use(m.middleware.ActorManagementAccessCheck)

		router.Get("/list", m.getGlobalUsers)
		router.Get("/{userID}", m.getUserProfileHandler)

		router.Patch("/{userID}/enable", m.enableGlobalUser)
		router.Patch("/{userID}/disable", m.disableGlobalUser)

		router.Patch("/{userID}/update", m.updateUserProfile)

		router.Delete("/{userID}/social-link/{socialLinkID}", m.removeUserSocialLink)
	})

	return mux
}

func NewUserServiceMux(params userServiceMuxParams) services.Mux {
	log.Info().
		Str("service", serviceName).
		Msg("new service mux")
	return &userServiceMux{
		service:    newUserService(params),
		middleware: params.Middleware(),
	}
}
