package types

import (
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/jackc/pgx/v5/pgtype"
)

type Permission struct {
	Key               string
	ServiceName       string
	Name              string
	Description       pgtype.Text
	RequiredResources []db_actions.GlobalPermissionResourceType
}
