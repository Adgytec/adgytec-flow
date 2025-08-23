package helpers

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/core"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
	"github.com/google/uuid"
)

func GetActorDetailsFromContext(ctx context.Context) (core.ActorDetails, error) {
	var zero core.ActorDetails

	// empty actor id and actor type are also considered errors
	// and is part of ErrInvalidActorDetails
	actorID, actorIDOk := GetContextValue(ctx, ActorIDKey)
	if !actorIDOk {
		return zero, app_errors.ErrInvalidActorID
	}

	actorType, actorTypeOk := GetContextValue(ctx, ActorTypeKey)
	if !actorTypeOk {
		return zero, app_errors.ErrInvalidActorType
	}

	actorUUID, actorUUIDErr := uuid.Parse(actorID)
	if actorUUIDErr != nil {
		return zero, app_errors.ErrInvalidActorID
	}

	actorTypeValue := core.ActorType(actorType).Value()
	if actorTypeValue == core.ActorTypeUnknown {
		return zero, app_errors.ErrInvalidActorType
	}

	return core.ActorDetails{
		ID:   actorUUID,
		Type: actorTypeValue,
	}, nil
}

// as both the actor id and actor type are closely releated
// if any one of them causes any error than both the values are considered invalid
func GetActorIdFromContext(ctx context.Context) (uuid.UUID, error) {
	actorDetails, actorDetailsErr := GetActorDetailsFromContext(ctx)
	return actorDetails.ID, actorDetailsErr
}

func GetActorTypeFromContext(ctx context.Context) (core.ActorType, error) {
	actorDetails, actorDetailsErr := GetActorDetailsFromContext(ctx)
	return actorDetails.Type, actorDetailsErr
}
