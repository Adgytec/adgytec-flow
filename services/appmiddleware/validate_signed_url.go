package appmiddleware

import (
	"context"
	"net/http"
	"net/url"

	"github.com/Adgytec/adgytec-flow/utils/payload"
)

func (pc *appMiddlewarePC) ValidateSignedURL(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signedURLValidationErr := pc.service.validateSignedURL(r.URL)
		if signedURLValidationErr != nil {
			payload.EncodeError(w, signedURLValidationErr)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (pc *appMiddlewarePC) ValidateSignedURLWithActor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signedURLValidationErr := pc.service.validateSignedURLWithActor(r.Context(), r.URL)
		if signedURLValidationErr != nil {
			payload.EncodeError(w, signedURLValidationErr)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *appMiddleware) validateSignedURL(signedURL *url.URL) error {
	return s.auth.ValidateSignedURL(signedURL)
}

func (s *appMiddleware) validateSignedURLWithActor(ctx context.Context, signedURL *url.URL) error {
	return s.auth.ValidateSignedURLWithActor(ctx, signedURL)
}
