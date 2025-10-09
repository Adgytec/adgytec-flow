package auth

import (
	"context"
	"net/url"

	"github.com/Adgytec/adgytec-flow/utils/actor"
)

func (a *authCommon) validateSignedURL(signedURL url.URL, baseQuery map[string]string) error {
	return nil
}

func (a *authCommon) ValidateSignedURL(signedURL url.URL) error {
	return a.validateSignedURL(signedURL, nil)
}

func (a *authCommon) ValidateSignedURLWithActor(ctx context.Context, signedURL url.URL) error {
	actorID, actorErr := actor.GetActorIdFromContext(ctx)
	if actorErr != nil {
		return actorErr
	}

	// create query and add actor details to it
	query := make(map[string]string)
	query[queryKeyActor] = actorID.String()

	return a.validateSignedURL(signedURL, query)
}
