package database

import (
	"context"

	"github.com/Adgytec/adgytec-flow/database/db"
)

type Tx interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type Database interface {
	Queries() *db.Queries
	WithTransaction(ctx context.Context) (*db.Queries, Tx, error)
}

type DatabaseWithShutdown interface {
	Database
	Shutdown()
}
