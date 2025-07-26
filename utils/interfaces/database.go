package interfaces

import db_actions "github.com/Adgytec/adgytec-flow/database/actions"

type IDatabase interface {
	Queries() *db_actions.Queries
}
