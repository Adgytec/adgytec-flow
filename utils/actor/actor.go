package actor

import (
	"context"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/google/uuid"
)

func GetActorDetailsFromContext(ctx context.Context) (ActorDetails, error) {
	var zero ActorDetails

	// empty actor id and actor type are also considered errors
	actorID, actorIDOk := ctx.Value(ActorKeyID).(string)
	if !actorIDOk {
		return zero, ErrInvalidActorID
	}

	actorType, actorTypeOk := ctx.Value(ActorKeyType).(string)
	if !actorTypeOk {
		return zero, ErrInvalidActorType
	}

	actorUUID, actorUUIDErr := uuid.Parse(actorID)
	if actorUUIDErr != nil {
		return zero, ErrInvalidActorID
	}

	actorTypeValue := db.GlobalActorType(actorType)
	if !actorTypeValue.Valid() {
		return zero, ErrInvalidActorType
	}

	return ActorDetails{
		ID:   actorUUID,
		Type: actorTypeValue,
	}, nil
}

// as both the actor id and actor type are closely related
// if any one of them causes any error than both the values are considered invalid
func GetActorIdFromContext(ctx context.Context) (uuid.UUID, error) {
	actorDetails, actorDetailsErr := GetActorDetailsFromContext(ctx)
	return actorDetails.ID, actorDetailsErr
}

func GetActorTypeFromContext(ctx context.Context) (db.GlobalActorType, error) {
	actorDetails, actorDetailsErr := GetActorDetailsFromContext(ctx)
	return actorDetails.Type, actorDetailsErr
}
