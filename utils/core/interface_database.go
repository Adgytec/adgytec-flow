package core

import (
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/jackc/pgx/v5"
)

type IDatabase interface {
	Queries() *db_actions.Queries
	NewTransaction() (pgx.Tx, error)
}
