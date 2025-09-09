package core

import "net/http"

// MiddlewarePC provides multiple middleware funcs to use with individual service mux
// Its the responsibility of individual service to correctly call the respective middleware
// As API key can be generated for both a single organization and for management purposes
// need multiple handler for api keys to correctly verify for the required action
type MiddlewarePC interface {
	// verify and adds actor details to request context
	ValidateAndGetActorDetailsFromHttpRequest(http.Handler) http.Handler

	// ensures actor type user can only access the next routes
	EnsureActorTypeUserOnly(http.Handler) http.Handler

	// check auth status of the actor type user
	ValidateActorTypeUserGlobalStatus(http.Handler) http.Handler

	// validates organization current status and check if request method is allowed
	CheckOrganizationStatusAndRequestMethod(http.Handler) http.Handler

	// check if current actor is part of management access
	ActorManagementAccessCheck(http.Handler) http.Handler

	// this checks whether actor is part of the organization whose resources its working on
	// these methods also checks whether the actor status for that organization, but its only used for individual organiztion related methods
	ActorOrganizationAccessCheck(http.Handler) http.Handler
	ActorOrganizationManagementAccessCheck(http.Handler) http.Handler
}
