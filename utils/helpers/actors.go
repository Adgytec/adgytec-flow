package helpers

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/core"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
	"github.com/google/uuid"
)

func GetActorDetailsFromContext(ctx context.Context) (core.ActorDetials, error) {
	var zero core.ActorDetials

	actorID, actorIDOk := GetContextValue(ctx, ActorIDKey)
	actorType, actorTypeOk := GetContextValue(ctx, ActorTypeKey)
	if actorIDOk && !actorTypeOk {
		return zero, app_errors.ErrInvalidActorDetails
	}

	actorUUID, actorUUIDErr := uuid.Parse(actorID)
	if actorUUIDErr != nil {
		return zero, app_errors.ErrInvalidActorDetails
	}

	actorTypeValue := core.ActorType(actorType).Value()
	if actorTypeValue == core.ActorTypeUnknown {
		return zero, app_errors.ErrInvalidActorDetails
	}

	return core.ActorDetials{
		ID:   actorUUID,
		Type: actorTypeValue,
	}, nil
}
