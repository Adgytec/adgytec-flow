package actor

import (
	"context"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/google/uuid"
)

func GetActorDetailsFromContext(ctx context.Context) (ActorDetails, error) {
	var zero ActorDetails

	// empty actor id and actor type are also considered errors
	actorID, actorIDOk := ctx.Value(ActorKeyID).(uuid.UUID)
	if !actorIDOk {
		return zero, ErrInvalidActorID
	}

	actorType, actorTypeOk := ctx.Value(ActorKeyType).(db.GlobalActorType)
	if !actorTypeOk || !actorType.Valid() {
		return zero, ErrInvalidActorType
	}

	return ActorDetails{
		ID:   actorID,
		Type: actorType,
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

// NewSystemActorContext returns a context with a "system" actor.
// This is useful during startup, migrations, or any operation
// that requires a transaction but has no real actor.
func NewSystemActorContext(ctx context.Context) context.Context {
	actorIDCtx := context.WithValue(ctx, ActorKeyID, uuid.Nil)
	return context.WithValue(actorIDCtx, ActorKeyType, db.GlobalActorTypeSystem)
}
