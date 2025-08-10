package core

import (
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/jackc/pgx/v5"
)

type IDatabase interface {
	Queries() db_actions.Querier
	NewTransaction() (pgx.Tx, error)
}

type IDatabaseWithShutdown interface {
	IDatabase
	Shutdown()
}
