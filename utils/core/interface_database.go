package core

import (
	"context"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/jackc/pgx/v5"
)

type Database interface {
	Queries() *db_actions.Queries
	NewTransaction(context.Context) (pgx.Tx, error)
}

type DatabaseWithShutdown interface {
	Database
	Shutdown()
}
