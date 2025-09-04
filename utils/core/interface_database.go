package core

import (
	"context"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/jackc/pgx/v5"
)

type Database interface {
	Queries() *db.Queries
	NewTransaction(context.Context) (pgx.Tx, error)
}

type DatabaseWithShutdown interface {
	Database
	Shutdown()
}
