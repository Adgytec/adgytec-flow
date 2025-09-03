package iam

import (
	"fmt"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
)

func (s *accessManagement) validateActorType(
	currentActorType db_actions.GlobalActorType,
	requiredActorType db_actions.GlobalAssignableActorType,
) error {
	switch string(requiredActorType) {
	case string(db_actions.GlobalAssignableActorTypeAll), string(currentActorType):
		return nil
	default:
		return &app_errors.PermissionDeniedError{
			Reason: fmt.Sprintf("The action requires actor type '%s' but the current actor type is '%s'.", requiredActorType, currentActorType),
		}
	}
}
