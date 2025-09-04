package iam

import (
	"fmt"

	"github.com/Adgytec/adgytec-flow/database/db"
)

func (s *iamService) validateActorType(
	currentActorType db.GlobalActorType,
	requiredActorType db.GlobalAssignableActorType,
) error {
	switch string(requiredActorType) {
	case string(db.GlobalAssignableActorTypeAll), string(currentActorType):
		return nil
	default:
		return &PermissionDeniedError{
			Reason: fmt.Sprintf("The action requires actor type '%s' but the current actor type is '%s'.", requiredActorType, currentActorType),
		}
	}
}
